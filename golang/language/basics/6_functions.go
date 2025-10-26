package basics

import (
	"errors"
	"fmt"
)

func Main_functions() {
	// 4: Create a function that returns result or an error
	res, err := divide(64, 8)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	// 5: Named return values
	a, b := getCoordinates()
	fmt.Println(a, b)

	// 7: Sum variables
	s := sumAll(1, 2, 3)
	fmt.Println(s)
}

func divide(divident, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("no division to zero")
	}
	return divident / divisor, nil
}

func getCoordinates() (x, y int) {
	x, y = 1, 2
	return
}

func sumAll(numbers ...int) int {
	s := 0
	for _, v := range numbers {
		s += v
	}
	return s
}

// Old exercises

func VariadicSum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
		fmt.Println("===", sum, num)
	}

	return sum
}

func NamedReturn(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func Anonym() {
	result := func(a, b int) int {
		return a + b
	}(3, 4)

	multiply := func(a, b int) int {
		return a * b
	}

	fmt.Println("Result:", result)
	fmt.Println("Multiply:", multiply(4, 5))

	// Results: 7, 20
}

func Closure1() {
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()

	fmt.Println("Counter:", counter()) // 1
	fmt.Println("Counter:", counter()) // 2
	fmt.Println("Counter:", counter()) // 3
}
