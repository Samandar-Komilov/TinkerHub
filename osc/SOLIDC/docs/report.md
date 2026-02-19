# SOLIDC Report (Problem -> Solution)

This report explains what was changed to finish SOLID-focused architecture in this project, why each change was needed, and what issue it solves.

## 1. SRP (Single Responsibility Principle)

### Problem
- High-level flow, transport details, and object wiring were mixed.
- `controller.c` knew concrete formatter/transport implementations.
- If policy changed, controller and wiring were coupled.

### What was done
1. Kept `Logger` responsible only for `format -> send` (`src/logger.c`).
2. Moved concrete wiring to composition root (`src/main.c`).
3. Kept policy in controller (`src/controller.c`) via `should_log_on_network` and app-flow methods.

### Without this
- Policy updates and infrastructure updates would affect the same module.
- Harder testing and more regression risk.

### Why this matches SRP
- Each module now has one reason to change:
  - `logger.c`: logging pipeline behavior
  - `controller.c`: transaction policy/flow
  - `main.c`: dependency wiring

---

## 2. OCP (Open/Closed Principle)

### Problem
- Core flow should not be edited for new format/transport combinations.
- Earlier controller had hardcoded concrete objects.

### What was done
1. Preserved abstraction-based logger (`Formatter`, `Sender`).
2. Made high-level flow consume injected context (`AppContext`) instead of hardcoded concrete components.
3. Added capability composition (`Transport` contains optional `flushable/connectable`) so new capabilities are extension points, not mandatory edits in logger/controller.

### Without this
- Adding/changing transport/formatter pair would require editing high-level flow code.

### Why this matches OCP
- High-level behavior modules are closed for modification and open to extension through new implementations and wiring.

---

## 3. LSP (Liskov Substitution Principle)

### Problem
- Different implementations could return inconsistent results.
- Partial sends and truncation risks could violate caller assumptions.

### What was done
1. Defined stronger behavioral contracts in `src/include/interfaces.h` comments.
2. Added defensive checks in `log_transaction` (`src/logger.c`):
   - null dependency checks
   - invalid length checks (`n < 0 || n >= MAX_BUFFER_SIZE`)
3. Normalized formatter semantics:
   - `text_formatter.c`: checks null, checks truncation
   - `json_formatter.c`: checks allocation/object failures and truncation
4. Normalized sender semantics:
   - `tcp_transport.c`: loops until full payload is sent
   - `udp_transport.c`: success only when sent length equals requested length
   - `disk_transport.c`: strict full-write check

### Without this
- Swapping one formatter/transport for another could silently break behavior.

### Why this matches LSP
- Any implementation can be substituted while preserving expected success/failure and size semantics.

---

## 4. ISP (Interface Segregation Principle)

### Problem
- A single large transport interface would force clients to depend on methods they do not use.

### What was done
1. Split transport roles in `src/include/interfaces.h`:
   - `Sender`
   - `Flushable`
   - `Connectable`
2. Kept `Logger` dependent only on `Sender` (`src/include/logger.h`).
3. Added optional capabilities by implementation:
   - disk provides `Flushable`
   - tcp provides `Connectable`

### Without this
- `Logger` would depend on `connect/flush/disconnect` even though it only sends.

### Why this matches ISP
- Clients now depend only on the methods they actually use.

---

## 5. DIP (Dependency Inversion Principle)

### Problem
- Controller depended on concrete low-level modules.
- Business flow could not be tested independently of infrastructure.

### What was done
1. Added `AppContext` abstraction in `src/include/controller.h`.
2. Refactored controller to consume injected dependencies:
   - `app_context_start`
   - `process_transaction_with_ctx`
   - `app_context_stop`
3. Moved concrete selection to `src/main.c`:
   - `TEXT_FORMATTER + DISK_TRANSPORT.sender`
   - `JSON_FORMATTER + TCP_TRANSPORT.sender`

### Without this
- High-level code changes required when low-level transport/formatter changed.

### Why this matches DIP
- High-level logic depends on abstractions (`Logger`, `Sender`, policy fn), while low-level details are wired externally.

---

## Added practical capabilities

- `flush` capability for disk transport (`src/transports/disk_transport.c`)
- `connect/disconnect` capability for tcp transport (`src/transports/tcp_transport.c`)
- lifecycle hooks in controller to use optional capabilities (`app_context_start` / `app_context_stop`)

These were added to demonstrate ISP + DIP cleanly without introducing runtime plugin complexity.

---

## Build verification

Compiled successfully with strict flags:

`gcc -Wall -Wextra -Wpedantic -Werror -Isrc/include -Isrc/json src/*.c src/formatters/*.c src/transports/*.c src/json/*.c -o bin/trlog -lm`
