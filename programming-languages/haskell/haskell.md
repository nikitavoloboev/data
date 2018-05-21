# [Haskell](https://www.haskell.org)
A language I want to learn.

## Notes
- [Kinds are the types of types. Typeclasses are more like predicates or relations between types.](https://www.reddit.com/r/haskell/comments/8cdiql/isnt_a_typeclass_just_a_type_of_a_type/)
- The fewer type parameters in a class, the better. Can you turn any into associated types? Can you split the class into two classes? Can you hive off some of the parameters into a superclass?
- Don't worry about strictness until it's time to optimise.
- Intuition about optimisation tends to be bad. Before profiling, limit yourself to reasoning about complexity classes.
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
- [Simon Peyton Jones - Closer to Nirvana](https://www.youtube.com/watch?v=xmjvOLlCdFU)
- [Curry-Howard-Lambek correspondence](https://wiki.haskell.org/Curry-Howard-Lambek_correspondence)
- [Theoretical foundations](https://wiki.haskell.org/Category:Theoretical_foundations)
- [FP course](https://github.com/data61/fp-course)
- [Russian Haskell book FAQ](https://www.ohaskell.guide/haskell-faq.html)
- [An opinionated guide to Haskell in 2018](https://lexi-lambda.github.io/blog/2018/02/10/an-opinionated-guide-to-haskell-in-2018/)
- [Please Don't Learn Category Theory to Learn Haskell](https://jozefg.bitbucket.io/posts/2013-10-14-please-dont-learn-cat-theory.html)
- [How do you guys get anything done?](https://www.reddit.com/r/haskell/comments/5wb5qw/how_do_you_guys_get_anything_done/)
- [If Haskell is so great, why hasn't it taken over the world? And the curious case of Go.](https://pchiusano.github.io/2017-01-20/why-not-haskell.html)