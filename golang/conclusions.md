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


### Generics
Soon...

## Error Handling

Errors are values in Go and returned from functions. The same is used in C programs by they way.

Go has built-in `error` interface with method `Error()`. We can create new errors using `errors.New()`

Sentinel Errors are like our enums of errors in C programs. We were using them in our nginx-clone project, for example `PARSE_ERROR=-712`. Sentinel errors can be put in package level variables or consts. Sentinel errors should be rare.

Since `error` is an interface, you can create your own errors by creating a new type that implements the `error` interface. For example:
```go
type Status int

const (
    InvalidLogin Status = iota + 1
    NotFound
)

type StatusErr struct {
    Status
    Status
    Message string
}
func (se StatusErr) Error() string {
    return se.message
}

// Which can be used as follows
func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
    err := login(uid, pwd)
    if err != nil {
        return nil, StatusErr{
            Status:
            InvalidLogin,
            Message: fmt.Sprintf("invalid credentials for user %s", uid),
        }
    }
    data, err := getData(file)
    if err != nil {
        return nil, StatusErr{
            Status:
            NotFound,
            Message: fmt.Sprintf("file %s not found", file),
        }
    }
    return data, nil
}
```
[!] If you are using your own error type, be sure you don’t return an uninitialized instance.

**Wrapping Errors**
Wrapping and Unwrapping errors, Is and As, Wrapping with defer - research
```go
// Wrapping errors
func ReadConfig() error {
	return errors.New("file not found")
}

func InitApp() error {
	if err := ReadConfig(); err != nil {
		return fmt.Errorf("InitApp failed: %w", err)
	}
	return nil
}
```

- `errors.Is(err, target)` - checks if the error chain contains a sentinel.
- `errors.As(err, &target)` - checks if error chain contains a specific error type.

We can wrap errors with defer, ensuring while the function is finishing with error, we can give additional context:
```go
func doStuff() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("doStuff failed: %w", err)
		}
	}()
	return errors.New("low level failure")
}

fmt.Println(doStuff()) // prints: doStuff failed: low level failure
```

**Panic and Recover** looks much like try-except in Python, but not really. 
Panic stops the normal execution. Recover catches a panic inside a defer and can help us to make it error as value and return to the caller instead of stopping the whole program flow.
```go
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()
	return a / b, nil // panic if b == 0
}

func main() {
	res, err := safeDivide(10, 0)
	fmt.Println("res:", res, "err:", err)
}
```
If we use `recover()` everywhere, we might hide bugs unintentionally. There are cases where `panic()` has to be and some cases where we can `recover()` it.

[!] `recover()` should be used inside of `defer` to ensure that it is executed even if a panic occurs later in the code.

[!] `recover()` function recovers the panic IN THE SAME GOROUTINE as the panic.

## Modules, Packages and Imports

Go code is organized into three main units: repositories, modules and packages.

To be a module, the directory must have `go.mod` file.

While importing, always use absolute path. Using relative path works, but it is not advised.

The name of the package is determined by its **package clause**, not its import path.

**Package names should be descriptive.**  
- Don’t create two functions in a `util` package called ExtractNames and FormatNames. If you do, every time you use these functions, they will be referred to as `util.ExtractNames` and `util.FormatNames`, and that `util` package tells you nothing about what the functions do. 
- It’s better to create one function called `Names` in a package called `extract` and a second function called `Names` in a package called `format`. Then you can call with `extract.Names()` and `format.Names()`.
- Don’t name your function `ExtractNames` when it is in the `extract` package.

You can override the name of the package if two names clash:
```go
import (
    crand "crypto/rand"
    "encoding/binary"
    "fmt"
    "math/rand"
)
```

You need to comment your code using `godoc`. Make sure you comment your code properly. At the very least, any exported identifier should have a comment. Go linting tools such as golint and golangci-lint can report missing comments on exported identifiers. 

Circular dependencies might occur, if your depedencies depend on each other. In that case, we might:
- merge the two packages
- move the dependent code in one of the package to the other

