# Bash

It's very useful to use [ShellCheck](https://www.shellcheck.net) to check any shell code you write.

## Notes

-   The proper she-bang for Bash is `#!/usr/bin/env bash`.

## Code

```bash
# Check if command is in PATH
checkDep() {
     path=`command -v ${1}` && echo "${1} found at ${path}" || { echo "${1} not found" >&2 ; exit 1; }
}
```

```bash
# Get output of command. https://www.cyberciti.biz/faq/unix-linux-bsd-appleosx-bash-assign-variable-command-output/
# i.e. save output of date to var now
now=$(date)
```

```md
File testing

![](https://i.imgur.com/QGkHbPm.png)
```

```bash
# Pipe output to file.
# i.e. pipe output of ls to output.txt
ls > output.txt
```

```bash
# Check if no arguments passed
if [ $# -eq 0 ] then ... fi
```

```bash
# Check if file does not exist
# Putting ! before makes it a not statement. Spaces before and after [] are important.
if [ ! -f ~/Desktop/file.txt ]; then
    echo "File not found!"
fi
```

```bash
# Source vs ./
# Runs the script as an executable file, launching a new shell to run it
./script

# Reads and executes commands from filename in the current shell environment
source script
```

```bash
# need to wrap the cd command inside () to run it in scope of the cd
# i.e. pod install will be run inside ios dir
(cd ios && pod install)
```

## Links

