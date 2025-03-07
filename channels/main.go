package main

import (
	"fmt"
	"time"
)

// "Don't communicate by sharing memory, share memory by communicating." - Rob Pike
//
// In traditional concurrency models, communication is done by sharing memory, i.e. multiple processes access
// the same memory space. This leads to very complex and error prone code, as you have to manage the locking
// of the memory space.
//
// Go inverts this idea and instead, shares memory BY communicating. In other words, Go limits access to
// resources by simply passing around the resources themselves (rahter than passing around the
// "lock"/"key"/"mutex"). This *implicitly* guarantees that only one goroutine can access the data at any
// given time.
//
// Channels are somewhat like mailboxes. Go routines can put data into them, and they can take data out of
// them. By default channels are bidirectional, i.e can send/receive values on them, but you can type them as
// read only channels `<-chan`, or send only channels `chan<-` for increased type safety.

func main() {
	// This line initialises a channel of type int. This means that we can send and receive integers on
	// this channel. We could have chosen any type here, but we chose int for simplicity.
	ch := make(chan int)

	// Spin up a goroutine to send data on the channel.
	fmt.Println("main:      Starting goroutine to send data on channel...")
	go send(ch)

	// Recieve a value from the channel. This is kinda like taking an item out of the "mailbox". This will
	// block until a value is sent on the channel - think about a kid waiting patiently for the mailman.
	fmt.Println("main:      Waiting for data")
	value := <-ch
	fmt.Printf("main:      Recieved the first value: %d\n", value)

	// We can also range over the channel to "drain" it. The for loop will block until a value is sent on
	// the channel, and will continue to read from the channel until it is closed.
	for i := range ch {
		fmt.Printf("main:      Recieved value: %d\n", i)
	}
	fmt.Println("main:      Done! The channel is closed.")
}

func send(ch chan<- int) {
	fmt.Println("goroutine: Sending the first value on the channel and then waiting for main to recieve it...")
	ch <- 42
	time.Sleep(1 * time.Second)

	fmt.Println("goroutine: Sending a bunch of values on the channel...")
	for i := range 10 {
		// Here, we send integers. You can think of this as putting the integers into the "mailbox".
		fmt.Printf("goroutine: Sending value %d on the channel and then waiting for main to recieve it...\n", i)
		ch <- i
		time.Sleep(1 * time.Second)
	}

	// Note the type of the channel here. We are using a send only channel, which means that we can only
	// send data. This line will not compile.
	// valueFromChannel := <-ch

	// Close the channel. This indicates that we are done putting data into it (note that the data we put
	// in is still there!). Once a channel is closed, it can never be opened again.
	fmt.Println("goroutine: Sleeping for another second before closing the channel...")
	time.Sleep(1 * time.Second)
	close(ch)
}
