Below is a **2-hour forced-learning session**.
No OOP. No frameworks. Works in **C-level thinking**.
Each block introduces **pain first**, then the principle as a *constraint that removes the pain*.
This structure comes directly from how the principles were derived in industry failures, not pedagogy.

Sources grounding the structure:

* Robert C. Martin, *Design Principles and Design Patterns*
* Parnas, “On the Criteria To Be Used in Decomposing Systems into Modules”
* *Clean Architecture* — Martin
* *A Philosophy of Software Design* — John Ousterhout
* Linux kernel design docs (procedural, C-based modularity)

---

# 2-Hour SOLID Session (Constraint-Driven)

## Setup (5 minutes)

Mental rule:
You are designing **for change under ignorance**.
Assume future requirements are hostile, frequent, and partially contradictory.

No code yet.

---

## Block 1 — S: Single Responsibility (25 minutes)

### Problem

Design a **transaction logging system** in C.

Responsibilities mixed:

* Format log entries
* Decide *when* to log
* Write to disk
* Rotate files
* Handle failures

Now simulate this change:

> “Logs must also be sent over the network in JSON, but disk logs stay text.”

### Forced failure

You will realize:

* Every change touches the same functions
* Logic changes and policy changes are interwoven
* Testing requires filesystem + network even for formatting

### Constraint learned

**SRP = one axis of change per module**
Not “one job”. One *stakeholder-driven reason to change*.

Parnas defined this in 1972:

> Modules should hide decisions likely to change.

Source: Parnas, 1972.

### Rewrite rule (still no code)

Split by *reasons to change*:

* Log policy (when, what)
* Log formatting
* Log delivery mechanisms

That is SRP without OOP.

---

## Block 2 — O: Open–Closed (25 minutes)

### Problem

You now support log outputs:

* File
* Syslog
* Network
* Stdout

Initial C-style solution:

```c
switch (log_type) { ... }
```

### Forced failure

Add a new output:

* Every addition requires modifying existing logic
* High-risk files are repeatedly edited
* Regression surface grows

### Constraint learned

**Stable code must not be edited for extension**.

OCP predates OOP. It is about **linkage direction**, not inheritance.

Source: Martin, *Design Principles and Design Patterns*.

### Rewrite rule

Invert control using:

* Function pointers
* Tables of behavior
* Registration at startup

You now *add files*, not edit core logic.

That is OCP in C.

---

## Block 3 — L: Liskov Substitution (20 minutes)

### Problem

You define a “Logger” abstraction:

* Some loggers fail on network timeout
* Some block
* Some drop messages silently

Caller assumes:

* Logging is non-blocking
* Logging never fails fatally

### Forced failure

Replace one logger with another → system behavior breaks without compiler errors.

### Constraint learned

**Substitution must preserve behavioral contracts, not signatures**.

LSP is about **semantic compatibility**.

Source: Barbara Liskov, 1987.

### Rewrite rule

Document and enforce:

* Timing guarantees
* Failure modes
* Side effects

If behavior differs, it is **not the same abstraction**.

Go’s tiny interfaces and Linux kernel “expectations” follow this rule.

---

## Block 4 — I: Interface Segregation (20 minutes)

### Problem

You expose a single “Logger API”:

* init
* flush
* rotate
* close
* stats
* configure

Most callers only need `log(message)`.

### Forced failure

* Callers depend on functions they don’t use
* Changes to rotation break unrelated code
* Testing requires stubbing irrelevant functions

### Constraint learned

**Dependency creates coupling even if unused**.

Source: Martin, ISP definition.

### Rewrite rule

Split APIs by **caller needs**, not implementer convenience:

* Logging interface
* Lifecycle interface
* Maintenance interface

Linux kernel uses this heavily via separate structs of function pointers.

---

## Block 5 — D: Dependency Inversion (20 minutes)

### Problem

Core business logic calls:

* `fopen`
* `write`
* `send`
* `printf`

### Forced failure

* Impossible to test without real IO
* Business rules change when infrastructure changes
* Porting to another OS touches core logic

### Constraint learned

**Policy must not depend on mechanisms**.

Source: Martin, *Clean Architecture*.

### Rewrite rule

* Core logic depends on abstract capabilities
* Infrastructure implements them
* Wiring happens at program boundary (`main`)

This is how kernels, compilers, and databases are structured.

---

## Final Integration (5 minutes)

### What you should now see

SOLID is not:

* OOP
* Classes
* Java interviews

SOLID is:

* **Change isolation theory**
* **Coupling direction control**
* **Semantic stability constraints**

Every principle answers one question:

> “What kind of change should *not* break this?”

---

## Why this makes code professional

Professional code:

* Survives requirement volatility
* Admits replacement without surgery
* Localizes damage

SOLID is the vocabulary for enforcing that in advance.

This is not optional at scale.
Languages only change how early you feel the pain.
