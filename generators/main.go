package main

import (
	"fmt"
	"time"
)

/*
Utilising goroutines and channels, we can simulate a python like
generator
*/

// yielder is a simple function that returns a channel that can be
// used to communicate with it internally.  Similar to pythons `yield`
// or `yield from`.  The returned channel is a receive only channel.
func yielder(message string) <-chan string {
	downstream := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			downstream <- fmt.Sprintf("message: %s number %d", message, i)
			time.Sleep(500 * time.Millisecond)
		}
		close(downstream)
	}()
	return downstream
}

func main() {
	oneCh := yielder("one")
	twoCh := yielder("two")
	for i := 0; i < 10; i++ {
		fmt.Println(<-oneCh)
		fmt.Println(<-twoCh)
	}

}
