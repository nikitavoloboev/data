# [Crafting interpreters](http://craftinginterpreters.com)
## Chapter 1
- You can implement a compiler in any language, including the same language it compiles, a process called “self-hosting”.
- You can’t compile it using itself yet, but if you have another compiler for your language written in some other language, you use that one to compile your compiler once.
	- Now you can use the compiled version of your own compiler to compile future versions of itself and you can discard the original one compiled from the other compiler.
	- This is called “bootstrapping” from the image of pulling yourself up by your own bootstraps.

## [Chapter 2 - map of the territory](http://www.craftinginterpreters.com/a-map-of-the-territory.html)
![](https://i.imgur.com/w4R6Mkl.jpg)

- A parser takes the flat sequence of tokens and builds a tree structure that mirrors the nested nature of the grammar.
	- These trees have a couple of different names — “parse tree” or “abstract syntax tree” — depending on how close to the bare syntactic structure of the source language they are.
		- In practice, language hackers usually call them “syntax trees”, “ASTs”, or often just “trees”.