# Concurrency

Concurrency is the composition of independently executing computations. [Real-Time Interrupt-driven Concurrency](https://rtic.rs/1/book/en/) is a great read.

## Notes

- [Race condition occurs when two or more threads can access shared data and they try to change it at the same time.](https://stackoverflow.com/questions/34510/what-is-a-race-condition/34550#34550)
- [Concurrency is when two or more tasks can start, run, and complete in overlapping time periods. It doesn't necessarily mean they'll ever both be running at the same instant. For example, multitasking on a single-core machine.](https://stackoverflow.com/questions/1050222/what-is-the-difference-between-concurrency-and-parallelism)
- [Concurrency is not Parallelism. Concurrency enables making progress on many tasks, without having to run A to completion to start work on B. Parallelism is actually running A and B simultaneously.](https://twitter.com/ktosopl/status/1400443844743684099)

## Links

- [The Deadlock Empire](https://deadlockempire.github.io/) - Interactive exercises.
- [Structured Concurrency](http://250bpm.com/blog:137)
- [A Separation Logic for Concurrent Randomized Programs](http://www.cs.cmu.edu/~rwh/papers/prob-conc/paper.pdf)
- [Lock-free programming for the masses (2016)](http://kcsrk.info/ocaml/multicore/2016/06/11/lock-free/)
- [CspExamples.jl](https://github.com/NHDaly/CspExamples.jl) - Julia implementations for the Example problems in Hoare's 1978 paper, "Communicating Sequential Processes".
- [I'm not feeling the async pressure (2020)](https://lucumr.pocoo.org/2020/1/1/async-pressure/) ([Lobsters](https://lobste.rs/s/xylmdn/i_m_not_feeling_async_pressure))
- [A (brief) retrospective on transactional memory (2010)](http://joeduffyblog.com/2010/01/03/a-brief-retrospective-on-transactional-memory/)
- [Concurrent programming, with examples (2020)](https://begriffs.com/posts/2020-03-23-concurrent-programming.html?hn=1) ([HN](https://news.ycombinator.com/item?id=22672128))
- [Deterministic Parallel Fixpoint Computation (2019)](https://arxiv.org/pdf/1909.05951.pdf)
- [Fearless concurrency: how Clojure, Rust, Pony, Erlang and Dart let you achieve that. (2019)](https://sites.google.com/a/athaydes.com/renato-athaydes/posts/fearlessconcurrencyhowclojurerustponyerlanganddartletyouachievethat) ([HN](https://news.ycombinator.com/item?id=19241427))
- [A Concurrency Cost Hierarchy (2020)](https://travisdowns.github.io/blog/2020/07/06/concurrency-costs.html)
- [Lockless Algorithms for Mere Mortals](https://lwn.net/SubscriberLink/827180/a1c1305686bfea67/) ([HN](https://news.ycombinator.com/item?id=23983508))
- [Structured Concurrency (2020)](https://ericniebler.com/2020/11/08/structured-concurrency/) ([HN](https://news.ycombinator.com/item?id=25032133))
- [Why Concurrency Is Hard (2020)](https://medium.com/oreillymedia/why-concurrency-is-hard-f93104cad54b) ([HN](https://news.ycombinator.com/item?id=25024901))
- [SOMns](https://github.com/smarr/SOMns) - Dynamic, class-based, object-oriented language in the tradition of Smalltalk and Self. Newspeak for Concurrency Research.
- [Concurrency Kit](http://concurrencykit.org/) - Concurrency primitives, safe memory reclamation mechanisms and non-blocking data structures for the research, design and implementation of high performance concurrent systems. ([Code](https://github.com/concurrencykit/ck))
- [Lock-free vs. wait-free concurrency (2010)](https://rethinkdb.com/blog/lock-free-vs-wait-free-concurrency)
- [Concurrent Expression Problem (2021)](https://matklad.github.io/2021/04/26/concurrent-expression-problem.html)
- [The Free Lunch Is Over: A Fundamental Turn Toward Concurrency in Software (2005)](http://www.gotw.ca/publications/concurrency-ddj.htm) ([HN](https://news.ycombinator.com/item?id=27649343))
- [What is structured concurrency? (2021)](https://oleb.net/2021/structured-concurrency/) ([Lobsters](https://lobste.rs/s/hsivmm/what_is_structured_concurrency))
- [Advanced Topics in Concurrency Course (2020)](https://github.com/anthezium/CS510-Advanced-Topics-in-Concurrency)
- [Syndicated Actors for Conversational Concurrency](https://syndicate-lang.org/) ([Lobsters](https://lobste.rs/s/bghtup/syndicated_actors_for_conversational))
- [Concepts of Concurrency (1997)](http://www.cs.nott.ac.uk/~pszgmh/con.pdf)
- [On parallelism and concurrency (2021)](https://inside.java/2021/11/30/on-parallelism-and-concurrency/)
- [KAIST Concurrency and Parallelism Laboratory](https://cp.kaist.ac.kr/) ([GitHub](https://github.com/kaist-cp))
- [Concurrency in modern programming languages](https://deepu.tech/concurrency-in-modern-languages/) ([Code](https://github.com/deepu105/concurrency-benchmarks))
- [Communicating Sequential Processes (CSP) (1978)](http://usingcsp.com/) ([PDF](https://www.cs.cmu.edu/~crary/819-f09/Hoare78.pdf)) ([Go Code](https://github.com/thomas11/csp))
- [Seamless, Fearless, and Structured Concurrency (2022)](https://verdagon.dev/blog/seamless-fearless-structured-concurrency)
- [Resources on concurrent PL semantics? (2022)](https://www.reddit.com/r/ProgrammingLanguages/comments/swxb61/resources_on_concurrent_pl_semantics/)
- [KAIST CS431: Concurrent Programming](https://github.com/kaist-cp/cs431)
- [Structured Concurrency](https://250bpm.com/blog:71/) ([HN](https://news.ycombinator.com/item?id=30720980))
- [Structured (Synchronous) Concurrency](https://fsantanna.github.io/structured-concurrency.html)
- [A Flexible Type System for Fearless Concurrency (2022)](https://www.cs.cornell.edu/andru/papers/gallifrey-types/) ([Reddit](https://www.reddit.com/r/ProgrammingLanguages/comments/tyiztq/a_flexible_type_system_for_fearless_concurrency/))
- [Exponential Backoff And Jitter (2022)](https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/)
- [Fearless concurrency at a discount? (2022)](https://newsletter.papersyoumightlove.pl/archive/fearless-concurrency-at-a-discount/)
- [Notes on concurrency bugs (2016)](https://danluu.com/concurrency-bugs/) ([HN](https://news.ycombinator.com/item?id=31514822))
- [Handling Concurrency Without Locks (2022)](https://hakibenita.com/django-concurrency) ([Lobsters](https://lobste.rs/s/abr8ls/handling_concurrency_without_locks))
- [Concurrent approaches to concurrency (2020)](https://susliko.dev/posts/concurrency/)
