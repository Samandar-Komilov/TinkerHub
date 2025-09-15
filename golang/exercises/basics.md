# Go Practice Exercises - Basics

### Variables and Data Types
> - `var` vs `:=`
> - zero values
> - `const` and `iota`
> - scope and shadowing
> - integers, floats, complex, boolean, type conversion

1. Declare an integer using `var` and another using `:=`, then print both.  
2. Declare an uninitialized float variable and print its zero value.  
3. Declare a constant string and print it.  
4. Swap two variables without using a temporary variable.  
5. Declare a boolean variable and print its opposite value.  
6. Use `iota` to create three constants representing `Low`, `Medium`, `High`.  
7. Declare a complex number and print its real and imaginary parts.  
8. Convert an integer to a float64 and vice versa.  
9. Declare a variable inside a function and shadow it in a nested block.
10. Declare a variable of type `rune` and print it as a character and Unicode value.  
11. Use `iota` to define constants for binary flags (`Read=1`, `Write=2`, `Execute=4`).  
12. Declare multiple variables in a single line using `var` and `:=`.  
13. Declare an untyped constant and assign it to both `int` and `float64` variables.  
14. Write a function that returns multiple values (sum and product of two integers).  
15. Declare a zero-value struct and print its fields.  
16. Use type inference to declare a variable without explicitly stating its type.  
17. Write a function that takes an empty interface and checks if it’s an integer.  
18. Declare a custom type `Celsius` (float64) and convert it to `Fahrenheit`.  
19. Print the type of a variable using reflection (`reflect.TypeOf`).  
20. Declare a constant using an expression (`const x = 10 / 2`) and print it.  
  

### Arrays and Slices
> - arrays
> - slices
> - capacity and growth
> - make()
> - conversions: slice -> array and array -> slice

1. Declare an array of 3 integers and print it.  
2. Declare a slice of strings and append two elements to it.  
3. Print the length and capacity of a slice after appending elements.  
4. Create a slice from an existing array and modify it.  
5. Use `make()` to create a slice with length 5 and capacity 10.  
6. Copy elements from one slice to another using `copy()`.  
7. Slice a given slice to get the first 3 elements.  
8. Slice a given slice to get elements from index 2 to the end.  
9. Append one slice to another.  
10. Remove the first element from a slice.  
11. Remove the last element from a slice.  
12. Remove an element at a specific index from a slice.  
13. Check if a slice is empty.  
14. Reverse a slice of integers.  
15. Concatenate two slices into a new slice.  
16. Find the index of a specific value in a slice.  
17. Replace all occurrences of a value in a slice.  
18. Convert an array to a slice and vice versa.  
19. Create a slice with `...` (ellipsis) from a list of values.  
20. Use a slice as a stack (push and pop operations).  
21. Use a slice as a queue (enqueue and dequeue operations).  
22. Filter a slice to keep only even numbers.  
23. Double each element in a slice of integers.  
24. Compare two slices to check if they are equal.  
25. Split a slice into two parts at a given index.  
26. Insert a new element at a specific index in a slice.  
27. Replace a subslice within a slice.  
28. Flatten a 2D slice into a 1D slice.  
29. Create a slice of slices (2D slice) and print it.  
30. Extract every second element from a slice.  
31. Rotate a slice left by `n` positions.  
32. Check if a slice contains any duplicates.  
33. Remove all duplicates from a slice.  
34. Convert a slice of integers to a slice of strings.  
35. Sum all elements in a slice of floats.  
36. Find the maximum value in a slice of integers.  
37. Sort a slice of strings in alphabetical order.  
38. Shuffle a slice randomly.  
39. Convert a slice into a fixed-size array (handle size mismatch).  
40. Write a function that takes a slice and returns a new slice with squared values.


### Strings, Runes and Bytes
> - strings
> - runes
> - bytes
> - raw string literals and interpreted string literals

