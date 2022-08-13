---
title: HTML
---

# [HTML](https://developer.mozilla.org/kab/docs/Web/HTML)

## Notes

- As a general rule, if there will be only one such element on the page, you can use an id. Otherwise, use a class.
- [`<form method=post\>` is good to communicate a mutation.](https://twitter.com/ryanflorence/status/1370403183822921731)
- [When I work on a website, I set the body to contenteditable="true". This allows you to edit all text on the page, which I find helpful to improve the content, and also enables spell checking.](https://www.reddit.com/r/programming/comments/tzmqua/those_html_attributes_you_never_use/)
- `span is an inline div, it doesn't go to a new line`.
- `use span for text, div for everything else (buttons, lines etc.)`
- [HTML file inputs can access photos, videos, and audio captured directly by the device's camera](https://twitter.com/housecor/status/1529437325792628736)

## Code

```html
<!-- Embed email inside link -->
<a href="mailto:name@gmail.com">Email me</a>
```

```html
<!-- Add image to the side with link -->
<!-- Useful in GitHub readmes. Can adjust width or add height. -->
<img
  src="https://i.imgur.com/hZe2AUY.png"
  align="right"
  width="70"
/>](https://nodejs.org)
```

```html
<!-- Breaking space. Can use in readme to separate things. -->
&nbsp;
```

```html
<!-- Collapsable content -->
<details>
  <summary>Table of Contents</summary>

  * [About](#about) * [Screenshots](#screenshots) * ..
</details>
```

```html
<!-- line break -->
<br />
```

## Examples

```html
<!-- to see difference of span with inline -->
<p>
Hi<span> I'm inline</span>
</p>
<p>
Hi<div>oh no break</div>
</p>
```

## Links

- [HTML reference](https://htmlreference.io/)
- [HTML Standard](https://html.spec.whatwg.org/multipage/) ([Twitter](https://twitter.com/htmlstandard))
- [DOM Enlightenment](http://domenlightenment.com/) - Exploring the relationship between JavaScript and the modern HTML DOM.
- [Virtual DOM is pure overhead (2018)](https://svelte.dev/blog/virtual-dom-is-pure-overhead) ([HN](https://news.ycombinator.com/item?id=19950253)) ([HN](https://news.ycombinator.com/item?id=27675371))
- [Folding the DOM (2019)](https://www.joshwcomeau.com/posts/folding-the-dom/)
- [About HTML semantics and front-end architecture (2012)](http://nicolasgallagher.com/about-html-semantics-front-end-architecture/)
- [Simple HTML form for your website](https://www.staticforms.xyz/)
- [HEAD](https://htmlhead.dev/) - List of everything that _could_ go in the head of your document.
- [HTML Living Standard](https://html.spec.whatwg.org/) ([Multipage version](https://html.spec.whatwg.org/multipage/)) ([Code](https://github.com/whatwg/html))
- [Learn to Code HTML & CSS book](https://learn.shayhowe.com/html-css/)
- [Native HTML Autocomplete with datalist](https://codepen.io/tejask/pen/OJJBLrq)
- [HTML isn’t done! (Chrome Dev Summit 2019)](https://www.youtube.com/watch?v=ZFvPLrKZywA)
- [A History of HTML Parsing at Cloudflare (2019)](https://blog.cloudflare.com/html-parsing-1/)
- [html-tags](https://github.com/sindresorhus/html-tags) - List of standard HTML tags.
- [hast](https://github.com/syntax-tree/hast) - Hypertext Abstract Syntax Tree format.
- [Optimal Virtual DOM (2019)](https://blog.kabir.sh/posts/optimal-virtual-dom.html)
- [tags-input](https://github.com/developit/tags-input) - Input element with tag input.
- [HTML attributes to improve your users' two factor authentication experience (2020)](https://www.twilio.com/blog/html-attributes-two-factor-authentication-autocomplete) ([HN](https://news.ycombinator.com/item?id=22022106))
- [Native HTML Elements with CSS superpowers](https://github.com/equinusocio/native-elements)
- [A Complete Guide to Links and Buttons (2020)](https://css-tricks.com/a-complete-guide-to-links-and-buttons/)
- [pup](https://github.com/ericchiang/pup) - Parsing HTML at the command line.
- [How to Build HTML Forms Right: Semantics (2020)](https://stegosource.com/how-to-build-html-forms-right-semantics/)
- [sal](https://github.com/mciastek/sal) - Performance focused, lightweight scroll animation library.
- [MVP.css](https://andybrewer.github.io/mvp/) - Minimalist stylesheet for HTML elements. No class names, no frameworks, just semantic HTML. ([Code](https://github.com/andybrewer/mvp))
- [HTML DOM](https://htmldom.dev/) - Common tasks of managing HTML DOM with vanilla JavaScript. ([Code](https://github.com/phuoc-ng/html-dom)) ([HN](https://news.ycombinator.com/item?id=22758218))
- [hyper(HTML)](https://github.com/WebReflection/hyperHTML) - Fast & Light Virtual DOM Alternative available for Node.js and NativeScript too.
- [Open Graph Image as a Service](https://github.com/vercel/og-image) - Serverless service that generates dynamic Open Graph images that you can embed in your <meta\> tags.
- [Responsive Images the Simple Way (2020)](https://cloudfour.com/thinks/responsive-images-the-simple-way/)
- [sahtml-query](https://github.com/mdevils/sahtml-query) - Fast and low memory query API for HTML (node.js).
- [HTML is more complicated than you think (2020)](https://www.tempertemper.net/blog/html-is-more-complicated-than-you-think)
- [X GIF](https://github.com/geelen/x-gif) - Custom element for flexible GIF playback.
- [HTMLHint](https://htmlhint.com/) - Static code analysis tool you need for your HTML. ([Code](https://github.com/htmlhint/HTMLHint))
- [Can I Include](https://caninclude.glitch.me/) - Understand which tag you can include in another.
- [htmx](https://htmx.org/) - High power tools for HTML. ([Code](https://github.com/bigskysoftware/htmx)) ([HN](https://news.ycombinator.com/item?id=23330881)) ([Lobsters](https://lobste.rs/s/cgtqit/htmx_high_power_tools_for_html)) ([Lobsters 2](https://lobste.rs/s/sk9cvw/htmx_high_power_tools_for_html))
- [What if we'd had better html-in-js syntax all along? (2020)](https://leontrolski.github.io/dom-syntax.html) ([Lobsters](https://lobste.rs/s/9im5wq/what_if_we_d_had_better_html_js_syntax_all)) ([HN](https://news.ycombinator.com/item?id=23142300))
- [PostHTML](https://github.com/posthtml/posthtml) - Tool for transforming HTML/XML with JS plugins.
- [All Visible HTML Elements](https://elements.xz.style/) - Simple list of all visible and style-able HTML elements. The list itself does not style any of these elements. ([Code](https://github.com/xz/elements))
- [pdf2htmlEX](https://github.com/pdf2htmlEX/pdf2htmlEX/) - Convert PDF to HTML without losing text or format.
- [htmlparser2](https://github.com/fb55/htmlparser2) - Forgiving html and xml parser.
- [Awesome lit-html](https://github.com/web-padawan/awesome-lit-html)
- [Grid.js](https://gridjs.io/) - Free and open-source HTML table plugin written in TypeScript. ([HN](https://news.ycombinator.com/item?id=23420091)) ([Code](https://github.com/grid-js/gridjs)) ([Grid.js Website code](https://github.com/grid-js/website))
- [keen-slider](https://github.com/rcbyr/keen-slider) - HTML touch slider carousel with the most native feeling.
- [HTML Tutorial for Beginners 101 (Including HTML5 Tags) (2020)](https://websitesetup.org/html-tutorial-beginners/)
- [ElaAdmin HTML5 Admin Dashboard Template](https://github.com/puikinsh/ElaAdmin)
- [HTML5 Boilerplate](https://html5boilerplate.com/) - Professional front-end template for building fast, robust, and adaptable web apps or sites. ([Code](https://github.com/h5bp/html5-boilerplate))
- [H5BP projects](https://h5bp.org/)
- [HTML Road Guide](https://lyty.dev/html/index.html)
- [Why Forking HTML Into A Static Language Doesn't Make Sense (2020)](https://robert.ocallahan.org/2020/05/why-forking-html-into-static-language.html) ([Lobsters](https://lobste.rs/s/zztnen/why_forking_html_into_static_language))
- [Building GitHub-style Hovercards with Stimulus and HTML-over-the-wire (2020)](https://boringrails.com/articles/hovercards-stimulus/)
- [Slingcode](https://slingcode.net/) - Personal computing platform in a single html file.
- [ThebeLab](https://github.com/minrk/thebelab) - Turns your static HTML pages into interactive ones, powered by a kernel. ([Docs](https://thebelab.readthedocs.io/en/latest/))
- [HTML+CSS Tutorial](https://github.com/cassidoo/HTML-CSS-Tutorial)
- [drag-drop](https://github.com/feross/drag-drop) - HTML5 drag & drop for humans.
- [HTML APIs: What They Are And How To Design A Good One (2017)](https://www.smashingmagazine.com/2017/02/designing-html-apis/)
- [Creating a sitemap.xml from a pile of HTML files](https://www.christopherbiscardi.com/creating-a-sitemap-xml-from-a-pile-of-html-files)
- [DOM Expressions](https://github.com/ryansolid/dom-expressions) - Fine-Grained Runtime for Performant DOM Rendering.
- [Form Design](https://gerireid.com/forms.html)
- [xm](https://github.com/giuseppeg/xm) - Extensible HTML. ([Docs](https://giuseppegurgone.com/xm/))
- [DOMPurify bypass: XSS via HTML namespace confusion (2020)](https://research.securitum.com/mutation-xss-via-mathml-mutation-dompurify-2-0-17-bypass/) ([HN](https://news.ycombinator.com/item?id=24703230))
- [The usability of HTML elements (2020)](https://shkspr.mobi/blog/2020/10/the-usability-of-html-elements/) ([HN](https://news.ycombinator.com/item?id=24729006))
- [Awfice](https://zserge.com/posts/awfice/) - Collection of tiny office suite apps. ([Article](https://zserge.com/posts/awfice/)) ([HN](https://news.ycombinator.com/item?id=24752546))
- [croquis.js](https://github.com/disjukr/croquis.js) - HTML5 drawing tool library.
- [Emmet](https://github.com/emmetio/emmet) - Web-developer’s toolkit for boosting HTML & CSS code writing. ([Web](https://emmet.io/))
- [This page is a truly naked, brutalist HTML quine](https://secretgeek.github.io/html_wysiwyg/html.html) ([HN](https://news.ycombinator.com/item?id=24824977))
- [Konva](https://github.com/konvajs/konva) - HTML5 Canvas JavaScript framework that extends the 2d context by enabling canvas interactivity for desktop and mobile applications. ([Web](https://konvajs.org/))
- [Typed HTML](https://github.com/nicojs/typed-html) - TypeSafe HTML templates using TypeScript. No need to learn a template library.
- [he](https://github.com/mathiasbynens/he) - Robust HTML entity encoder/decoder written in JavaScript.
- [bluemonday](https://github.com/microcosm-cc/bluemonday) - HTML sanitizer implemented in Go. It is fast and highly configurable.
- [vanilla-colorful](https://github.com/web-padawan/vanilla-colorful) - Tiny color picker custom element for modern web apps.
- [HTML forms, Databases, Integration tests (2020)](https://www.lpalmieri.com/posts/2020-08-31-zero-to-production-3-5-html-forms-databases-integration-tests/)
- [Tocbot](https://github.com/tscanlin/tocbot) - Build a table of contents from headings in an HTML document. ([Web](http://tscanlin.github.io/tocbot/))
- [async-htm-to-string](https://github.com/voxpelli/async-htm-to-string) - Renders a htm tagged template asyncly into a string.
- [HTML and CSS simple tips and tricks for your website](https://thomasorus.com/html-tips.html)
- [Headers.CSS](https://headers-css.vercel.app/) - Blueprint HTML and CSS for 17+ website headers. ([Code](https://github.com/shadeed/headers-css))
- [minify-html](https://github.com/wilsonzlin/minify-html) - Fast and smart HTML + JS minifier, available for Rust, Node.js, Python, Java, and Ruby.
- [HTML & CSS Is Hard | Friendly web development tutorial](https://www.internetingishard.com/html-and-css/)
- [imperative-html](https://github.com/johnnesky/imperative-html) - Small JavaScript library for creating HTML (and SVG) elements in a web browser.
- [lit-html-server](https://github.com/popeindustries/lit-html-server) - Render lit-html templates on the server as strings or streams (and in the browser too).
- [hypertag](https://github.com/AndreasPizsa/hypertag) - HTML tag parser built for speed.
- [DomStorm](https://domstorm.skepticfx.com/) - Dashboard for interesting DOM tricks/techniques. ([Code](https://github.com/skepticfx/domstorm))
- [Alt vs Figcaption (2020)](https://thoughtbot.com/blog/alt-vs-figcaption)
- [Ultralight](https://ultralig.ht/) - Next-Generation HTML Renderer for Desktop Apps and Games. ([Code](https://github.com/ultralight-ux/Ultralight))
- [The Zen of index.html (2020)](https://hugodaniel.com/posts/using-just-an-index-to-develop-a-web-app/)
- [PostHTML](https://github.com/posthtml/posthtml-expressions) - Use variables, JS-like expressions, and even markup-powered logic in your HTML. ([Web](https://posthtml.org/#/))
- [Hotwire](https://hotwire.dev/) - Alternative approach to building modern web applications without using much JavaScript by sending HTML instead of JSON over the wire. ([Tweet](https://twitter.com/dhh/status/1341420143239450624)) ([HN](https://news.ycombinator.com/item?id=25507942)) ([Lobsters](https://lobste.rs/s/nyxq4o/hotwire_html_over_wire)) ([Article](https://delitescere.medium.com/hotwire-html-over-the-wire-2c733487268c)) ([Why Hotwire Could be the Future of Front-end Dev](https://wizville.fr/en/blog/a-product-guys-take-on-hotwire-vs-vue-react/)) ([The time is right for Hotwire](https://world.hey.com/dhh/the-time-is-right-for-hotwire-ecdb9b33)) ([Lobsters](https://lobste.rs/s/lifxol/time_is_right_for_hotwire)) ([Hotwire in Action](https://github.com/asyraffff/Hotwire-in-action)) ([Reddit](https://www.reddit.com/r/rails/comments/sjrh1g/opinions_on_hotwire/)) ([Hotwire Newsletter](https://masilotti.com/hotwire/))
- [entities](https://github.com/fb55/entities) - Encode & decode HTML & XML entities with ease & speed.
- [HTML5 still doesn't replicate what mattered about Flash (2020)](https://twitter.com/larsiusprime/status/1344404336252768257) ([HN](https://news.ycombinator.com/item?id=25587765))
- [The unreasonable effectiveness of simple HTML (2021)](https://shkspr.mobi/blog/2021/01/the-unreasonable-effectiveness-of-simple-html/) ([HN](https://news.ycombinator.com/item?id=25915313))
- [Managing focus in the shadow DOM (2021)](https://nolanlawson.com/2021/02/13/managing-focus-in-the-shadow-dom/)
- [rehype](https://github.com/rehypejs/rehype) - HTML processor powered by plugins part of the unified collective. ([Awesome](https://github.com/rehypejs/awesome-rehype))
- [linkedom](https://github.com/WebReflection/linkedom) - Triple-linked lists based DOM implementation.
- [bin2png](https://github.com/lovasoa/bin2png) - Embed binary data inside an HTML file in an efficient way.
- [litehtml](https://github.com/litehtml/litehtml) - Fast and lightweight HTML/CSS rendering engine. ([Web](http://www.litehtml.com/))
- [Gorillas’ nav: a case study (2021)](https://kittygiraudel.com/2021/03/13/gorillas-nav-a-case-study/)
- [Under-Engineered Select Menus (2021)](https://adrianroselli.com/2021/03/under-engineered-select-menus.html)
- [Parsed HTML Rewriter](https://github.com/worker-tools/parsed-html-rewriter) - DOM-based implementation of Cloudflare Worker's HTMLRewriter.
- [htmlq](https://github.com/mgdm/htmlq) - Like jq, but for HTML. Uses CSS selectors to extract bits content from HTML files. ([HN](https://news.ycombinator.com/item?id=28441880)) ([Lobsters](https://lobste.rs/s/9ljgez/htmlq_like_jq_for_html))
- [Beautiful PDFs from HTML](https://pdf.math.dev/) ([Code](https://github.com/ashok-khanna/pdf)) ([HN](https://news.ycombinator.com/item?id=26691626))
- [SelectMadu](https://github.com/pavish/select-madu) - Replacement for the select menu, with support for searching, multiple selections, async data loading and more.
- [Sortable Table Columns (2021)](https://adrianroselli.com/2021/04/sortable-table-columns.html)
- [My current HTML Boilerplate (2021)](https://www.matuzo.at/blog/html-boilerplate/) - Summarized which elements I use for the basic structure of a HTML document in 2021, with explanations why. ([HN](https://news.ycombinator.com/item?id=26952557))
- [Ammonia](https://github.com/rust-ammonia/ammonia) - Repair and secure untrusted HTML.
- [Modeling and Reasoning about DOM Events](https://world.cs.brown.edu/~sk/Publications/Papers/Published/lckqk-model-reason-dom-events/paper.pdf)
- [html-parsed-element](https://github.com/WebReflection/html-parsed-element) - Base custom element class with a reliable `parsedCallback` method.
- [focus-trap](https://github.com/focus-trap/focus-trap) - Trap focus within a DOM node.
- [Incremental DOM](https://github.com/google/incremental-dom) - Library for building up DOM trees and updating them in-place when data changes.
- [The Humble <img\> Element And Core Web Vitals (2021)](https://www.smashingmagazine.com/2021/04/humble-img-element-core-web-vitals/)
- [Update Canvas 2D API](https://github.com/fserb/canvas2D)
- [Google Docs will now use canvas based rendering (2021)](https://workspaceupdates.googleblog.com/2021/05/Google-Docs-Canvas-Based-Rendering-Update.html) ([HN](https://news.ycombinator.com/item?id=27129858)) ([Lobsters](https://lobste.rs/s/uqb3kj/google_docs_will_now_use_canvas_based))
- [DOM Events](https://domevents.dev/) - Learn about the DOM Event system through exploration.
- [The Button Cheat Sheet](https://www.buttoncheatsheet.com/) ([HN](https://news.ycombinator.com/item?id=27276547))
- [HTML Parser Book](https://htmlparser.info/) - Idiosyncrasies of the HTML parser. ([Code](https://github.com/zcorpan/html-parser-book)) ([HN](https://news.ycombinator.com/item?id=27311627))
- [Declarative Shadow DOM](https://github.com/mfreed7/declarative-shadow-dom)
- [Your Ultimate Guide to Understanding DOM Events](https://egghead.io/courses/the-ultimate-guide-for-understanding-dom-events-6c0c0d23)
- [Stampino](https://github.com/justinfagnani/stampino) - Fast and extremely powerful HTML template system.
- [Million](https://github.com/aidenybai/million) - Lightweight compiler-augmented Virtual DOM. ([HN](https://news.ycombinator.com/item?id=30935188)) ([Web](https://millionjs.org/)) ([HN](https://news.ycombinator.com/item?id=32291822))
- [The right tag for the job: why you should use semantic HTML (2021)](https://localghost.dev/2021/06/the-right-tag-for-the-job-why-you-should-use-semantic-html/)
- [html5parser](https://github.com/acrazing/html5parser) - Super fast and tiny HTML5 parser.
- [templ](https://github.com/a-h/templ) - Strongly typed HTML templating language that compiles to Go code, and has great developer tooling.
- [Capsid](https://github.com/capsidjs/capsid) - Declarative DOM programming library based on TypeScript decorators.
- [html5ever](https://github.com/servo/html5ever) - High-performance browser-grade HTML5 parser.
- [HTML Recipes](https://htmlrecipes.dev/) - Collection of quick copy HTML snippets for a variety of common scenarios. ([Code](https://github.com/5t3ph/htmlrecipes))
- [xPath Finder](https://github.com/trembacz/xpath-finder) - Click on any element to get the xPath.
- [Marko](https://github.com/marko-js/marko) - Declarative, HTML-based language that makes building web apps fun. ([Web](https://markojs.com/)) ([Marko Tags](https://dev.to/ryansolid/introducing-the-marko-tags-api-preview-37o4))
- [Why are there no PUT and DELETE methods on HTML forms?](https://softwareengineering.stackexchange.com/questions/114156/why-are-there-no-put-and-delete-methods-on-html-forms) ([Lobsters](https://lobste.rs/s/an3e64/why_are_there_no_put_delete_methods_on_html))
- [Improving responsiveness in text inputs (2021)](https://nolanlawson.com/2021/08/08/improving-responsiveness-in-text-inputs/)
- [Does shadow DOM improve style performance? (2021)](https://nolanlawson.com/2021/08/15/does-shadow-dom-improve-style-performance/)
- [Don't attach tooltips to document.body](https://atfzl.com/don-t-attach-tooltips-to-document-body) ([HN](https://news.ycombinator.com/item?id=28230977))
- [Accessible autocomplete](https://github.com/alphagov/accessible-autocomplete) - Autocomplete component, built to be accessible.
- [DOM Testing Library](https://github.com/testing-library/dom-testing-library) - Simple and complete DOM testing utilities that encourage good testing practices. ([Docs](https://testing-library.com/docs/)) ([Discord](https://discord.gg/testing-library))
- [On the <dl\> (2021)](https://benmyers.dev/blog/on-the-dl/)
- [LazyHTML](https://github.com/cloudflare/lazyhtml) - HTML5-compliant parser and serializer than enables building transformation pipeline in a pluggable manner.
- [tiny-vdom](https://github.com/aidenybai/tiny-vdom) - Smallest possible virtual DOM implementation.
- [7GUIs in Vanilla HTML, CSS, JavaScript](https://7guis.bradwoods.io/) ([HN](https://news.ycombinator.com/item?id=28600804))
- [ct.css](https://github.com/csswizardry/ct) - Diagnostic CSS snippet that exposes potential performance issues in your page’s <head\> tags. ([Web](https://csswizardry.com/ct/))
- [Get Your Head Straight Talk](https://speakerdeck.com/csswizardry/get-your-head-straight)
- [html-to-image](https://github.com/bubkoo/html-to-image) - Generates an image from a DOM node using HTML5 canvas and SVG.
- [Safe DOM manipulation with the Sanitizer API (2021)](https://web.dev/sanitizer/)
- [HTML Sanitizer API](https://wicg.github.io/sanitizer-api/) ([Code](https://github.com/WICG/sanitizer-api)) ([Article](https://portswigger.net/daily-swig/google-mozilla-close-to-finalizing-sanitizer-api-for-chrome-and-firefox-browsers))
- [HTML Segmentator](https://github.com/fabiospampinato/html-segmentator) - Small library for splitting an HTML string into its top-level sections. Based on htmlparser2.
- [Embetter](https://github.com/cacheflowe/embetter) - Embed 3rd-party media with lazy-loaded iframes.
- [Inline Everything Cheat Sheet](https://cacheflowe.github.io/inline-everything/) - Examples for inlining different file types. ([Code](https://github.com/cacheflowe/inline-everything))
- [HTML Compendium](https://github.com/xdesro/html-compendium) - Omnibus HTML file you can test your CSS cascade and formatting with.
- [writable-dom ](https://github.com/marko-js/writable-dom) - Utility to stream HTML content into a live document.
- [Tasty HTML buttons](https://www.htmhell.dev/26-tasty-buttons/)
- [HTMHell](https://www.htmhell.dev/) - Collection of bad practices in HTML, copied from real websites.
- [HTML Tips & Tricks](https://www.htmhell.dev/tips/)
- [The HTML video element needs to go back on the drawing board (2021)](https://www.ctrl.blog/entry/html-responsive-video.html) ([HN](https://news.ycombinator.com/item?id=29024868))
- [Percollate](https://github.com/danburzo/percollate) - Command-line tool that turns web pages into beautifully formatted PDF, EPUB, or HTML files.
- [In Defence Of Dialog (2021)](https://whistlr.info/2021/in-defence-of-dialog/)
- [python](https://github.com/byteface/domonic) - Generate HTML with python 3.
- [rehype-prism-plus](https://github.com/timlrx/rehype-prism-plus) - Rehype plugin to highlight code blocks in HTML with Prism (via refractor) with line highlighting and line numbers.
- [A Clever Sticky Footer Technique (2021)](https://css-tricks.com/a-clever-sticky-footer-technique/)
- [Canvas Txt](https://github.com/geongeorge/Canvas-Txt) - Better way to render text on HTML5 Canvas.
- [sanitize-html](https://github.com/apostrophecms/sanitize-html) - Simple HTML sanitizer with a clear API.
- [Cross-fading any two DOM elements is currently impossible (2021)](https://jakearchibald.com/2021/dom-cross-fade/) ([Tweet](https://twitter.com/jaffathecake/status/1462803936679763971))
- [html5gum](https://github.com/untitaker/html5gum) - WHATWG-compliant HTML tokenizer in Rust.
- [Super OG Images](https://superblog.ai/og) - API to create opengraph images programmatically. ([Code](https://github.com/s-kris/super-og-images))
- [hq](https://github.com/mitsuhiko/hq) - jq but for HTML.
- [mailgen](https://github.com/eladnava/mailgen) - Node.js package that generates clean, responsive HTML e-mails for sending transactional mail.
- [domutils](https://github.com/fb55/domutils) - Utilities for working with htmlparser2's DOM.
- [carota](https://github.com/danielearwicker/carota) - Simple, flexible rich text rendering/editing on HTML Canvas.
- [html2text](https://github.com/jugglerchris/rust-html2text) - Rust library to render HTML as text.
- [xj](https://idiomdrottning.org/xj) - HTML to JSON. ([HN](https://lobste.rs/s/dmyqry/xj_html_json))
- [parse5](https://github.com/inikulin/parse5) - HTML parsing/serialization toolset for Node.js.
- [Entire website in a single HTML file](https://css-tricks.com/a-whole-website-in-a-single-html-file/) ([HN](https://news.ycombinator.com/item?id=29668260))
- [HTML Standard FAQ](https://github.com/whatwg/html/blob/main/FAQ.md)
- [Div Divisiveness (2022)](https://www.scottohara.me/blog/2022/01/20/divisive.html) ([HN](https://news.ycombinator.com/item?id=30047311))
- [Hibiki HTML](https://github.com/dashborg/hibiki) - Powerful new web framework for creating modern, dynamic, frontend applications without JavaScript, that can be fully scripted and controlled by backend code. ([HN](https://news.ycombinator.com/item?id=30103906))
- [DOM is not slow. Misconception that the DOM is slow comes from misuse.](https://twitter.com/dannymoerkerke/status/1488528544909115392)
- [Introducing the Dialog Element (2022)](https://webkit.org/blog/12209/introducing-the-dialog-element/)
- [PageCrypt](https://github.com/MaxLaumeister/PageCrypt) - Client-side password-protection for HTML.
- [details-utils](https://github.com/zachleat/details-utils) - Suite of utilities to add more features to the <details\> element.
- [Replace JavaScript Dialogs With New HTML Dialog (2022)](https://css-tricks.com/replace-javascript-dialogs-html-dialog-element/)
- [rasterizeHTML.js](https://github.com/cburgmer/rasterizeHTML.js) - Renders HTML into the browser's canvas.
- [tl](https://github.com/y21/tl) - Fast, zero-copy HTML Parser written in Rust.
- [AcadHomepage](https://github.com/RayeRen/acad-homepage.github.io) - Modern and Responsive Academic Personal Homepage.
- [Pikaso](https://github.com/pikasojs/pikaso) - Seamless and headless HTML5 Canvas library.
- [metatags](https://github.com/microlinkhq/metatags) - Ensure your HTML is previewed beautifully across social networks.
- [Explain the first 10 lines of Twitter's source code to me (2022)](https://css-tricks.com/explain-the-first-10-lines-of-twitter-source-code/) ([HN](https://news.ycombinator.com/item?id=30468793))
- [Obscure HTML Tags](https://devapt.com/html-tags) ([HN](https://news.ycombinator.com/item?id=30469695))
- [It's always been you, Canvas2D (2022)](https://developer.chrome.com/blog/canvas2d/) ([HN](https://news.ycombinator.com/item?id=30553585))
- [html5lib](https://github.com/html5lib/html5lib-python) - Pure-python library for parsing HTML.
- [HTML-friendly spaCy Tokenizer](https://github.com/pmbaumgartner/spacy-html-tokenizer)
- [Building a loading bar component (2022)](https://web.dev/building-a-loading-bar-component/)
- [Hotwire Example Template](https://github.com/thoughtbot/hotwire-example-template) - Collection of branches that transmit HTML over the wire.
- [djLint](https://github.com/Riverside-Healthcare/djLint) - Find common formatting issues and reformat HTML templates.
- [worker-tools/html](https://github.com/worker-tools/html) - HTML templating and streaming response library for Service Worker-like environments such as Cloudflare Workers.
- [HTML Modules Explainer](https://github.com/WICG/webcomponents/blob/gh-pages/proposals/html-modules-explainer.md)
- [Fast HTML Parser](https://github.com/taoqf/node-html-parser) - Very fast HTML parser, generating a simplified DOM, with basic element query support.
- [html2runes](https://github.com/spacecowboy/html2runes) - HTML to Text converter program written in Rust.
- [One way to archive web pages](https://twitter.com/masukomi/status/1507460071265280001)
- [Hyntax](https://github.com/mykolaharmash/hyntax) - Straightforward HTML parser for JavaScript.
- [Those HTML Attributes You Never Use (2022)](https://www.smashingmagazine.com/2022/03/html-attributes-you-never-use/) ([HN](https://news.ycombinator.com/item?id=30887445)) ([Reddit](https://www.reddit.com/r/programming/comments/tzmqua/those_html_attributes_you_never_use/))
- [Picture perfect images with the modern img element (2022)](https://stackoverflow.blog/2022/03/28/picture-perfect-images-with-the-modern-element/)
- [How To Build A Progressively Enhanced, Accessible, Filterable And Paginated List (2022)](https://www.smashingmagazine.com/2022/04/accessible-filterable-paginated-list-11ty-alpinejs/)
- [Designing Better Breadcrumbs (2022)](https://www.smashingmagazine.com/2022/04/designing-better-breadcrumbs/)
- [HTML to Text](https://github.com/html-to-text/node-html-to-text) - Advanced converter that parses HTML and returns beautiful text.
- [Why the GOV.UK Design System team changed the input type for numbers (2022)](https://technology.blog.gov.uk/2020/02/24/why-the-gov-uk-design-system-team-changed-the-input-type-for-numbers/)
- [HTML Rewriter](https://github.com/worker-tools/html-rewriter) - WASM-based implementation of Cloudflare's HTML Rewriter for use in Deno, browsers, etc.
- [HTMLJS-Parser](https://github.com/marko-js/htmljs-parser) - HTML parser recognizes content and string placeholders and allows JavaScript expressions as attribute values.
- [Open Graph Image as a Service](https://dev.zaubrik.com/portrait/) ([Code](https://github.com/Zaubrik/portrait))
- [jtml](https://github.com/github/jtml) - Write HTML in JavaScript, using template-tags.
- [Broken Link Checker](https://github.com/stevenvachon/broken-link-checker) - Find broken links, missing images, etc within your HTML.
- [Under-Engineered Multi-Selects (2022)](https://adrianroselli.com/2022/05/under-engineered-multi-selects.html)
- [Progressive Enhancement and HTML Forms: use FormData (2022)](https://www.bram.us/2022/04/22/progressive-enhancement-and-html-forms-use-formdata/)
- [Lucid](https://github.com/chrisdone/lucid) - Clear to write, read and edit DSL for writing HTML.
- [HTML5 Parser](https://github.com/kovidgoyal/html5-parser) - Fast C based HTML 5 parsing for python.
- [hast-util-from-html](https://github.com/syntax-tree/hast-util-from-html) - hast utility that turns HTML into a syntax tree.
- [Media Element Syncer](https://github.com/tjenkinson/media-element-syncer) - Synchronize two or more HTML5 media elements.
- [tabbable](https://github.com/focus-trap/tabbable) - Find descendants of a DOM node that are in the tab order.
- [Two levels of customising \<selectmenu\> (2022)](https://hidde.blog/custom-select-with-selectmenu/)
- [Muffet](https://github.com/raviqqe/muffet) - Fast website link checker in Go.
- [Avoiding \<img\> layout shifts: aspect-ratio vs width & height attributes](https://jakearchibald.com/2022/img-aspect-ratio/)
- [Able Player](https://github.com/ableplayer/ableplayer) - Fully accessible cross-browser HTML5 media player.
- [Tabler](https://github.com/tabler/tabler) - Free and open-source HTML Dashboard UI Kit built on Bootstrap. ([HN](https://news.ycombinator.com/item?id=32278397))
- [HTML Parser](https://github.com/mathiversen/html-parser) - Simple and general purpose html/xhtml parser lib/bin, using Pest.
