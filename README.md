# Go Concurrency

This repository houses various golang concurrency pattern implementations.  If you are new to 
concurrency in go, the optimal order to review the solutions is outlined below:

-----

## Patterns

* [01 - Goroutines](goroutines/main.go)
* [02 - WaitGroups](waitgroups/main.go)
* [03 - Channels](channels/main.go)
* [04 - Generators](generators/main.go)
* [05 - Fan In](fan_in/main.go)
* [06 - Restore Sequence](restore_sequence/main.go)
* [07 - Select Timeout](select_timeout/main.go)
* [08 - Quit Signal](quit_signal/main.go)
* [09 - Daily Chain](daisy_chain/main.go)
* [10 - Google Search Synchronous](google_search_synchronous/main.go)
* [11 - Google Search Asynchronous](google_search_asynchronous/main.go)
* [12 - Google Search Async Filtered](google_search_async_timeout/main.go)
* [18 - Worker Pool](workerpool/main.go)

-----

## Pattern Overviews

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
* ...
* `Worker Pool` - Fan out work to multiple works to complete in parallel.



## Learning Materials TODO

More materials to include in the repository, things I personally need to learn more about:

 - indepth channels
 - go scheduler, runtime and os thread specifics
 - daisy chain is confusing!

## Attribution notice

This repository was forked from [symonk's repo][fork]. I intend to pretty
heavily modify it to suit my own needs, but the original repo is proving to be a
great starting point. Shout outs to [symonk][symonk] for laying out the base
framework for this project!

[fork]: https://github.com/symonk/go-concurrency-deep-dive
[symonk]: https://github.com/symonk
