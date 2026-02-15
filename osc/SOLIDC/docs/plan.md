# SOLID in C: A 2-Hour Practical Plan

This plan is designed to guide you through learning SOLID principles using C, proving they are fundamental to software engineering and not just OOP concepts.

**Prerequisites:**
- C compiler (gcc/clang)
- Text editor
- No frameworks. No OOP libraries. Just C.

---

## Phase 0: Mental Setup (5 Minutes)

**Goal:** Establish the correct mindset for the exercise.

- [ ] **Adopt the Mindset**: ALL future requirements will be hostile, contradictory, and frequent. You are designing for *change under ignorance*.
- [ ] **The "Pain First" Rule**: Do not apply a principle until you feel the pain of *not* using it.

---

## Phase 1: Single Responsibility Principle (SRP) (25 Minutes)

**Context**: You are building a transaction logging system.

### Step 1: The Initial Implementation (Bad)
- [ ] Create a `logger.c` / `logger.h`.
- [ ] Implement a function `void log_transaction(const char* user, double amount)`.
    - This function should:
    - Format the string (e.g., "User X sent Y").
    - Open a file.
    - Write to the file.
    - Handle write errors.

### Step 2: The Forced Failure (Pain)
- [ ] **New Requirement**: "Logs must also be sent over the network in JSON format, but disk logs must remain as text."
- [ ] Try to implement this in your existing `log_transaction` function.
    - *Notice how you have to touch code that handles file IO just to add JSON formatting.*
    - *Notice how policies (WHAT to log) are mixed with mechanisms (HOW to log).*

### Step 3: Refactor (The Fix)
- [ ] **Goal**: Split by *reason to change*.
- [ ] Create separate modules/functions for:
    - [ ] `formatter.c`: Logic for formatting log messages (Text vs JSON).
    - [ ] `transport.c`: Logic for writing to disk vs sending to network.
    - [ ] `controller.c`: Logic for *when* to log (policy).
- [ ] **Check**: Can you change the log format without touching the file writing code?

**Why:** SRP is about isolating change. If a module has two reasons to change (format vs storage), it will break for two different stakeholders.

---

## Phase 2: Open-Closed Principle (OCP) (25 Minutes)

**Context**: Promoting the logging system to support multiple output targets.

### Step 1: The Initial Implementation (Bad)
- [ ] Modify your logger to take a type enum: `enum LogType { FILE, SYSLOG, STDOUT }`.
- [ ] Use a `switch` statement in your main logging function to handle each case.

### Step 2: The Forced Failure (Pain)
- [ ] **New Requirement**: Add a `NETWORK` logging type.
- [ ] Implement it.
    - *Notice you have to modify the core `switch` statement.*
    - *Notice that every new feature requires editing stable, tested code.*
    - *This is a regression risk.*

### Step 3: Refactor (The Fix)
- [ ] **Goal**: Add new code without editing old code.
- [ ] Define a `struct Logger` with function pointers (e.g., `void (*log)(char*)`).
- [ ] Implement distinct "subclasses" (struct instances) for File, Syslog, Stdout.
    - `file_logger.c`
    - `stdout_logger.c`
- [ ] The core logic should just call `logger->log(msg)` without knowing implementation details.
- [ ] **Check**: Can you add `network_logger.c` without touching the main loop?

**Why:** Stable code should be closed for modification but open for extension. In C, this is done via function pointers (tables of behavior).

---

## Phase 3: Liskov Substitution Principle (LSP) (20 Minutes)

**Context**: The Logger abstraction is now in place.

### Step 1: The Initial Implementation (Bad)
- [ ] Create a "Network Logger" that:
    - Fails if the network is down.
    - Blocks for 5 seconds on timeout.
- [ ] Create a "File Logger" that:
    - Never fails (silently drops on error).
    - Returns immediately.

### Step 2: The Forced Failure (Pain)
- [ ] Swap the File Logger with the Network Logger in your main application.
- [ ] Run the application.
    - *Notice the main application freezes (blocks).*
    - *Notice the main application crashes or errors out unexpectedly.*
    - *The "Logger" interface lied. It promised "logging", but implemented it differently.*

