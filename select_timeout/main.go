package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Select Timeout is a mechanism for waiting for one (or many) goroutines
to push values onto channels but with the capability to time out the
operation using a select.  Select essentially allows waiting on multiple
channels and most things often expose an underlying channel (such as tickers)
etc.

This case below gives a goroutine n amount of time to be finished, otherwise
something else is fired instead.  Selects will pick an available channel
if more than one are available at random and will block indefinitely when
no channels are ready and it is not being ran in an infinite for loop with
a default case.
*/

// boring returns a channel to interact with it.
// it randomly sleeps to demonstrate sometimes the select
// will use it when ready, else fall back to the duration lapse
// monitoring channel
func boring() <-chan string {
	downstream := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			downstream <- fmt.Sprintf("iteration: %d", i)
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		}
		close(downstream)
	}()

	return downstream
}

func main() {
	ch := boring()
	t := time.NewTicker(time.Second)

loop:
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("All values exhausted.")
				break loop
			}
			fmt.Printf("iteration in the time limit: %s\n", msg)
		case <-t.C:
			// TODO: This is cool, but I don't think symonk fully grasped the idea here. This is
			// going to "tick" on *every* second, regardless of what's going on in the goroutine.
			fmt.Println("took more than a second")
		}
	}
	fmt.Println("finished...")
}
