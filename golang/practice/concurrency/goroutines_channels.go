package concurrency

import (
	"fmt"
	"time"
)

func Main_gorochannel() {
	// ex1
	// go numPrint(1, 2)
	// charPrint()

	// ex2 and 3
	// ch1 := make(chan int)
	// go sender1(ch1)
	// go sender2(ch1)
	// fmt.Println(<-ch1)
	// fmt.Println(<-ch1)
	// fmt.Println(<-ch1)
	// fmt.Println(<-ch1)
	// fmt.Println(<-ch1)

	// ex4
	// ch2 := make(chan string)
	// fmt.Println(<-ch2)

	// ex5: for-range stops when channel is closed.
	ch2 := make(chan int)
	go sender1(ch2)
	for num := range ch2 {
		fmt.Println(num)
	}
}

// ex1
func numPrint(from int, till int) {
	for i := from; i <= till; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func charPrint() {
	str := "ab"
	for _, v := range str {
		fmt.Printf("Char: %c\n", v)
		time.Sleep(500 * time.Millisecond)
	}
}

func sender1(ch chan int) {
	for i := range 5 {
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}

func sender2(ch chan int) {
	for i := 10; i < 12; i++ {
		ch <- i
	}
}
