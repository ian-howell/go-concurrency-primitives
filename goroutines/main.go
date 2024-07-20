package main

import "time"

func main() {
	/*
		In order to run a function asynchronously in golang the `go`
		keyword can be used.  The go runtime is complex and does all
		sorts of magic with OS threads etc, more on that later.
	*/

	// This function is blocking, it will block the program for 5 seconds
	// as its blocking the main thread
	synchronous()

	// This function will not block, the go runtime will multiplex it
	// However, because it is not blocking it will not prevent main() from
	// exiting.  main will exit instantly, regardless of the sleep.
	// more on other concurrency primities to handle these cases later.
	go asynchronous()

	/*
		Here we technically, leaked a goroutine but the program will terminate in 5 seconds:

		go run goroutines/main.go  0.05s user 0.06s system 1% cpu 5.059 total
	*/
}

func synchronous() {
	time.Sleep(5 * time.Second)
}

func asynchronous() {
	time.Sleep(1 * time.Hour)
}
