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
	for range 10 { // bonus! check out the new for-range loop introduced in go1.22
		// Add the goroutines to the waitgroup
		wg.Add(1)

		// Sleep for a random amount of time. Concurrently.
		go func() {
			sleepRandom(5000)

			// Mark the goroutine as done
			wg.Done()
		}()
	}

	// wg.Wait will block until *all* goroutines have called Done
	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("All goroutines finished!")
}

// sleepRandom sleeps a random number of milliseconds between 0 and n
func sleepRandom(n int) {
	d := time.Duration(rand.Intn(n)) * time.Millisecond
	fmt.Println("Sleeping for", d)
	time.Sleep(d)
	fmt.Println("Woke up after", d)
}
