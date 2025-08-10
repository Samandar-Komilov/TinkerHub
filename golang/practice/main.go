package main

import (
	"fmt"
)

func main() {
	// basics.Loop1to10()

	// var strings = []string{"apple", "orange", "I", "have", "money"}
	// basics.Iterstrings(strings)
	// basics.Switch1(strings[1])

	var data = [3]string{"Eshmat", "Tashkent", "Uzbekistan"}
	fmt.Println(data)
	var ints []int
	fmt.Printf("%p\n", &ints)
	ints = append(ints, 10, 20, 30)

	fmt.Printf("%p\n", &ints)
}
