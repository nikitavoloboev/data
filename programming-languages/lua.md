# [Lua](https://www.lua.org)

## Notes

- [Lua was created explicitly to be an embedded scripting language. It's evolution has very tightly focused on that usage, and that continues to be the most common way it is used. It's language as library, a small, pure ANSI C library with no external dependencies. It's easy to include in a project and extend, while remaining a powerful, expressive language that performs well. It's fantastic for that purpose.](https://www.reddit.com/r/lua/comments/hg9g3p/if_someone_asks_me_why_lua_i_show_him_this_link/)
- [Lua is a language I really want to love. I like the emphasis on simplicity and minimalism, and the Scheme-like semantics, which mix imperative and functional styles, really hits a sweet spot IMO. LuaJIT is a crazy impressive feat of software engineering. However, there are some specific issues which hold Lua back IMO. First, as LuaJIT author Mike Pall famously noted, the Lua authors constantly break compatibility between releases. Lua is really several different, incompatible languages (Lua 5.1, 5.2, etc). LuaJIT is still at Lua 5.1, IIRC. Second, there are a bunch of minor nitpicks (1-based-indexing, anyone?) which turn off a bunch of people. Lastly, because Lua is so minimal and focused on portability, people end up reimplementing their own abstractions (such as object systems) from scratch, further fracturing the ecosystem. I think there's a space for a new project, which takes LuaJIT as a starting point and addresses some of the issues I described. It would also be great if this hypothetical new language had better support for Unicode and concurrency.](https://news.ycombinator.com/item?id=23851393)

## Links

- [Luacheck](https://github.com/luarocks/luacheck) - Tool for linting and static analysis of Lua code.
- [Fennel](https://fennel-lang.org/) - Programming language that brings together the speed, simplicity, and reach of Lua with the flexibility of a lisp syntax and macro system. ([Code](https://github.com/bakpakin/Fennel)) ([Rationale](https://fennel-lang.org/rationale)) ([Lobsters](https://lobste.rs/s/6bphbw/fennel_programming_language_rationale)) ([HN](https://news.ycombinator.com/item?id=24390904)) ([FennelConf](https://conf.fennel-lang.org/))
- [rlua](https://github.com/kyren/rlua) - High level bindings between Rust and Lua.
- [How to Lua and C - a short novel (2018)](https://sepisoad.com/blog/how%20to%20lua%20and%20c%20-%20a%20short%20novel.html)
- [Code Formatter for Lua](https://github.com/trixnz/lua-fmt) - Pretty-printer for Lua code, written in TypeScript and deeply inspired by prettier.
- [Lua VM in pure Go](https://github.com/Shopify/go-lua)
- [GopherLua](https://github.com/yuin/gopher-lua) - VM and compiler for Lua in Go.
- [Awesome Lua](https://github.com/LewisJEllis/awesome-lua)
- [Embedding Lua Tutorial](https://github.com/davepoo/EmbeddingLuaTutorial)
- [Lunatic Python](https://github.com/bastibe/lunatic-python) - Two-way bridge between Python and Lua.
- [Interesting things about the Lua interpreter (2020)](https://thesephist.com/posts/lua/)
- [Interactive Lua development with Fennel (2018)](https://technomancy.us/189) ([Lobsters](https://lobste.rs/s/1xlmb3/interactive_lua_development_with_fennel))
- [LuaJIT](https://github.com/LuaJIT/LuaJIT) - Just-In-Time (JIT) compiler for the Lua programming language. ([Web](http://luajit.org/))
- [LuaJIT Bytecode Optimizations](http://wiki.luajit.org/Optimizations)
- [Lua in Rust](https://github.com/lonng/lua-rs) - Pure Rust implementation of Lua compiler.
- [Typed Lua](https://github.com/andremm/typedlua) - Optional Type System for Lua.
- [Teal](https://github.com/teal-language/tl) - Compiler for Teal, a typed dialect of Lua.
- [pprint.lua](https://github.com/jagt/pprint.lua) - Lua pretty printer.
- [Using Lua As A Serialization Format (2020)](https://mkhan45.github.io/2020/06/16/using-lua-as-a-serialization-format.html) ([HN](https://news.ycombinator.com/item?id=23539332)) ([Lobsters](https://lobste.rs/s/dttksl/using_lua_as_serialization_format))
- [Lua Integration (2020)](https://mkhan45.github.io/2020/06/12/lua-integration.html)
- [LuaJIT Language Toolkit](https://github.com/franko/luajit-lang-toolkit) - Implementation of the Lua programming language written in Lua itself.
- [Antifennel](https://git.sr.ht/~technomancy/antifennel) - Turn Lua code into Fennel code.
- [HN: Lua 5.4.0 (2020)](https://news.ycombinator.com/item?id=23686297)
- [Languages that compile to Lua](https://github.com/hengestone/lua-languages)
- [Making the Fennel compiler self-hosting with another compiler (2020)](https://technomancy.us/192)
- [How to Plan a Luau: Augmenting Lua’s Syntax With Types (2020)](https://medium.com/roblox-tech-blog/how-to-plan-a-luau-augmenting-luas-syntax-with-types-7751a790f0d8) ([HN](https://news.ycombinator.com/item?id=24448364))
- [Croissant](https://github.com/giann/croissant) - Lua REPL and debugger implemented in Lua.
- [Fengari](https://github.com/fengari-lua/fengari) - Lua VM written in JS ES6 for Node and the browser. ([Web](https://fengari.io/))
- [Lua Style Guide](https://github.com/Olivine-Labs/lua-style-guide)
- [busted](https://github.com/Olivine-Labs/busted) - Elegant Lua unit testing.
- [RxLua](https://github.com/bjornbytes/RxLua) - Reactive Extensions for Lua.
- [LTUI](https://github.com/tboox/ltui) - Cross-platform terminal ui library based on Lua.
- [Raymarching with Fennel and LÖVE (2020)](https://andreyorst.gitlab.io/posts/2020-10-15-raymarching-with-fennel-and-love/)
- [Moonshine](https://github.com/gamesys/moonshine) - Lightweight Lua VM for the browser.
- [A Look at the Design of Lua (2018)](https://cacm.acm.org/magazines/2018/11/232214-a-look-at-the-design-of-lua/fulltext) ([HN](https://news.ycombinator.com/item?id=18327661))
- [lua-fsm](https://github.com/unindented/lua-fsm) - Simple finite-state machine implementation for Lua.
- [plenary.nvim](https://github.com/nvim-lua/plenary.nvim) - All the lua functions I don't want to write twice.
- [LuaUnit](https://github.com/bluebird75/luaunit) - Popular unit-testing framework for Lua, with an interface typical of xUnit libraries.
- [LuaFormatter](https://github.com/Koihik/LuaFormatter) - Reformats your Lua source code.
- [Heart](https://github.com/Hyperspace-Logistics/heart) - High performance Lua web server with a simple, powerful API.
- [Lua, a misunderstood language (2021)](https://andregarzia.com/2021/01/lua-a-misunderstood-language.html) ([Lobsters](https://lobste.rs/s/novtvd/lua_misunderstood_language)) ([HN](https://news.ycombinator.com/item?id=25796852))
- [Lua and Python (2020)](https://lwn.net/Articles/812122/) ([HN](https://news.ycombinator.com/item?id=25794374)) ([Lobsters](https://lobste.rs/s/2lpxqj/lua_python))
- [luapower](https://luapower.com/) - LuaJIT distribution. ([GitHub](https://github.com/luapower))
- [Penlight](https://github.com/lunarmodules/Penlight) - Useful pure Lua modules, focusing on input data handling, functional programming and OS path management.
- [StyLua](https://github.com/JohnnyMorganz/StyLua) - Opinionated Lua code formatter.
- [Full Moon](https://github.com/Kampfkarren/full-moon) - Lossless Lua 5.1 parser.
- [Ravi](https://github.com/dibyendumajumdar/ravi) - Dialect of Lua, featuring limited optional static typing, JIT and AOT compilers. ([Web](http://ravilang.github.io/))
- [LuaJIT compiler explorer](https://luajit.me/) ([Code](https://github.com/rapidlua/luajit.me))
- [LadyLua](https://github.com/tongson/LadyLua) - Batteries-included static Lua 5.1 interpreter. ([HN](https://news.ycombinator.com/item?id=26738006))
