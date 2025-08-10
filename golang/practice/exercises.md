## Variables and Data Types

1.  Declare a variable `x` of type `int` and assign it the value `42`. Print its value.
2.  Declare two variables, `name` (a `string`) and `age` (an `int`), and initialize them in a single line. Print a sentence using these variables.
3.  Declare a constant `pi` with the value `3.14159`. Print its value.
4.  Use the short variable declaration operator `:=` to declare and initialize a variable `isCool` of type `bool` with the value `true`. Print its value.
5.  Declare a variable `pointerToInt` as a pointer to an integer. Initialize it to point to a new integer with the value `100`. Print both the memory address and the value it points to.
6.  Create a variable `f` of type `float64` and assign it `3.5`. Then, declare a new variable `i` of type `int` and assign it the result of converting `f` to an integer. Print both values.
7.  Create a program that declares a variable `y` of type `int` and a variable `z` of type `int32`. Try to assign `y` to `z` and observe the compiler error. Then, fix it by using an explicit type conversion.
8.  Declare a constant `x` with the value `10` and a constant `y` with the value `20`. Use `iota` to assign these values without explicitly writing them.
9.  Create a program that declares a variable `a`, an `int`. Assign it the value `100`. Then, declare a variable `b` and initialize it with the address of `a`. Print both the value of `a` and the value stored at the address in `b`.
10. Declare a variable of type `complex128` with a real part of `2.5` and an imaginary part of `1.5`. Print its value.

## Functions

11. Write a function named `printGreeting` that takes a `string` argument (a name) and prints "Hello, [name]!". Call it with your name.
12. Create a function named `add` that takes two `int` arguments and returns their sum. Call the function and print the result.
13. Write a function named `calculator` that takes two `int` arguments and returns two `int` values: their sum and their product.
14. Write a function named `variadicSum` that takes a variable number of `int` arguments and returns their sum. Call it with `1, 2, 3, 4, 5`.
15. Create a function `namedReturn` that takes two `int` arguments and returns their sum and difference using named return values.
16. Write a function `swap` that takes two `*int` arguments and swaps the values they point to.
17. Write a function `makeMultiplier` that takes an `int` argument `x` and returns a function that takes an `int` argument `y` and returns `x * y`.
18. Write a function `deferExample` that prints `A`, then `C`. Use `defer` to print `B` between them.
19. Create a function `applyFunc` that takes an `int` and a function as arguments. The function argument should take an `int` and return an `int`. `applyFunc` should return the result of calling the function with the integer argument.
20. Define a `struct` named `Circle` with a `radius` field of type `float64`. Write a method for `Circle` named `Area` that calculates and returns the circle's area.

## Loops and Conditionals

21. Use a `for` loop to print the numbers from 1 to 10.
22. Write an infinite `for` loop that prints "Looping..." and then uses `break` to exit the loop after 5 iterations.
23. Write a `for` loop that iterates over a slice of strings and prints each string.
24. Use a `for` loop to calculate the sum of numbers from 1 to 100.
25. Use a `for` loop to print even numbers between 1 and 20.
26. Write an `if-else` statement that checks if a number `n` is positive, negative, or zero.
27. Write an `if` statement with a short declaration that checks if a value returned from a function is an error.
28. Write a `switch` statement that checks the value of a string variable `fruit`. Print a different message for "apple," "banana," and a default case.
29. Create a `switch` statement that checks a number `grade` and prints "Pass" if `grade` is 60 or above, and "Fail" otherwise. Use a `case` without an expression.
30. Write a `switch` statement with multiple cases that checks the day of the week. For "Monday" through "Friday," print "Weekday." For "Saturday" and "Sunday," print "Weekend."

## Arrays and Slices

31. Declare an array of 3 strings and initialize it with your name, city, and country. Print the entire array.
32. Declare a slice of integers and append the numbers `10, 20, 30` to it. Print the slice.
33. Create a slice `s` with a length of 5 and a capacity of 10 using `make`. Print the length and capacity.
34. Given a slice `s1 := []int{1, 2, 3, 4, 5}`, create a new slice `s2` that contains elements from index 1 to 3 of `s1`. Print `s2`.
35. Given a slice `s := []int{10, 20, 30}`, append the value `40` to it. Print the new slice.
36. Create a slice literal `numbers := []int{1, 2, 3, 4, 5}`. Use a `for-range` loop to print each element and its index.
37. Given a slice `s := []int{10, 20, 30, 40, 50}`, remove the element at index 2 (value `30`). Print the resulting slice.
38. Write a function that takes a slice of integers and returns a new slice containing the squares of each number.
39. Create an array of 4 integers and then create a slice that references the entire array. Modify an element in the slice and observe if the change is reflected in the array.
40. Declare a `nil` slice of integers. Check if it's `nil` and print a message accordingly.

## Mixed and Advanced Concepts

41. Write a function that takes a slice of integers and returns two values: the count of even numbers and the count of odd numbers.
42. Create a map where keys are `string` (names) and values are `int` (ages). Add three entries and then print the age of one person.
43. Define a `struct` named `Book` with `Title` (string), `Author` (string), and `Pages` (int) fields. Create an instance of `Book` and print its details.
44. Create a method `String()` for the `Book` struct that returns a formatted string representation of the book.
45. Write a function that takes an interface and uses a `switch` statement to determine its underlying type (e.g., `string`, `int`, `bool`).
46. Create a goroutine that prints numbers from 1 to 5. The `main` function should print "Hello" and then exit. Use a `time.Sleep` to ensure you see the goroutine's output.
47. Write a program that uses a channel to communicate between two goroutines. One goroutine should send the value `42` to the channel, and the other should receive it and print it.
48. Create a `struct` named `Rectangle` with `width` and `height` (both `float64`). Define an `interface` named `Shape` with a single method `Area() float64`. Make `Rectangle` satisfy the `Shape` interface.
49. Write a program that uses a `defer` statement to print "Exiting program." at the very end of the `main` function, even if an error or panic occurs.
50. Write a program that declares a `map[string]int`. Try to access a key that doesn't exist and use the "comma ok" idiom to check for its presence.