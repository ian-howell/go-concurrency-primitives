package main

import (
	"fmt"
	"time"

	"math/rand"
)

/*
Building on the synchronous google search.  All options are
asynchronously fetched and presented as one.

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

func GoogleSearch(query string) (results []Result) {
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

	for i := 0; i < 4; i++ {
		results = append(results, <-ch)
	}
	return
}

func main() {
	now := time.Now()
	results := GoogleSearch("concert tickets")
	elapsed := time.Since(now)
	fmt.Println(results)
	fmt.Println(elapsed)
}
