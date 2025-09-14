# Golang Practice Quick

## Standard Library
- slices
- strings
- maps
- sort

## Basics


### Packages and Modules, Compilation

**The "main" Function's Special Role:** In Go, func main() is the entry point of an executable program. When you compile a Go program, the linker looks for a single, unique func main() to begin execution. A package that contains a main function is called the main package.
**Packages are Compilation Units:** A Go package is the primary unit of code organization and compilation. All files within a single package are compiled together as a single entity. They share a single namespace, and all declarations (functions, types, variables) in one file are visible to all other files in that same package.

**Can I build a `.go` file as standalone, like I do in C?** The answer is no, not in the way you might think of it in C. In Go, the build process is always *package-centric*:
- The Go compiler compiles packages, not individual files. When you run go build, it takes all the .go files in a directory that belong to a single package and compiles them together into a single object file.
- After compilation, the linker combines these object files with any other packages your program depends on, creating a final executable. This is where the main function becomes important—the linker needs a single entry point to create an executable.

**When I run `go mod tidy` after importing third party package, where does it save the package after download? Is there any virtual environment for Go?** By default, Go modules are stored in the $GOPATH/pkg/mod directory. This directory is not a global space, but rather a cache of modules that are used by your Go projects. Go doesn't have a direct equivalent to Python's virtualenv, but it does have a concept of a "module cache" that serves a similar purpose. The module cache is a directory that stores copies of dependencies required by your Go projects.

**How to name packages in Go? If everything is package-centric and there have to be a single `main.go` file with main package, how can I place multiple files in the same directory as `main.go`?**
The simplest and most important rule for naming Go packages is: a package's name should be the same as the directory it resides in. There are a few key exceptions and best practices:
- The main Package: This is the only exception to the rule. An executable program must reside in a directory with a package named main, and this package contains the func main() entry point. The directory name itself can be anything (e.g., cmd/myapp), but the package name must be main.
- Lowercase Naming: Package names should be all lowercase and should not contain underscores (_). Use short, concise, and descriptive names. For example, use http instead of http_handler and server instead of my_go_server.
In terms of many files in the same directory as `main.go`, let's see this example:
```
myapp/
├── main.go
├── server.go
└── routes.go
```
In short, all files in this directory must have `package main`. They are all compiled together into a single executable (recall that there is package unit, not file in Go). The functions and variables declared in one file (e.g., startServer() in server.go) are directly accessible to all other files in the same package (e.g., main.go) without needing an import statement.

**Quite confusing... Can you explain the packages and modules in Go?**
Well, let's consider this topic a bit more in depth.
- A package is a collection of source files in the same directory that all start with the same package declaration.
- A module is a collection of one or more packages, defined by a single go.mod file at its root. The go.mod file defines the module's import path (its name) and its external dependencies. The module is the unit of versioning and dependency management. When you use go get to install a third-party library, you are pulling in a module.

There is more:
- Module Root: The `go.mod` file at the top-level directory (`my-web-app`) declares this entire directory as a module. Its name is the import path `github.com/myuser/my-web-app`.
- Package Naming: Each subdirectory is a package, and its name matches the directory name.
- Importing Packages: To use code from one package in another, you must import it. The import path is relative to the module root.

**What are Exported and Unexported code in Go?**
In Go, a name is exported if it begins with a capital letter. For example, ``Pizza`` is an exported name, as is `Pi`, which is exported from the `math` package. `pizza` and `pi` do not start with a capital letter, so they are not exported.


### Functions, Variables and Data Types
**What are Named Return Values?**
Go's return values may be named. If so, they are treated as variables defined at the top of the function:
```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```
A return statement without arguments returns the named return values. This is known as a "naked" return.

**We already know simple variable declarations, multiple initializations, they work like in Python. But what is short variable declaration?**
Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type:
```go
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```
Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

The data types include: bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, byte, rune, float32, float64, complex64, complex128, string and error.

More about strings later...

