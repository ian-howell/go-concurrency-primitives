## Go Concurrency

This repository houses various golang concurrency pattern implementations.  If you are new to 
concurrency in go, the optimal order to review the solutions is outlined below:

-----

### Patterns

* [01 - Goroutines](goroutines/main.go)
* [02 - Channels](channels/main.go)
* [03 - Generators](generators/main.go)
* [04 - Fan In](fan_in/main.go)
* [05 - Restore Sequence](restore_sequence/main.go)
* [06 - Select Timeout](select_timeout/main.go)
* [07 - Quit Signal](quit_signal/main.go)
* [08 - Daily Chain](daisy_chain/main.go)
* [09 - Google Search Synchronous](google_search_synchronous/main.go)
* [10 - Google Search Asynchronous](google_search_asynchronous/main.go)
* [11 - Google Search Async Filtered](google_search_async_timeout/main.go)

-----

### Pattern Overviews

* `Goroutines` - A basic implementation of goroutines, running things asynchronously.
* `Channels` - A communication mechanism between goroutines, sharing memory by communicating.
* `Generators` - Python like generators, a function returns a channel to interact with it.
* `Fan In` - Merging multiple goroutines results into a single channel for processing.
* `Restore Sequence` - Fan in with a twist, allowing order of fanned channel results to be controlled.
* `Select Timeout` - Giving a routine a period of time to finish else doing something else.
* `Quit Signal` - Signally to a goroutine that it is time to stop potentially early.
* `Daisy Chain` - Many goroutines, passing along messages like chinese whispers - confusing!
* `Google Search Synchronous` - A synchronous implementation of basic google searching. 
* `Google Search Asynchronous` - Goroutines per search type and filtering in results.
* `Google Search Asynchronous Filtered` - Omitting results that took too long.



### Learning Materials TODO

More materials to include in the repository, things I personally need to learn more about:

 - indepth channels
 - go scheduler, runtime and os thread specifics
 - daisy chain is confusing!