1. Declare a string variable and print its length.  
2. Iterate over a string and print each character as a rune.  
3. Convert a string to a slice of bytes and print it.  
4. Convert a slice of bytes back to a string and print it.  
5. Declare a raw string literal (with backticks) and print it.  
6. Declare an interpreted string literal (with double quotes) and print it.  
7. Concatenate two strings using `+` and print the result.  
8. Use `strings.Contains()` to check if a string contains a substring.  
9. Use `strings.ReplaceAll()` to replace all occurrences of a substring.  
10. Split a string into a slice of words using `strings.Fields()`.  
11. Join a slice of strings into a single string using `strings.Join()`.  
12. Convert a string to uppercase using `strings.ToUpper()`.  
13. Trim whitespace from a string using `strings.TrimSpace()`.  
14. Check if a string starts with a prefix using `strings.HasPrefix()`.  
15. Check if a string ends with a suffix using `strings.HasSuffix()`.  
16. Count the number of occurrences of a substring in a string.  
17. Extract a substring from a string using slicing.  
18. Convert an integer to a string using `strconv.Itoa()`.  
19. Convert a string to an integer using `strconv.Atoi()` (handle error).  
20. Repeat a string `n` times using `strings.Repeat()`.  
21. Compare two strings lexicographically using `strings.Compare()`.  
22. Pad a string with spaces to a fixed width.  
23. Reverse a string by converting it to a slice of runes.  
24. Check if a string is a palindrome (reads the same backward).  
25. Replace the first occurrence of a substring in a string.  
26. Split a string by a custom delimiter using `strings.Split()`.  
27. Convert a string to a slice of runes and print each rune.  
28. Count the number of runes (Unicode characters) in a string.  
29. Remove all occurrences of a substring from a string.  
30. Write a function that takes a string and returns a new string with vowels removed.  


### Maps
> - maps
> - comma-ok idiom

1. Declare an empty map from `string` to `int`.  
2. Initialize a map with key-value pairs and print it.  
3. Add a new key-value pair to an existing map.  
4. Delete a key from a map.  
5. Check if a key exists in a map using the "comma-ok" idiom.  
6. Iterate over a map and print all key-value pairs.  
7. Get the value for a key, returning a default if it doesn’t exist.  
8. Count the frequency of words in a string using a map.  
9. Merge two maps into a new map.  
10. Check if two maps are equal (have the same key-value pairs).  
11. Copy a map to a new map (avoid reference issues).  
12. Clear all entries from a map.  
13. Implement a set using a map (keys only, no values).  
14. Check if a slice contains duplicates using a map.  
15. Find the intersection of two slices using a map.  
16. Group strings in a slice by their length using a map.  
17. Convert a map into a slice of keys.  
18. Convert a map into a slice of values.  
19. Sort the keys of a map and print them in order.  
20. Write a function that inverts a map (keys become values and vice versa).  


### Structs
> - structs
> - JSON tags
> - struct embedding

1. Define a simple struct for a Person with Name and Age fields
2. Create an instance of the Person struct and print it
3. Add a method to Person that returns a greeting string
4. Create a nested Address struct within Person
5. Initialize a struct using field names explicitly
6. Initialize a struct without field names (positional)
7. Compare two struct instances for equality
8. Add JSON tags to Person fields for marshaling
9. Marshal a Person struct to JSON
10. Unmarshal JSON into a Person struct
11. Create an anonymous struct and use it
12. Embed Address struct within Person (struct embedding)
13. Access embedded struct fields directly
14. Add methods to embedded struct
15. Create interface that your struct satisfies
16. Use pointer receiver for struct method
17. Create slice of struct instances
18. Sort slice of structs by different fields
19. Create map with struct values
20. Pass struct to function by value vs by pointer
21. Modify embedded struct field
22. Add tags for XML encoding
23. Marshal struct to XML
24. Create constructor function for your struct
25. Implement Stringer interface for your struct
26. Use struct as map key
27. Create zero-value struct instance
28. Add validation to struct fields
29. Create method with value receiver
30. Compare performance of value vs pointer receivers
31. Create interface that accepts your struct
32. Embed multiple structs
33. Handle naming conflicts in embedded structs
34. Create struct with private and public fields
35. Access unexported fields from same package
36. Create struct with function field
37. Pass struct to goroutine
38. Create channel of struct type
39. Add tags for database operations
40. Benchmark struct field access methods

