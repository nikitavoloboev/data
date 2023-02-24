---
title: Reactive programming
---

# [Reactive programming](https://en.wikipedia.org/wiki/Reactive_programming)

[Super Charging Fine-Grained Reactive Performance](https://dev.to/modderme123/super-charging-fine-grained-reactive-performance-47ph) & [Hands-on Introduction to Fine-Grained Reactivity](https://dev.to/ryansolid/a-hands-on-introduction-to-fine-grained-reactivity-3ndf) are great reads.

Love [Solid](../programming-languages/javascript/js-libraries/solid.md)'s reactivity model. [Voby](https://github.com/vobyjs/voby) is interesting too.

## Notes

- [Event streams (like plain observables) emits values. Signals (and behavior subjects) have values. Queues (channels and some buffered and published observables) hold/store values. For example seminal paper on functional reactive programming defines two primitives- event streams and behaviors. They have different laws and this different capabilities. Queues are even more powerful but provide less guarantees.](https://www.reddit.com/r/reactjs/comments/yk6iy3/react_vdom_vs_preact_signal_performance_flame/)
- [Signals are not just about 'optimizing performance'. They do for ordinary data what React only does for DOM trees: provide a composable, consistent, managed way to derive it from your root application state, and to respond to changes over time.](https://twitter.com/djsheldrick/status/1629036972944633856)

## Links

- [Introduction to Reactive Programming you've been missing](https://gist.github.com/staltz/868e7e9bc2a7b8c1f754)
- [What is Reactive Programming? (2010)](http://paulstovell.com/blog/reactive-programming)
- [The essence and origins of FRP](https://github.com/conal/talk-2015-essence-and-origins-of-frp) - Keynote talk for LambdaJam 2015.
- [xstream](https://github.com/staltz/xstream) - Extremely intuitive, small, and fast functional reactive stream library for JavaScript.
- [CONNECTIVE](https://connective.dev/) - Agent-based reactive programming library for typescript.
- [RxViz](https://rxviz.com/) - Animated playground for Rx Observables. ([Code](https://github.com/moroshko/rxviz))
- [Explaining Streams to Rich Harris (2019)](https://johnlindquist.com/explaining-streams-to-rich-harris)
- [Towards a unified theory of reactive UI (2019)](https://raphlinus.github.io/ui/druid/2019/11/22/reactive-ui.html) ([HN](https://news.ycombinator.com/item?id=21607818))
- [Streams for reactive programming (2020)](https://surma.dev/things/streams-for-reactive-programming/)
- [David Khourshid — The visual future of reactive applications with statecharts (2020)](https://www.youtube.com/watch?v=o84Xw8qiTCw)
- [RxFeedback](https://github.com/NoTests/RxFeedback.swift) - Universal system operator and architecture for RxSwift.
- [reactive.how](https://reactive.how/) - Learn RxJS operators and Reactive Programming principles. ([Code](https://github.com/cedricss/reactive.how))
- [Kairo.js](https://github.com/3Shain/kairo) - Refined reactive programming pattern for web applications, the framework over frameworks.
- [A Hands-on Introduction to Fine-Grained Reactivity (2021)](https://dev.to/ryansolid/a-hands-on-introduction-to-fine-grained-reactivity-3ndf)
- [FRP Guides](https://github.com/HeinrichApfelmus/frp-guides) - Tutorials, guidelines, examples, patterns and half-baked ideas on functional reactive programming (FRP).
- [ShapeRankai](https://github.com/f5devcentral/shapeRank) - Targeted at data analytics, machine learning and reactive programming. Purely functional and statically typed. ([Talk](https://www.youtube.com/watch?v=vMO-CFlbYf8))
- [Main aspects of reactivity (2021)](https://dev.to/ninjin/main-aspects-of-reactivity-58co)
- [cellx](https://github.com/Riim/cellx) - Ultra-fast implementation of reactivity for JS.
- [Разбираемся в сортах реактивности](https://github.com/nin-jin/slides/blob/master/reactivity/readme.md) ([Tweet](https://twitter.com/andrey_sitnik/status/1460956157854752768))
- [Reactive](https://github.com/YousefED/reactive) - Super simple, yet powerful and performant library for State Management / Reactive Programming.
- [General Theory of Reactivity](https://github.com/kriskowal/gtor)
- [Scramjet](https://github.com/scramjetorg/framework-js) - Simple reactive stream programming framework in TypeScript.
- [Hands-on Introduction to Fine-Grained Reactivity (2021)](https://dev.to/ryansolid/a-hands-on-introduction-to-fine-grained-reactivity-3ndf)
- [Awesome Observables](https://github.com/sindresorhus/awesome-observables)
- [Building a Reactive Library from Scratch (2021)](https://dev.to/ryansolid/building-a-reactive-library-from-scratch-1i0p)
- [LiteObservable](https://github.com/lxsmnsyc/LiteObservable) - Cold Observables for JS in a lightweight fashion.
- [observable](https://github.com/ally-ui/observable) - Tiny observable implementation.
- [reactive-box](https://github.com/re-js/reactive-box) - Minimalistic, fast, and highly efficient reactivity.
- [Flimsy](https://github.com/fabiospampinato/flimsy) - Single-file <1kb min+gzip simplified implementation of the reactive core of Solid, optimized for clean code.
- [S.js](https://github.com/adamhaile/S) - Simple, Clean, Fast Reactive Programming in JavaScript. ([Tweet](https://twitter.com/RyanCarniato/status/1604764288052142080))
- [Kefir](https://github.com/kefirjs/kefir) - Reactive Programming library for JavaScript.
- [ArrowJS](https://github.com/justin-schroeder/arrow-js) - Reactivity without the framework. ([Tweet](https://twitter.com/jpschroeder/status/1593326779179073537)) ([Docs](https://www.arrow-js.com/))
- [Spred](https://github.com/art-bazhin/spred) - Simple and fast JavaScript reactive programming library.
- [Super Charging Fine-Grained Reactive Performance (2022)](https://dev.to/modderme123/super-charging-fine-grained-reactive-performance-47ph)
- [Reactively](https://github.com/modderme123/reactively) - Library for fine grained reactive programming.
- [Reactivity & Rendering (2022)](https://www.youtube.com/watch?v=R5AcOtxIdMk)
- [Learning Leptos: Build a fine-grained reactive system in 100 lines of code (2022)](https://www.youtube.com/watch?v=GWB3vTWeLd4)
- [Reactor.js](https://github.com/fynyky/reactor.js) - Simple reactive programming without a framework.
- [sprae](https://github.com/dy/sprae) - DOM hydration with reactive attributes.
- [The Future of Fine-Grained Reactivity w/ Milo (2023)](https://www.youtube.com/watch?v=A2TiLvCDKSg)
- [Cont-Signal](https://github.com/rkirov/cont-signal) - Continuations based Signals JS framework.
- [Проектируем идеальную систему реактивности](https://page.hyoo.ru/#!=eh2o9_cl9nuy/View%22eh2o9_cl9nuy%22.Details=Reactive%20ReactJS)
- [JS Reactivity Benchmark](https://github.com/modderme123/js-reactivity-benchmark)
