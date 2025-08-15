package basics

import (
	"fmt"
	"slices"
)

func Main_arrslice() {
	// 1
	var arr = [...]int{1, 2, 3}
	fmt.Println(arr)
	//2
	var strings1 = []string{"hello", "hi"}
	strings1 = append(strings1, "why", "are", "you")
	fmt.Println(strings1)
	// 3
	fmt.Println(len(strings1), cap(strings1))
	// 4
	slc1 := arr[:]
	slc1 = append(slc1, 10, 20)
	fmt.Println(slc1)
	// 5
	slc2 := make([]int, 5, 10)
	fmt.Println(slc2, len(slc2), cap(slc2))
	// 6
	copy(slc2, slc1)
	fmt.Println(slc2)
	// 9
	slc2 = append(slc2, slc2...)
	fmt.Println(slc2)
	// Removing elements
	// 10
	slc2 = append(slc2[:0], slc2[1:]...)
	fmt.Println(slc2)
	// 11
	slc2 = append(slc2[:len(slc2)-1], slc2[len(slc2):]...)
	fmt.Println(slc2)
	// 12
	slc2 = append(slc2[:5], slc2[6:]...)
	fmt.Println(slc2)
	// 13
	if slc2 == nil {
		fmt.Println("Slice is nil")
	}
	// 14
	slices.Reverse(slc2)
	fmt.Println(slc2)
}
