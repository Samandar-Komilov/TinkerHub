# SOLIDC Design Analysis and Next Plan

## Scope
This document verifies current implementation of:
- Single Responsibility Principle (SRP)
- Open/Closed Principle (OCP)

And provides an implementation plan for:
- Liskov Substitution Principle (LSP)
- Interface Segregation Principle (ISP)
- Dependency Inversion Principle (DIP)

Code reviewed:
- `SOLIDC/src/include/interfaces.h`
- `SOLIDC/src/include/logger.h`
- `SOLIDC/src/logger.c`
- `SOLIDC/src/controller.c`
- `SOLIDC/src/formatters/*.c`
- `SOLIDC/src/transports/*.c`
- `SOLIDC/src/main.c`

Compile verification:
- `gcc -Wall -Wextra -Wpedantic -Werror -Isrc/include -Isrc/json src/*.c src/formatters/*.c src/transports/*.c src/json/*.c -o bin/trlog -lm` (passes)

---

## 1) Verification: Single Responsibility Principle (SRP)

### What is good (implemented)

1. `Transaction` DTO only holds data.
   - `src/include/dto.h` has no behavior, only structure.

2. Formatting is separated from transport.
   - `Formatter` interface: `int (*format)(...)`
   - `Transport` interface: `int (*send)(...)`
   - Implementations are separate files (`text_formatter.c`, `json_formatter.c`, `disk_transport.c`, `tcp_transport.c`, `udp_transport.c`).

3. `Logger` orchestrates only "format then send".
   - `src/logger.c` does exactly one workflow.

4. `controller.c` handles policy (when/where to log), not formatting/transport internals.

### SRP gaps (small but important)

1. `JSON_FORMATTER` currently does both serialization and heap/resource management.
   - `cJSON` object lifecycle and string buffer conversion are mixed in one function.
   - This is acceptable for now, but if complexity grows (schema versions, validation, redaction), this function will violate SRP quickly.

2. `main.c` mixes user interaction and application flow.
   - Input parsing, loop control, and transaction dispatch are all in `main`.
   - Not critical yet, but it will become hard to test or replace input source.

### SRP verdict
SRP is largely implemented at architecture level (good separation of concerns). Remaining issues are local and can be addressed incrementally.

---

## 2) Verification: Open/Closed Principle (OCP)

### What is good (implemented)

1. Extension through new formatter/transport types is already supported.
   - Add new implementation of `Formatter` or `Transport` without editing `logger.c`.

2. `Logger` depends on abstractions (`Formatter`, `Transport`) and remains unchanged when adding new behaviors.

3. `controller.c` composes concrete pairs (`DISK_TEXT_LOGGER`, `NETWORK_JSON_LOGGER`) instead of embedding formatting/sending logic.

### OCP gaps (where modifications are still required)

1. `components.h` requires editing to register each new concrete implementation.
   - Every new formatter/transport currently forces header modification.

2. `controller.c` must be changed to use newly added components.
   - Composition is hardcoded; there is no runtime config/registration.

3. `config.h` uses compile-time macros only.
   - Switching behavior by environment/user profile requires source change + rebuild.

### OCP verdict
OCP is implemented in the core logger workflow, but not yet in composition/configuration layers. The architecture is extensible, but not fully closed to modification at higher levels.

---

## 3) Plan to implement LSP

## Goal
Any implementation of a `Formatter` or `Transport` should be safely substitutable without breaking caller expectations.

## Current risk
The interfaces do not define strict behavior contracts. Different implementations may return inconsistent status/length semantics.

## Implementation steps

1. Define explicit contracts in `interfaces.h`.
   - `Formatter::format` contract:
     - Returns `>= 0` number of bytes written (excluding `\0`).
     - Returns `< 0` on error.
     - Must never write beyond `buf_size`.
     - Must produce data suitable for transport as raw bytes.
   - `Transport::send` contract:
     - Returns `0` on success, `-1` on failure.
     - Must attempt to send exactly `len` bytes.

2. Normalize all implementations to these contracts.
   - `json_format` currently relies on `snprintf`; guard against truncation explicitly.
   - all transport implementations should treat partial sends as failure unless they loop to completion.

3. Add contract tests (substitutability tests).
   - A single shared test routine should run against every formatter and transport instance.
   - Same inputs, same expected behavior class.

4. Add defensive checks in `log_transaction`.
   - Validate `lg`, `formatter`, `transport`, and `t` pointers.
   - Reject impossible formatter return values (`n >= MAX_BUFFER_SIZE`).

## Example LSP contract snippet

