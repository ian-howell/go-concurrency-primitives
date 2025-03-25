package main

import (
	"fmt"
	"time"
)

func worker(id int, taskQ <-chan int, results chan<- int) {
	for n := range taskQ {
		fmt.Printf("worker %d processed task %d\n", id, n)
		time.Sleep(500 * time.Millisecond)
		results <- n * 2
	}
}

func main() {
	const tasks = 8
	jobs := make(chan int)
	results := make(chan int)
	for i := 0; i < tasks; i++ {
		go worker(i, jobs, results)
	}

	for i := 0; i < tasks; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < tasks; i++ {
		fmt.Printf("result %d\n", <-results)
	}
	close(results)
	fmt.Println("finished")
}
