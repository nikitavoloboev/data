# [CSS](https://developer.mozilla.org/en-US/docs/Web/CSS)

[Bulma](https://bulma.io) framework seems nice. [Understanding static and relative positioning](https://webplatform.github.io/docs/tutorials/static_and_relative_positioning/) and [MDN](https://developer.mozilla.org/en-US/docs/Learn/CSS/Introduction_to_CSS) are great intros.

## Notes

- ID's are unique.
  - Each element can have only one ID.
  - Each page can have only one element with that ID.
- Classes are NOT unique.
  - You can use the same class on multiple elements.
  - You can use multiple classes on the same element.
- [The DOM is just a tree of objects. CSS selects parts of that tree and applies attributes to those objects. It's all just code, it's just a shorthand for adding a bunch of attributes to a bunch of objects.](https://www.reddit.com/r/javascript/comments/9jpvon/do_you_even_need_a_css_framework/)
- [Always design your CSS format from the inside out](https://www.quora.com/What-is-the-best-way-to-prevent-divs-from-overlapping-I-have-3-divs-The-First-div-changes-its-size-and-overlaps-the-second-one-which-is-a-set-of-images)
  - format the elements within their container so that they look correct regardless of the size of the container.
  - similarly format those containers within their own containers
  - repeat until `<body>` is the containers
  - never use absolute widths (px, in, cm, etc.) for anything
- body tag takes up the whole width and height of the browser screen.

## Links

- [Code guide by @mdo](http://codeguide.co/)
- [CSS reference](https://cssreference.io/)
- [CSS Blocks](https://github.com/linkedin/css-blocks) - High performance, maintainable stylesheets.
- [Opticss](https://github.com/linkedin/opticss) - CSS Optimizer.
- [CSS Puns](https://saijogeorge.com/css-puns/)
- [Magic of CSS](https://adamschwartz.co/magic-of-css/)
- [Min](https://github.com/owenversteeg/min) - World's smallest (995 bytes) CSS framework.
- [Emotion](https://emotion.sh/) - Performant and flexible CSS-in-JS library.
- [Linaria](https://github.com/callstack/linaria#readme) - Zero-runtime CSS in JS library.
- [Why We Use Styled Components at Decisiv](https://medium.com/@_alanbsmith/why-we-use-styled-components-at-decisiv-a8ac6e1507ac)
- [The case for CSS modules - Mark Dalgleish](https://www.youtube.com/watch?v=zR1lOuyQEt8)
- [Michael Chan - Inline Styles: themes, media queries, contexts, & when it's best to use CSS (2015)](https://www.youtube.com/watch?v=ERB1TJBn32c)
- [Understanding the CSS box model for inline elements](https://hacks.mozilla.org/2015/03/understanding-inline-box-model/)
- [astroturf](https://github.com/4Catalyzer/astroturf) - An "artificial" css-in-js for those that want it all.