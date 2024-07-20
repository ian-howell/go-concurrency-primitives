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

-----

### Pattern Overviews

`Goroutines` - A basic implementation of goroutines, running things asynchronously.
`Channels` - A communication mechanism between goroutines, sharing memory by communicating.
`Generators` - Python like generators, a function returns a channel to interact with it.
`Fan In` - Merging multiple goroutines results into a single channel for processing.
`Restore Sequence` - Fan in with a twist, allowing order of fanned channel results to be controlled.
`Select Timeout` - Giving a routine a period of time to finish else doing something else.


### Learning Materials TODO

More materials to include in the repository, things I personally need to learn more about:

 - indepth channels
 - go scheduler, runtime and os thread specifics