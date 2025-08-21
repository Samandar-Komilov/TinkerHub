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
