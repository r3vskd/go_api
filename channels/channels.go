package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Worker %d sending %d\n", id, i)
		ch <- i
		time.Sleep(time.Millisecond * 500)
	}
	close(ch) // Close the channel when done
}

func receiver(ch <-chan int) {
	for val := range ch {
		fmt.Printf("Received: %d\n", val)
	}
}

func main() {
	// Unbuffered channel
	fmt.Println("Using unbuffered channel:")
	unbufferedChannel := make(chan int)
	go worker(1, unbufferedChannel)
	receiver(unbufferedChannel)

	fmt.Println("\nUsing buffered channel:")
	bufferedChannel := make(chan int, 3)
	go worker(2, bufferedChannel)
	receiver(bufferedChannel)

	fmt.Println("\nUsing multiple workers:")
	multiChannel := make(chan int)
	for i := 1; i <= 3; i++ {
		go worker(i, multiChannel)
	}
	go func() {
		time.Sleep(time.Second * 3)
		close(multiChannel)
	}()
	for val := range multiChannel {
		fmt.Printf("Main received: %d\n", val)
	}
}
