# [Conflict-free replicated data type](https://en.wikipedia.org/wiki/Conflict-free_replicated_data_type)

## Notes

- [Distributed state is so fundamentally complex that I think we actually need CRDTs (or something like them) to reason about it effectively. And certainly to build reliable systems. The abstraction of a single, global, logical truth is so nice and tidy and appealing, but it becomes so leaky that I think all successful systems for distributed state will abandon it beyond a certain scale.](https://lobste.rs/s/9fufgr/i_was_wrong_crdts_are_future)

## Links

- [James Long - CRDTs for Mortals (2019)](https://www.dotconferences.com/2019/12/james-long-crdts-for-mortals)
- [Applied Monotonicity: A Brief History of CRDTs in Riak (2019)](http://christophermeiklejohn.com/erlang/lasp/2019/03/08/monotonicity.html)
- [Why CRDT didn't work out as well for collaborative editing xi-editor](https://github.com/xi-editor/xi-editor/issues/1187#issuecomment-491473599) ([HN](https://news.ycombinator.com/item?id=19886883))
- [How Figma’s multiplayer technology works (2019)](https://www.figma.com/blog/how-figmas-multiplayer-technology-works/)
- [HN thread on CRDT (2019)](https://news.ycombinator.com/item?id=21464189)
- [An API for data that changes over time (2019)](https://josephg.com/blog/api-for-changes/)
- [Statecraft](https://github.com/josephg/statecraft) - Protocol and set of tools for interacting with data that changes over time.
- [Full implementation of CRDTs using hybrid logical apps and a demo app that uses it](https://github.com/jlongster/crdt-example-app)
- [Swarm.js](https://github.com/gritzko/swarm) - JavaScript replicated model (M of MVC) library.
- [To OT or CRDT, that is the question (2020)](https://www.tiny.cloud/blog/real-time-collaboration-ot-vs-crdt/) ([HN](https://news.ycombinator.com/item?id=22039950))
- [CRDTs wiki](https://crdt.tech/) ([Code](https://github.com/ept/crdt-website))
- [Awesome CRDT](https://github.com/alangibson/awesome-crdt)
- [Local-first packages & explorations](https://github.com/jaredly/local-first)
- [Automerge in JS](https://github.com/automerge/automerge) - JSON-like data structure (a CRDT) that can be modified concurrently by different users, and merged again automatically. ([Tweet](https://twitter.com/steveruizok/status/1421865156724805639))
- [Automerge in Rust](https://github.com/alexjg/automerge-rs)
- [Peter Bourgon on CRDTs and State at the Edge (2020)](https://overcast.fm/+GdnXKIjWQ)
- [Moving Elements in List CRDTs](https://martin.kleppmann.com/papers/list-move-papoc20.pdf)
- [Chronofold: a data structure for versioned text (2020)](https://arxiv.org/abs/2002.09511v4)
- [roshi](https://github.com/soundcloud/roshi) - Large-scale CRDT set implementation for timestamped events.
- [CRDT benchmarks](https://github.com/dmonad/crdt-benchmarks) - Collection of reproducible benchmarks.
- [pony-crdt](https://github.com/jemc/pony-crdt) - Delta-State Convergent Replicated Data Types (ẟ-CRDTs) for the Pony language.
- [Jylis](https://github.com/jemc/jylis) - Distributed in-memory database for Conflict-free Replicated Data Types (CRDTs), built for speed, scalability, availability, and ease of use.
- [An introduction to Conflict-Free Replicated Data Types (2020)](https://lars.hupel.info/topics/crdt/01-intro) ([HN](https://news.ycombinator.com/item?id=23737639))
- [CRDTs: The Hard Parts (2020)](https://www.youtube.com/watch?v=x7drE24geUw) ([Abstract & References](https://martin.kleppmann.com/2020/07/06/crdt-hard-parts-hydra.html)) ([HN](https://news.ycombinator.com/item?id=23802208))
- [CRDTs in a Nutshell (2020)](https://amattn.com/p/riaks_two_contentions_and_crdts.html) ([Lobsters](https://lobste.rs/s/ipbe60/crdts_nutshell))
- [A First Replicating Type (2020)](https://appdecentral.com/2020/07/22/a-first-replicating-type/)
- [Conflict-Free Replicated Data Types (CRDTs) in Swift (2020)](https://appdecentral.com/2020/07/12/conflict-free-replicated-data-types-crdts-in-swift/)
- [Local-first software: You Own Your Data, in spite of the Cloud (2019)](https://www.inkandswitch.com/media/local-first/local-first.pdf) ([HN](https://news.ycombinator.com/item?id=24027663))
- [Are CRDTs suitable for shared editing? (2020)](https://blog.kevinjahns.de/are-crdts-suitable-for-shared-editing/)
- [Room Service](https://www.roomservice.dev/) - Faster-Than-Light Multiplayer Engine. ([Docs](https://www.roomservice.dev/docs))
- [I was wrong. CRDTs are the future (2020)](https://josephg.com/blog/crdts-are-the-future/) ([HN](https://news.ycombinator.com/item?id=24617542)) ([Lobsters](https://lobste.rs/s/9fufgr/i_was_wrong_crdts_are_future)) ([Reddit](https://www.reddit.com/r/rust/comments/j1hb3a/i_was_wrong_crdts_are_the_future/))
- [A Pragmatic Approach to Live Collaboration (2020)](https://hex.tech/blog/a-pragmatic-approach-to-live-collaboration)
- [Local First (2020)](https://brandur.org/nanoglyphs/014-local-first)
- [A Bluffers Guide to CRDTs in Riak](https://gist.github.com/russelldb/f92f44bdfb619e089a4d)
- [bft-crdts](https://github.com/davidrusu/bft-crdts) - Byzantine Fault Tolerant CRDT's and other Eventually Consistent Algorithms. In Rust.
- [Large-Scale Geo-Replicated Conflict-free Replicated Data Types](https://www.gsd.inesc-id.pt/~ler/reports/carlosbartolomeu-midterm.pdf)
- [CRDT notes](https://github.com/pfrazee/crdt_notes)
- [syncpad](https://github.com/Nishimura-Katsuo/syncpad) - CRDT-based collaborative source code editor (featuring the Monaco editor).
- [Verifying Strong Eventual Consistency in δ-CRDTs](https://github.com/ttaylorr/thesis)
- [Reactive CRDT](https://github.com/YousefED/reactive-crdt) - Easy-to-use library for building collaborative applications that sync automatically.
- [Conflict-free Reinterpretable Ordered Washed Data (CROWD)](https://github.com/hyoo-ru/crowd.hyoo.ru) - Like CRDT but better. ([Demo](https://crowd.hyoo.ru/))
- [crdt-tree](https://github.com/codesandbox/crdt-tree) - Implementation of a tree Conflict-Free Replicated Data Type in TS.
- [CRDT Tutorial for Beginners](https://github.com/ljwagerfield/crdt)
- [CRDT optimizations (2021)](https://bartoszsypytkowski.com/crdt-optimizations/)
- [Pure operation-based CRDTs (2021)](https://bartoszsypytkowski.com/pure-operation-based-crdts/)
- [Operation-based CRDTs: JSON document (2021)](https://bartoszsypytkowski.com/operation-based-crdts-json-document/)
- [Text CRDT Prototype](https://github.com/josephg/text-crdt-rust) - Prototype of a simple high performance CRDT for text. Its loosely based off automerge.
- [Faster CRDTs: An Adventure in Optimization (2021)](https://josephg.com/blog/crdts-go-brrr/) ([HN](https://news.ycombinator.com/item?id=28017204))
- [Chronofold](https://github.com/dkellner/chronofold) - Conflict-free replicated data structure (a.k.a. CRDT) for versioned text.
- [What to use to build collaborative web application today (2021)](https://twitter.com/tmcw/status/1433436431658196997)
- [HN: CRDT resources (2021)](https://news.ycombinator.com/item?id=28998767)
- [How do CRDTs solve distributed data consistency challenges? (2021)](https://ably.com/blog/crdts-distributed-data-consistency-challenges)
