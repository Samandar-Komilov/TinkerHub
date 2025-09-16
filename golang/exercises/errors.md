# Errors in Go

---

## 1. Errors are values

* In Go, an error is just a value of type `error` (an interface).
* Functions return `(T, error)`. No exceptions, no hidden control flow.

```go
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by zero", a)
	}
	return a / b, nil
}

func main() {
	res, err := Divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", res)
}
```

**Why?**
You can treat errors like any other value: store them, wrap them, compare them. They’re explicit.

**Exercise 1:**
Write a function `OpenFile(name string)` that returns either `"file opened: name"` or an error if name is empty.
Print error and result separately.

---

## 2. Sentinel Errors

* A **sentinel error** is a predefined error value you can compare with `==`.

```go
var ErrNotFound = errors.New("not found")

func LookupUser(id int) (string, error) {
	if id == 0 {
		return "", ErrNotFound
	}
	return "Alice", nil
}

func main() {
	_, err := LookupUser(0)
	if err == ErrNotFound {
		fmt.Println("user not found")
	}
}
```

**Why?**
They’re like constants for common failures.
But overusing them couples code to specific error values → prefer `errors.Is` for flexibility.

**Exercise 2:**
Define `ErrInvalidAge` and `ErrTooOld`. Write `ValidateAge(age int)` that returns one of them depending on age. In `main`, check error type and print message.

---

## 3. Wrapping and Unwrapping Errors

* When propagating errors, you often want to add context without losing the root cause.
* `fmt.Errorf("context: %w", err)` wraps the error.

```go
func ReadConfig() error {
	return errors.New("file not found")
}

func InitApp() error {
	if err := ReadConfig(); err != nil {
		return fmt.Errorf("InitApp failed: %w", err)
	}
	return nil
}

func main() {
	err := InitApp()
	fmt.Println(err) // prints "InitApp failed: file not found"
}
```

**Why?**
Context helps debugging without losing the original error chain.

**Exercise 3:**
Create `LoadUser()` that calls `QueryDB()`. In `QueryDB`, return a plain error. In `LoadUser`, wrap it with `fmt.Errorf("LoadUser: %w", err)`. Print both plain and wrapped error.

---

## 4. errors.Is and errors.As

* `errors.Is(err, target)` → checks if error chain contains a sentinel.
* `errors.As(err, &target)` → checks if error chain contains a specific error type.

```go
var ErrConfig = errors.New("config missing")

type ParseError struct {
	File string
}
func (e ParseError) Error() string { return "parse error in " + e.File }

func Load() error {
	return fmt.Errorf("load failed: %w", ParseError{"config.yaml"})
}

func main() {
	err := Load()

	var pe ParseError
	if errors.As(err, &pe) {
		fmt.Println("as -> file:", pe.File)
	}

	if errors.Is(err, ErrConfig) {
		fmt.Println("is -> config missing")
	}
}
```

**Why?**

* `Is` → good for sentinels.
* `As` → good for custom error types.

**Exercise 4:**
Define a custom error `PermissionError{User string}`.
Return it wrapped from a function.
In `main`, use `errors.As` to extract the `User` and print it.

---

## 5. Wrapping Errors with defer

* Sometimes you want to handle or wrap errors right before a function returns.

```go
func doStuff() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("doStuff failed: %w", err)
		}
	}()
	return errors.New("low-level failure")
}

func main() {
	fmt.Println(doStuff()) // doStuff failed: low-level failure
}
```

**Why?**
Defer is perfect for cleaning up and enriching errors consistently.

**Exercise 5:**
Write `ProcessFile()` that uses `defer` to wrap any returned error with `"ProcessFile error: ..."`. Return both nil and error cases.

---

## 6. Panic and Recover

* **panic** → stops normal execution (like exceptions in other langs).
* **recover** → catches a panic inside a `defer`.

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

**Why not recover everywhere?**

* If you recover globally, you may hide bugs.
* Idiomatic Go: use panic for truly exceptional states (programmer mistakes, unrecoverable failures).
* Convert panic → error only at **boundaries** (e.g., library function → return error).

**Exercise 6:**
Write `ParseConfig(text string)` that panics if text is empty.
Wrap it with `SafeParseConfig` that recovers and returns an error instead.
Test both empty and valid cases.

