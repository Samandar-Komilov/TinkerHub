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
	// ch2 := make(chan int)
	// go sender1(ch2)
	// for num := range ch2 {
	// 	fmt.Println(num)
	// }

	// ex6, 8
	ch3 := make(chan int, 2)
	ch3 <- 1
	ch3 <- 2
	// ch3 <- 3
	fmt.Println(<-ch3)
	ch3 <- 3
	fmt.Println(<-ch3)
	// ex8
	fmt.Println(len(ch3), cap(ch3))
	fmt.Println(<-ch3)

	// ex11
	ch4 := make(chan int)
	go sender1(ch4)
	for i := range ch4 {
		fmt.Println("Received:", i)
	}

	// ex12
	v, ok := <-ch4
	if !ok {
		fmt.Println("Channel already closed. V:", v)
	}

	// ex13

	// ex14
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered panic:", err)
		}
	}()

	ch4 <- 4
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
		time.Sleep(300 * time.Millisecond)
	}
	close(ch)
}

func sender2(ch chan int) {
	for i := 10; i < 12; i++ {
		ch <- i
	}
}
