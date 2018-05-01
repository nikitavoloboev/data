# [Go](https://golang.org)
Go is currently my favourite programming language.

I mostly love the tooling around it like VS Code and its [Go plugin](https://github.com/Microsoft/vscode-go). The powerful  `go` command line tool and the rich ecosystem of libraries and tools that people have built.

Go promotes composition over inheritance.

## Commenting Go code
- Comments documenting declarations should be full sentences.
- Comments should begin with the name of the thing being described and end in a period.

## Error checking
- You can `log.Fatal(err)` when playing with code.
	- In actual applications you need to decide what you need to do with each error response - bail immediately, pass it to the caller, show it to the user, log it and continue, etc ...

## Vendoring
[vg](https://github.com/getstream/vg) and [dep](https://github.com/golang/dep) are pretty good.

## Notes
- I can call functions from anywhere if they are in the same package.
- The package “main” tells the Go compiler that the package should compile as an executable program instead of a shared library.
- Note that any type implements the empty interface interface{} because it doesn't have any methods and all types have zero methods by default.
- Simply pushing my source code to GitHub, makes it `go get`table.
- Interface types represent fixed sets of methods.
- the prevailing wisdom in Go is to use a flat directory structure and only create new directories when you are building self-contained functionality.
- `go get` = download the source code to your PC.
- `go install` = download, build, and put it in the path so you can use it.
- I don't need to comment all functions as some are self describing. I do need to comment exported functions however.
- I [can put my Go projects not in my GOPATH if I am not planning to release them publically as library or cli tool](https://golang.org/doc/code.html#GOPATH). i.e. My Alfred workflows live in `~/dev/alfred` and not in my GOPATH and are still writen in Go.
- `go run` looks at the paths you specify.  `go build` takes the package(s) you specify and uses GOPATH to find their location.
- [What works really well is to use Go to create an HTTP API and leave the rest (html templates etc) to a client-side frontend (e.g. React). That gives you the best of both worlds: a fast and lightweight backend in Go with a rich frontend in JS.](https://www.reddit.com/r/golang/comments/8eeolx/recommended_web_framework_for_go/)

## Links
- [Effective Go](https://golang.org/doc/effective_go.html)
- [How to Write Go Code](https://golang.org/doc/code.html)
- [Go proverbs](https://go-proverbs.github.io/)
- [Go internals](https://github.com/teh-cmc/go-internals#readme)
- [Notes on Go](https://brandur.org/go)
- [Avoiding complexity in Go](https://bradgignac.com/2014/09/24/avoiding-complexity-with-go.html)
- [Gopher reading list](https://github.com/enocom/gopher-reading-list#readme)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Performance without the event loop (2015)](https://dave.cheney.net/2015/08/08/performance-without-the-event-loop)
- [7 common mistakes in Go and when to avoid them by Steve Francia](https://www.youtube.com/watch?v=29LLRKIL_TI)
- [Performance without the event loop (2015)](https://dave.cheney.net/2015/08/08/performance-without-the-event-loop)
- [What I learned in 2017 Writing Go](https://www.commandercoriander.net/blog/2017/12/31/writing-go/)
- [Golang interfaces](https://blog.chewxy.com/2018/03/18/golang-interfaces/)
- [Go in 5 minutes](https://github.com/arschles/go-in-5-minutes)
- [Learn Go with tests](https://github.com/quii/learn-go-with-tests)
- [Using Instruments to profile Go programs](https://rakyll.org/instruments/)
- [Design Philosophy On Packaging](https://www.ardanlabs.com/blog/2017/02/design-philosophy-on-packaging.html)
- [Go best practices, six years in](https://peter.bourgon.org/go-best-practices-2016/#repository-structure)
- [Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)
- [Style guideline for Go packages](https://rakyll.org/style-packages/)