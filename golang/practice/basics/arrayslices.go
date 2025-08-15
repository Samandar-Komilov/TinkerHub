package basics

import "fmt"

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
}
