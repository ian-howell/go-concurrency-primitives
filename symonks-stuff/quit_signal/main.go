package main

import (
	"fmt"
	"time"
)

func boring(quit chan string, msg string) <-chan string {
	downstream := make(chan string)
	go func() {
		defer close(downstream)
		for i := 0; i < 10; i++ {
			select {
			case <-quit:
				fmt.Println("early exit...")
				quit <- "here is some final state"
				return
			case downstream <- fmt.Sprintf("Message %s: Called %d", msg, i):
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	return downstream
}

func main() {
	quit := make(chan string)
	goro := boring(quit, "foo")

	// do 8, still 2 available on the channel to read
	for i := 0; i < 8; i++ {
		fmt.Println(<-goro)
	}

	// stop early and have the goroutine gracefully exit.
	quit <- "stop!"
	fmt.Printf("final goroutine state was: %s\n", <-quit)

}
