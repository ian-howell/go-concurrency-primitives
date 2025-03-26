package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a buffered channel with a capacity of 3
	ch := make(chan int, 3)

	// Preload the channel with some values
	for i := 1; i <= 3; i++ {
		ch <- i
		fmt.Printf("Pushed %d into the channel\n", i)
	}

	// Channel now looks something like this:
	// OUTPUT <- [ 1 2 3 ] <- INPUT

	go func() {
		// This will block until the channel is read
		ch <- 4
		fmt.Println("Pushed 4 into the channel")

		close(ch)
	}()

	fmt.Println("Sleeping to make a point")
	fmt.Println()
	time.Sleep(1 * time.Second)

	for num := range ch {
		fmt.Printf("Received %d from the channel\n", num)
	}
}
