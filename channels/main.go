package main

import "fmt"

/*
Channels are the pipes that connect concurrent go routines.  Go has a good philosophy
in that most languages communicate by sharing memory (think threading models like java,
python etc).  This is extremely error prone, whereas go is built around a module of share
memory BY communicating.  Rather than mutexing or thread local etc go encourages the use
of channels to pass references of data between goroutines, this guarantees implicitly that
only one goroutine can access the data at one given time.

Channels are a synchonisation primitive that hold data references and control locking/mutexing
implicitly. They can be buffered or by default (non buffered).  buffered cases should be scarcely
used unless the use case is clear.

by default channels are bidirectional, i.e can send/receive values on them, but you can type them
as read only channels `<-chan`, or send only channels `chan<-` for increased type safety.

There are many charactistics of send/receive operations on channels in terms of blocking etc.
You can find more on that in the channel deep dive.
*/
func main() {
	ch := make(chan int)
	// asynchronously push values into the channel
	// this is an anonymous function, called asynchronously.
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	// Range over the channel here to print all the numbers
	// this will block until the channel is closed by the
	// goroutine after it has published all of it's numbers
	for n := range ch {
		fmt.Println(n)
	}

}
