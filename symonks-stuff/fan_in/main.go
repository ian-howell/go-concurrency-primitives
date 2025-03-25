package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Fanning in allows us to funnel multiple channels through a single new channel.
Allowing all work to be processed in one central place and handle getting
results from many goroutines.


This example spawns three seperate generator goroutines, each with their own channe
for communication and using fanIn, filters those 3 channels into one seperate one
for streamlining the results in main().

This simplifies logic, error handling and is much easier to reason with rather than
having a fourth channel that each routine publishes their results too etc.
*/

func main() {
	ch1 := boring("foo", 100*time.Millisecond)
	ch2 := boring("bar", 1000*time.Millisecond)
	ch3 := boring("baz", 2000*time.Millisecond)
	merged := fanIn(ch1, ch2, ch3)

	for v := range merged {
		fmt.Println(v)
	}

}

func boring(msg string, delay time.Duration) <-chan string {
	out := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			out <- fmt.Sprintf("message %s: call %d", msg, i)
			time.Sleep(delay)
		}
		close(out)
	}()
	return out
}

// fanIn combines multiple channels into a single channel
func fanIn(chans ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	wg.Add(len(chans))
	downstream := make(chan string)
	handle := func(c <-chan string) {
		defer wg.Done()
		for v := range c {
			downstream <- v
		}
	}

	go func() {
		wg.Wait()
		close(downstream)
	}()

	for _, c := range chans {
		go handle(c)
	}
	return downstream
}
