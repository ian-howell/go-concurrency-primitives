package main

import (
	"fmt"
	"runtime"
	"sync"
)

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
	scaled := make([]<-chan int, runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		scaled[i] = squared(genChan)
	}
	for val := range merge(scaled...) {
		fmt.Println(val)
	}
}

// merge combines multiple channels into a single channel
// a fan out approach.
func merge(chans ...<-chan int) chan int {
	var wg sync.WaitGroup
	wg.Add(len(chans))
	downstream := make(chan int)
	consumer := func(c <-chan int) {
		for n := range c {
			downstream <- n
		}
		wg.Done()
	}
	for _, ch := range chans {
		go consumer(ch)
	}

	go func() {
		wg.Wait()
		close(downstream)
	}()

	return downstream
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
