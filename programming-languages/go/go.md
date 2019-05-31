# [Go](https://golang.org)

Go is currently my favorite programming language.

I mostly love the tooling around it like [VS Code](../../text-editors/vs-code/vs-code.md) and its [Go plugin](https://github.com/Microsoft/vscode-go). The powerful `go` command line tool and the rich ecosystem of libraries and tools that people have built.

Go promotes composition over inheritance.

## Commenting Go code

- Comments documenting declarations should be full sentences.
- Comments should begin with the name of the thing being described and end in a period.

## Error checking

- You can `log.Fatal(err)` when playing with code.
  - In actual applications you need to decide what you need to do with each error response - bail immediately, pass it to the caller, show it to the user, log it and continue, etc ...

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
- [What works well is to use Go to create an HTTP API and leave the rest (html templates etc) to a client-side frontend (e.g. React). That gives you the best of both worlds: a fast and lightweight backend in Go with a rich frontend in JS.](https://www.reddit.com/r/golang/comments/8eeolx/recommended_web_framework_for_go/)
- This is a core concept in Go’s type system; instead of designing our abstractions in terms of what kind of data our types can hold, we design our abstractions in terms of what actions our types can execute.

## Links

- [Effective Go](https://golang.org/doc/effective_go.html)
- [How to Write Go Code](https://golang.org/doc/code.html)
- [Go proverbs](https://go-proverbs.github.io/)
- [Go internals](https://github.com/teh-cmc/go-internals#readme)
- [Go 101](https://go101.org/article/101.html) - Great book.
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
- [Share memory by communicating](https://blog.golang.org/share-memory-by-communicating)
- [Go Build Template](https://github.com/thockin/go-build-template)
- [Go for Industrial Programming](https://peter.bourgon.org/go-for-industrial-programming/) - Great insights.
- [Sum Types in Go](http://www.jerf.org/iri/post/2917)
- [Gophercises](https://gophercises.com/) - Free coding exercises for budding gophers.
- [Joy Compiler](https://mat.tm/joy/) - Translate idiomatic Go into concise JavaScript that works in every browser.
- [The Go Type System for newcomers](https://rakyll.org/typesystem/)
- [Web Assembly and Go: A look to the future (2018)](https://brianketelsen.com/web-assembly-and-go-a-look-to-the-future/)
- [How I Learned to Stop Worrying and Love Golang](https://corte.si/posts/code/go/golang-practicaly-beats-purity/index.html)
- [List of advices and tricks in the Go's world](https://github.com/cristaloleg/go-advices#readme)
- [Golang challenge](http://golang-challenge.org/)
- [dep2nix](https://github.com/nixcloud/dep2nix) - Using golang/dep to create a deps.nix file for go projects to package them for nixpkgs.
- [GopherCon 2018 - Rethinking Classical Concurrency Patterns](https://about.sourcegraph.com/go/gophercon-2018-rethinking-classical-concurrency-patterns/) - [Slides](https://drive.google.com/file/d/1nPdvhB0PutEJzdCq5ms6UI58dp50fcAN/view)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout#readme)
- [Building Web Apps with Go](https://codegangsta.gitbooks.io/building-web-apps-with-go/content/)
- [Golang Monorepo](https://github.com/flowerinthenight/golang-monorepo#readme) - Example of a golang-based monorepo.
- [Lorca](https://github.com/zserge/lorca) - Build cross-platform modern desktop apps in Go + HTML5.
- [Data Structures with Go Language](https://github.com/Dentrax/Data-Structures-with-Go)
- [go-hardware](https://github.com/rakyll/go-hardware) - Directory of hardware related libs, tools, and tutorials for Go.
- [HN: Go 2, here we go (2018)](https://news.ycombinator.com/item?id=18561587)
- [gotestsum](https://github.com/gotestyourself/gotestsum) - Runs tests, prints friendly test output and a summary of the test run.
- [Nice example of a small Go CLI tool](https://github.com/htdvisser/did)
- [GolangCI-Lint](https://github.com/golangci/golangci-lint) - Linters Runner for Go. 5x faster than gometalinter. Nice colored output. Can report only new issues. Fewer false-positives. Yaml/toml config.
- [Nice VS Code Go snippets](https://github.com/tj/vscode-snippets)
- [revive](https://github.com/mgechev/revive) - Fast, configurable, extensible, flexible, and beautiful linter for Go.
- [Going Infinite, handling 1M websockets connections in Go (2019)](https://github.com/eranyanay/1m-go-websockets)
- [Awesome Go Linters](https://github.com/golangci/awesome-go-linters#readme)
- [Leaktest](https://github.com/fortytw2/leaktest) - Goroutine Leak Detector.
- [Practical Go: Real world advice for writing maintainable Go programs (2018)](https://dave.cheney.net/practical-go/presentations/qcon-china.html) ([HN](https://news.ycombinator.com/item?id=19218097))
- [The State of Go: Feb 2019](https://speakerdeck.com/campoy/the-state-of-go-feb-2019?slide=38)
- [gomacro](https://github.com/cosmos72/gomacro) - Interactive Go interpreter and debugger with generics and macros.
- [go2ll-talk](https://github.com/pwaller/go2ll-talk) - Live coding a basic Go compiler with LLVM in 20 minutes.
- [Thoughts on Go performance optimization](https://github.com/dgryski/go-perfbook#readme)
- [Go Package Store](https://github.com/shurcooL/Go-Package-Store) - App that displays updates for the Go packages in your GOPATH.
- [Why are my Go executable files so large? Size visualization of Go executables using D3 (2019)](https://science.raphael.poss.name/go-executable-size-visualization-with-d3.html)
- [goweight](https://github.com/jondot/goweight) - Tool to analyze and troubleshoot a Go binary size.
- [Go Patterns](https://github.com/tmrts/go-patterns#readme) - Curated list of Go design patterns, recipes and idioms.
- [Go Developer Roadmap](https://github.com/Alikhll/golang-developer-roadmap#readme) - Roadmap to becoming a Go developer in 2019.
- [Clear is better than clever (2019)](https://dave.cheney.net/paste/clear-is-better-than-clever.pdf) ([HN](https://news.ycombinator.com/item?id=19837981))
