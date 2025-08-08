package basics

import "fmt"

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