### Composite Types: Arrays and Slices

- [Everything about Golang Arrays](https://www.kelche.co/blog/go/golang-arrays/)
- [Go Slices and its Reallocation Strategy](https://medium.com/@arjun.devb25/understanding-gos-slice-data-structure-and-its-growth-pattern-48fe6dd914b4)
- [make() function usages](https://www.zetcode.com/golang/builtins-make/)
	- includes the usage of `make()` with slices, maps and channels
- [Array to Slice and Slice to Array Conversions](https://labex.io/tutorials/go-how-to-slice-arrays-correctly-418936)


### Maps
- [Go Maps Official Blog](https://go.dev/blog/maps)


### Structs
- [GoByExample Official](https://gobyexample.com/structs)
- [GoByExample Struct Embedding](https://gobyexample.com/struct-embedding)
- [Anonymous Structs](https://www.willem.dev/articles/anonymous-structs/)
- [Struct Field Validations](https://leapcell.io/blog/go-validation-complex-structs)


## Pointers, Custom Types and Interfaces

### Pointers

C had `NULL = (*void) 0`, but `nil` in Go is different. `nil` is a separate type which represent the absence of the value. It is defined in universal block, meaning we can shadow it if we create a new variable named `nil`.

Go has allocator function similar to `calloc()` in C: `new()` and `newArray()`.

Pointer to string `*string` cannot be assigned to a string literal, but to only pointer to a string on heap. We need to create a separate function that returns address of the string which is allocated on heap in background due to stack escape:
```go
type person struct {
    FirstName  string
    MiddleName *string
    LastName   string
}

p := person{
    FirstName: "Eshmat",
    MiddleName: "Gishmatovich", // this won't compile
    MiddleName: stringp("Gishmatovich") // but this do
    LastName: "Toshmatov"
}

func stringp(s string) *string {
    return &s
}
```

Idiomatic Go prefers immutability: write a function that returns a new instance instead of accepting a pointer to mutable variable and mutating it. However, if you have at least one pointer mutating method in some struct, follow the same rule to all of the methods.

Be careful while using pointers in Go, the more you work with pointers, the more Garbage Collector suffers.

Avoid using maps for input parameters, instead use structs.

Use slices as buffers

In Go, stack size can change while program is running. Because each goroutine has its own "stack" which is in heap, and managed by go compiler.

Garbage Collector suffers if you use (slices of) pointers to structs instead of slice of structs. Because each pointer is located in distinct position and while iterating through a single slice, even though it feels linear, we have to go and find the actual data based on addresses.

Go's Garbage Collector works on **low latency** mode, not the biggest garbage first mode. This ensures that we don't have to wait for the GC to run before program execution. It is also more efficient in terms of memory usage.


### Types and Interfaces

Go types are beyond structs: we can create a type from a function!
```go
// Anonymous function is being assigned to a variable in the function type
type Score int
type Converter func(string) Score

func main() {
    fmt.Println("Hello, 世界")
    var c Converter = func(s string) Score {
        return Score(len(s))
    }
    result := c("hello")
    fmt.Println(result) // prints: 5
}
```

Abstract and Concrete types are strictly separated in Go which is mixed in classic OOP languages like Java.

Type definitions and method implementations should be in the same file.

If one of the methods use pointer receivers, it is better to use pointer receivers for all to be consistent.

No need to write getter-setters in Go.

Method have to support `nil` instances. For example, when you write a `Tree` data structure, you have to support the case when `Tree` is empty (`nil`).

**Functions versus Methods** When to use which?
- Methods, if your function depends on external data which is not given in parameters list.
- Functions, if your function depends only on its parameters.
Any time your logic depends on values that are configured at startup or changed while your program is running, those values should be stored in a struct and that logic should be implemented as a method.

Package level state should be immutable in Go.

**Type Declarations are not Inheritance in Go.** Declaring a type based on another type looks a bit like inheritance, but it isn’t. The two types have the same underlying type, but that’s all. There is no hierarchy between these types. In languages with inheritance, a child instance can be used anywhere the parent instance is used. The child instance also has all the methods and data structures of the parent instance. That’s not the case in Go. You can’t assign an instance of type HighScore to a variable of type Score or vice versa without a type conversion, nor can you assign either of them to a variable of type int without a type conversion.

**Types are documentation in Go.** That's why we might benefit from using `Percentage` custom type instead of `int` everywhere which is vague: why is it integer?

**No enums, but iota in Go.** Go does not have enums, but we can benefit from `iota`. The iota-based enumerations only make sense when you care about being able to differentiate between a set of values, and don’t particularly care what the value is behind the scenes. If you care about the value, each constant has a specific value that is also used somewhere else, you should write constants by hand.
```go
const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)
```
Here, we don't necessarily need what value is `Spam`, but we should know it is not `Personal` for example. So, such usages are preferred only in internal code.

**Embedding is not Inheritance in Go!**
In traditional OOP, inheritance creates an "is-a" relationship with dynamic dispatch. The Puppy has-a Dog, it doesn't is-a Dog. The embedded type becomes a field, and Go automatically promotes its fields and methods to the outer type (following example).

**There is no dynamic dispatch in Go.** The methods on the embedded field have no idea they are embedded, so they stay "loyal" to their struct. This means:
```go
// Define some types
type Animal struct {
    name string
}

func (a Animal) Speak() string {
    return "Some generic animal sound"
}

func (a Animal) Info() string {
    return fmt.Sprintf("Animal name: %s", a.name)
}

type Dog struct {
    Animal  // Embedding, NOT inheritance
    breed string
}

func (d Dog) Speak() string {
    return "Woof!"
}

// Puppy embeds Dog (which embeds Animal)
type Puppy struct {
    Dog
    age int
}

func (p Puppy) Speak() string {
    return "Yip yip!"
}

func main() {
    fmt.Println("=== Embedding Example ===")
    
    puppy := Puppy{
        Dog: Dog{
            Animal: Animal{name: "Rex"},
            breed:  "Golden Retriever",
        },
        age: 3,
    }
    
    fmt.Printf("Puppy name: %s\n", puppy.name)       // Promoted field
    fmt.Printf("Puppy breed: %s\n", puppy.breed)     // Promoted field
    fmt.Printf("Puppy age: %d\n", puppy.age)
    
    // Method calls - which one gets called?
    fmt.Printf("Puppy speaks: %s\n", puppy.Speak())     // Puppy's method
    fmt.Printf("Puppy info: %s\n", puppy.Info())        // Animal's method (promoted)
    
    // Direct access to embedded types
    fmt.Printf("Dog speaks: %s\n", puppy.Dog.Speak())     // Dog's method
    fmt.Printf("Animal speaks: %s\n", puppy.Animal.Speak()) // Animal's method
    
    fmt.Println("\n=== No Dynamic Dispatch ===")
```

Why Go made this decision? Go's creators wanted to avoid the complexity of inheritance hierarchies:
```go
// In inheritance-based languages, this can be confusing:
// Which method gets called? What's the method resolution order?
grandparent.method() // ??

// In Go, it's always explicit:
puppy.Animal.Speak()  // Clearly calls Animal's method
puppy.Dog.Speak()     // Clearly calls Dog's method  
puppy.Speak()         // Clearly calls Puppy's method (most specific)
```


**Implicit Interfaces.** While Go’s concurrency model (which we cover in Chapter 10) gets all of the publicity, the real star of Go’s design is its implicit interfaces, the only abstract type in Go.
- A concrete type does not declare that it implements an interface. If the method set for a concrete type contains all of the methods in the method set for an interface, the concrete type implements the interface.
Debate: No interfaces, but duck typing (Python, Ruby, JS) vs explicit interfaces (Java, C++). Why Go chose middle?
**RESEARCH: 201-205 pages**


You can also embed interfaces in interfaces.