### Conditionals and loops
> - if-else
> - switch
> - for (4 ways)

1. Basic if-else with simple condition
2. If with short statement before condition
3. If-else if-else chain
4. Switch with constant cases
5. Switch with expression cases
6. Switch with no condition (like if-else chain)
7. Switch with fallthrough
8. Type switch with interface{}
9. Basic for loop (single condition)
10. Traditional for loop (init; condition; post)
11. Infinite for loop with break
12. For-range loop over slice
13. For-range loop over map
14. For-range loop ignoring index/value
15. Nested loops with labels and break/continue


### Functions and Closures
> - functions
> - variadic functions
> - multiple and named return values
> - anonymous functions
> - closures
> - call by value and call by reference

1.  **Create a function `greet()` that prints "Hello, World!" and call it from `main`.**
2.  **Write a function `printSum(a int, b int)` that calculates and prints the sum of two integers.**
3.  **Create a function `isEven(num int) bool` that returns `true` if the number is even.**
4.  **Write a function `divide(dividend, divisor float64) (float64, error)` that returns the result and an error if divisor is zero.**
5.  **Create a function `getCoordinates() (x, y int)` with named returns for x and y, assign values to them, and return.**
6.  **Call the `divide` function from exercise 4 but ignore the error value using an underscore `_`.**
7.  **Write a function `sumAll(numbers ...int) int` that sums a variable number of integers.**
8.  **Create a function `printAll(values ...interface{})` that prints each value and its type.**
9.  **Create a slice of strings `names := []string{"Alice", "Bob"}`. Pass it to a variadic function `printNames(names ...string)`.**
10. **Assign an anonymous function `func() { fmt.Println("I'm anonymous") }` to a variable `f` and then call `f()`.**
11. **Write an Immediately Invoked Function Expression (IIFE) that prints "Running now!" as soon as it's defined.**
12. **Write a function `makeMultiplier(factor int) func(int) int` that returns a function which multiplies its input by the factor.**
13. **Create a function `counter() func() int` that returns a closure. Each call to the closure should increment and return a counter variable.**
14. **Write a closure `func(prefix string) func(string)` that returns a function which prints the prefix and a message passed to it.**
15. **Create two different counters using the `counter()` function from exercise 13. Demonstrate they have independent state.**
16. **Write a function `newCounter() func() int` where the closure modifies a variable in the outer function's scope.**
17. **Create a function `apply(numbers []int, op func(int) int) []int` that applies the operation `op` to each number in the slice.**
18. **Write a function `createGreeter(greeting string) (func(string), func(string))` that returns two functions: one that shouts the greeting and one that whispers it.**
19. **Write a function `tryToChange(val int)` that tries to modify its integer argument. Show that the original variable in `main` is unchanged.**
20. **Write a function `changeForReal(ptr *int)` that modifies the integer value at the given pointer.**
21. **Create a struct `Rectangle` with `Width` and `Height`. Write a method `Area() int` for it. Compare calling `rect.Area()` vs. calling a standalone function `calcArea(rect Rectangle) int`.**
22. **Write a recursive function `factorial(n int) int` to calculate the factorial of n.**
23. **Write a function `readFile(filename string)` that uses `defer` to ensure a "file read attempted" message is printed at the end.**
24. **Write a function that has three `defer` statements printing numbers 1, 2, 3. Show they run in LIFO (Last-In-First-Out) order.**
25. **Write a function `getFileInfo(name string) (size int, err error)` with named returns. Use `defer` to log the size and error before returning.**
26. **Write a function `trickyDefer() (result int)`. Inside, set `result = 10`, then use a `defer` to change `result++`. See what it returns.**
27. **Create a function `riskyOperation()` that panics with a message "Something went terribly wrong!".**
28. **Write a function `safeCall(f func())` that calls `f` and uses `recover` to catch any panic, printing the panic value.**
29. **Write a function `getStringLength(val interface{}) (int, error)` that uses a type assertion to get the length if `val` is a string, or returns an error otherwise.**
30. **Write a generic function `getFirst[T any](slice []T) T` that returns the first element of any type of slice.**
31. **Write a generic function `areEqual[T comparable](a, b T) bool` that checks if two values of any comparable type are equal.**
32. **Modify the `divide` function from exercise 4 to return a clear error message like "cannot divide by zero".**
33. **Call the `divide` function and handle the error by printing it, or print the result if there is no error.**
34. **Write a function `makeAdder() func(int) int` that returns a function which adds a specific number to its input.**
35. **Write a function `processRequest(ctx context.Context, data string)` that checks if `ctx.Done()` is received and aborts early.**
36. **Write a function `updateSlice(s []string, index int, value string)` that modifies the element at the given index.**
37. **Write a function `updateMap(m map[string]int, key string, value int)` that sets a key-value pair in the map.**
38. **Write a function `sendMessage(ch chan<- string, msg string)` that sends a message into the channel.**
39. **Write a function `describeShape(s Shape)` that takes an interface parameter and calls its `Area()` method.**
40. **Write a function `doubleIt(num *int)` that modifies the integer pointed to by `num` by multiplying it by 2.**
41. **Write a function `createCopy(s []int) []int` that returns a new slice with the same elements as the input.**
42. **Use the `testing` package to write a benchmark for a function that calculates the nth Fibonacci number.**
43. **Write a function `logValues(values ...interface{})` that uses a type switch to handle `int`, `string`, and `bool` differently.**
44. **Write a function `acceptAnything(val interface{})` that can be called with any type of argument.**
45. **Write a function `createPerson(name string, age int) Person` that returns a struct of type `Person`.**
46. **Write a function `createNumber() *int` that returns a pointer to a new integer.**
47. **Write a function `getReader() io.Reader` that returns an interface type.**
48. **Write a function `createResultChannel() chan int` that returns a new channel of integers.**
49. **Write a function `timeFunction(f func()) time.Duration` that measures and returns how long the function `f` takes to execute.**