### Step 3: Refactor (The Fix)
- [ ] **Goal**: Enforce semantic compatibility.
- [ ] Define the contract strictly:
    - "Loggers must be non-blocking."
    - "Loggers must handle their own errors."
- [ ] Wrap the Network Logger to spool to a meaningful buffer (non-blocking) or fail silently (to match the contract).
- [ ] **Check**: Can you swap loggers blindly without the main program caring?

**Why:** LSP is about trust. If I point a pointer to your struct, it must behave exactly as the interface implies, regardless of what's under the hood.

---

## Phase 4: Interface Segregation Principle (ISP) (20 Minutes)

**Context**: The Logger API is growing.

### Step 1: The Initial Implementation (Bad)
- [ ] Create a "Mega Logger" struct/interface:
    ```c
    struct Logger {
        void (*log)(char*);
        void (*rotate)();
        void (*close)();
        void (*set_config)(char*);
        void (*get_stats)();
    };
    ```
- [ ] Pass this struct to a simple function `notify_admin(struct Logger* l)` that only needs to *log* a message.

### Step 2: The Forced Failure (Pain)
- [ ] Implement a simple `ConsoleLogger` that doesn't support rotation or config.
- [ ] Pass it to `notify_admin`.
    - *Notice you have to implement dummy functions for `rotate`, `close`, etc., just to satisfy the struct.*
    - *Notice that if you change the `rotate` signature, `notify_admin` might need recompiling even though it doesn't use it.*

### Step 3: Refactor (The Fix)
- [ ] **Goal**: Callers should not depend on what they don't use.
- [ ] Split the interface:
    - `struct ILog { void (*log)(char*); };`
    - `struct ILifecycle { void (*rotate)(); void (*close)(); };`
- [ ] Change `notify_admin` to take `struct ILog*`.
- [ ] **Check**: Does `notify_admin` care if the logger supports rotation?

**Why:** Depending on unused things creates unnecessary coupling.

---

## Phase 5: Dependency Inversion Principle (DIP) (20 Minutes)

**Context**: The core business logic.

### Step 1: The Initial Implementation (Bad)
- [ ] Write a high-level function `process_payment` in `payment.c`.
- [ ] Inside `process_payment`, directly call `fopen`, `fprintf`, or specific functions from your concrete `file_logger.c`.
    - `process_payment` -> depends on -> `file_logger` -> depends on -> `stdio`.
    
### Step 2: The Forced Failure (Pain)
- [ ] **Test**: Try to write a unit test for `process_payment` that runs in memory without creating files.
    - *You can't. It's hardcoded to call `fopen`.*
- [ ] **Port**: Try to move this code to an embedded platform without a filesystem.
    - *You can't. The policy (payment) depends on the mechanism (file).*

### Step 3: Refactor (The Fix)
- [ ] **Goal**: Policy should not depend on mechanism. Both should depend on abstractions.
- [ ] Define an interface (struct of function pointers) `struct IPaymentLogger` at the *high level* (in `payment.h`).
- [ ] Have the low-level `file_logger` *implement* that interface.
- [ ] Inject the implementation into `process_payment` (e.g., via argument or init).
    - `process_payment` calls logic on the interface.
- [ ] **Check**: Can you pass a "FakeLogger" for testing?

**Why:** This inverts the dependency direction. Source code dependencies point against the flow of control.

---

## Final Review

If you followed the steps, you have:
1.  **SRP**: Separated formatting from IO.
2.  **OCP**: Added new features by adding files, not editing switch statements.
3.  **LSP**: Ensured all loggers behave consistently.
4.  **ISP**: Ensured clients only see what they need.
5.  **DIP**: Decoupled high-level policy from low-level detail.

**Conclusion**: You did all this in C. No classes. No inheritance. No "OOP". SOLID is about **change management**, engineered through **dependency control**.
