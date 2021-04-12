# Category theory

Category theory is an abstract branch of mathematics used to model mathematical "objects", and mappings between them.

In order to define a category you define the mathematical objects involved (such as "Sets" for instance), and the mappings (such as "Functions" for instance). So there is a category of mathematical sets and their functions.

Now the categories themselves could also be considered mathematical objects, and we could look at mappings between categories as well, which would also be a category.

In more detail a Category consists of:

- Objects.
- Mappings (called morphisms) between those objects.
- A binary composition operator, which can combine those mappings such that a mapping from object A to object B, and a mapping from object B to object C can be combined into a single mapping from object A to object C.
- There should be a mapping that acts as an identity under composition, such that composing that mapping with another mapping M, gives you back the mapping M.
- Composition should be associative.

Anything that can be represented by the above list can be analyzed as a category, that is why it is considered so abstract.

## Notes

- The Curry Howard isomorphism. The idea that logic and programming are just two sides of the same coin: types are propositions and programs are proofs.
- [Originally, category theory traces it roots to algebraic topology. What you find is that you can associate to topological spaces (often simplicial complexes or other triangulated spaces) an algebraic object (often a group, abelian group, ring, vectorspace, or module). And the continuous maps then "lift" to the corresponding algebraic homomorphism.](https://www.reddit.com/r/math/comments/5z5tqg/what_are_you_working_on/dew677u/)
- [Look at the relationships between things, rather than the components of a thing](https://www.reddit.com/r/explainlikeimfive/comments/8rsvqw/eli5_what_is_category_theory/).

## Links

- [An invitation to category theory (2018)](http://chalkdustmagazine.com/features/an-invitation-to-category-theory/) - Just what is category theory? Tai-Danae Bradley explains.
- [What is Applied Category Theory? (2018)](https://arxiv.org/abs/1809.05923) ([HN](https://news.ycombinator.com/item?id=26011025))
- [Category Theory nLab](https://ncatlab.org/nlab/show/category+theory)
- [Applied Category Theory – Online Course](https://johncarlosbaez.wordpress.com/2018/03/26/seven-sketches-in-compositionality/)
- [Functional Programming and Category Theory [Part 1] - Categories and Functors](http://nikgrozev.com/2016/03/14/functional-programming-and-category-theory-part-1-categories-and-functors/)
- [Seven Sketches in Compositionality: An Invitation to Applied Category Theory](http://math.mit.edu/~dspivak/teaching/sp18/) ([PDF](http://math.mit.edu/~dspivak/teaching/sp18/7Sketches.pdf)) ([Paper](https://arxiv.org/abs/1803.05316))
- [Category Theory for Programmers PDF](https://github.com/hmemcpy/milewski-ctfp-pdf)
- [Category Theory for Programmers: The Preface](https://bartoszmilewski.com/2014/10/28/category-theory-for-programmers-the-preface/) ([Notes](https://github.com/rpeszek/notes-milewski-ctfp-hs))
- [Learn You Some Category Theory](https://jozefg.bitbucket.io/posts/2013-10-22-category-theory-in-haskell.html)
- [The Catsters YouTube channel](https://www.youtube.com/user/TheCatsters) - Videos on category theory.
- [ELI 5: Category Theory](https://www.reddit.com/r/explainlikeimfive/comments/8rsvqw/eli5_what_is_category_theory/)
- [A Perspective on Higher Category Theory](https://golem.ph.utexas.edu/category/2010/03/a_perspective_on_higher_catego.html)
- [Category theory foundations 1.0 course](https://www.youtube.com/watch?v=BF6kHD1DAeU&list=PLGCr8P_YncjVjwAxrifKgcQYtbZ3zuPlb)
- [Why category theory matters: a functional programmer’s perspective](https://www.researchgate.net/publication/282868977_Why_category_theory_matters_a_functional_programmer's_perspective)
- [Category Theory in Coq](https://github.com/jwiegley/category-theory) - Encodes category theory in Coq, with the primary aim being to allow representation and manipulation of categorical terms, as well realization of those terms in various target categories.
- [Formalization of category theory in cubical Agda](https://github.com/fredefox/cat)
- [Applied Category Theory Course: Databases](https://johncarlosbaez.wordpress.com/2018/06/06/applied-category-theory-course-databases/)
- [Category theory library for Agda](https://github.com/jmchapman/categories)
- [Conexus Act](https://conexus.ai/ventures) - Looking for promising ventures that want to utilize category theory to pursue the latest breakthroughs in Logistics, Pharma, Aerospace, Manufacturing, Transportation and more.
- [From design patterns to category theory (2017)](https://blog.ploeh.dk/2017/10/04/from-design-patterns-to-category-theory/) ([HN](https://news.ycombinator.com/item?id=20151055))
- [Fully formal ETCS](https://ncatlab.org/nlab/show/fully+formal+ETCS)
- [The Yoneda Perspective: an object is completely determined by its relationships to other objects (2017)](https://www.math3ma.com/blog/the-yoneda-lemma)
- [Understanding Yoneda (2013)](https://bartoszmilewski.com/2013/05/15/understanding-yoneda/)
- [Idris category theory](https://github.com/statebox/idris-ct) - Formally verified category theory library.
- [Category Theory and Lambda Calculus Thesis (2018)](https://mroman42.github.io/ctlc/ctlc.pdf)
- [Computational Category Theory](https://pdfs.semanticscholar.org/3f99/553ca06ce451c5b76479c96e191ad69f3e04.pdf)
- [A categorical semantics for causal structure (2019)](https://arxiv.org/pdf/1701.04732.pdf)
- [A categorical view of computational effects (2019)](http://www.math.jhu.edu/~eriehl/lambda.pdf)
- [Awesome Applied Category Theory](https://github.com/statebox/awesome-applied-ct)
- [Ecats](https://github.com/mbernat/ecats) - Editor for category theory.
- [Lambda World 2019 - A categorical view of computational effects - Emily Riehl](https://www.youtube.com/watch?v=Ssx2_JKpB3U)
- [Theoretical Computer Science for the Working Category Theorist](http://www.sci.brooklyn.cuny.edu/~noson/TCStext.html)
- [Notes on Applied Category Theory (2018)](https://www.math3ma.com/blog/notes-on-act)
- [Applied Category Theory MIT lectures (2019)](https://www.youtube.com/playlist?list=PLhgq-BqyZ7i5lOqOqqRiS0U5SwTmPpHQ5)
- [Categorical Logic Notes](https://github.com/awodey/CatLogNotes)
- [Programming with Categories lectures (2020)](https://www.youtube.com/watch?v=NUBEB9QlNCM)
- [Aspects of categorical recursion theory (2020)](https://arxiv.org/abs/2001.05778) ([HN](https://news.ycombinator.com/item?id=22083147))
- [Compiling to categories (2017)](http://conal.net/papers/compiling-to-categories/)
- [OCaml version of Category Theory For Programmers](https://github.com/ArulselvanMadhavan/ocaml-ctfp)
- [Recursion schemes, categorically](https://hackmd.io/@oli-kitty/recursion-schemes-categorically)
- [Todd Trimble: Geometry of regular relational calculus (2020)](https://www.youtube.com/watch?v=RonyrB0kLew)
- [Bartosz Milewski's publications](https://github.com/BartoszMilewski/Publications)
- [Computational Category Theory in Python](http://www.philipzucker.com/computational-category-theory-in-python-i-dictionaries-for-finset/) ([HN](https://news.ycombinator.com/item?id=23058551))
- [HN: Applied Category Theory (2019)(https://news.ycombinator.com/item?id=23048149)
- [Definition of a monoidal category summarised in string diagrams](https://twitter.com/NathanielVirgo/status/1262019641720832001)
- [Statebox Category Theory Course](https://training.statebox.org/)
- [Category Theory -- The math behind hyper-convergence automation (2019)](https://multix.sfo2.cdn.digitaloceanspaces.com/Category%20Theory%20TechBBQ.pdf)
- [Pittsburgh Functional Programming Book Club featuring Category Theory for Programmers](https://github.com/chiroptical/ctfp-book-club) -
- [Applied Category Theory 2020](https://act2020.mit.edu/) ([2020 Recordings](https://www.youtube.com/playlist?list=PLCOXjXDLt3pYot9VNdLlZqGajHyZUywdI)) ([2020 Tutorials](https://www.youtube.com/playlist?list=PLCOXjXDLt3pYPE63bVbsVfA41_wa3sZOh))
- [MIT Categories Seminar - Eugenia Cheng: Distributive laws for Lawvere theories (2020)](https://www.youtube.com/watch?v=YCJAzV1g6Xo)
- [AlgebraicJulia: Applied Category Theory in Julia | James Fairbanks (2020)](https://www.youtube.com/watch?v=7zr2qnud4XM)
- [Why does the "working mathematician" need category theory? (2020)](https://www.reddit.com/r/math/comments/i7erja/why_does_the_working_mathematician_need_category/)
- [F-Algebras (2017)](https://bartoszmilewski.com/2017/02/28/f-algebras/)
- [Topology: A Categorical Approach (2020)](https://topology.pubpub.org/) - Graduate-level textbook that presents basic topology from the perspective of category theory. ([Lobsters](https://lobste.rs/s/cq30tt/topology_categorical_approach))
- [MIT: Programming with Categories (2020)](http://brendanfong.com/programmingcats.html) ([HN](https://news.ycombinator.com/item?id=24353976)) ([Draft](http://brendanfong.com/programmingcats_files/cats4progs-DRAFT.pdf))
- [n-Category Cafe category theory posts](https://golem.ph.utexas.edu/category/)
- [Category Theory for the Sciences (2014)](https://mitpress.mit.edu/books/category-theory-sciences)
- [Categorical Probability and Statistics (2020)](http://perimeterinstitute.ca/personal/tfritz/2019/cps_workshop/) ([Videos](https://www.youtube.com/playlist?list=PLaILTSnVfqtIebAXFOcee9MvAyBwhIMyr))
- [Charity](http://pll.cpsc.ucalgary.ca/charity1/www/home.html) - Categorical programming language. ([Code](https://github.com/mietek/charity-lang))
- [A monoid is a category with just one object. so what's the problem? (2019)](https://k-bx.github.io/articles/boring-monoid-category.html)
- [Epidemiological Modeling With Structured Cospans (2020)](https://johncarlosbaez.wordpress.com/2020/10/19/epidemiological-modeling-with-structured-cospans/)
- [Garlandus blog](https://garlandus.co/) - Essays on some of the mathematics behind computing, from classical logic to the monads of category theory, by way of Goettingen.
- [Logical Relations as Types: Proof-Relevant Parametricity for Program Modules (2020)](http://www.jonmsterling.com/pdfs/lrat.pdf)
- [Category Theory Resources](https://github.com/prathyvsh/category-theory-resources)
- [Applied Category Theory: Notes on generative effects](https://github.com/joshvera/generative-effects)
- [Executing Categories (2020)](https://www.philipzucker.com/i-gave-a-talk-on-executing-categories/)
- [App for interactive category theory diagrams and notebooks](https://github.com/AviCraimer/category-theory-diagrams)
- [Yoneda lemma](https://ncatlab.org/nlab/show/Yoneda+lemma)
- [Category Theory introduction lectures](https://www.youtube.com/playlist?list=PLRadJIh6Wu9m60eT0tXjngv7nebUgraaA)
- [The significance of the Curry-Howard isomorphism (2019)](https://richardzach.org/2019/11/01/the-significance-of-the-curry-howard-isomorphism/)
- [Basic Category Theory for Computer Scientists](https://mitpress.mit.edu/books/basic-category-theory-computer-scientists)
- [Categorical Systems Theory Book (2020)](https://github.com/DavidJaz/DynamicalSystemsBook/blob/master/book/DynamicalBook.pdf)
- [Emily Riehl Research](http://www.math.jhu.edu/~eriehl/) - Associate professor in the department of mathematics at Johns Hopkins University working on a variety of topics in category theory related to homotopy theory.
- [David Jaz Myers](http://davidjaz.com/) - Interested in (higher) category theory, homotopy type theory, and cohesion.
- [P-Func: Math for the Mothership (2020)](https://github.com/dspivak/poly/blob/main/Book-Poly.pdf)
- [David Jaz Myers: Paradigms of composition (2020)](https://www.youtube.com/watch?v=50s62D5Ah-M)
- [ACT 2020 Recordings](https://www.youtube.com/playlist?list=PLCOXjXDLt3pYot9VNdLlZqGajHyZUywdI)
- [David Jaz Myers: Open dynamical systems, trajectories and hierarchical planning (2020)](https://www.youtube.com/watch?v=3FxeY5DbPn0)
- [Quiver](https://q.uiver.app/) - Modern commutative diagram editor. ([Code](https://github.com/varkor/quiver)) ([HN](https://news.ycombinator.com/item?id=25213201))
- [Applied Category Theory 2021](https://johncarlosbaez.wordpress.com/2021/01/02/applied-category-theory-2021-adjoint-school/)
- [Category Theory Illustrated](https://boris-marinov.github.io/category-theory-illustrated/) ([Code](https://github.com/boris-marinov/category-theory-illustrated))
- [Papers on aspects of Generalised Algebraic Theories, Contextual Categories and Mathematical Theory Of Data](https://github.com/JohnWCartmell/Theory)
- [Emily Riehl is rewriting the foundations of higher category theory (2021)](https://www.quantamagazine.org/emily-riehl-conducts-the-mathematical-orchestra-from-the-middle-20200902/) ([HN](https://news.ycombinator.com/item?id=26334535))
- [Emily Pillmore: Type Arithmetic and the Yoneda Lemma (2020)](https://www.youtube.com/watch?v=aXS5HZ_1fNQ)
- [Categorical logic from a categorical point of view](https://github.com/mikeshulman/catlog)
- [Category Theory Illustrated - Orders](https://boris-marinov.github.io/category-theory-illustrated/04_order/)
- [tikzcd-editor](https://tikzcd.yichuanshen.de/) - Simple visual editor for creating commutative diagrams. ([Code](https://github.com/yishn/tikzcd-editor))
- [The Categorical Machine: Part 1 (Applied Category Theory) (2020)](https://www.youtube.com/watch?v=hU8lG-R67Qg) - Compare Functional Programming, Lambda Calculus, and Categories (Cats).
- [Category Theoretic Approaches to Machine Learning](https://github.com/bgavran/Category_Theory_Machine_Learning)
- [Towards Categorical Foundations of Learning (2021)](https://www.brunogavranovic.com/posts/2021-03-03-Towards-Categorical-Foundations-Of-Neural-Networks.html)

## Images

![](https://i.imgur.com/4Qcz4tc.png)
