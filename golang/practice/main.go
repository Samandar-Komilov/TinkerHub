package main

import "fmt"

func main() {
	cnt := 0

	for {
		fmt.Println("Looping...")
		cnt++
		if cnt == 5 {
			break
		}
	}
}
