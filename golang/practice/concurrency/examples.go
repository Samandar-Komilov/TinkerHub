package concurrency

import (
	"fmt"
	"time"
)

func Main_coroexamples() {
	// example 1
	go worker(1)
	go worker(2)

	time.Sleep(1 * time.Second)
	fmt.Println("Done")

	// example 2
	ch1 := make(chan int)
	go func(num int) {
		ch1 <- num
	}(42)
	val := <-ch1
	fmt.Println("Data read from channel:", val)

	// example 3
	ch2 := make(chan string, 2)
	ch2 <- "A"
	ch2 <- "B"

	fmt.Println(<-ch2)
	ch2 <- "C"
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	// example 4
	ch3 := make(chan int, 3)

	go func() {
		for i := range 3 {
			ch3 <- i
		}
		close(ch3)
	}()

	for v := range ch3 {
		fmt.Println("Got:", v)
	}
}

func worker(id int) {
	for i := range 3 {
		fmt.Printf("Worker %d: %d\n", id, i)
		time.Sleep(200 * time.Millisecond)
	}
}
