package basics

import "fmt"

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
