# Go
Go is currently my favourite language.

I mostly love the tooling around it like VS Code and its [Go plugin](https://github.com/Microsoft/vscode-go). The powerful  `go` command line tool and the rich ecosystem of libraries and tools that people have built.

Go promotes composition over inheritance. 

## Commenting Go code
- Comments documenting declarations should be full sentences
- Comments should begin with the name of the thing being described and end in a period

# Snippets
- [read file line by line](https://gist.github.com/69824f3c2f29ae5fc6519452e2c89a4d)

# Links
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [15: Performance without the event loop](https://dave.cheney.net/2015/08/08/performance-without-the-event-loop)

# Notes
- I can call functions from anywhere if they are in the same package
- The package “main” tells the Go compiler that the package should compile as an executable program instead of a shared library
- Note that any type implements the empty interface interface{} because it doesn't have any methods and all types have zero methods by default