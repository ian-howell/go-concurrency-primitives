package main

import (
	"fmt"
	"time"
)

func main() {
	// In order to run a function asynchronously in golang the `go` keyword can be used. The go runtime is
	// complex and does all sorts of magic with OS threads etc.

	// This function will not block; the Go runtime will multiplex it. However, because it is not
	// blocking, it will not prevent main from exiting. main will exit instantly, regardless of what the
	// function is doing.
	go asynchronous()

	// This function is blocking, it will block the program for 5 seconds as its blocking the main thread
	fmt.Println("This function is blocking!")
	synchronous()

	// We're technically leaking a goroutine here, but since this is not a long running program, the Go
	// runtime will clean up after us when main exits.
}

func synchronous() {
	fmt.Println("synchronous: Sleeping for 5 seconds...")
	time.Sleep(5 * time.Second)
	fmt.Println("synchronous: Done!")
}

func asynchronous() {
	for {
		fmt.Println("asynchronous: Sleeping for 1 second...")
		time.Sleep(1 * time.Second)
	}
	// Unreachable...
	fmt.Println("asynchronous: Done!")
}
