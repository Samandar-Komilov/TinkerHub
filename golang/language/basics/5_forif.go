package basics

import "fmt"

func Loop1to10() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func Iterstrings(slice []string) {
	for _, str := range slice {
		fmt.Println(str)
	}
}

func Switch1(fruit string) {
	switch fruit {
	case "apple":
		fmt.Println("This is apple!")
	case "banana":
		fmt.Println("This is banana!")
	default:
		fmt.Println("This is another fruit!")
	}
}