### Pointers: with structs, maps, slices and functions
> - pointers
> - struct pointers
> - map pointers
> - slice pointers
> - pointers in functions: call by reference
> - memory management and garbage collection

1. Declare an integer and create a pointer to it. Print the pointer address and dereferenced value.  
2. Swap two integers using pointers.  
3. Declare a pointer to a float64, allocate memory using `new()`, and assign a value.  
4. Pass an integer to a function by reference and modify it.  
5. Return a pointer to a local variable from a function (is this safe?).  
6. Create a `Person` struct and pass it to a function by pointer to modify fields.  
7. Write a method with a pointer receiver that updates a struct field.  
8. Compare passing a struct by value vs. by pointer in terms of performance.  
9. Initialize a struct using `&` shorthand and modify its fields.  
10. Create a slice of struct pointers and modify elements.  
11. Pass a slice to a function and modify an element (does it change the original?).  
12. Pass a map to a function and modify a key-value pair.  
13. Append to a slice inside a function (does it affect the original?).  
14. Return a pointer to a slice from a function.  
15. Compare modifying slices vs. modifying slice pointers in functions.  
16. Create a pointer to an array element and modify it.  
17. Check if a pointer is `nil` before dereferencing.  
18. Assign a pointer to another pointer and verify if they point to the same memory.  
19. Use `runtime.GC()` to force garbage collection and observe behavior.  
20. Create a memory leak by holding unnecessary pointers (e.g., global slice references).  
21. Implement a linked list using pointers.  
22. Create a function that returns a pointer to a new `[]int` slice.  
23. Pass a pointer to an interface and modify it.  
24. Use `unsafe.Pointer` to convert between different pointer types (carefully!).  
25. Benchmark struct method calls with value vs. pointer receivers.  
26. Write a function that takes a pointer to a pointer (`**int`) and modifies the underlying value.  
27. Create a circular reference between two structs using pointers (does GC clean it up?).  
28. Simulate a "pointer to a pointer" in Go (e.g., `var pp **int`).  
29. Use a pointer to a function and call it.  
30. Implement a simple reference counter using pointers.  
31. Explain why `&[]int{1,2,3}[0]` is invalid in Go.  
32. Modify a struct field via reflection using a pointer.  
33. Check if two pointers point to the same struct instance.  
34. Pass a pointer to a goroutine and observe concurrency issues.  
35. Compare `new()` vs. `make()` in terms of memory allocation.  
36. Benchmark passing large structs by value vs. by pointer.  
37. Pre-allocate a slice of pointers to avoid repeated allocations.  
38. Use a pointer to a large array to avoid copying.  
39. Implement a simple object pool using pointers.  
40. Observe memory usage changes when pointers are held vs. released.  


