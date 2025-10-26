# Types, Methods and Interfaces

The reason why we allocated a separate file for this topic is, interfaces and type system is the core reason which makes Go different from other languages. Everyone can think about goroutines that the language primarily benefits from, but that's not the case.


## The Big Picture Exercises

**1. Methods with Value vs Pointer Receivers**
* Define a type `Point {x, y int}`.
* Add methods `Move(dx, dy int)` with **value receiver** and `Shift(dx, dy int)` with **pointer receiver**.
* Call both methods and print results.
* See how value receivers work on copies, pointer receivers mutate in place.
* Compare with C structs (you’d pass pointers manually) and with Python/Java where *all methods are pointer-like*.

**2. Nil-Friendly Methods**
* Define a binary tree type:

  ```go
  type Node struct {
      Val int
      Left, Right *Node
  }
  ```
* Add `Insert(v int)` method on `*Node`.
* Make it so that calling `var root *Node; root.Insert(10)` creates the root automatically.
* Realize Go encourages **nil-safe methods** (unlike Java’s NPE).
* Compare with Python where you’d get `AttributeError` unless you check explicitly.

**3. Embedding vs Inheritance**
* Define `type Logger struct {}` with method `Log(msg string)`.
* Define `type Server struct { Logger }`.
* Show that `Server` can call `s.Log("hi")`.
* Then try assigning a `Server` to a `Logger` variable — it won’t work.
* Composition ≠ inheritance.
* Compare with Java: `class Server extends Logger`.
* Compare with Python: multiple inheritance vs composition.

**4. Implicit Interfaces**
* Define `type Stringer interface { String() string }`.
* Implement `String()` for your `Point`.
* Show that you never “declare” a type implements an interface — it just does.
* This is Go’s *duck typing but compile-time checked*.
* Compare:

  * Java → must `implements Stringer`.
  * Python → any `__str__` works (no compile-time check).
  * C → manual function pointers in structs.

**5. Accept Interfaces, Return Structs**
* Write `func PrintShape(s fmt.Stringer)` which takes an interface.
* But `func NewPoint(x, y int) Point` returns a concrete `Point`.
* Callers depend on **interfaces at input**, but you keep **control over output types**.
* Compare: In Java, APIs often return interfaces; in Go, concrete is preferred.

**6. Interfaces and Nil**
* Write `var s fmt.Stringer`.
* Print `s == nil` → true.
* Now assign `var p *Point; s = p`. Print `s == nil` → false, but `p == nil`.
* Interface values are a tuple: `(type, value)`. Nil inside still gives a non-nil interface.
* Compare: In Python `None` is just `None`. In Java, `null` means null for everything.

**7. Empty Interface**
* Write a function `func Dump(v interface{})` that prints type using `fmt.Printf("%T\n", v)`.
* Call it with `int`, `string`, `Point`, etc.
* Empty interface means “I accept anything”, but loses type info (you’ll need type assertions).
* Compare: `void*` in C, `Object` in Java, `Any` in Python.

**8. Type Assertion and Type Switch**
* Extend `Dump(v interface{})` to:
  * Use `v.(int)` to check if it’s an int.
  * Use `switch v := v.(type)` to branch on types.
* Type-safe runtime inspection.
* Compare:
  * C: `void*` + manual casting (dangerous).
  * Java: `instanceof`.
  * Python: `isinstance()`.

**9. Function Types as Interfaces**
* Define `type HandlerFunc func(int)`.
* Define `type Handler interface { Handle(int) }`.
* Write an adapter so a `HandlerFunc` satisfies `Handler`.
* Functions can be used as objects (bridges to interfaces).
* Compare:
  * Java → functional interfaces (lambdas).
  * Python → first-class functions naturally.
  * C → function pointers.

**10. Dependency Injection with Implicit Interfaces**
* Define `type Database interface { GetUser(id int) string }`.
* Implement `RealDB` and `MockDB` (for testing).
* Write a function `func HandleRequest(db Database)`.
* See why Go’s implicit interfaces make testing & DI super easy — no boilerplate `implements`.
* Compare: Java → must create explicit mocks or use frameworks. Python → easy, but no compile-time guarantees. C → would need function pointers everywhere.

