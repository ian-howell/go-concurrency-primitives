package main

import "fmt"

/*
In the context of software, upstream is a system that sends data
into the collaborative system.  Whereas downstream is a system
that receives data from the collaborative system.

Pipelining is when various stages are chained together
Each stage receives data through an upstream inbound channel
Each stage pushes its transformed/handled data through a downstream
channel via outbound channels
*/
func main() {
	genChan := numInRange(1, 1_000_000)
	sq := squared(genChan)
	for number := range sq {
		fmt.Println(number)
	}

	// since both the upstream and downstream channels are of the same type
	// they can composed n times
	for n := range squared(squared(squared(numInRange(1, 10)))) {
		fmt.Println(n)
		/*
			1
			256
			6561
			65536
			390625
			1679616
			5764801
			16777216
			43046721
		*/
	}

}

// generateNumbersInRange asynchronously generates numbers in a range
func numInRange(inclusiveLow, exclusiveHigh int) <-chan int {
	downstream := make(chan int)
	go func() {
		for i := inclusiveLow; i < exclusiveHigh; i++ {
			downstream <- i
		}
		close(downstream)
	}()
	return downstream
}

// squareroot asynchronously squares the numbers
func squared(upstream <-chan int) <-chan int {
	downstream := make(chan int)
	go func() {
		for n := range upstream {
			downstream <- n * n
		}
		close(downstream)
	}()
	return downstream
}
