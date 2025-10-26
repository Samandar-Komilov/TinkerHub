package basics

import "fmt"

func Main_stringrunes() {
	// 1
	s1 := "Hello, 世界!"
	fmt.Println(s1, len(s1))
	// 2
	for i, r := range s1 {
		fmt.Printf("Character %d: %c\n", i, r)
	}
	// 3

}
