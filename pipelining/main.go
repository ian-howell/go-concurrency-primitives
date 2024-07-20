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
	done := make(chan struct{})
	defer close(done)

	numbers := numInRange(1, 1000)
	scaled := make([]<-chan int, runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		scaled[i] = squared(done, numbers)
	}
	for val := range merge(done, scaled...) {
		fmt.Println(val)
	}
}

// merge combines multiple channels into a single channel
// a fan out approach.
func merge(done <-chan struct{}, chans ...<-chan int) chan int {
	var wg sync.WaitGroup
	downstream := make(chan int)
	consumer := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case downstream <- n:
			case <-done:
				return
			}
		}
	}
	wg.Add(len(chans))
	for _, ch := range chans {
		go consumer(ch)
	}

	go func() {
		wg.Wait()
		close(downstream)
	}()

	return downstream
}

// numInRange creates a channel full of integers in the range
func numInRange(inclusiveLow, exclusiveHigh int) <-chan int {
	downstream := make(chan int, exclusiveHigh-inclusiveLow)
	for i := inclusiveLow; i < exclusiveHigh; i++ {
		downstream <- i
	}
	return downstream
}

// squareroot asynchronously squares the numbers
func squared(done <-chan struct{}, upstream <-chan int) <-chan int {
	downstream := make(chan int)
	go func() {
		defer close(downstream)
		for n := range upstream {
			select {
			case downstream <- n * n:
			case <-done:
				return
			}
		}
	}()
	return downstream
}
