package main

import (
	"fmt"
	"time"

	"math/rand"
)

/*
Building on the asynchronous google searching.  This implementation
filters out slow results from searching that exceed a particular
threshold and presents the client with results it could find in time
*/

type Result string
type Search func(kind string) Result

var (
	Web      = doSearch("web")
	Image    = doSearch("image")
	Video    = doSearch("video")
	Shopping = doSearch("shopping")
)

func doSearch(query string) Search {
	return func(kind string) Result {
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		return Result(fmt.Sprintf("result %s for query %q\n", query, kind))
	}
}

func GoogleSearch(query string, timeout time.Duration) (results []Result) {
	ch := make(chan Result, 4)

	go func() {
		ch <- Web(query)
	}()

	go func() {
		ch <- Image(query)
	}()

	go func() {
		ch <- Video(query)
	}()

	go func() {
		ch <- Shopping(query)
	}()

	ticker := time.NewTicker(timeout)
	defer ticker.Stop()
	for i := 0; i < 4; i++ {
		select {
		case result := <-ch:
			results = append(results, result)
		case <-ticker.C:
			fmt.Printf("timed out fetching results for %q\n", query)
			continue
		}
	}
	return
}

func main() {
	now := time.Now()
	results := GoogleSearch("concert tickets", time.Second)
	elapsed := time.Since(now)
	fmt.Println(results)
	fmt.Println(elapsed)
}
