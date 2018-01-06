# Go
Go is currently my favourite language.

I mostly love the tooling around it like VS Code and its [Go plugin](https://github.com/Microsoft/vscode-go). The powerful  `go` command line tool and the rich ecosystem of libraries and tools that people have built.

Go promotes composition over inheritance. 

## Commenting Go code
- Comments documenting declarations should be full sentences
- Comments should begin with the name of the thing being described and end in a period

## Error checking
- You can `log.Fatal(err)` when playing with code.
	- In actual applications you need to decide what you need to do with each error response - bail immediately, pass it to the caller, show it to the user, log it and continue, etc ...

## Notes
- I can call functions from anywhere if they are in the same package
- The package “main” tells the Go compiler that the package should compile as an executable program instead of a shared library
- Note that any type implements the empty interface interface{} because it doesn't have any methods and all types have zero methods by default
- Simply pushing my source code to GitHub, makes it `go get`table.
- Interface types represent fixed sets of methods
- the prevailing wisdom in Go is to use a flat directory structure and only create new directories when you are building self-contained functionality
- `go get` = download the source code to your PC.
- `go install` = download, build, and put it in the path so you can use it.
- I don't need to comment all functions as some are self describing. I do need to comment exported functions however.

## Links
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Performance without the event loop (2015)](https://dave.cheney.net/2015/08/08/performance-without-the-event-loop)
- [7 common mistakes in Go and when to avoid them by Steve Francia](https://www.youtube.com/watch?v=29LLRKIL_TI)
- [Performance without the event loop (2015)](https://dave.cheney.net/2015/08/08/performance-without-the-event-loop)
- [What I learned in 2017 Writing Go](https://www.commandercoriander.net/blog/2017/12/31/writing-go/)