**11. Object-Oriented? But Not Really**
* Create a `Shape` interface with `Area()` and `Perimeter()`.
* Implement `Circle` and `Rectangle`.
* Then add a `Square` **without** touching the interface.
* Go lets you extend “hierarchies” without editing the base — no rigid inheritance tree.
* Compare:
  * Java → must extend classes or implement interfaces explicitly.
  * Python → multiple inheritance works, but messy.
  * C → function pointers & structs.

---

# 🌐 Project: Mini Web App Simulator

### Concept

We simulate requests flowing through:

1. **Middleware** (function type implementing interface).
2. **Handlers** (DI: choose different storage/loggers).
3. **Storage backends** (extensible via interfaces, no base “class” editing).

---

### Step 1. Define Core Interfaces

```go
package main

import (
	"fmt"
)

// Request is a simple struct instead of *http.Request
type Request struct {
	Path string
	Body string
}

// Handler processes requests
type Handler interface {
	Handle(req Request)
}
```

---

### Step 2. Function Type as Handler (Exercise 9)

```go
// HandlerFunc is a function that can act like a Handler
type HandlerFunc func(Request)

func (f HandlerFunc) Handle(req Request) {
	f(req) // just call the function
}
```

Now any plain function with signature `func(Request)` can be treated as a `Handler`.
(Aha: like `http.HandlerFunc` in the stdlib.)

---

### Step 3. Storage Interface (Exercise 10: DI)

```go
// Storage is an interface for saving data
type Storage interface {
	Save(data string)
}

// RealStorage saves to "real DB"
type RealStorage struct{}

func (RealStorage) Save(data string) {
	fmt.Println("Saving to DB:", data)
}

// MockStorage is for testing
type MockStorage struct{}

func (MockStorage) Save(data string) {
	fmt.Println("Mock save:", data)
}
```

---

### Step 4. A Request Handler That Uses Storage (DI in Action)

```go
// SaveHandler depends on Storage, but doesn’t create it itself
type SaveHandler struct {
	storage Storage
}

func (h SaveHandler) Handle(req Request) {
	h.storage.Save(req.Body)
}
```

You can now inject either `RealStorage` or `MockStorage`.

---

### Step 5. Extensible Interfaces (Exercise 11)

```go
// Logger is independent; anyone can implement it
type Logger interface {
	Log(msg string)
}

type ConsoleLogger struct{}

func (ConsoleLogger) Log(msg string) { fmt.Println("LOG:", msg) }

type FileLogger struct{}

func (FileLogger) Log(msg string) { fmt.Println("File log:", msg) }
```

Notice: we didn’t touch the `Logger` interface when we added `FileLogger`.
Any new type that has `Log(string)` can be plugged in.

---

### Step 6. Middleware Using Function Types

```go
// Middleware wraps a Handler
type Middleware func(Handler) Handler

// Example middleware: logging
func LoggingMiddleware(logger Logger) Middleware {
	return func(next Handler) Handler {
		// Return a HandlerFunc (function as handler)
		return HandlerFunc(func(req Request) {
			logger.Log("Received request: " + req.Path)
			next.Handle(req)
		})
	}
}
```

---

### Step 7. Putting It Together

```go
func main() {
	// Choose DI storage implementation
	storage := RealStorage{} // or MockStorage{}

	// Choose DI logger implementation
	logger := ConsoleLogger{} // or FileLogger{}

	// Core handler
	saveHandler := SaveHandler{storage: storage}

	// Wrap handler with middleware
	loggedHandler := LoggingMiddleware(logger)(saveHandler)

	// Simulate request
	req := Request{Path: "/submit", Body: "Hello, Go!"}
	loggedHandler.Handle(req)
}
```

---

# ✅ What This Shows

* **9. Functions as interfaces** → `HandlerFunc` lets raw functions behave like objects. Middleware is just a function wrapping a function.
* **10. Dependency injection** → `SaveHandler` doesn’t know if it’s using real DB or mock; you inject the `Storage`. Similarly for loggers.
* **11. Extensible interfaces** → You can add `FileLogger`, `CloudLogger`, `MemoryLogger` without ever touching the `Logger` interface.

---

# 🎯 Exercises in This Project

1. **Add a new middleware** that rejects requests if `req.Body == ""`.
2. **Add a new storage** type (`InMemoryStorage`) that stores values in a slice. Inject it and print the slice after a few requests.

---

👉 Do you want me to extend this into a **toy framework** (like a stripped-down `net/http`) so that later, when you actually learn `net/http`, you’ll recognize the design instantly?