```c
/* interfaces.h */
typedef struct Formatter {
    /*
     * Contract:
     * - returns bytes written [0, buf_size)
     * - returns -1 on failure
     */
    int (*format)(const Transaction *t, char *buf, size_t buf_size);
} Formatter;

typedef struct Transport {
    /* Contract: return 0 success, -1 failure */
    int (*send)(const char *msg, size_t len);
} Transport;
```

## Example LSP violation to avoid
A formatter that returns success but writes invalid/non-deterministic length metadata, causing transport to read garbage. This breaks substitutability because caller assumes stable contract.

## LSP acceptance criteria
- Swapping `TEXT_FORMATTER` with `JSON_FORMATTER` never changes error semantics.
- Swapping `TCP_TRANSPORT` with `UDP_TRANSPORT` preserves success/failure contract.
- `log_transaction` behavior remains consistent for all implementations.

---

## 4) Plan to implement ISP

## Goal
Clients should depend only on methods they use.

## Current risk
Today interfaces are small (good), but growth pressure can cause “fat interface” anti-pattern, for example adding `flush`, `connect`, `set_level`, `close` into `Transport` for only one backend.

## Implementation steps

1. Freeze small interfaces and split by behavior before growth.
   - Keep `TransportSend` minimal.
   - Add optional capability interfaces separately.

2. Introduce focused interfaces.

```c
typedef struct Sender {
    int (*send)(const char *msg, size_t len);
} Sender;

typedef struct Flushable {
    int (*flush)(void);
} Flushable;

typedef struct Connectable {
    int (*connect)(void);
    int (*disconnect)(void);
} Connectable;
```

3. Update high-level modules to depend only on required role.
   - `Logger` should use only `Sender`.
   - A network session manager (if added) can use `Connectable`.

4. Use capability composition in concrete implementations.
   - TCP transport can expose `Sender + Connectable`.
   - Disk transport can expose `Sender + Flushable`.

## Example ISP usage

```c
typedef struct Logger {
    const Formatter *formatter;
    const Sender *sender;   /* not full transport super-interface */
} Logger;
```

## ISP acceptance criteria
- No client includes methods it does not call.
- Adding `flush` support does not force touching `Logger`.
- Disk/UDP/TCP remain independently evolvable.

---

## 5) Plan to implement DIP

## Goal
High-level policy modules depend on abstractions, not concrete details.

## Current state
Partially good:
- `logger.c` already depends on `Formatter` and `Transport` abstractions.

Not yet complete:
- `controller.c` hardcodes concrete globals (`TEXT_FORMATTER`, `DISK_TRANSPORT`, etc.).
- Composition is compile-time, not injected.

## Implementation steps

1. Introduce an application wiring object (`AppContext`).

```c
typedef struct AppContext {
    Logger disk_logger;
    Logger network_logger;
    int (*should_log_on_network)(const Transaction *t);
} AppContext;
```

2. Make controller depend on injected abstractions.

```c
void process_transaction_with_ctx(const AppContext *ctx, const Transaction *t) {
    (void)log_transaction(&ctx->disk_logger, t);
    if (ctx->should_log_on_network(t)) {
        (void)log_transaction(&ctx->network_logger, t);
    }
}
```

3. Move concrete selection to composition root (`main.c` or `bootstrap.c`).
   - main wires `TEXT_FORMATTER + DISK_TRANSPORT`, `JSON_FORMATTER + TCP_TRANSPORT`.
   - alternative builds/tests can wire stubs or mocks.

4. Add test doubles for formatter/transport.
   - fake formatter returns fixed message.
   - fake transport captures last sent bytes.
   - enables unit testing `process_transaction_with_ctx` without network/disk.

5. Optional next step: runtime configuration.
   - choose formatter/transport by config file or CLI.
   - high-level logic unchanged.

## Example DIP outcome
After DIP, switching from TCP to UDP requires only composition change, not controller policy change.

## DIP acceptance criteria
- `controller.c` includes no concrete component headers.
- High-level flow is testable with fakes only.
- Concrete implementations are selected in one composition place.

---

## 6) Recommended implementation order

1. LSP contracts + normalization
   - lowest risk, improves correctness baseline.

2. DIP injection refactor
   - removes high-level concrete coupling, unlocks testing.

3. ISP capability split (only where needed)
   - apply incrementally when new behaviors appear.

This sequence minimizes regressions and avoids premature interface fragmentation.

---

## 7) Minimal concrete backlog

1. `interfaces.h`: add precise behavior contracts in comments.
2. `logger.c`: add null/length defensive checks.
3. `json_formatter.c`: handle allocation/truncation failures explicitly.
4. `controller.[ch]`: add context-based processing API.
5. `main.c`: move object wiring to composition root.
6. `tests/` (new): add contract tests for all formatters/transports and policy tests with fakes.

