# Regex

I use [regex101](https://regex101.com) every time I want write a new regex.

Small tip, [regex101](https://regex101.com) has night mode.

![](https://i.imgur.com/ZVm6HVX.png)

[Rulex](https://rulex-rs.github.io/) is nice.

## Notes

- `.*` = will select all occurrences
- A regex is useful for validating simple patterns and for finding patterns in text. For anything beyond that it’s almost certainly a terrible choice.

## Code

[Match if it has notes or twitter in input](https://regex101.com/r/eKyP11/2)
`^(?=.*(?:notes|twitter)).*`

Match all that don't have notes or twitter in input
`^(?!._(?:notes|twitter))._`

## Links

- [RegexLearn](https://regexlearn.com/) - Step by step, from zero to advanced. ([HN](https://news.ycombinator.com/item?id=29251702)) ([Code](https://github.com/aykutkardas/regexlearn.com))
- [Learn regex the easy way](https://github.com/ziishaned/learn-regex)
- [Learn regex](https://github.com/zeeshanu/learn-regex) - Great reference.
- [DebuggexBeta](https://debuggex.com/)
- [RegexOne](https://regexone.com/) - Learn Regular Expressions with simple, interactive exercises.
- [Regular Expression in Python HOWTO](https://docs.python.org/3.8/howto/regex.html#regex-howto) - Good intro.
- [Hyperscan](https://github.com/intel/hyperscan) - High-performance regular expression matching library.
- [Oniguruma](https://github.com/kkos/oniguruma) - Modern and flexible regular expressions library.
- [Regex for Designers and UX Writers (2019)](https://daneden.me/2019/11/23/regex-for-designers-and-writers/)
- [Regexes vs Combinatorial Parsing (2019)](http://khanlou.com/2019/12/regex-vs-combinatorial-parsing/)
- [RegExr](https://regexr.com/) - Learn, Build & Test RegEx. ([HN](https://news.ycombinator.com/item?id=30103085))
- [regexgen](https://npm.runkit.com/regexgen) - Generate regular expressions that match a set of strings. ([Code](https://github.com/devongovett/regexgen))
- [Regular Expressions - Computerphile (2020)](https://www.youtube.com/watch?v=528Jc3q86F8)
- [Regex Tester](https://extendsclass.com/regex-tester.html) - Visual regex tester.
- [I hate Regex - Regex cheat sheet](https://ihateregex.io/) ([HN](https://news.ycombinator.com/item?id=22200584)) ([Code](https://github.com/geongeorge/i-hate-regex))
- [grex](https://github.com/pemistahl/grex) - Command-line tool and library for generating regular expressions from user-provided test cases.
- [Regular Expressions for Regular Folk](https://refrf.shreyasminocha.me/) ([HN](https://news.ycombinator.com/item?id=23042079))
- [regHEX](https://github.com/kitten/reghex) - Magical sticky regex-based parser generator.
- [tiny-regex-c](https://github.com/kokke/tiny-regex-c) - Small portable regex in C.
- [Structural Regular Expressions - Rob Pike](http://doc.cat-v.org/bell_labs/structural_regexps/se.pdf) ([Lobsters](https://lobste.rs/s/1aocan/structural_regular_expressions_1987))
- [Regexes vs Combinatorial Parsing (2020)](https://khanlou.com/2019/12/regex-vs-combinatorial-parsing/) ([HN](https://news.ycombinator.com/item?id=23331499))
- [Regex Crossword](https://regexcrossword.com/)
- [RegEx Crossword 2](http://jimbly.github.io/regex-crossword/)
- [libfsm](https://github.com/katef/libfsm) - DFA regular expression library & friends.
- [Python re(gex) book](https://leanpub.com/py_regex) - Learn Python Regular Expressions step by step from beginner to advanced levels with hundreds of examples and exercises. ([Code](https://github.com/learnbyexample/py_regular_expressions))
- [JavaScript RegExp](https://github.com/learnbyexample/learn_js_regexp) - Example based guide to mastering JavaScript regexp. ([Article](https://learnbyexample.github.io/cheatsheet/javascript/javascript-regexp-cheatsheet/))
- [Compile time regular expressions](https://github.com/hanickadot/compile-time-regular-expressions)
- [A Visual Guide to Regular Expression (2020)](https://amitness.com/regex/)
- [Emulating regexp lookarounds in GNU sed (2020)](https://learnbyexample.github.io/sed-lookarounds/)
- [Try Regex](http://tryregex.com/) - Interactive regex tutorial. ([Code](https://github.com/callumacrae/tryregex))
- [Three Great Videos About Regex Derivatives (2020)](http://www.oilshell.org/blog/2020/12/regex-videos.html)
- [Super Expressive](https://github.com/francisrstokes/super-expressive) - Zero-dependency JavaScript library for building regular expressions in (almost) natural language. ([HN](https://news.ycombinator.com/item?id=25857753))
- [Stephen Kleene's 1951 Paper Where He Introduces Regular Expressions](https://www.rand.org/content/dam/rand/pubs/research_memoranda/2008/RM704.pdf)
- [Implementing Regular Expressions](https://swtch.com/~rsc/regexp/) ([Lobsters](https://lobste.rs/s/nvxhdz/implementing_regular_expressions))
- [Avoid RegEx (2021)](https://robinwinslow.uk/how-to-regex#avoid-coding-in-regex-if-you-can)
- [RegEx Pal](https://www.regexpal.com/) - Test your regular expressions quickly.
- [Regular expression compilation visualized](https://compiler.org/reason-re-nfa/src/index.html) ([HN](https://news.ycombinator.com/item?id=26537734))
- [regex101](https://regex101.com/) - Build, test, and debug regex. ([Issues](https://github.com/firasdib/Regex101))
- [igrepper](https://github.com/igoyak/igrepper) - Interactive grepping using ncurses and rust.
- [Some useful regular expressions for programmers (2021)](https://lemire.me/blog/2021/04/22/some-useful-regular-expressions-for-programmers/) ([HN](https://news.ycombinator.com/item?id=26910717))
- [Bling Fire](https://github.com/microsoft/BlingFire) - Lightning fast Finite State machine and REgular expression manipulation library.
- [Let’s Build a Regex Engine (2019)](https://kean.blog/post/lets-build-regex) ([HN](https://news.ycombinator.com/item?id=27287987))
- [GNU BRE/ERE cheatsheet and differences between grep, sed and awk (2021)](https://learnbyexample.github.io/gnu-bre-ere-cheatsheet/)
- [The Best Regex Trick (2014)](http://rexegg.com/regex-best-trick.html) ([HN](https://news.ycombinator.com/item?id=27774584))
- [The Regular Expression Edition (2021)](https://whyisthisinteresting.substack.com/p/the-regular-expression-edition) - On code, early neural networks, and once discredited AI pioneers.
- [REXS](https://github.com/uellenberg/REXS) - Language for writing regular expressions.
- [Regex cheatsheet](https://remram44.github.io/regex-cheatsheet/regex.html) ([Code](https://github.com/remram44/regex-cheatsheet))
- [Regex Legends: The People Behind the Magic](https://blog.stevenlevithan.com/archives/regex-legends)
- [Regexly](https://regexly.js.org/) - WYSIWYG Regex playground for those who JavaScript. ([Code](https://github.com/NeekSandhu/regexly))
- [Orchestra](https://github.com/pouyakary/Orchestra) - Visual language that compiles into RegExp.
- [UltimateRegexResource](https://github.com/GoldinGuy/UltimateRegexResource) - Ultimate collection of regex syntax and resources.
- [Regex Generator](https://regex-generator.olafneumann.org/) - Generate regular expressions from sample texts. ([Code](https://github.com/noxone/regex-generator))
- [Regex-like domain-specific language for pattern matching syntax trees](https://github.com/fkohlgrueber/pattern-matching)
- [rxe](https://github.com/mtrencseni/rxe) - Literate and composable regular expressions.
- [Recognize Regex Easily (2020)](https://dev.to/jamland/recognize-regex-easily-4k5i)
- [Regex literals optimization (2020)](https://nitely.github.io/2020/11/30/regex-literals-optimization.html)
- [A DFA for submatches extraction (2020)](https://nitely.github.io/2020/01/19/a-dfa-for-submatches-extraction.html)
- [regex-cache](https://github.com/jonschlinkert/regex-cache) - Memoize the results of a call to the RegExp constructor, avoiding repetitious runtime compilation.
- [tree-sitter-regex](https://github.com/tree-sitter/tree-sitter-regex) - Regex grammar for tree-sitter.
- [rr4r](https://github.com/yutannihilation/rr4r) - Rust regex for R.
- [Regulex](https://jex.im/regulex/) - JavaScript Regular Expression Parser & Visualizer. ([Code](https://github.com/CJex/regulex))
- [Regexes are Cool and Good (2022)](https://buttondown.email/hillelwayne/archive/regexes-are-cool-and-good/) ([Lobsters](https://lobste.rs/s/wq3n73/regexes_are_cool_good))
- [RegEx Library](https://uibakery.io/regex-library) - Curated list of useful regular expressions for different programming languages.
- [A Regular Expression Matcher](https://www.cs.princeton.edu/courses/archive/spr09/cos333/beautiful.html)
- [Melody](https://github.com/yoav-lavi/melody) - Language that compiles to regular expressions and aims to be more easily readable and maintainable. ([HN](https://news.ycombinator.com/item?id=30358554)) ([Reddit](https://www.reddit.com/r/programming/comments/stvxxa/melody_a_language_that_compiles_to_regular/))
- [Alternative Regex Syntax](https://github.com/oilshell/oil/wiki/Alternative-Regex-Syntax) ([Lobsters](https://lobste.rs/s/molbhc/alternative_regex_syntax))
- [Python support for regular expressions (2022)](https://lwn.net/SubscriberLink/885682/ebb44eea5667f358/)
- [A regular expression to check for prime numbers (2007)](https://www.noulakaz.net/2007/03/18/a-regular-expression-to-check-for-prime-numbers/) ([HN](https://news.ycombinator.com/item?id=30564287))
- [Rulex](https://github.com/rulex-rs/rulex) - New, portable, regular expression language. ([Reddit](https://www.reddit.com/r/rust/comments/tbu6s3/announcing_rulex_a_new_regular_expression_language/)) ([Web](https://rulex-rs.github.io/)) ([HN](https://news.ycombinator.com/item?id=31690878))
- [rare](https://github.com/zix99/rare) - Real time regex-extraction and aggregation into common formats such as histograms, bar graphs, numerical summaries, tables, and more.
- [Regex Vis](https://regex-vis.com/) - Regex visualizer & editor. ([Code](https://github.com/Bowen7/regex-vis)) ([HN](https://news.ycombinator.com/item?id=31307123))
- [I-Regexp](https://github.com/cabo/iregexp) - Interoperable Regexp Format.
- [compose-regexp.js](https://github.com/compose-regexp/compose-regexp.js) - Build and compose maintainable regular expressions in JavaScript.
- [flexlint](https://github.com/dalance/flexlint) - Flexible linter with rules defined by regular expression.
- [Animating Regular Expressions With Python and Graphviz (2022)](https://betterprogramming.pub/animating-regular-expressions-with-python-and-graphviz-e0df447b827a)
- [Web tool to evaluate rust regular expressions](https://2fd.github.io/rust-regex-playground/) ([Code](https://github.com/2fd/rust-regex-playground))
- [super-regex](https://github.com/sindresorhus/super-regex) - Make a regular expression time out if it takes too long to execute.
- [patterns-finder](https://github.com/benouinirachid/patterns-finder) - Simple, Fast, Powerful and Easily extensible python package for extracting patterns from text, with over than 60 predefined Regular Expressions.
- [Utilties for the Developer. Regex, HTTP echo. Diffing](https://utils.zest.dev/regex) ([Code](https://github.com/zestcreative/utility))
- [Swift Regex](https://swiftregex.com/) - Online tool to learn, build and test Swift Regex Regex and Regex Builder. ([Code](https://github.com/SwiftFiddle/swiftregex))
- [regexparam](https://github.com/lukeed/regexparam) - Tiny (394B) utility that converts route patterns into RegExp.
- [AutoRegex](https://www.autoregex.xyz/) - Convert from English to RegEx with Natural Language Processing. ([HN](https://news.ycombinator.com/item?id=32032134))
- [magic-regexp](https://github.com/danielroe/magic-regexp) - Compiled-away, type-safe, readable RegExp alternative.
- [PRegEx](https://github.com/manoss96/pregex) - Programmable Regular Expressions.
- [Rex](https://github.com/areknawo/Rex) - JS Library for writing complex RegExps with help of auto completion.
