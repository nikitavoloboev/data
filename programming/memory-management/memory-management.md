# Memory management

## Links

- [What Every Programmer Should Know About Memory (2007)](https://people.freebsd.org/~lstewart/articles/cpumemory.pdf) ([Web](https://lwn.net/Articles/250967/))
- [Memory Management Reference](https://www.memorymanagement.org/) - Resource for programmers and computer scientists interested in memory management and garbage collection.
- [Memory Pool System](https://www.ravenbrook.com/project/mps/) - Flexible and adaptable memory manager. Among its many advantages are an incremental garbage collector with very low pause times, and an extremely robust implementation. ([Code](https://github.com/Ravenbrook/mps))
- [A unified theory of garbage collection (2004)](https://www.researchgate.net/publication/221321424_A_unified_theory_of_garbage_collection)
- [Memory Bandwidth Napkin Math (2020)](https://www.forrestthewoods.com/blog/memory-bandwidth-napkin-math/)
- [Isolation Alloc](https://github.com/struct/isoalloc) - New general purpose secure memory allocator that implements an isolation security strategy to mitigate memory safety issues.
- [Broom](https://github.com/zesterer/broom) - Ergonomic tracing garbage collector that supports mark 'n sweep garbage collection.
- [Writing a Memory Allocator (2019)](http://dmitrysoshnikov.com/compilers/writing-a-memory-allocator/) ([HN](https://news.ycombinator.com/item?id=25402841))
- [Writing a Pool Allocator (2019)](http://dmitrysoshnikov.com/compilers/writing-a-pool-allocator/)
- [Writing a Mark-Sweep Garbage Collector (2020)](http://dmitrysoshnikov.com/compilers/writing-a-mark-sweep-garbage-collector/)
- [rpmalloc](https://github.com/mjansson/rpmalloc) - Cross platform lock free thread caching 16-byte aligned memory allocator implemented in C.
- [memleax](https://github.com/WuBingzheng/memleax) - Debugs memory leak of a running process by attaching it, without recompiling or restarting.
- [An introduction to virtual memory (2020)](https://www.internalpointers.com/post/introduction-virtual-memory) ([HN](https://news.ycombinator.com/item?id=23096747))
- [Visualizing memory management in Rust (2020)](https://deepu.tech/memory-management-in-rust/)
- [Memory Allocators 101 - Write a simple memory allocator (2015)](https://arjunsreedharan.org/post/148675821737/memory-allocators-101-write-a-simple-memory)
- [gperftools](https://github.com/gperftools/gperftools) - Collection of a high-performance multi-threaded malloc() implementation, plus some pretty nifty performance analysis tools.
- [Memory Consistency Models: A Tutorial (2016)](https://www.cs.utexas.edu/~bornholt/post/memory-models.html) ([HN](https://news.ycombinator.com/item?id=23546316))
- [Modern garbage collection (2016)](https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e) - A look at the Go GC strategy.
- [Virtual Memory Tricks (2017)](https://ourmachinery.com/post/virtual-memory-tricks/)
- [Reference counting vs GC (2020)](https://twitter.com/mraleph/status/1305209967641796611)
- [Memory allocator showdown (2020)](https://blog.janestreet.com/memory-allocator-showdown/)
- [Understanding memory initialization patterns](http://tasvideos.org/Nach/MemoryInit.html)
- [Custom Allocators Demystified (2020)](https://slembcke.github.io/2020/10/12/CustomAllocators.html) ([HN](https://news.ycombinator.com/item?id=24762840))
- [MMTk](https://github.com/mmtk/mmtk-core) - Framework for the design and implementation of memory managers. In Rust.
- [Quantifying the Performance of Garbage Collection vs. Explicit Memory Management (2005)](https://people.cs.umass.edu/~emery/pubs/gcvsmalloc.pdf)
- [The Lost Art of Structure Packing](http://catb.org/esr/structure-packing/) - Technique for reducing the memory footprint of programs in compiled languages with C-like structures - manually repacking these declarations for reduced size.
- [klox](https://github.com/dkopko/klox) - Proof-of-concept demonstration of O(1) garbage collection. ([HN](https://news.ycombinator.com/item?id=25161666)) ([Lobsters](https://lobste.rs/s/z96jmk/experimental_o_1_garbage_collector))
- [herdtools7](https://github.com/herd/herdtools7) - Tool suite to test weak memory models.
- [Sound Garbage Collection for C using Pointer Provenance (2020)](https://2020.splashcon.org/details/splash-2020-oopsla/52/Sound-Garbage-Collection-for-C-using-Pointer-Provenance)
- [Memory Management in Lobster](http://aardappel.github.io/lobster/memory_management.html)
- [Generational references faster than reference counting (2021)](https://vale.dev/blog/generational-references) ([Lobsters](https://lobste.rs/s/sglvcc/generational_references_2_3x_faster_than))
- [When allocators are hoarding your precious memory (2021)](https://www.algolia.com/blog/engineering/when-allocators-are-hoarding-your-precious-memory/)
- [Writing a Simple Garbage Collector in C (2020)](https://maplant.com/gc.html) ([Lobsters](https://lobste.rs/s/dd8k4k/writing_simple_garbage_collector_c_2020))
- [A look into Automatic Reference Counting (2021)](https://neelbakshi.medium.com/a-look-into-automatic-reference-counting-b17e9539d34f)
- [Papers for heap memory analysis and leak detection (2021)](https://github.com/grin-compiler/ghc-whole-program-compiler-project/issues/5)
