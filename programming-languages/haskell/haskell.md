# [Haskell](https://www.haskell.org)

[The book of types](http://reasonablypolymorphic.com/types.pdf) is a pretty great intro to type systems and Type-Level Programming in Haskell.

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