**Gracefully renaming your API**  
We can use type aliasing if we need to rename the existing data, like structs to be seen differently from outside world.

**Working with Modules**
We need to include the full path of the module we are importing, usually they are the paths of github repository.


**Versioning your Module**
If you publish your own module:  
- Start with `v0.x` for experimental.
- `v1.x` → stable API.
- Breaking change → bump to `v2` and change import path (`/v2`).

**Organizing our module**  
There is not a standard way of doing this, but there are a number of precious resources that show different ways of doing it:
- [Alex Edwards Blog](https://www.alexedwards.net/blog/11-tips-for-structuring-your-go-projects)
- [Smart Byte Labs](https://medium.com/@smart_byte_labs/organize-like-a-pro-a-simple-guide-to-go-project-folder-structures-e85e9c1769c2)
- [Melkey Youtube](https://youtu.be/dxPakeBsgl4?si=EzkGMvqAHCXn802U)


---

# Golang Practice Deeper

## Concurrency

### Goroutines, Channels and selection
We should use concurrency during I/O bound tasks most of the time. It is not directly related to "making code faster", but managing resources more efficiently.

**Goroutines.** A goroutine is a function running concurrently.
```go
import (
	"fmt"
	"time"
)

func worker(id int) {
	for i := range 3 {
		fmt.Printf("Worker %d: %d\n", id, i)
		time.Sleep(2 * time.Second)
	}
}

func Main_concurrency1() {
	go worker(1)
	go worker(2)

	time.Sleep(10 * time.Second)
	fmt.Println("Done")
}
```
This code runs 2 independent goroutines each running its own for loop and sleeping for 2 seconds. It will run for 10 seconds and then print "Done".

**Channels.** Channels are typed pipes for communication between channels. They are the core concept behind "don't share memory, communicate instead". Channels are built-in type in Go, just like slices and maps and can be created using `make()` function.
```go
ch := make(chan int) // unbuffered

go func() {
    ch <- 42 // send
}()

val := <-ch // receive
fmt.Println(val) // 42
```
By default channels are unbuffered. Every write to an open, unbuffered channel causes the writing goroutine to pause until another goroutine reads from the same channel. Likewise, a read from an open, unbuffered channel causes the reading goroutine to pause until another goroutine writes to the same channel. This means you cannot write to or read from an unbuffered channel without at least two concurrently running goroutines.

**Buffered Channels.** These channels buffer a limited number of writes without blocking. If the buffer fills before there are any reads from the channel, a subsequent write to the channel pauses the writing goroutine until the channel is read. Just as writing to a channel with a full buffer blocks, reading from a channel with an empty buffer also blocks. A buffered channel is created by specifying the capacity of the buffer when creating the channel:
```go
ch2 := make(chan string, 2)
ch2 <- "A"
ch2 <- "B"

fmt.Println(<-ch2)
ch2 <- "C"          // if I would put that after previous 2 writes, deadlock would happen. So, at least 1 read should occur before I write again.
fmt.Println(<-ch2)
fmt.Println(<-ch2)
```

**for-range over channels.** You can also read from channels using `for-range` loops. 
```go
ch3 := make(chan int, 3)

go func() {
    for i := range 3 {
        ch3 <- i
    }
    close(ch3)
}()

for v := range ch3 {
    fmt.Println("Got:", v)
}
```
When you finish writing to a channel, you should `close()` the channel. After closing, any attempts to write to the channel or close the channel again will panic. Interestingly, attempting to read from a closed channel always succeeds. We can use comma-ok idiom to check if the channel is closed or not.

**What is the difference between buffered and unbuffered channels? Why we need them both?**   
- Unbuffered channels for Synchronization. 
- Buffered channels for Producer-Consumer pattern.

Let's consider the following scenarios one by one.
(1) An application that receives user requests to process and upload a file. The main goroutine handles the user's HTTP request, but it needs to ensure the file is successfully moved to a final, secure location by a separate "file mover" goroutine before it can return a successful response to the user.
```go
package main

import (
	"fmt"
	"time"
)

// A function that simulates moving a file to a final destination.
// It sends a confirmation signal when done.
func moveFile(source, destination string, done chan struct{}) {
	fmt.Printf("File mover: Starting to move file from %s to %s...\n", source, destination)

	// Simulate a time-consuming but crucial operation
	time.Sleep(3 * time.Second)

	fmt.Println("File mover: File move complete. Sending confirmation.")

	// Send an empty struct to signal completion.
	// This send operation will block until a receiver is ready.
	done <- struct{}{}
}

func main() {
	fmt.Println("Main: User request received. Spawning file mover goroutine.")

	// Create an unbuffered channel for synchronization.
	// It has a capacity of 0.
	done := make(chan struct{})

	// Start the background goroutine to move the file.
	go moveFile("/tmp/upload.dat", "/var/secure/data.dat", done)

	fmt.Println("Main: Waiting for file move confirmation...")

	// Block here, waiting for a signal on the 'done' channel.
	// This will block the main function until the 'moveFile' goroutine sends a value.
	<-done

	fmt.Println("Main: File move confirmed! Returning success to the user.")
}

/* Output:
Main: User request received. Spawning file mover goroutine.
Main: Waiting for file move confirmation...
File mover: Starting to move file from /tmp/upload.dat to /var/secure/data.dat...
File mover: File move complete. Sending confirmation.
Main: File move confirmed! Returning success to the user.
*/
```

(2) A web crawler. A single, fast-crawling routine (the producer) identifies new URLs to visit. These URLs are then passed to a pool of slower worker goroutines (the consumers) that handle the actual work of downloading the content. The producer is often much faster than the consumers.
```go
package main

import (
	"fmt"
	"time"
)

// The producer goroutine: finds new URLs to crawl.
func producer(tasks chan string) {
	urls := []string{"url-1", "url-2", "url-3", "url-4", "url-5"}

	for _, url := range urls {
		fmt.Printf("Producer: Found new URL: %s\n", url)
		tasks <- url // Add the URL to the buffered channel.
		// The producer continues immediately, even if no consumer is ready.
		// It only blocks if the buffer is full.
	}
	close(tasks) // Close the channel when all tasks are produced.
}

// The consumer goroutine: downloads and processes the content from a URL.
func consumer(id int, tasks chan string) {
	for url := range tasks {
		fmt.Printf("Consumer %d: Processing URL %s...\n", id, url)
		// Simulate slow, time-consuming work
		time.Sleep(2 * time.Second)
		fmt.Printf("Consumer %d: Finished with URL %s.\n", id, url)
	}
}

func main() {
	// Create a buffered channel with a capacity of 3.
	// This acts as a queue for tasks.
	tasks := make(chan string, 3)

	// Launch a producer goroutine.
	go producer(tasks)

	// Launch 2 consumer goroutines.
	go consumer(1, tasks)
	go consumer(2, tasks)

	// Wait long enough for all tasks to be processed.
	// In a real application, you'd use a sync.WaitGroup.
	time.Sleep(15 * time.Second)
	fmt.Println("Main: All tasks are likely done.")
}

/*
Producer: Found new URL: url-1
Producer: Found new URL: url-2
Producer: Found new URL: url-3
Producer: Found new URL: url-4
Consumer 1: Processing URL url-1...
Consumer 2: Processing URL url-2...
Producer: Found new URL: url-5
Consumer 1: Finished with URL url-1.
Consumer 1: Processing URL url-3...
Consumer 2: Finished with URL url-2.
Consumer 2: Processing URL url-4...
Consumer 1: Finished with URL url-3.
Consumer 1: Processing URL url-5...
Consumer 2: Finished with URL url-4.
Consumer 1: Finished with URL url-5.
Main: All tasks are likely done.
*/
```

**Channel Multiplexing - `select()`**  


### Concurrency Best Practices and Patterns


## Standard Library


## The Context


## Writing Tests