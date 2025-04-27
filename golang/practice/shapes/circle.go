package shapes

import "fmt"

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// 1. Variadic functions

func Sum(nums ...int) {
	for i, num := range nums {
		fmt.Println(i, num)
	}
}
