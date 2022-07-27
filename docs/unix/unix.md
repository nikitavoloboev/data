---
title: Unix
---

# [Unix](http://en.wikipedia.org/wiki/Unix)

[Unix philosophy is to make everything composable with lots of tiny individual programs doing one thing and communicating through text outputs.](https://twitter.com/chigbarg/status/1466497684953440259)

## Notes

- `/dev/tty` is a synonym for the controlling terminal (if any) of the current process.
- [Major theme of The Art of Unix Programming is to separate policy (what to do) from mechanism (how to do it). Policy is things like business rules that change rapidly in response to the real world; often policy for daemons like Nginx are entirely defined in config files. Mechanism on the other hand changes much slower, typically you write once and then only need minor bug fixes and other maintenance. Often this looks like library code such as a HTTP request parser. It’s good to write integration tests to verify policy code, and unit tests to verify mechanism.](https://lobste.rs/s/rgqg1v/what_functions_why_functions)

## Links

- [Shell scripts and tools](https://yoshuawuyts.gitbooks.io/knowledge/content/unix/unix.html)
- [Stream editing](https://yoshuawuyts.gitbooks.io/knowledge/content/unix/streams.html)
- [AT&T Archives: The UNIX Operating System](https://www.youtube.com/watch?v=tc4ROCJYbm0&t=4m8s) - Awesome lookback on history of UNIX.
- [cicada](https://github.com/mitnk/cicada) - Simple Unix shell written in Rust.
- [Test your sysadmin skills](https://github.com/trimstray/test-your-sysadmin-skills)
- [The History of Unix, Rob Pike (2018)](https://www.youtube.com/watch?v=_2NI6t2r_Hs)
- [For the Love of Pipes (2019)](https://blog.jessfraz.com/post/for-the-love-of-pipes/) ([HN](https://news.ycombinator.com/item?id=18967249))
- [VCF East 2019 -- Brian Kernighan interviews Ken Thompson (2019)](https://www.youtube.com/watch?time_continue=3&v=EY6q5dv_B-o)
- [Awesome UNIX](https://github.com/sirredbeard/Awesome-UNIX) - All the UNIX and UNIX-Like: Linux, BSD, macOS, Illumos, 9front, and more.
- [How Unix Works: Become a Better Software Engineer (2019)](https://neilkakkar.com/unix.html)
- [Unix Toolbox](http://cb.vu/unixtoolbox.xhtml) - Collection of Unix/Linux/BSD commands and tasks which are useful for IT work or for advanced users. ([HN](https://news.ycombinator.com/item?id=10022729))
- [Using ptrace To Call a Userspace Function](https://github.com/eklitzke/ptrace-call-userspace)
- [EOF is not a character (2020)](https://ruslanspivak.com/eofnotchar/) ([HN](https://news.ycombinator.com/item?id=22557412))
- [How are Unix pipes implemented? (2020)](https://toroid.org/unix-pipe-implementation)
- [Unix Considered Harmful](https://zge.us.to/unix-harmful.html)
- [Time on Unix (2020)](https://venam.nixers.net/blog/unix/2020/05/02/time-on-unix.html)
- [The Deprecated nix API (2020)](https://www.bitquabit.com/post/deprecated-nix-api/) ([Lobsters](https://lobste.rs/s/ojiulv/deprecated_nix_api))
- [Bell Labs Employee List](http://cm.bell-labs.co/who/)
- [The beauty of Unix pipelines (2020)](https://prithu.xyz/posts/unix-pipeline/) ([HN](https://news.ycombinator.com/item?id=23420786))
- [Relational pipes](https://relational-pipes.globalcode.info/v_0/index.xhtml) - Open data format designed for streaming structured data between two processes.
- [Is Sudo Almost Useless? (2020)](https://security.stackexchange.com/questions/232924/is-sudo-almost-useless) ([HN](https://news.ycombinator.com/item?id=23468456))
- [tildeverse](https://tildeverse.org/) - Loose association of like-minded tilde communities. ([tilde news](https://tilde.news/))
- [SDF Public Access UNIX System](https://sdf.org/)
- [Dennis Ritchie's Dissertation](https://minnie.tuhs.org/pipermail/tuhs/2020-August/021937.html)
- [Discovering Dennis Ritchie’s Lost Dissertation (2020)](https://computerhistory.org/blog/discovering-dennis-ritchies-lost-dissertation/) ([HN](https://news.ycombinator.com/item?id=23582070))
- [Pipe is a thing (2020)](https://blog.8-p.info/en/2020/06/16/pipe/)
- [The Unix-Haters Handbook (1994)](https://web.mit.edu/~simsong/www/ugh.pdf)
- [When Unix learned to reboot(2) (2020)](http://bsdimp.blogspot.com/2020/07/when-unix-learned-to-reboot2.html) ([Lobsters](https://lobste.rs/s/e0e0qe/when_unix_learned_reboot_2))
- [Modular Synthesis and UNIX (2020)](https://nora.codes/post/modular-synthesis-and-unix/) ([HN](https://news.ycombinator.com/item?id=24023727)) ([Lobsters](https://lobste.rs/s/cbtcax/modular_synthesis_unix))
- [Cron Helper](https://cron.help/) - Cron syntax for us humans.
- [File handling in Unix: tips, traps and outright badness (2020)](https://rachelbythebay.com/w/2020/08/11/files/) ([HN](https://news.ycombinator.com/item?id=24129113))
- [History of UNIX Design and Interfaces](https://github.com/penberg/unix-history)
- [Advanced Programming in the UNIX Environment](https://stevens.netmeister.org/631/) ([Lobsters](https://lobste.rs/s/zyt4hk/cs631_advanced_programming_unix)) ([Code](https://github.com/jschauma/cs631apue)) ([Rust Code](https://github.com/philippkeller/apue-rust))
- [Hurd, seL4, thoughts](https://nalaginrut.com/archives/2019/12/11/hurd%2c%20sel4%2c%20thoughts) ([Lobsters](https://lobste.rs/s/5bfhrj/hurd_sel4_thoughts))
- [Advanced Editing on Unix (Kernighan)](http://maibriz.de/unix/ultrix/etc/ae.pdf)
- [The UNIX Time-Sharing System (1974)](https://chsasank.github.io/classic_papers/unix-time-sharing-system.html) ([HN](https://news.ycombinator.com/item?id=24797312))
- [Web version of Lions' Commentary on UNIX 6th Edition](https://warsus.github.io/lions-/)
- [Cronie](https://github.com/cronie-crond/cronie) - Standard UNIX daemon crond that runs specified programs at scheduled times and related tools.
- [Unix Recovery Legend (1986)](https://www.ee.ryerson.ca/~elf/hack/recovery.html) ([HN](https://news.ycombinator.com/item?id=25491790))
- [Unix History Repository](https://github.com/dspinellis/unix-history-repo) - Continuous Unix commit history from 1970 until today.
- [Chain loading, not preloading: the dynamic linker as a virtualization vector (2021)](https://www.cs.kent.ac.uk/people/staff/srk21/blog/2021/01/04/)
- [USENIX](https://www.usenix.org/) - Advanced Computing Systems Association.
- [Supervisor](https://github.com/Supervisor/supervisor) - Client/server system that allows its users to monitor and control a number of processes on UNIX-like operating systems. ([Docs](http://supervisord.org/))
- [Ask HN: What are the bad parts of Unix? (2021)](https://news.ycombinator.com/item?id=26604833)
- [Detailed Behaviors of Unix Signal (2021)](https://www.dyx.name/posts/essays/signal.html)
- [Unix Philosophy with an Example (2019)](https://massimo-nazaria.github.io/blog/2019/03/02/unix-philosophy-with-an-example.html) ([Lobsters](https://lobste.rs/s/0zuri5/unix_philosophy_with_example))
- [Dark Side Of Posix Apis (2021)](https://vorner.github.io/2021/01/03/dark-side-of-posix-apis.html)
- [Ghosts of Unix Past](https://lwn.net/Articles/411845/) ([HN](https://news.ycombinator.com/item?id=27183784))
- [Evolution of the Unix System Architecture: An Exploratory Case Study (2019)](https://ieeexplore.ieee.org/document/8704965)
- [Program design in the UNIX environment (1984)](http://harmful.cat-v.org/cat-v/unix_prog_design.pdf) ([Lobsters](https://lobste.rs/s/a2k2pp/program_design_unix_environment_1984))
- [Unix tooling - join, don't extend (2021)](https://qmacro.org/2021/07/21/unix-tooling-join,-don't-extend/)
- [Fun with Unix domain sockets](https://simonwillison.net/2021/Jul/13/unix-domain-sockets/)
- [Posix Permissions Are Weird (2021)](https://paulcavallaro.com/blog/posix-permissions-are-weird/)
- [Unix – The Hole Hawg of Operating Systems (1999)](http://www.team.net/mjb/hawg.html) ([HN](https://news.ycombinator.com/item?id=28015229))
- [Unix and Microservice Platforms (2021)](https://blog.deref.io/unix-and-microservice-platforms/) ([HN](https://news.ycombinator.com/item?id=28039542))
- [On Unix composability](https://p.janouch.name/text/on-unix-composability.html)
- [Bringing the Unix Philosophy to the 21st Century: Make JSON a default output option. (2019)](https://blog.kellybrazil.com/2019/11/26/bringing-the-unix-philosophy-to-the-21st-century/) ([HN](https://news.ycombinator.com/item?id=28266193)) ([Reddit](https://www.reddit.com/r/programming/comments/pa4cbb/bringing_the_unix_philosophy_to_the_21st_century/))
- [I’m not sure that UNIX won (2021)](https://rubenerd.com/im-not-sure-that-unix-won/) ([Lobsters](https://lobste.rs/s/ezqjv5/i_m_not_sure_unix_won))
- [Unix as IDE (2012)](https://blog.sanctum.geek.nz/series/unix-as-ide/)
- [.plan](https://plan.cat/) ([Lobsters](https://lobste.rs/s/woxgih/plan)) ([HN](https://news.ycombinator.com/item?id=29248368))
- [Why V7 Unix matters so much (2021)](https://utcc.utoronto.ca/~cks/space/blog/unix/V7WhyItMattersSoMuch) ([HN](https://news.ycombinator.com/item?id=29308246))
- [The Open Group Base Specifications](https://pubs.opengroup.org/onlinepubs/9699919799.2018edition/)
- [Writing Your Own Sudo (2022)](https://rtpg.co/2022/02/13/your-own-sudo.html)
- [The early days of Unix at Bell Labs (2022)](https://www.youtube.com/watch?v=ECCr_KFl41E)
- [Unix Text Processing Book](https://github.com/larrykollar/Unix-Text-Processing) ([HN](https://news.ycombinator.com/item?id=30396667))
- [The history (sort of) of service management in Unix (2022)](https://utcc.utoronto.ca/~cks/space/blog/unix/ServiceManagementHistory)
- [Cron best practices (2016)](https://blog.sanctum.geek.nz/cron-best-practices/) ([HN](https://news.ycombinator.com/item?id=30636872))
- [Basics of the Unix Philosophy](https://homepage.cs.uri.edu/~thenry/resources/unix_art/ch01s06.html)
- [makeself](https://github.com/megastep/makeself) - Make self-extractable archives on Unix.
- [Advanced Programming in the UNIX Environment](https://people.cs.nctu.edu.tw/~chuang/courses/unixprog/) ([Code](https://github.com/ss8651twtw/Unix-Programming))
- [A Brief Introduction to termios (2009)](https://blog.nelhage.com/2009/12/a-brief-introduction-to-termios/)
- [daemonize](https://github.com/bmc/daemonize) - Command line utility to run a program as a Unix daemon.
- [Scripts to extend the Unix environment](https://github.com/thrig/scripts)
- [Unix command line conventions over time (2022)](https://blog.liw.fi/posts/2022/05/07/unix-cli/) ([HN](https://news.ycombinator.com/item?id=31293032))
- [Dropping privileges (2022)](https://blog.habets.se/2022/03/Dropping-privileges.html)
- [Replay Cron Events With Cron::Sequencer (2022)](https://leejo.github.io/2022/06/02/cron_sequencer/)
- [ugm](https://github.com/ariasmn/ugm) - Terminal based UNIX user and group viewer.
