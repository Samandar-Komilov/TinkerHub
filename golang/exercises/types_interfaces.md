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

