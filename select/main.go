package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Create channels which will recieve odds and evens
	oddCh := make(chan int)
	evenCh := make(chan int)

	// Kick off goroutines to filter the numbers
	go odds(nums, oddCh)
	go evens(nums, evenCh)

	numRecieved := 0
	for numRecieved < len(nums) {
		// The select will block until one of the channels recieves a value
		select {
		case num := <-oddCh:
			fmt.Printf("%d is odd!\n", num)
			numRecieved++
		case num := <-evenCh:
			fmt.Printf("%d is even!\n", num)
			numRecieved++
			// default:
			// 	fmt.Println("No values are ready")
			// 	time.Sleep(1 * time.Second)
		}
	}
}

func odds(nums []int, c chan int) {
	for _, num := range nums {
		if num%2 == 1 {
			c <- num
		}
		randSleep()
	}
}

func evens(nums []int, c chan int) {
	for _, num := range nums {
		if num%2 == 0 {
			c <- num
		}
		randSleep()
	}
}

func randSleep() {
	d := time.Duration(rand.Intn(100))
	time.Sleep(d * time.Millisecond)
}
