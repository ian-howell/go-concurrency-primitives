package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
A naive implementation of a synchronous google search.
This repository builds on this throughout multiple phases.
See google_search_two for concurrency improvements.

*/

type Result string
type Search func(query string) Result

var (
	Web      = doSearch("web")
	Image    = doSearch("image")
	Video    = doSearch("video")
	Shopping = doSearch("shopping")
)

// doSearch simulates searching the web for a result
func doSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func GoogleSearch(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	results = append(results, Shopping(query))
	return
}

func main() {
	now := time.Now()
	results := GoogleSearch("concert tickets")
	elapsed := time.Since(now)
	fmt.Println(results)
	fmt.Println(elapsed)

}
