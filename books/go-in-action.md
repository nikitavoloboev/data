# [Go in action](https://www.goodreads.com/book/show/22727352-go-in-action)
# Chapter 2 - quick start
- init() functions get called before main functions.
- All variables in Go are initialized to their zero value. Pointer's zero value is nil.
- Once the main function returns, the program terminates. Any goroutines that were launched and are still running at this time will also be terminated by the Go run-time.
- When you write concurrent programs, it’s best to cleanly terminate any goroutines that were launched prior to letting the main function return.

# Chapter 3 - packages
- The convention for naming your package is to use the name of the directory containing it.