# [Haskell](https://www.haskell.org)

[Haskell book](http://haskellbook.com/) & [Thinking with types](https://leanpub.com/thinking-with-types) are great intros to Haskell, type systems and Type-Level Programming in Haskell.

## Notes

- [Kinds are the types of types. Typeclasses are more like predicates or relations between types.](https://www.reddit.com/r/haskell/comments/8cdiql/isnt_a_typeclass_just_a_type_of_a_type/)
- The fewer type parameters in a class, the better. Can you turn any into associated types? Can you split the class into two classes? Can you hive off some of the parameters into a superclass?
- Don't worry about strictness until it's time to optimize.
- Intuition about optimization tends to be bad. Before profiling, limit yourself to reasoning about complexity classes.
- Don't judge haskell based on these definitions. They appear strange and crazy, but when you actually do stuff most of them turn out to be quite intuitive.
  - Advice is to ignore these things. Don't read a million monad tutorials. Just play around and code something, you don't have to understand the monad-definition before you can use the do-notation. Try to ignore them. After a while you get an intuition and then the definitions will make sense.
- Objects are a way to generalize reusability and composability (Interheritance, Encapsulation, Polymorphism) Haskell goes down a different route of reusability and composability that draws more from math than from object models. The real benefit to learning haskell is learning to think in that way rather than an object oriented or imperative way.
- In Haskell, you could say that the type system is a syntactic method for classifying program phrases according to the kinds of values they compute.
- Because Haskell (usually) only evaluates into [WHNF](https://wiki.haskell.org/Weak_head_normal_form) if you produce data constructors while doing the infinite recursion you can use infinite recursion.
- If, while forcing a thunk, you see the same thunk being forced again, you can report infinite recursion.
- [GADTs](https://en.wikibooks.org/wiki/Haskell/GADT) might give you a radically simple way to express things that rudimentary type systems like Go's cant express.
- Haskell's type system lets you express a lot of domain specific invariants and logic in a statically verifiable manner.
- Go is as type safe as haskell, the language doesn't allow invalid operations on data types. What people mean with "weak" is often expressivity. Go's enums sucks and it doesn't have ADTs or generic.
  - When people talk about strong type system, they usually just mean expressivity, not safety.

## Links

- [How I Learned to Stop Worrying and Love the Type System](http://reasonablypolymorphic.com/blog/love-types/)
- [Simon Peyton Jones - Closer to Nirvana](https://www.youtube.com/watch?v=xmjvOLlCdFU)
- [Curry-Howard-Lambek correspondence](https://wiki.haskell.org/Curry-Howard-Lambek_correspondence)
- [Theoretical foundations](https://wiki.haskell.org/Category:Theoretical_foundations)
- [FP course](https://github.com/data61/fp-course)
- [Russian Haskell book FAQ](https://www.ohaskell.guide/haskell-faq.html)
- [An opinionated guide to Haskell in 2018](https://lexi-lambda.github.io/blog/2018/02/10/an-opinionated-guide-to-haskell-in-2018/)
- [Please Don't Learn Category Theory to Learn Haskell](https://jozefg.bitbucket.io/posts/2013-10-14-please-dont-learn-cat-theory.html)
- [How do you guys get anything done?](https://www.reddit.com/r/haskell/comments/5wb5qw/how_do_you_guys_get_anything_done/)
- [If Haskell is so great, why hasn't it taken over the world? And the curious case of Go.](https://pchiusano.github.io/2017-01-20/why-not-haskell.html)
- [Haskell IDE Engine](https://github.com/haskell/haskell-ide-engine) - Engine for haskell ide-integration.
- [Of Ideas and men](http://reasonablypolymorphic.com/blog/ideas-and-men/)
- [What is the track to mastering Haskell and where would it lead me professionally?](https://www.quora.com/profile/Edward-Kmett)
- [Nix and Haskell in production](https://github.com/Gabriel439/haskell-nix)
- [Revisiting Combinators by Edward Kmett](https://www.youtube.com/watch?v=PA1Fc7DNKtA)
- [Haskell Programming From First Principles - Follow-up Resources](https://github.com/pushcx/hpffp-resources)
- [Quchen's articles](https://github.com/quchen/articles)
- [Haskell Design Patterns?](https://www.reddit.com/r/haskell/comments/5r271m/haskell_design_patterns/)
- [Haskell companies](https://github.com/erkmos/haskell-companies#readme) - Curated list of companies using Haskell in industry.
- [Snack](https://github.com/nmattia/snack#readme) - Nix-based incremental build tool for Haskell projects.
- [What exactly makes the Haskell type system so revered (vs say, Java)?](https://softwareengineering.stackexchange.com/questions/279316/what-exactly-makes-the-haskell-type-system-so-revered-vs-say-java)
- [HN: Introducing Haskell to a Company](https://news.ycombinator.com/item?id=18118874)
- [Haskell's kind system - a primer](https://diogocastro.com/blog/2018/10/17/haskells-kind-system-a-primer/)
- [Fused effects](https://github.com/robrix/fused-effects) - Fast, flexible, fused effect system for Haskell.
- [Examples of Dependently-typed programs in Haskell](https://github.com/sweirich/dth)
- [haskell-lsp](https://github.com/alanz/haskell-lsp) - Haskell library for the Microsoft Language Server Protocol.
- [Lambda World 2018 - Opening Keynote by Edward Kmett](https://www.youtube.com/watch?v=HGi5AxmQUwU)
- [Haskell Source Extensions](https://github.com/haskell-suite/haskell-src-exts) - Package for handling and manipulating Haskell source code.
- [Hedgehog](https://github.com/hedgehogqa/haskell-hedgehog) - Modern property-based testing system, in the spirit of QuickCheck.
- [Haskeleton: a Haskell project skeleton](http://taylor.fausak.me/2014/03/04/haskeleton-a-haskell-project-skeleton/)
- [Haskell ITMO course at CTDHaskell ITMO course at CTD](https://github.com/jagajaga/FP-Course-ITMO#readme)
- [Google's Haskell training: 101 and 102](https://github.com/google/haskell-trainings)
- [brittany](https://github.com/lspitzner/brittany) - Haskell source code formatter.
- [Stack](https://github.com/commercialhaskell/stack) - Cross-platform program for developing Haskell projects.
- [List of Foundational Haskell Papers](https://github.com/cohomolo-gy/haskell-resources#readme)
- [hpack](https://github.com/sol/hpack) - Modern format for Haskell packages.
- [Haskell Code Explorer](https://github.com/alexwl/haskell-code-explorer) - Web application for exploring and understanding Haskell codebases.
- [static-haskell-nix](https://github.com/nh2/static-haskell-nix) - Easily build most Haskell programs into fully static Linux executables.
- [Real World Haskell by Bryan O’Sullivan updated to 2019](https://github.com/tssm/up-to-date-real-world-haskell#readme)
- [Higher-order Type-level Programming in Haskell (2019)](https://www.microsoft.com/en-us/research/uploads/prod/2019/03/ho-haskell-5c8bb4918a4de.pdf) ([HN](https://news.ycombinator.com/item?id=19412667))
- [Why Haskell?](https://github.com/github/semantic/blob/master/docs/why-haskell.md)
- [Property-Based Testing in a Screencast Editor: Introduction (2019)](https://wickstrom.tech/programming/2019/03/02/property-based-testing-in-a-screencast-editor-introduction.html)
- [Haskell Fan Site](http://www-cs-students.stanford.edu/~blynn/haskell/) ([HN](https://news.ycombinator.com/item?id=20255694))
- [An opinionated beginner's guide to Haskell in mid 2019](https://github.com/theindigamer/not-a-blog/blob/master/opinionated-haskell-guide-2019.md)
- [Why I (as of June 22 2019) think Haskell is the best general purpose language](http://www.philipzucker.com/why-i-as-of-june-22-2019-think-haskell-is-the-best-general-purpose-language-as-of-june-22-2019/)
- [Ormolu](https://github.com/tweag/ormolu) - Formatter for Haskell source code.
- [Koka](https://github.com/koka-lang/koka) - Function-oriented language with effect inference.
- [Write You a Haskell 2](https://github.com/JKTKops/Write-You-a-Haskell-2) - Continuation of Steven Diehl's Write You a Haskell.
- [Verification with Refinement Types](https://ranjitjhala.github.io/CAV19-tutorial/00-outline.html)
- [Hoogle](https://github.com/ndmitchell/hoogle) - Haskell API search engine.
- [Haskell Papers](https://mitchellwrosen.github.io/haskell-papers/)
- [Simon Peyton Jones how GHC type inference engine actually works (2019)](https://www.youtube.com/watch?v=x3evzO8O9e8)
- [Chris Penner - Comonads by Example (2019)](https://www.youtube.com/watch?v=HOmOQnQGtPU&list=PLcAu_kKy-krxDD1WwRX_9rc0knAFK3nHs&index=7)