### OOP in Go: struct methods
> - struct methods
> - pseudo-inheritance using embedding
> - pointer and value receivers

1. Define a `Rectangle` struct with `Area()` and `Perimeter()` methods.  
2. Add a `Scale(factor float64)` method with a pointer receiver.  
3. Compare method calls on a struct vs. a struct pointer.  
4. Write a method that modifies a struct’s private field (same package).  
5. Define a method on a struct that returns another struct.  
6. Create an `Animal` struct and embed it in a `Dog` struct.  
7. Override an embedded method in the child struct.  
8. Access an embedded struct’s fields directly.  
9. Embed multiple structs and resolve naming conflicts.  
10. Use embedded structs to simulate "mixins."  
11. Write a method that must use a pointer receiver (modifies state).  
12. Write a method that can use a value receiver (read-only).  
13. Benchmark method calls with pointer vs. value receivers.  
14. Mix pointer and value receivers in the same struct (when is it allowed?).  
15. Explain why some interfaces require pointer receivers.  
16. Implement a `String()` method for custom printing.  
17. Use struct embedding to implement a "base class" pattern.  
18. Create an interface that your struct satisfies.  
19. Simulate polymorphism using interfaces and embedded structs.  
20. Implement a simple "factory" function for struct initialization.  
21. Compose structs to avoid deep inheritance.  
22. Use embedded structs to add "traits."  
23. Implement a fluent interface using pointer receivers.  
24. Create a method that returns `interface{}`.  
25. Simulate "private" methods using lowercase naming.  
26. Implement a simple cache using struct methods.  
27. Create a counter struct with thread-safe methods.  
28. Define a struct with methods that can be chained.  
29. Use embedding to extend a third-party struct.  
30. Compare Go’s OOP approach with classical inheritance.  


### Interfaces (moved to a separate file)


### Generics
> - generic functions
> - generic types and interfaces
> - type constraints and inference

1. Write a generic function that swaps two values of any type.  
2. Create a generic `Stack` data structure that works with any type.  
3. Implement a generic `Max` function that returns the larger of two comparable values.  
4. Write a generic `Filter` function that takes a slice and a predicate function.  
5. Define a generic `Map` function that applies a transformation to each element in a slice.  
6. Create a generic `Contains` function that checks if a value exists in a slice.  
7. Implement a generic `Keys` function that extracts all keys from a map.  
8. Write a generic `Merge` function that combines two maps.  
9. Define a generic `BinarySearch` function for sorted slices.  
10. Create a generic `Pair` struct that holds two values of different types.  
11. Write a generic `Reduce` function that accumulates values from a slice.  
12. Implement a generic `Set` type using a map with empty struct values.  
13. Define a generic `Optional` type that can hold a value or be empty.  
14. Write a generic `Chunk` function that splits a slice into smaller chunks.  
15. Create a generic `LinkedList` with type-safe nodes.  
16. Implement a generic `Cache` that stores key-value pairs.  
17. Write a generic `Zip` function that combines two slices into pairs.  
18. Define a generic `Clamp` function that restricts a value within a range.  
19. Create a generic `Queue` with enqueue and dequeue operations.  
20. Write a generic `Equal` function that checks if two slices are identical.  


### Error Handling
> - `error` interface
> - wrapping and unwrapping errors
> - `panic` and `recover`
> - stack traces and debugging

