package main

import (
	"fmt"
	"sync"
	"time"
)

// Message encapsulates a message to send along with
// some throttling-esque state to manage ordering
type Message struct {
	msg  string
	wait chan bool
}

// String implements fmt.Stringer
func (m Message) String() string {
	return fmt.Sprintf("message: %s", m.msg)
}

// here we can guarantee the order of messages.  first and second will
// always be printed in order, one at a time for the cases in which we
// rely on it.
func main() {
	merged := fanIn(boring("foo"), boring("bar"), boring("baz"))

	for i := 0; i < 10; i++ {
		first := <-merged
		second := <-merged
		third := <-merged
		fmt.Println(first)
		fmt.Println(second)
		fmt.Println(third)

		first.wait <- true
		second.wait <- true
		third.wait <- true
	}

	fmt.Println("finished restoring the sequence.")
}

// boring returns a read only channel to communicate with it.
// this time it is responsible for pushing Message objects onto
// the downstream channel, each time it publishes a message it
// waits for the waiting channel to receive a value.  This allows
// to keep things in order when required.
func boring(msg string) <-chan Message {
	out := make(chan Message)
	wait := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			out <- Message{
				msg:  fmt.Sprintf("message %s: count %d", msg, i),
				wait: wait,
			}
			time.Sleep(500 * time.Millisecond)
			<-wait
		}
	}()
	return out
}

func fanIn(chans ...<-chan Message) <-chan Message {
	out := make(chan Message)
	var wg sync.WaitGroup
	wg.Add(len(chans))
	handle := func(c <-chan Message) {
		defer wg.Done()
		for msg := range c {
			out <- msg
		}
	}

	for _, c := range chans {
		go handle(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
