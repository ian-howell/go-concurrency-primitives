package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Waitgroups are a way to wait for a collection of goroutines to finish executing.
// They are similar to the counting semaphore pattern, but they are more flexible and easier to use.

func main() {
	// Initialise a waitgroup
	wg := &sync.WaitGroup{}

	// Add one to the waitgroup. This indicates that we are waiting for one goroutine to finish.
	wg.Add(1)
	go func() {
		fmt.Println("Sleeping for 1 second...")
		time.Sleep(1 * time.Second)
		fmt.Println("Woke up after 1 second")

		// Mark the goroutine as done
		wg.Done()
	}()

	// Wait for the goroutine to finish
	fmt.Println("Waiting for initial goroutine to finish...")
	wg.Wait()
	fmt.Println("Initial goroutine finished!")

	// Spin up a handful of goroutines
	for i := range 10 { // bonus! check out the new for-range loop introduced in go1.22
		// Add the goroutines to the waitgroup
		wg.Add(1)

		// Sleep for a random amount of time. Concurrently.
		go func(i int) {
			// Generate a random duration between 0 and 5 seconds
			d := time.Duration(rand.Intn(5000)) * time.Millisecond

			// Sleep!
			fmt.Printf("Goroutine %d:   Sleeping for %v\n", i, d)
			time.Sleep(d)

			// Mark the goroutine as done
			fmt.Printf("Goroutine %d:   Woke up after %v\n", i, d)
			wg.Done()
		}(i)
	}

	// wg.Wait will block until *all* goroutines have called Done
	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("All goroutines finished!")
}
