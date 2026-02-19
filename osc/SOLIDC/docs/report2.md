# SOLIDC Report 2: A Field Story of SOLID in C

I still remember projects from around 20 years ago where C systems started small, then requirements grew wild. The code didn’t fail all at once. It failed slowly: every change became expensive, every fix broke something else, and eventually the team feared touching core files.

This codebase had the same warning signs early, so we used SOLID deliberately. Here is the story, principle by principle.

## 1) SRP: One module, one reason to change

### The problem we faced
In old systems, we used to pack policy, wiring, and I/O behavior into one place. It felt fast for week one, but by month three every change was risky. In this project, the early shape had controller flow tightly tied to concrete implementations.

### The idea to solve it
Split responsibilities so each file answers one question only:
- policy: should we log this transaction to network?
- pipeline: how to format and send?
- wiring: which concrete formatter/transport pair do we run?

### How SRP realizes it
SRP forced us to separate roles clearly:
- `src/controller.c` handles flow and policy.
- `src/logger.c` only runs `format -> send`.
- `src/main.c` is the composition root for concrete wiring.

### Code showing why
- `src/logger.c` now has one job and validates contract boundaries.
- `src/main.c` builds `AppContext` with concrete objects.
- `src/controller.c` executes policy-driven orchestration with injected context.

### How we achieved it
We removed hardcoded concrete loggers from controller and introduced context-based processing APIs:
- `app_context_start`
- `process_transaction_with_ctx`
- `app_context_stop`

Now policy changes do not force transport/wiring edits, and wiring changes do not force policy edits.

---

## 2) OCP: Extend behavior without editing core policy

### The problem we faced
Years ago, adding a new output format usually meant editing the same central file everyone feared. This system had similar risk in composition logic.

### The idea to solve it
Core logic should be closed for modification, open for extension. Add new formatters/transports by implementing interfaces and wiring them, not by rewriting processing flow.

### How OCP realizes it
We preserved abstraction boundaries so high-level logic depends on contracts, not concrete details.

### Code showing why
- `src/include/interfaces.h`: formal abstraction points (`Formatter`, `Sender`, optional capabilities).
- `src/controller.c`: consumes abstractions via `AppContext`.
- `src/main.c`: selects concrete pairings (`TEXT+DISK`, `JSON+TCP`) without modifying controller flow.

### How we achieved it
By moving concrete selection to composition root and keeping controller/logger abstraction-driven, we can extend implementations with minimal impact to stable logic.

---

## 3) LSP: Swapping implementations must not change correctness

### The problem we faced
In older codebases, two modules could share the same function pointer signature but behave differently: one truncates silently, another returns partial success. Same shape, different promises. That breaks systems in production.

### The idea to solve it
Define and enforce behavioral contracts, not just type signatures.

### How LSP realizes it
LSP says any implementation of the same abstraction must be safely substitutable. That means return semantics and boundary handling must align.

### Code showing why
- `src/include/interfaces.h`: contracts documented for `format` and `send`.
- `src/logger.c`: rejects invalid states and impossible formatter lengths.
- `src/formatters/text_formatter.c`: handles null/truncation explicitly.
- `src/formatters/json_formatter.c`: handles cJSON allocation/object failures and truncation.
- `src/transports/tcp_transport.c`: sends full payload using loop.
- `src/transports/udp_transport.c`: success only if full datagram length is sent.
- `src/transports/disk_transport.c`: strict full-write semantics.

### How we achieved it
We normalized success/failure behavior across all implementations and enforced contracts at boundary points. Now swapping formatter/transport does not silently change correctness expectations.

---

## 4) ISP: Clients should depend only on what they use

### The problem we faced
I’ve seen huge transport interfaces grow over years: `send`, `flush`, `connect`, `disconnect`, `reconnect`, `stats`... then every consumer depended on all of it, even when it only needed one call.

### The idea to solve it
Split capabilities before the interface gets fat.

### How ISP realizes it
Use small role-specific interfaces:
- `Sender`
- `Flushable`
- `Connectable`

Clients take only what they need.

### Code showing why
- `src/include/interfaces.h`: role split done.
- `src/include/logger.h`: `Logger` depends only on `Sender`.
- `src/transports/disk_transport.c`: exposes `DISK_FLUSHABLE`.
- `src/transports/tcp_transport.c`: exposes `TCP_CONNECTABLE`.

### How we achieved it
We kept logger lean and exposed optional capabilities in composed `Transport` objects. So lifecycle behavior exists where needed, but does not leak into every client.

---

## 5) DIP: High-level policy should not know low-level details

### The problem we faced
The classic failure pattern: business logic directly names concrete implementations. Then changing infrastructure forces business-layer edits. Tests become slow and brittle.

### The idea to solve it
Invert dependencies: high-level modules depend on abstractions, low-level modules plug into them.

### How DIP realizes it
`AppContext` carries abstract dependencies (loggers, policy function, optional lifecycle capabilities). Controller consumes that context, not concrete globals.

### Code showing why
- `src/include/controller.h`: defines `AppContext` and context-based APIs.
- `src/controller.c`: no concrete formatter/transport usage.
- `src/main.c`: concrete wiring only in one place.

### How we achieved it
We refactored from hardcoded static loggers in controller to injected context:
- before: controller chose concrete components directly.
- now: controller executes policy using injected abstractions.

That made the high-level flow portable and testable independent of transport/formatter implementations.

---

## Final reflection
Twenty years ago, the hard lesson was this: architecture debt is not loud at the beginning. It grows quietly as coupling. SOLID is not about decoration; it is about protecting change.

In this codebase, the result is practical:
- clearer ownership of responsibilities,
- safer substitution behavior,
- smaller client contracts,
- and high-level flow independent from infrastructure details.

That is exactly the kind of design that survives evolving requirements.
