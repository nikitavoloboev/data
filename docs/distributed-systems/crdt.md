---
title: Conflict-free replicated data type
---

# [Conflict-free replicated data type](https://en.wikipedia.org/wiki/Conflict-free_replicated_data_type)

[Liveblocks](https://github.com/liveblocks/liveblocks) is a neat application of CRDTs. [Yjs](https://github.com/yjs/yjs) is nice too. [CRDT wiki](https://crdt.tech/) is great.

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
- [CRDTs wiki](https://crdt.tech/) ([Code](https://github.com/ept/crdt-website)) ([HN](https://news.ycombinator.com/item?id=30983770))
- [Awesome CRDT](https://github.com/alangibson/awesome-crdt)
- [Local-first packages & explorations](https://github.com/jaredly/local-first)
- [Automerge in JS](https://github.com/automerge/automerge) - JSON-like data structure (a CRDT) that can be modified concurrently by different users, and merged again automatically. ([Tweet](https://twitter.com/steveruizok/status/1421865156724805639)) ([HN](https://news.ycombinator.com/item?id=30412550)) ([Docs](https://automerge.org/docs/hello/)) ([HN](https://news.ycombinator.com/item?id=30881016))
- [Automergeable](https://github.com/jeffa5/automergeable)
- [Automerge in Rust](https://github.com/alexjg/automerge-rs)
- [Automerge persistent](https://github.com/jeffa5/automerge-persistent)
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
- [I was wrong. CRDTs are the future (2020)](https://josephg.com/blog/crdts-are-the-future/) ([HN](https://news.ycombinator.com/item?id=24617542)) ([Lobsters](https://lobste.rs/s/9fufgr/i_was_wrong_crdts_are_future)) ([Reddit](https://www.reddit.com/r/rust/comments/j1hb3a/i_was_wrong_crdts_are_the_future/)) ([HN](https://news.ycombinator.com/item?id=31049883))
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
- [How do CRDTs solve distributed data consistency challenges? (2021)](https://ably.com/blog/crdts-distributed-data-consistency-challenges) ([HN](https://news.ycombinator.com/item?id=29025724))
- [LWW-element-graphs study](https://github.com/agravier/crdt-study) ([Comment](https://news.ycombinator.com/item?id=29026978))
- [go-ds-crdt](https://github.com/ipfs/go-ds-crdt) - Distributed go-datastore implementation using Merkle-CRDTs.
- [ldb](https://github.com/vitorenesduarte/ldb) - Replication of CRDTs in Erlang.
- [CRDTs meet Redux](https://github.com/HerbCaudill/crdx)
- [Peritext](https://github.com/inkandswitch/peritext) - CRDT for asynchronous rich-text collaboration, where authors can work independently and then merge their changes. ([Article](https://www.inkandswitch.com/peritext/))
- [Automerge: a new foundation for collaboration software (2021)](https://www.youtube.com/watch?v=Qytg0Ibet2E) ([Tweet](https://twitter.com/martinkl/status/1465013736167833601)) ([HN](https://news.ycombinator.com/item?id=29501465))
- [SyncedStore CRDT](https://github.com/yousefed/SyncedStore) - Easy-to-use library for building live, collaborative applications that sync automatically. ([Docs](https://syncedstore.org/docs/)) ([HN](https://news.ycombinator.com/item?id=29483913))
- [valtio-yjs](https://github.com/dai-shi/valtio-yjs) - Proxy state library for ReactJS and VanillaJS. yjs is an implementation of CRDT algorithm (which allows to merge client data without server coordination).
- [Naisho](https://github.com/SerenityNotes/naisho) - Architecture for end-to-end encrypted CRDTs.
- [The limits of conflict-free replicated data types (2021)](http://pepijndevos.nl/2021/12/18/the-limits-of-conflict-free-replicated-data-types.html)
- [delta-crdts](https://github.com/peer-base/js-delta-crdts) - Delta State-based CRDTs in JavaScript.
- [Y-Py](https://github.com/y-crdt/ypy) - Python binding for Y-CRDT.
- [CRDT: Conflict-free Replicated Data Types (2018)](https://medium.com/@amberovsky/crdt-conflict-free-replicated-data-types-b4bfc8459d26)
- [Making CRDTs Byzantine Fault Tolerant](https://martin.kleppmann.com/papers/bft-crdt-papoc22.pdf) ([HN](https://news.ycombinator.com/item?id=30560573))
- [A highly-available move operation for replicated trees](https://martin.kleppmann.com/papers/move-op.pdf) ([HN](https://news.ycombinator.com/item?id=30811072))
- [Testing CRDTs in Rust, From Theory to Practice (2022)](https://www.ditto.live/blog/testing-crdts-in-rust-from-theory-to-practice) ([Lobsters](https://lobste.rs/s/uzuizm/testing_crdts_rust_from_theory_practice))
- [CRDTs](https://github.com/orbitdb/crdts) - Library of Conflict-Free Replicated Data Types for JavaScript.
- [CRDT Swift](https://github.com/bluk/CRDT) - Convergent and Commutative Replicated Data Types implementation in Swift.
- [Python CRDT](https://github.com/anshulahuja98/python3-crdt) - Python library for CRDTs (Conflict-free Replicated Data types).
- [Designing Data Structures for Collaborative Apps (2022)](https://mattweidner.com/2022/02/10/collaborative-data-design.html)
- [Programming Local-first Software (2022)](https://2022.ecoop.org/home/plf-2022)
- [checked-automerge](https://github.com/jeffa5/checked-automerge) - Attempt at model checking Automerge.
- [Vaxine](https://github.com/vaxine-io/vaxine) - Rich-CRDT database platform for backend applications. ([Web](https://vaxine.io/))
- [Yboard](https://github.com/felipeleivav/yboard) - Multiplayer desktop-like workspace based on CRDT. ([HN](https://news.ycombinator.com/item?id=31891608))
- [Automerge Chat Demo](https://github.com/okdistribute/automerge-chat-demo)
