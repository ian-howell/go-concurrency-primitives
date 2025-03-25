package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Create channels which will recieve odds and evens
	oddCh := make(chan int)
	evenCh := make(chan int)

	// Kick off goroutines to filter the numbers
	go odds(nums, oddCh)
	go evens(nums, evenCh)

	for range nums {
		// The select will block until one of the channels recieves a value
		select {
		case num := <-oddCh:
			fmt.Printf("%d is odd!\n", num)
		case num := <-evenCh:
			fmt.Printf("%d is even!\n", num)
		}
	}
}

func odds(nums []int, c chan int) {
	for _, num := range nums {
		if num%2 == 1 {
			c <- num
		}
	}
}

func evens(nums []int, c chan int) {
	for _, num := range nums {
		if num%2 == 0 {
			c <- num
		}
	}
}
