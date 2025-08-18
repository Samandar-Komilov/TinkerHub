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
	// 15
	slc3 := append(slc1, slc2...)
	fmt.Println(slc3)
	// 16
	for i, v := range slc3 {
		if v == 3 {
			fmt.Println("Value 3 found at: ", i+1)
			break
		}
	}
	// 17
	for i, v := range slc3 {
		if v == 3 {
			slc3[i] = -1
		}
	}
	fmt.Println(slc3)
	// 18
	// Array to slice: arr[:] -> makes slice
	// Slice to array is a bit weird

	// 19
	slc4 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(slc4)
	// 20: Stack
	slc2 = push(slc2, 800)
	fmt.Println(slc2)
	var popped int
	slc2, popped = pop(slc2)
	fmt.Println(slc2, popped)

	// 21: Queue
	slc1 = enqueue(slc1, 1000)
	fmt.Println(slc1)
	slc1 = dequeue(slc1)
	fmt.Println(slc1)

	// 22: easy
	// 23: easy
	// 24: check if equal slices
	is_equal := slices.Equal(slc1, slc2)
	fmt.Println(is_equal)
	// 25: easy
	// 26: insert to specific index
	slc1 = insertEfficient(slc1, -10, 0)
	fmt.Println(slc1)
}

// 20
func push(slice []int, val int) []int {
	return append(slice, val)
}

func pop(slice []int) ([]int, int) {
	ln := len(slice)
	return slice[:ln-1], slice[ln-1]
}

// 21
func enqueue(slice []int, val int) []int {
	return append(slice, val)
}

func dequeue(slice []int) []int {
	return slice[1:]
}

// 26
func insertInefficient(slice []int, val int, idx int) []int {
	return append(append(slice[:idx], val), slice[idx+1:]...)
}

func insertEfficient(slice []int, val int, idx int) []int {
	newslice := append(slice, 0)
	copy(newslice[idx+1:], newslice[idx:])
	newslice[idx] = val
	return newslice
}

// 27: Replace a subslice within a slice

// 28: Flatten 2D slice into 1D slice

// 31: Rotate a slice by k positions

// 32: Check if a slice contains duplicates

// 33: Remove all duplicates from a slice

// 34: Convert a slice of integers to a slice of strings

// 37: Sort a slice of strings in alphabetical order

// 38: Shuffle a slice of integers randomly

// 39: Convert a slice into a fixed-size array
