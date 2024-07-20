package main

import "fmt"

// f takes a value from the right channel, increments it by 1 and
// stores it the left channel to simulate data flowing around the
// chain
func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	const n = 1000
	leftmost := make(chan int)
	left := leftmost
	right := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	go func(c chan int) {
		c <- 1
	}(right)
	fmt.Println(<-leftmost)

}