-   [Parameter expansion](http://wiki.bash-hackers.org/syntax/pe)
-   [Defensive BASH Programming](http://www.kfirlavi.com/blog/2012/11/14/defensive-bash-programming/)
-   [Safe ways to do things in Bash](https://github.com/anordal/shellharden/blob/master/how_to_do_things_safely_in_bash.md) ([HN](https://news.ycombinator.com/item?id=17057596))
-   [Pure Bash Bible](https://github.com/dylanaraps/pure-bash-bible) - Collection of pure bash alternatives to external processes. ([HN](https://news.ycombinator.com/item?id=21013150)) ([HN](https://news.ycombinator.com/item?id=30363506))
-   [Bash Infinity](https://github.com/niieani/bash-oo-framework) - Modern boilerplate / framework / standard library for Bash.
-   [Funky](https://github.com/bbugyi200/funky) - Takes shell functions to the next level by making them easier to define, more flexible, and more interactive.
-   [create-bash-script](https://github.com/nikita-skobov/create-bash-script) - Bash script designed to create other bash scripts with argument parsing.
-   [Bash-LSP](https://github.com/bash-lsp/bash-language-server) - Language server for Bash. ([HN](https://news.ycombinator.com/item?id=26663135))
-   [Bash Reference Manual](https://tiswww.case.edu/php/chet/bash/bashref.html)
-   [Bash Guide](https://mywiki.wooledge.org/BashGuide) ([Lobsters](https://lobste.rs/s/onsous/bash_guide))
-   [How To Automate Basic Development Tasks with Bash](https://jmulholland.com/how-to-automate-basic-development-tasks-with-bash)
-   [Mini Gitbook for a presentation about BASH](https://erkanerol.github.io/bash-lingua-non-grata/#/) ([Code](https://github.com/erkanerol/bash-lingua-non-grata))
-   [Bash $\* and $@ (2017)](https://eklitzke.org/bash-$%2A-and-$@) ([HN](https://news.ycombinator.com/item?id=22027809))
-   [Understanding Bash: Elements of Programming (2018)](https://www.linuxjournal.com/content/understanding-bash-elements-programming) ([HN](https://news.ycombinator.com/item?id=22052890))
-   [Bash tricks](https://github.com/SimonBaeumer/bash-tricks) - Simple bash tricks which will make your life easier. ([Reddit](https://www.reddit.com/r/bash/comments/eokp46/some_simple_bash_tipps/))
-   [Anybody can write good bash (with a little effort) (2020)](https://blog.yossarian.net/2020/01/23/Anybody-can-write-good-bash-with-a-little-effort) ([Lobsters](https://lobste.rs/s/y0nx8o/anybody_can_write_good_bash_with_little))
-   [Bash Quick References](https://shellmagic.xyz/)
-   [critic.sh](https://github.com/Checksum/critic.sh) - Dead simple testing framework for Bash with coverage.
-   [Bash Hackers Wiki](https://wiki.bash-hackers.org/) ([HN](https://news.ycombinator.com/item?id=22382686))
-   [Bats](https://github.com/bats-core/bats-core) - Bash Automated Testing System.
-   [HN: Renaming files with mv without typing the name two times](https://news.ycombinator.com/item?id=22859935)
-   [The first two statements of your BASH script should be.. (2020)](https://ashishb.net/all/the-first-two-statements-of-your-bash-script-should-be/) ([Lobsters](https://lobste.rs/s/ajoaje/first_two_statements_your_bash_script))
-   [Use Bash Strict Mode (Unless You Love Debugging)](http://redsymbol.net/articles/unofficial-bash-strict-mode/)
-   [Bash code](https://github.com/bminor/bash)
-   [Some Relatively Obscure Bash Tips (2020)](https://zwischenzugs.com/2020/05/09/some-relatively-obscure-bash-tips/) ([HN](https://news.ycombinator.com/item?id=23126305))
-   [A Bash Cheat Sheet: Top 25 most-used commands, and how to create custom commands (2020)](https://medium.com/better-programming/bash-cheat-sheet-top-25-commands-and-creating-custom-commands-75941dcdc450)
-   [Bash Env Variable Defaults](https://www.yesthatblog.com/post/0065-env-defaults/)
-   [Using Bash traps in your scripts (2020)](https://opensource.com/article/20/6/bash-trap)
-   [Help message for shell scripts (2020)](https://samizdat.dev/help-message-for-shell-scripts/) ([HN](https://news.ycombinator.com/item?id=23763166)) ([Lobsters](https://lobste.rs/s/5njqrb/help_message_for_shell_scripts))
-   [Supercharge Your Bash History (2020)](https://metaredux.com/posts/2020/07/07/supercharge-your-bash-history.html) ([Lobsters](https://lobste.rs/s/ruygyw/supercharge_your_bash_history))
-   [Import](https://import.pw/) - Simple and Fast Module System for Bash and Other Unix Shells. ([HN](https://news.ycombinator.com/item?id=23864909))
-   [Intelligent Bash (ibash)](https://rdmp.org/dmbcs/i-bash)
-   [Functional programming in bash](https://github.com/ssledz/bash-fun) ([Slides](https://ssledz.github.io/presentations/bash-fun.html#/))
-   [Bash Pitfalls](https://mywiki.wooledge.org/BashPitfalls) ([Lobsters](https://lobste.rs/s/1vqimp/bash_pitfalls))
-   [argbash](https://github.com/matejak/argbash) - Bash argument parsing code generator. ([HN](https://news.ycombinator.com/item?id=24636367)) ([HN](https://news.ycombinator.com/item?id=30737701))
-   [Bash Bracket Cheat Sheet](https://wizardzines.com/comics/brackets-cheatsheet/)
-   [Bash Error Handling](https://wizardzines.com/comics/bash-errors/) ([HN](https://news.ycombinator.com/item?id=24727495))
-   [ctypes.sh](https://github.com/taviso/ctypes.sh) - Foreign function interface for bash. ([HN](https://news.ycombinator.com/item?id=24738814))
-   [DevHints Bash Cheat Sheet](https://devhints.io/bash)
-   [sbang](https://github.com/spack/sbang) - Run scripts with very long shebang lines. ([Lobsters](https://lobste.rs/s/rin6rc/sbang_lets_you_run_scripts_with_very_long))
-   [Oh My Bash](https://github.com/ohmybash/oh-my-bash) - Open source, community-driven framework for managing your bash configuration.
-   [An Awful Edge Case in Bash's set -e (2020)](http://jbrot.com/blog/dash_e_problems.html)
-   [Variant](https://github.com/mumoshu/variant2) - Turn your bash scripts into a modern, single-executable CLI app.
-   [Even faster bash startup (2020)](https://work.lisk.in/2020/11/20/even-faster-bash-startup.html)
-   [How To Use Bash Parameter Substitution Like A Pro (2020)](https://www.cyberciti.biz/tips/bash-shell-parameter-substitution-2.html)
-   [Introduction to Bash Scripting Book](https://github.com/bobbyiliev/introduction-to-bash-scripting)
-   [Bash Snippets](https://github.com/alexanderepstein/Bash-Snippets) - Collection of small bash scripts for heavy terminal users with no dependencies.
-   [Minimal safe Bash script template (2020)](https://betterdev.blog/minimal-safe-bash-script-template/) ([Lobsters](https://lobste.rs/s/yeloyn/minimal_safe_bash_script_template)) ([HN](https://news.ycombinator.com/item?id=25428621))
-   [Bash 5.1](https://lists.gnu.org/archive/html/info-gnu/2020-12/msg00003.html) ([HN](https://news.ycombinator.com/item?id=25492551))
-   [Batsh](https://batsh.org/) - Language that compiles to Bash and Windows Batch. ([Code](https://github.com/batsh-dev-team/Batsh))
-   [Bash HTTP monitoring dashboard](https://raymii.org/s/software/Bash_HTTP_Monitoring_Dashboard.html)
-   [Parallel bash](https://github.com/Akianonymus/parallel-bash) - Parallel processing of commands in pure bash along with the support of functions.
-   [Learn Bash in 27 minutes](https://github.com/pforret/LearnBashQuickly)
-   [Better BASHing Through Technology (2020)](https://andydote.co.uk/2020/08/28/better-bashing-through-technology/)
-   [bashful](https://github.com/wagoodman/bashful) - Use a yaml file to stitch together commands and bash snippets and run them with a bit of style.
-   [bash_unit](https://github.com/pgrange/bash_unit) - Bash unit testing.
-   [Elegant bash conditionals (2021)](https://timvisee.com/blog/elegant-bash-conditionals/) ([Lobsters](https://lobste.rs/s/nao13f/elegant_bash_conditionals))
-   [Bash lambda](https://github.com/spencertipping/bash-lambda) - Anonymous functions and FP stuff for bash.
-   [How to navigate directories faster with bash (2015)](https://mhoffman.github.io/2015/05/21/how-to-navigate-directories-with-the-shell.html) ([HN](https://news.ycombinator.com/item?id=26899531))
-   [Your `~/.bashrc` doesn't have to be a mess (2021)](https://write.as/bpsylevc6lliaspe) ([Lobsters](https://lobste.rs/s/r1tpld/your_bashrc_doesn_t_have_be_mess))
-   [Bash Function Names Can Be Almost Anything (2021)](https://blog.dnmfarrell.com/post/bash-function-names-can-be-almost-anything/) ([HN](https://news.ycombinator.com/item?id=27726699))
-   [Writing a Bash Builtin in C to Parse INI Configs (2021)](https://mbuki-mvuki.org/posts/2021-07-12-writing-a-bash-builtin-in-c-to-parse-ini-configs/) ([Lobsters](https://lobste.rs/s/6wjuk4/writing_bash_builtin_c_parse_ini_configs))
-   [Bashly](https://github.com/DannyBen/bashly) - Create beautiful bash scripts from simple YAML configuration. ([HN](https://news.ycombinator.com/item?id=28305479)) ([Docs](https://bashly.dannyb.co/))
-   [tree-sitter-bash](https://github.com/tree-sitter/tree-sitter-bash) - Bash grammar for tree-sitter.
-   [Bash Boilerplate](https://github.com/xwmx/bash-boilerplate) - Collection of Bash scripts for creating safe and useful command line programs.
-   [Bash-TPL](https://github.com/TekWizely/bash-tpl) - Smart, lightweight shell script templating engine, written in Bash.
-   [Bash functions are better than I thought (2021)](https://cuddly-octo-palm-tree.com/posts/2021-10-31-better-bash-functions/) ([Reddit](https://www.reddit.com/r/programming/comments/qjnzmn/underwhelmed_by_bash_functions_maybe_youre_using/)) ([HN](https://news.ycombinator.com/item?id=29058140))
-   [bash-completion](https://github.com/scop/bash-completion) - Programmable completion functions for bash.
-   [Bash patterns I use weekly](https://will-keleher.com/posts/5-Useful-Bash-Patterns.html) ([HN](https://news.ycombinator.com/item?id=29318751))
-   [How to write idempotent Bash scripts (2019)](https://arslan.io/2019/07/03/how-to-write-idempotent-bash-scripts/) ([HN](https://news.ycombinator.com/item?id=29483070)) ([Lobsters](https://lobste.rs/s/ixuahi/how_write_idempotent_bash_scripts))
-   [Mastering Bash and Terminal (2017)](https://www.blockloop.io/mastering-bash-and-terminal/) ([HN](https://news.ycombinator.com/item?id=13400350))
-   [bish](https://github.com/tdenniston/bish) - Language that compiles to Bash. Shell scripting with a modern feel.
-   [Sherver](https://github.com/remileduc/sherver) - Pure Bash lightweight web server. ([HN](https://news.ycombinator.com/item?id=29648135))
-   [Modern Bash (Zsh) Scripting](https://www.mulle-kybernetik.com/modern-bash-scripting/) ([Lobsters](https://lobste.rs/s/gwzjjw/modern_bash_scripting))
-   [NL2Bash: A Corpus and Semantic Parser for Natural Language Interface to the Linux Operating System (2018)](https://arxiv.org/abs/1802.08979) ([Code](https://github.com/TellinaTool/nl2bash))
-   [Ask HN: Let's build Check style for Bash? (2022)](https://news.ycombinator.com/item?id=30405177)
-   [Replicating Bash Argument Splitting (2021)](https://blog.dnmfarrell.com/post/replicating-bash-argument-splitting/)
-   [Creating a bash completion script (2018)](https://iridakos.com/programming/2018/03/01/bash-programmable-completion-tutorial) ([HN](https://news.ycombinator.com/item?id=30524440))
-   [5 Modern Bash Scripting Techniques That Only A Few Programmers Know (2022)](https://levelup.gitconnected.com/5-modern-bash-scripting-techniques-that-only-a-few-programmers-know-4abb58ddadad)
-   [bpkg](https://github.com/bpkg/bpkg) - Lightweight bash package manager. ([Web](https://bpkg.sh/))
-   [Bash One liners](https://github.com/onceupon/Bash-Oneliner) - Collection of handy Bash One-Liners and terminal tricks. ([HN](https://news.ycombinator.com/item?id=31250275))
-   [CLI argument parser for Bash scripts and functions](https://bashjazz.orion3.space/cli-args.html)
-   [Type-ish](https://github.com/Mythra/typeish) - Runtime type checker for bash functions, implemented entirely in bash.
-   [Bash, Pipes, & Socket SDK (2022)](https://socketsupply.co/blog/bash-pipes-and-socket-sdk/)
-   [Vercel Bash](https://github.com/importpw/vercel-bash) - Bash Runtime for Vercel serverless functions.
-   [Beautysh](https://github.com/lovesegfault/beautysh) - Bash beautifier for the masses.