1.  **Basic Error Creation:** Write a function `Divide(a, b float64) (float64, error)` that returns an error if `b` is zero. Use `errors.New` to create the error.
2.  **Error Checking:** Call your `Divide` function and check the error using a standard `if err != nil` block. Print the result or the error.
3.  **Sentinel Errors:** Create a package-level error variable `var ErrNotFound = errors.New("not found")`. Write a function `FindUser(id int) (*User, error)` that returns this error if a user doesn't exist.
4.  **Checking Sentinel Errors:** Call `FindUser` and use `errors.Is(err, ErrNotFound)` to check if the specific "not found" error occurred, handling it differently from other errors.
5.  **Error Wrapping:** Write a function `ReadFile(path string) ([]byte, error)` that attempts to read a file. If `os.Open` fails, wrap its error using `fmt.Errorf("could not open file: %w", err)`.
6.  **Unwrapping Errors:** Call your `ReadFile` function. Use `errors.Unwrap(err)` to get the original underlying error that was wrapped.
7.  **Custom Error Type:** Create a custom error type `ValidationError` struct with a `Field` string and `Msg` string. Implement the `Error() string` method for it.
8.  **Checking Custom Error Types:** Write a function that returns a `ValidationError`. Call it and use `errors.As(err, &valErr)` to check for and extract the specific `ValidationError`.
9.  **Panic and Recover:** Write a function that panics with a string message. In your `main` function, use `defer` and `recover()` to catch the panic, print the message, and prevent a crash.
10. **Defer with Error:** Write a function that opens a file and uses `defer` to close it. Ensure the file is closed even if an error occurs during processing.
11. **Multiple Errors:** Write a function that performs two operations that can fail (e.g., open a file and decode its JSON). Return the first error encountered, if any.
12. **Ignoring Errors (The Right Way):** Use the `os.Create` function. Intentionally ignore the error result using the blank identifier `_` because you are just practicing and assume the temp directory always exists. *(Note: This is rarely recommended in real code!)*
13. **Error Context:** Write a function that processes a configuration file. Wrap any errors (from reading or parsing) to add context, e.g., `fmt.Errorf("processing config: %w", err)`.
14. **Simple Retry Logic:** Write a function that calls a flaky `apiCall() error` function. If it returns an error, retry the call once before giving up.
15. **Type Assertion for Errors:** Call a function that returns a standard library error (e.g., `os.Open`). Use a type assertion to check if the error is of type `*os.PathError` and print its `Path` field.
16. **Nil Error Value:** Write a function `MaybeReturnsError() error` that conditionally returns an error or `nil`. In the caller, handle both cases correctly.
17. **Error Logging:** Instead of just printing an error from `ReadFile`, use `log.Printf("operation failed: %v", err)` to log it with a timestamp.
18. **Adding Context to Existing Error:** You are given an error `err` from an external library. Create a new error that provides more context: `fmt.Errorf("login failed for user %s: %w", username, err)`.
19. **Handling Specific Syscall Errors:** Call `os.Open` on a non-existent file. Use `os.IsNotExist(err)` to check for this specific condition and handle it gracefully.
20. **Cleanup on Error:** Write a function that creates a temporary file. If a later step in the function fails, use a deferred function to clean up (remove) the temporary file.


### Code Organization
> - modules and dependencies
> - packages: import rules, 3rd party usage
> - publishing modules

1. Initialize a new Go module with `go mod init`.  
2. Add an external dependency to your module.  
3. Upgrade a dependency to its latest version.  
4. Vendor dependencies using `go mod vendor`.  
5. Import a local package within the same module.  
6. Create a package with internal visibility (`internal/`).  
7. Publish a module to GitHub and use it in another project.  
8. Replace a dependency with a local fork in `go.mod`.  
9. Use build tags to conditionally compile code.  
10. Write a script to automate module version tagging.  

### Concurrency
> - goroutines
> - channels: select, buffered and unbuffered
> - worker pools
> - `sync` package: mutexes and waitgroups

### Testing and Benchmarking
> - `testing` package
> - table-driven tests
> - mocks and stubs

