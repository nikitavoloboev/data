---
title: JavaScript
---

# [JavaScript](https://developer.mozilla.org/bm/docs/Web/JavaScript)

[Modern JavaScript Tutorial](https://javascript.info/) is great.

## Notes

- [You can define and run the function straight away by enclosing it in parenthesis and adding empty () after](https://forum.keyboardmaestro.com/t/javascript-assistance/8331/2)
- In JavaScript, a type is a runtime tag describing the actual kind of data you have at runtime.
- A property is any value attached to an object. A method is a property that is a function.
- `this`, always references the owner of the function it is in, for this case — since it is now out of scope — the window/global object.
- [Can type `debugger` in JS/TS to activate debugger at that point](https://twitter.com/calebporzio/status/1151185995309690886).
- Can `console.log({ myConst })` to log the variable with the value as an object so you see the variable name & the value. As opposed to doing `console.log(myConst, 'myConst')`.
- [Periodic reminder that `await import('//dev.jspm.io/[pkg]')` allows you to import anything from npm in the browser instantly. I often use this in the console to test things out.](https://twitter.com/guybedford/status/1202022281633030145)
- [Between typescript and prettier, there is very little justification for eslint.](https://twitter.com/mweststrate/status/1296437898330419209)
- [JS proxy can observe a nested object for changes made to it. That means there's none of the traditional selecting or diffing.](https://twitter.com/0xca0a/status/1328756186288713730)
- [One optimization that turned out well: storing JavaScript strings as either UTF-16 or UTF-8, depending on whats inside. It turns out this is also what JavaScript engines do, so it makes converting from a string in a JavaScript VM to a string in Bun's AST fast.](https://twitter.com/jarredsumner/status/1441876861106024449)
- [Stop using strings concatenation to create and edit URLs in JS. All major browsers have great built-in functionality for URL parsing & manipulation.](https://twitter.com/Steve8708/status/1509280798423191554)
- [Use arrow functions for 1-line functions only, otherwise named functions are always the best choice: better stack trace, better tracing and more readable.](https://twitter.com/ManuEomm/status/1523222905437319168)

## Code

```js
// Click on buttons in a page
// https://twitter.com/brian_lovin/status/1240662440666222597

let buttons = document.getElementsByClassName("unfollow");

for (let [i, v] of [...buttons].entries()) {
  setTimeout(() => {
    buttons[i].click();
  }, i * 500);
}
```

```js
// Go to specific URL
window.location.href = "https://www.google.com";
```

```js
// Speed up video playback. Put in console of open page.
$("video").playbackRate = 2;
```

```js
// click on all the expand buttons on https://github.com/
var inputs = document.querySelectorAll("svg[aria-label=Expand]");
var buttons = Array.from(inputs).map((e) => e.parentElement.parentElement);
buttons.forEach((b) => b.click());
```

```js
// start/close PiP video from currently open tab
var video = document.querySelectorAll("video")[0];
if (
  video !== undefined &&
  video.webkitSupportsPresentationMode &&
  typeof video.webkitSetPresentationMode === "function"
) {
  video.webkitSetPresentationMode(
    video.webkitPresentationMode === "picture-in-picture"
      ? "inline"
      : "picture-in-picture"
  );
}
```

## Links

- [JavaScript. The Core](http://dmitrysoshnikov.com/ecmascript/javascript-the-core-2nd-edition/)
- [Modern JavaScript Tutorial](https://javascript.info/) ([HN](https://news.ycombinator.com/item?id=25333350)) ([HN](https://news.ycombinator.com/item?id=31390742))
- [You don't know JS books](https://github.com/getify/You-Dont-Know-JS)
- [JS in 14 minutes](https://jgthms.com/javascript-in-14-minutes/)
- [2017 JavaScript Rising Stars](https://risingstars.js.org/2017/en/)
- [ES6 features](https://github.com/lukehoban/es6features)
- [The Definitive Guide to Object-Oriented JavaScript](https://www.youtube.com/watch?v=PMfcsYzj-9M) - Amazing video to understand JS inheritance & objects.
- [JavaScript: The Core](http:J//dmitrysoshnikov.com/ecmascript/javascript-the-core-2nd-edition/)
- [JavaScript is Good, Actually](https://ashfurrow.com/blog/javascript-is-good-actually/) ([HN](https://news.ycombinator.com/item?id=17079952))
- [JavaScript Algorithms](https://github.com/trekhleb/javascript-algorithms) - Algorithms and data structures implemented in JavaScript with explanations and links to further readings. ([HN](https://news.ycombinator.com/item?id=28933618))
- [EC6 Features](http://es6-features.org/)
- [Clean Code concepts adapted for JavaScript](https://github.com/ryanmcdermott/clean-code-javascript)
- [Ask HN: “Expert Level” JavaScript questions?](https://news.ycombinator.com/item?id=17324538)
- [Benefits of prototypal inheritance over classical?](https://stackoverflow.com/questions/2800964/benefits-of-prototypal-inheritance-over-classical/16872315#16872315)
- [Pax](https://github.com/nathan/pax) - Fastest JavaScript bundler in the galaxy.
- [Philip Roberts: What the heck is the event loop anyway? (2014)](https://www.youtube.com/watch?v=8aGhZQkoFbQ)
- [Jake Archibald: In The Loop (2018)](https://www.youtube.com/watch?v=cCOL7MC4Pl0)
- [Yonatan Kra - The Event Loop and your code (2020)](https://www.youtube.com/watch?v=Nqx3rtv_dko)
- [BundlePhobia](https://bundlephobia.com/) - Find the cost of adding a npm package to your bundle. ([Code](https://github.com/pastelsky/bundlephobia)) ([Tweet](https://twitter.com/jsunderhood/status/1318204704988561409))
- [An Overview of JavaScript Testing in 2018](https://medium.com/welldone-software/an-overview-of-javascript-testing-in-2018-f68950900bc3)
- [Introduction to ES6 Promises – The Four Functions You Need To Avoid Callback Hell](http://jamesknelson.com/grokking-es6-promises-the-four-functions-you-need-to-avoid-callback-hell/)
- [Nice ES6/Promises/React cheat sheets](http://jamesknelson.com/thank-you-for-subscribing/)
- [JavaScript Visualizer](https://tylermcginnis.com/javascript-visualizer/) - Tool for visualizing Execution Context, Hoisting, Closures, and Scopes in JavaScript.
- [WallabyJS](https://wallabyjs.com/docs/) - Integrated continuous testing tool for JavaScript.
- [ES6 features](http://es6-features.org/)
- [The State of JavaScript - The State of the Web (2018)](https://www.youtube.com/watch?v=i5R7giitymk)
- [A Quick Tour Of ES6 (Or, The Bits You’ll Actually Use)](http://jamesknelson.com/es6-the-bits-youll-actually-use/)
- [JavaScript on the Desktop, Fast and Slow (2018)](https://medium.com/@felixrieseberg/javascript-on-the-desktop-fast-and-slow-2b744dfb8b55)
- [ES6 for humans](https://github.com/metagrover/ES6-for-humans)
- [33 concepts every JavaScript developer should know](https://github.com/leonardomso/33-js-concepts)
- [Design Patterns JS](https://github.com/fbeline/Design-Patterns-JS) - All the 23 (GoF) design patterns implemented in JavaScript.
- [Standard Library Proposal](https://github.com/tc39/proposal-javascript-standard-library)
- [30 seconds of code](https://github.com/30-seconds/30-seconds-of-code) - Curated collection of useful JavaScript snippets that you can understand in 30 seconds or less.
- [puppet-run](https://github.com/andywer/puppet-run) - Run anything JavaScript in a headless Chrome from your command line.
- [Yalc](https://github.com/whitecolor/yalc) - Better workflow than npm | yarn link for package authors.
- [ECMAScript proposals](https://github.com/tc39/proposals)
- [FromJS](https://github.com/mattzeunert/fromjs) - See where each character on the screen came from in code.
- [RunJS](https://projects.lukehaas.me/runjs/) - Scratchpad for your thoughts, a playground for your creativity.
- [Pragmatic, balanced FP in JavaScript book](https://github.com/getify/Functional-Light-JS)
- [Pack](https://github.com/pikapkg/pack) - Helps you build amazing packages without the hassle.
- [Learning JavaScript (2016)](https://mafinto.sh/blog/learning-javascript.html)
- [@pika/web](https://github.com/pikapkg/web) - Install npm dependencies that run directly in the browser. No Browserify, Webpack or import maps required.
- [Sucrase](https://github.com/alangpierce/sucrase) - Super-fast alternative to Babel for when you can target modern JS runtimes. ([Web](https://sucrase.io/))
- [Airbnb JavaScript Style Guide](https://github.com/airbnb/javascript)
- [JavaScript Developer's Reading List](https://github.com/twhite96/js-dev-reads) - List of hand-picked books and articles for JavaScript developers.
- [Promisees](https://github.com/bevacqua/promisees) - Promise visualization playground for the adventurous.
- [promise-fun](https://github.com/sindresorhus/promise-fun) - Promise packages, patterns, chat, and tutorials.
- [Perflink](https://github.com/lukejacksonn/perflink) - JavaScript performance benchmarks that you can share via URL.
- [Mostly adequate guide to FP (in JavaScript)](https://mostly-adequate.gitbook.io/mostly-adequate-guide/) ([Code](https://github.com/MostlyAdequate/mostly-adequate-guide)) ([HN](https://news.ycombinator.com/item?id=22654135))
- [Volta](https://github.com/volta-cli/volta) - JavaScript Launcher. ([Web](https://volta.sh/)) ([HN](https://news.ycombinator.com/item?id=27023453))
- [Modern JS Cheat Sheet](https://github.com/mbeaudru/modern-js-cheatsheet)
- [Fastpack](https://fastpack.sh/) - Pack JavaScript fast & easy.
- [Reference implementation for the JavaScript Binary AST format](https://github.com/binast/binjs-ref)
- [Babel Handbook](https://github.com/jamiebuilds/babel-handbook/blob/master/translations/en/README.md)
- [List of (Advanced) JavaScript Questions](https://github.com/lydiahallie/javascript-questions)
- [Faster script loading with BinaryAST? (2019)](https://blog.cloudflare.com/binary-ast/)
- [recast](https://github.com/benjamn/recast) - JavaScript syntax tree transformer, nondestructive pretty-printer, and automatic source map generator.
- [Madge](https://github.com/pahen/madge) - Create graphs from your CommonJS, AMD or ES6 module dependencies.
- [npmfs](https://npmfs.com/) - JavaScript Package Inspector.
- [Fantasy Land Specification](https://github.com/fantasyland/fantasy-land) - Specification for interoperability of common algebraic structures in JavaScript.
- [Meriyah](https://github.com/meriyah/meriyah) - 100% compliant, self-hosted javascript parser with high focus on both performance and stability.
- [The cost of JavaScript in 2019](https://news.ycombinator.com/item?id=20317736) ([HN](https://news.ycombinator.com/item?id=20317736))
- [Poi](https://github.com/egoist/poi) - Zero-config bundler for JavaScript applications.
- [Advanced JavaScript Course](https://tylermcginnis.com/courses/advanced-javascript)
- [Jay](https://github.com/nikersify/jay) - Supercharged JavaScript REPL.
- [Data Structures and Algorithms in JavaScript](https://github.com/amejiarosario/dsa.js-data-structures-algorithms-javascript)
- [JavaScript & Node.js Testing Best Practices](https://github.com/goldbergyoni/javascript-testing-best-practices)
- [Just](https://github.com/microsoft/just) - Library that organizes build tasks for your JS projects.
- [ECMAScript (JS) specification](https://tc39.es/ecma262/) ([Code](https://github.com/tc39/ecma262)) ([Web version 2](https://read262.netlify.com))
- [André Staltz: Two Fundamental Abstractions - Uphill Conf 2018](https://www.youtube.com/watch?v=fdol03pcvMA)
- [JSMonday](http://www.jsmonday.dev/) - Weekly JS inspiration.
- [Chevrotain](https://github.com/SAP/chevrotain) - Parser Building Toolkit for JavaScript.
- [Comprehensive list of new ES features](https://github.com/daumann/ECMAScript-new-features-list)
- [Exploring JS: JavaScript books for programmers](https://exploringjs.com/)
- [JavaScript for Impatient Programmers book](https://exploringjs.com/impatient-js/toc.html) ([HN](https://news.ycombinator.com/item?id=23689280)) ([HN](https://news.ycombinator.com/item?id=29673206))
- [Exploring ES2018 and ES2019](https://exploringjs.com/es2018-es2019/toc.html)
- [Ecma TC39 GitHub](https://github.com/tc39)
- [TC39 Meeting Notes](https://tc39.es/tc39-notes/)
- [Mesh Spreadsheet](https://github.com/chrispsn/mesh/) - Visualise data and edit JavaScript code using a spreadsheet interface. ([Web](http://mesh-spreadsheet.com/))
- [Immutable JavaScript Data Structures with Immer (2019)](https://egghead.io/courses/immutable-javascript-data-structures-with-immer)
- [Immutability is Changing - From Immutable.js to Immer (2019)](https://www.youtube.com/watch?v=bFuRvcAEiHg)
- [Tenko](https://github.com/pvdz/tenko) - 100% spec compliant ES2020 JavaScript parser written in JS.
- [code-red](https://github.com/Rich-Harris/code-red) - Experimental toolkit for writing x-to-JavaScript compilers.
- [Reduce in JavaScript (2019)](https://yuanchuan.dev/2019/03/04/the-reduce-function.html)
- [Pika](https://www.pika.dev/registry) - New kind of package registry for the modern web.
- [Brian Holt: Futurist Code Bases: Integrating JS of the Future Today (2019)](https://www.youtube.com/watch?v=lQOWTXanWwU)
- [JS TLDR](https://js-tldr.info/) - Zen mode web-documentation. ([Code](https://github.com/RusinovAnton/js-tldr)) ([Article](https://medium.com/@rusinovantondev/js-tl-dr-zen-mode-web-docs-for-javascript-developers-cf45e0143a09))
- [Currying Functions in ES6 (2016)](https://sunjay.dev/2016/08/13/es6-currying)
- [Manipulating AST with JavaScript (2019)](https://lihautan.com/manipulating-ast-with-javascript/)
- [Is JavaScript Statically or Dynamically Scoped? (2018)](https://www.cs.cornell.edu/~asampson/blog/scope.html)
- [Fixed-point combinators in JavaScript: Memoizing recursive functions](http://matt.might.net/articles/implementation-of-recursive-fixed-point-y-combinator-in-javascript-for-memoization/)
- [runpkg](https://github.com/FormidableLabs/runpkg) - Lets you navigate any JavaScript package on npm thanks to unpkg.com.
- [What is this in JavaScript?](https://www.madebymike.com.au/writing/this-in-javascript/)
- [Beginner JavaScript course](https://beginnerjavascript.com/)
- [ES6 Cheat Sheet](https://github.com/DrkSephy/es6-cheatsheet)
- [JavaScript Visualized: Event Loop (2019)](https://dev.to/lydiahallie/javascript-visualized-event-loop-3dif)
- [JavaScript Visualized: Scope (Chain) (2019)](https://dev.to/lydiahallie/javascript-visualized-scope-chain-13pd)
- [JavaScript Visualized: Hoisting (2019)](https://dev.to/lydiahallie/javascript-visualized-hoisting-478h)
- [Fuzzilli](https://github.com/googleprojectzero/fuzzilli) - JavaScript Engine Fuzzer.
- [Deep JavaScript: Theory and techniques (2019)](https://exploringjs.com/deep-js/) ([HN](https://news.ycombinator.com/item?id=23552180))
- [JavaScript Adaption of Structure and Interpretation of Computer Programs](https://sicp.comp.nus.edu.sg/) ([HN](https://news.ycombinator.com/item?id=21822903)) ([Code](https://github.com/source-academy/sicp))
- [State of JS 2019](https://2019.stateofjs.com/) ([HN](https://news.ycombinator.com/item?id=21831747))
- [Cancelation without Breaking a Promise (2016)](https://medium.com/hackernoon/considering-cancelation-a96e0f3c2298) - Reflecting on what was so tricky about cancelable Promises, embracing functional purity as a solution.
- [ECMAScript Discussion Archives](https://esdiscuss.org/) ([Code](https://github.com/esdiscuss/esdiscuss.org))
- [What Is JavaScript Made Of? (2019)](https://overreacted.io/what-is-javascript-made-of/)
- [JavaScript Visualized: Prototypal Inheritance (2020)](https://dev.to/lydiahallie/javascript-visualized-prototypal-inheritance-47co)
- [Y: The Most Beautiful Idea in Computer Science explained in JavaScript (2018)](https://lucasfcosta.com/2018/05/20/Y-The-Most-Beautiful-Idea-in-Computer-Science.html)
- [2019 JavaScript Rising Stars](https://risingstars.js.org/2019/en/) ([Code](https://github.com/bestofjs/javascript-risingstars))
- [Best of JS](https://bestofjs.org/) - Best of JavaScript, HTML and CSS. ([Code](https://github.com/bestofjs/bestofjs-webui)) ([Web Timeline](https://bestofjs.org/timeline)) ([HN](https://news.ycombinator.com/item?id=24142462))
- [omggif](https://github.com/deanm/omggif) - JavaScript implementation of a GIF 89a encoder and decoder.
- [Sampling bias, FDR, and The State of JS (2020)](https://davidea.st/articles/sampling-bias-fdr-state-of-js)
- [JavaScript Visualized: Generators and Iterators (2020)](https://dev.to/lydiahallie/javascript-visualized-generators-and-iterators-e36)
- [jsep](https://github.com/soney/jsep) - JavaScript Expression Parser.
- [JS Tips & Tidbits](https://github.com/nas5w/javascript-tips-and-tidbits)
- [I have been underestimating JS (2020)](https://adlrocha.substack.com/p/adlrocha-i-have-been-underestimating) - Understanding V8 and NodeJS Steams.
- [Микроптимизации производительности и JavaScript (2020)](https://medium.com/devschacht/optimizations-and-javascript-f8e060d3eae3)
- [Taming the asynchronous beast with CSP channels in JavaScript (2014)](https://jlongster.com/Taming-the-Asynchronous-Beast-with-CSP-in-JavaScript)
- [Debounce vs Throttle: Definitive Visual Guide (2019)](https://redd.one/blog/debounce-vs-throttle)
- [GistLink](https://gist.link/) - Code apps or components. See them render as you type. Share your creations via URL.
- [source-map-explorer](https://github.com/danvk/source-map-explorer) - Analyze and debug space usage through source maps.
- [Diglett](https://github.com/oblador/diglett) - Keep your JS project lean by detecting duplicate dependencies.
- [Learn Vanilla JS Roadmap](https://learnvanillajs.com/roadmap/)
- [Learn JavaScript](https://learnjavascript.online/) - Easiest way to learn & practice modern JavaScript step by step.
- [Mini projects built with HTML5, CSS & JavaScript. No frameworks or libraries](https://github.com/bradtraversy/vanillawebprojects) ([HN](https://news.ycombinator.com/item?id=22231963))
- [IxJS](https://github.com/ReactiveX/IxJS) - Interactive Extensions for JavaScript.
- [Renovate](https://github.com/renovatebot/renovate) - Universal dependency update tool that fits into your workflows.
- [The ECMAScript Ecosystem (2020)](https://dev.to/laurieontech/the-ecmascript-ecosystem-2e13)
- [ESbuild](https://github.com/evanw/esbuild/) - Extremely fast JavaScript bundler and minifier written in Go. ([HN](https://news.ycombinator.com/item?id=22335707)) ([Architecture](https://github.com/evanw/esbuild/blob/master/docs/architecture.md)) ([serverless-esbuild](https://github.com/floydspace/serverless-esbuild)) ([Awesome](https://github.com/egoist/awesome-esbuild)) ([Web](https://esbuild.github.io/)) ([Esbuild plugins](https://github.com/remorses/esbuild-plugins)) ([Tweet](https://twitter.com/tannerlinsley/status/1445135850866499586)) ([HN](https://news.ycombinator.com/item?id=28860713)) ([Online Playground](https://github.com/egoist/play-esbuild))
- [Community plugins for esbuild](https://github.com/esbuild/community-plugins)
- [Why Is Esbuild Fast?](https://esbuild.github.io/faq/#why-is-esbuild-fast) ([HN](https://news.ycombinator.com/item?id=26154509))
- [bundless](https://github.com/remorses/bundless) - Dev server and bundler for esbuild. ([Web](https://bundless.vercel.app/))
- [esbuild-register](https://github.com/egoist/esbuild-register) - Transpile JSX, TypeScript and esnext features on the fly with esbuild.
- [JavaScript: Understanding the Weird Parts course (2015)](https://www.udemy.com/course/understand-javascript/)
- [Fastpack](https://github.com/fastpack/fastpack) - Pack JS code into a single bundle fast & easy.
- [guijs](https://github.com/Akryum/guijs) - App that helps you manage JS projects with a Graphical User Interface.
- [Rome Toolchain](https://github.com/rome/tools) - Linter, compiler, bundler, and more for JavaScript, TypeScript, HTML, Markdown, and CSS. ([Web](https://rome.tools/)) ([HN](https://news.ycombinator.com/item?id=22430682)) ([HN 2](https://news.ycombinator.com/item?id=24094377)) ([HN 3](https://news.ycombinator.com/item?id=24882413))
- [Bolt](https://github.com/boltpkg/bolt) - Super-powered JavaScript project management.
- [tiny-js](https://github.com/gfwilliams/tiny-js) - Aims to be an extremely simple (~2000 line) JavaScript interpreter.
- [JavaScript and TypeScript tooling overview](https://github.com/slikts/tooling)
- [Seafox](https://github.com/KFlash/seafox) - Blazing fast 100% spec compliant, self-hosted javascript parser written in Typescript.
- [Awesome JavaScript Learning](https://github.com/micromata/awesome-javascript-learning)
- [Awesome Promises](https://github.com/wbinnssmith/awesome-promises)
- [jscodeshift](https://github.com/facebook/jscodeshift) - Toolkit for running codemods over multiple JavaScript or TypeScript files.
- [React Workout: Reducers with Cassidy Williams (2020)](https://www.youtube.com/watch?v=sf4spiPynBE)
- [JavaScript: The First 20 Years (2020)](http://www.wirfs-brock.com/allen/posts/866)
- [Awesome Storybook](https://github.com/lauthieb/awesome-storybook)
- [QuickJS](https://github.com/bellard/quickjs) - Small and embeddable JavaScript engine. ([Web](https://bellard.org/quickjs/)) ([HN](https://news.ycombinator.com/item?id=24867103)) ([HN](https://news.ycombinator.com/item?id=30598026))
- [Test262: Official ECMAScript Conformance Test Suite](https://github.com/tc39/test262)
- [Hegel](https://github.com/JSMonk/hegel) - Advanced static type checker. ([Web](https://hegel.js.org/)) ([Intro to Hegel](https://blog.logrocket.com/introduction-to-hegel/))
- [NectarJS](https://github.com/NectarJS/nectarjs) - JS God mode. No VM. No Bytecode. No Garbage Collector. Full Compiled and Native binaries.
- [Eloquent JavaScript book (2018)](https://eloquentjavascript.net/) ([HN](https://news.ycombinator.com/item?id=22990926))
- [JS.coach](https://js.coach/) - Manually curated list of packages related to React, Vue, Webpack, Babel and PostCSS. ([Code](https://github.com/jscoach/client))
- [How to create a reactive state-based UI component with vanilla JS Proxies (2020)](https://gomakethings.com/how-to-create-a-reactive-state-based-ui-component-with-vanilla-js-proxies/)
- [Kite Autocomplete for JavaScript](https://www.kite.com/javascript/) ([Article](https://www.kite.com/blog/product/kite-launches-ai-powered-javascript-completions/))
- [Excalidraw: Cool JS Tricks Behind the Scenes - Christopher Chedeau (2020)](https://www.youtube.com/watch?v=fix2-SynPGE)
- [Cleaner async JavaScript code without the try/catch mess (2020)](https://davidwells.io/blog/cleaner-async-await-code-without-try-catch)
- [Shifty](https://github.com/jeremyckahn/shifty) - Tweening engine for JavaScript. It is a lightweight library meant to be encapsulated by higher level tools.
- [JS Bits](https://github.com/vasanthk/js-bits) - JavaScript concepts explained with code.
- [Binary-parser](https://github.com/keichi/binary-parser) - Parser builder for JavaScript that enables you to write efficient binary parsers in a simple and declarative manner.
- [estrella](https://github.com/rsms/estrella) - Light-weight runner for the esbuild compiler.
- [jsparagus](https://github.com/mozilla-spidermonkey/jsparagus) - JavaScript parser written in Rust.
- [Callbag](https://github.com/callbag/callbag) - Standard for JS callbacks that enables lightweight observables and iterables. ([Wiki](https://github.com/callbag/callbag/wiki))
- [JavaScript Standard Style](https://standardjs.com/) - JavaScript style guide, linter, and formatter. ([Code](https://github.com/standard/standard))
- [Boa](https://github.com/boa-dev/boa) - Experimental JavaScript lexer, parser and compiler written in Rust. ([Blog](https://boa-dev.github.io/)) ([Reddit](https://www.reddit.com/r/rust/comments/tfno83/boa_release_v014_a_javascript_engine_written_in/))
- [Understanding JavaScript Execution Context like never before (2020)](https://blog.greenroots.info/understanding-javascript-execution-context-like-never-before-ckb8x246k00f56hs1nefzpysq)
- [Causes of Memory Leaks in JavaScript and How to Avoid Them (2020)](https://www.ditdot.hr/en/causes-of-memory-leaks-in-javascript-and-how-to-avoid-them) ([Lobsters](https://lobste.rs/s/ar5avz/causes_memory_leaks_javascript_how_avoid))
- [UI.dev](https://ui.dev/) - Master the JavaScript Ecosystem.
- [Do Not Follow JavaScript Trends (2020)](https://pragmaticpineapple.com/do-not-follow-javascript-trends/) ([Lobsters](https://lobste.rs/s/wb3ma8/do_not_follow_javascript_trends)) ([HN](https://news.ycombinator.com/item?id=23538473))
- [Some things that can be avoided in JS for clearer code (2020)](https://twitter.com/buildsghost/status/1274042818219044864)
- [JS fundamentals and resources to learn them (2020)](https://twitter.com/Madisonkanna/status/1274424134139666432)
- [A little bit of plain JavaScript can do a lot (2020)](https://jvns.ca/blog/2020/06/19/a-little-bit-of-plain-javascript-can-do-a-lot/) ([Lobsters](https://lobste.rs/s/6umqjn/little_bit_plain_javascript_can_do_lot)) ([HN](https://news.ycombinator.com/item?id=23578319))
- [Memoization: What, Why, and How (2020)](https://kyleshevlin.com/memoization)
- [An Open Source Maintainer's Guide to Publishing npm Packages (2020)](https://formidable.com/blog/2020/publish-npm-packages/)
- [Robust Client-Side JavaScript (2020)](https://molily.de/robust-javascript/) ([HN](https://news.ycombinator.com/item?id=23612184))
- [Visualization of npm dependencies](https://npm.anvaka.com/#!/) ([Code](https://github.com/anvaka/npmgraph.an))
- [How to Learn JavaScript](https://sivers.org/learn-js) ([HN](https://news.ycombinator.com/item?id=23659531))
- [Google Closure Compiler](https://github.com/google/closure-compiler) - Tool for making JavaScript download and run faster.
- [JSConf](https://jsconf.com/) - Conferences for the JavaScript Community.
- [The history of Promises](https://samsaccone.com/posts/history-of-promises.html)
- [Skypack](https://www.skypack.dev/) - New kind of JavaScript delivery network. ([HN](https://news.ycombinator.com/item?id=23825798)) ([Introducing Skypack Discover](https://www.skypack.dev/blog/2020/10/introducing-skypack-discover/)) ([Docs](https://docs.skypack.dev/)) ([skypin](https://github.com/MarshallCB/skypin))
- [Openbase](https://openbase.io/) - Help developers choose the right JS package for any task - through user reviews and insights about packages' popularity, reliability, activity and more. ([HN](https://news.ycombinator.com/item?id=23833441))
- [Basho](https://github.com/bashojs/basho) - Shell Automation with Plain JavaScript. ([Docs](https://bashojs.org/))
- [What is the JS Event Loop and Call Stack?](https://gist.github.com/jesstelford/9a35d20a2aa044df8bf241e00d7bc2d0)
- [Starving the Event Loop with microtasks](https://gist.github.com/jesstelford/bbb30b983bddaa6e5fef2eb867d37678)
- [GPU.js](https://gpu.rocks/#/) - GPU accelerated JavaScript. ([HN](https://news.ycombinator.com/item?id=24027487)) ([HN](https://news.ycombinator.com/item?id=28797182))
- [The JavaScript Promise Tutorial (2020)](https://adrianmejia.com/promises-tutorial-concurrency-in-javascript-node/)
- [Underrated JS array methods (2020)](https://dev.to/assuncaocharles/underrated-array-methods-2mdj)
- [Javascript Generators, Meet XPath (2020)](https://jack.wrenn.fyi/blog/xpath-for-2020/) ([Lobsters](https://lobste.rs/s/5fsljg/javascript_generators_meet_xpath))
- [goja](https://github.com/dop251/goja) - ECMAScript 5.1(+) implementation in Go.
- [Guide to unit testing in JavaScript](https://github.com/mawrkus/js-unit-testing-guide)
- [How I wrote the fastest JavaScript memoization library (2017)](https://community.risingstack.com/the-worlds-fastest-javascript-memoization-library/)
- [JavaScript Playgrounds](https://unpkg.com/javascript-playgrounds@^1.0.0/public/index.html) - Interactive JavaScript sandbox. ([Code](https://github.com/dabbott/javascript-playgrounds))
- [Speakeasy JS](https://speakeasyjs.com/) - Weekly JavaScript meetup.
- [Elsa](https://github.com/elsaland/elsa) - Minimal, fast and secure QuickJS wrapper written in Go. ([HN](https://news.ycombinator.com/item?id=24626655))
- [quickjs-rs](https://github.com/theduke/quickjs-rs) - Rust wrapper for QuickJS.
- [RSLint](https://github.com/rslint/rslint) - Extremely fast JavaScript and TypeScript linter and Rust crate. ([Web](https://rslint.org/))
- [Beginner's Series to: JavaScript by Microsoft](https://www.youtube.com/playlist?list=PLlrxD0HtieHhW0NCG7M536uHGOtJ95Ut2) ([Code](https://github.com/microsoft/beginners-intro-javascript-node))
- [Please stop using CDNs for external Javascript libraries (2020)](https://shkspr.mobi/blog/2020/10/please-stop-using-cdns-for-external-javascript-libraries/) ([Lobsters](https://lobste.rs/s/mpznhm/please_stop_using_cdns_for_external)) ([HN](https://news.ycombinator.com/item?id=24745194))
- [ESM Hot Module Replacement (ESM-HMR) Spec](https://github.com/pikapkg/esm-hmr)
- [esbuild-js](https://github.com/marvinhagemeister/esbuild-js) - es-build implemented in JS.
- [Visual Guide to References in JavaScript (2020)](https://daveceddia.com/javascript-references/)
- [Modern JavaScript features you may have missed (2019)](https://www.breck-mckye.com/blog/2019/10/modern-javascript-features-you-may-have-missed/)
- [RegPack](https://github.com/Siorki/RegPack) - Self-contained packer for size-constrained JS code.
- [SciterJS](https://sciter.com/) - Embeddable HTML/CSS/JavaScript engine. Electron alternative. ([HN](https://news.ycombinator.com/item?id=24797423)) ([SDK](https://github.com/c-smile/sciter-sdk)) ([JS SDK](https://github.com/c-smile/sciter-js-sdk)) ([HN](https://news.ycombinator.com/item?id=29742670))
- [ESTree Spec](https://github.com/estree/estree) - Manipulate JavaScript source code.
- [Pattern Matching in JavaScript (2020)](https://kyleshevlin.com/pattern-matching)
- [How to chain methods in JS in order to write concise and readable code (2020)](https://medium.com/@laflamablanc/method-chaining-and-javascript-7d840d6e3687)
- [npmview](https://npmview.now.sh/) - Web application to view npm package files. ([Code](https://github.com/pd4d10/npmview))
- [Metadata Reflection API for JS](https://github.com/rbuckton/reflect-metadata)
- [SurviveJS](https://survivejs.com/) - Learn JavaScript. From apprentice to master.
- [Composing Closures and Callbacks in JavaScript (2020)](https://egghead.io/playlists/composing-closures-and-callbacks-in-javascript-1223)
- [CJS Module Lexer](https://github.com/guybedford/cjs-module-lexer) - Fast lexer to extract named exports via analysis from CommonJS modules.
- [JavaScript minification (2019)](http://neugierig.org/software/blog/2019/04/js-minifiers.html)
- [export-size](https://github.com/antfu/export-size) - Analysis bundle cost for each export of an ESM package.
- [ESM](https://github.com/postui/esm.sh) - Fast, global content delivery network ES Modules.
- [Benny](https://github.com/caderek/benny) - Dead simple benchmarking framework for JS/TS libs.
- [Functional Programming in JS - Composition (Currying, Lodash and Ramda) (2020)](https://11sigma.com/blog/functional-programming-in-js-part-i-composition)
- [Understanding Modules, Import and Export in JavaScript (2020)](https://www.taniarascia.com/javascript-modules-import-export/)
- [Intent to stop using 'null' in my JS code](https://github.com/sindresorhus/meta/issues/7) ([HN](https://news.ycombinator.com/item?id=24956156))
- [What Makes JavaScript JavaScript? Prototypal Inheritance (2020)](https://dmitripavlutin.com/javascript-prototypal-inheritance/)
- [ni](https://github.com/antfu/ni) - Use the right package manager. Detect whether to use npm/yarn/pnpm.
- [Making a modern JS library in 2020](https://pitayan.com/posts/modernest-lib-hello-world/)
- [JavaScript Interview Questions & Answers](https://github.com/sudheerj/javascript-interview-questions)
- [JS Operator Lookup](https://joshwcomeau.com/operator-lookup/) - Search JavaScript Operators.
- [The state of JavaScript at the end of 2020](https://www.ideamotive.co/javascript-business-guide) ([HN](https://news.ycombinator.com/item?id=25046293))
- [What the fuck JavaScript](https://github.com/denysdovhan/wtfjs) - List of funny and tricky JavaScript examples.
- [1loc](https://1loc.dev/) - JavaScript Utilities in 1 LOC. ([Code](https://github.com/1milligram/1loc))
- [Component Driven User Interfaces](https://www.componentdriven.org/) - Open standard for UI component examples based on JavaScript ES6 modules. ([Code](https://github.com/ComponentDriven/csf)) ([Website Code](https://github.com/ComponentDriven/componentdriven.org))
- [JavaScript Modern Interview Code Challenges](https://github.com/sadanandpai/javascript-code-challenges)
- [Building a Promise from Scratch (2020)](https://www.youtube.com/watch?v=CVzx-6fu0d8)
- [Tips and tricks for working with types in JavaScript](https://github.com/voxpelli/types-in-js)
- [Astring](https://github.com/davidbonnet/astring) - Tiny and fast JavaScript code generator from an ESTree-compliant AST.
- [EStimator.dev](https://estimator.dev/) - Calculate the size and performance impact of switching to modern JavaScript syntax. ([Code](https://github.com/GoogleChromeLabs/estimator.dev))
- [Publish, ship, and install modern JavaScript for faster applications (2020)](https://web.dev/publish-modern-javascript/)
- [Universal JavaScript Build and Packaging](https://github.com/mikeal/ipjs)
- [Maybe you don't need Rust and WASM to speed up your JS (2018)](https://mrale.ph/blog/2018/02/03/maybe-you-dont-need-rust-to-speed-up-your-js.html)
- [lage](https://github.com/microsoft/lage) - Task runner in JS monorepos. ([Web](https://microsoft.github.io/lage/))
- [Module Server](https://github.com/google/module-server) - System for efficient serving of CommonJS modules to web browsers.
- [How JavaScript works: exceptions + best practices for synchronous and asynchronous code (2021)](https://blog.sessionstack.com/how-javascript-works-exceptions-best-practices-for-synchronous-and-asynchronous-environments-39f66b59f012)
- [How JavaScript works: an overview of the engine, the runtime, and the call stack (2017)](https://blog.sessionstack.com/how-does-javascript-actually-work-part-1-b0bacc073cf) ([HN](https://news.ycombinator.com/item?id=27941979))
- [Source Map Visualization](http://evanw.github.io/source-map-visualization/) - Visualization of JavaScript source map data, which is useful for debugging problems with generated source maps. ([Code](https://github.com/evanw/source-map-visualization))
- [Manypkg](https://github.com/Thinkmill/manypkg) - Linter for package.json files in Yarn, Bolt or pnpm monorepos.
- [Putout](https://github.com/coderaiser/putout) - Pluggable and configurable code transformer with built-in eslint, babel plugins and jscodeshift codemods support. ([Editor](https://putout.cloudcmd.io/))
- [Prettier Plugin: Organize Imports](https://github.com/simonhaenisch/prettier-plugin-organize-imports)
- [A mostly complete guide to error handling in JavaScript (2020)](https://www.valentinog.com/blog/error/)
- [Awesome FP JS](https://github.com/stoeffel/awesome-fp-js)
- [Perflink](https://perf.link/) - JS Benchmarks.
- [JavaScript benchmark playground](https://jsbench.github.io/)
- [Element Worklet (2021)](https://jasonformat.com/element-worklet/)
- [MDN JS Code Examples](https://github.com/mdn/js-examples)
- [Understanding Hoisting in JavaScript (2021)](https://www.digitalocean.com/community/tutorials/understanding-hoisting-in-javascript)
- [JavaScript: The Good Parts Book (2008)](https://www.oreilly.com/library/view/javascript-the-good/9780596517748/) ([Notes](https://github.com/dwyl/Javascript-the-Good-Parts-notes)) ([Notes](https://github.com/ahmaazouzi/js_good_parts))
- [A Model for Reasoning About JavaScript Promises (2017)](http://www.franktip.org/pubs/oopsla2017promises.pdf)
- [JavaScript Minification Benchmarks](https://github.com/privatenumber/minification-benchmarks) ([HN](https://news.ycombinator.com/item?id=26048291))
- [Faster JavaScript Calls (2021)](https://v8.dev/blog/adaptor-frame) ([HN](https://news.ycombinator.com/item?id=26143648))
- [Streams — The definitive guide (2021)](https://web.dev/streams/)
- [Starlight](https://github.com/Starlight-JS/starlight) - JS engine in Rust.
- [Awesome JavaScript](https://github.com/sorrycc/awesome-javascript)
- [JavaScript and the next decade of data programming (2021)](http://benschmidt.org/post/2020-01-15/2020-01-15-webgpu/) ([Tweet](https://twitter.com/benmschmidt/status/1368988766174666753))
- [Glide](https://github.com/Myrannas/glide) - Interpreted JavaScript VM written entirely in safe rust.
- [stricter](https://github.com/atlassian/stricter) - Project-wide js-linting tool.
- [JavaScript Pass By Value Function Parameters](https://kentcdodds.com/blog/javascript-pass-by-value-function-parameters)
- [Use console.log() like a pro (2020)](https://markodenic.com/use-console-log-like-a-pro/) ([HN](https://news.ycombinator.com/item?id=26779800))
- [The complete guide to working with strings in modern JavaScript](https://www.baseclass.io/guides/string-handling-modern-js)
- [Natto.dev](https://natto.dev/) - Canvas for JavaScript. ([HN](https://news.ycombinator.com/item?id=26790285))
- [Kataw](https://github.com/kataw/kataw) - Insane fast JavaScript toolchain.
- [JS Tooling not in JS](https://github.com/RobinCsl/awesome-js-tooling-not-in-js) - Curated list of JavaScript tooling not written in JavaScript. ([HN](https://news.ycombinator.com/item?id=26872457))
- [How Slow is JavaScript Really? JavaScript vs C++ (Data Structures & Optimization) (2021)](https://www.youtube.com/watch?v=WLwTlC1R2sY)
- [Duktape](https://github.com/svaarala/duktape) - Embeddable Javascript engine with a focus on portability and compact footprint. ([Web](https://duktape.org/))
- [JavaScript for Data Science](http://js4ds.org/) ([Code](https://github.com/software-tools-in-javascript/js4ds/)) ([HN](https://news.ycombinator.com/item?id=26931296))
- [FIBJS](https://github.com/fibjs/fibjs) - JavaScript runtime built on Chrome's V8 JavaScript engine.
- [A JavaScript optimizing compiler (2021)](https://medium.com/leaningtech/a-javascript-optimizing-compiler-3fd3f49bd071)
- [In-depth exploration of JavaScript iteration protocols (2021)](https://www.nodejsdesignpatterns.com/blog/javascript-async-iterators/)
- [Elk](https://github.com/cesanta/elk) - Tiny JS engine for embedded systems. ([HN](https://news.ycombinator.com/item?id=28614092))
- [Endo](https://github.com/endojs/endo) - Distributed secure JavaScript sandbox, based on SES.
- [How JavaScript works: 3 types of polymorphism (2021)](https://blog.sessionstack.com/how-javascript-works-3-types-of-polymorphism-f10ff4992be1) ([HN](https://news.ycombinator.com/item?id=27074514))
- [Modern JavaScript: Everything you missed over the last 10 years (2020)](https://turriate.com/articles/modern-javascript-everything-you-missed-over-10-years) ([HN](https://news.ycombinator.com/item?id=27165954))
- [Rethinking JavaScript Infrastructure (2021)](https://cpojer.net/posts/rethinking-javascript-infrastructure)
- [2D Optics Demos in JavaScript (2021)](https://www.philipzucker.com/aesthetic-javascript-eduction/)
- [Beginner Web Dev](https://beginnerwebdev.com/) - Getting Started with JavaScript Course.
- [V8 Sparkplug – A non-optimizing JavaScript compiler (2021)](https://v8.dev/blog/sparkplug) ([HN](https://news.ycombinator.com/item?id=27304808))
- [JavaScript Notes and Reference](https://wesbos.com/javascript)
- [JavaScript Bytecode VM - YouTube](https://www.youtube.com/playlist?list=PLMOpZvQB55beChggmvk-sUm8X_vSezpqL)
- [Advanced console.log Tips and Tricks (2020)](https://medium.com/nmc-techblog/advanced-console-log-tips-tricks-fa3762930bca) ([HN](https://news.ycombinator.com/item?id=27499335))
- [Testing JavaScript Applications Book (2021)](https://www.manning.com/books/testing-javascript-applications) ([Code](https://github.com/lucasfcosta/testing-javascript-applications))
- [JavaScript Is Weird](https://jsisweird.com/) ([HN](https://news.ycombinator.com/item?id=27659590))
- [Temporal: Getting started with JavaScript's new date time API (2021)](https://2ality.com/2021/06/temporal-api.html) ([HN](https://news.ycombinator.com/item?id=27661667))
- [Dwitter](https://www.dwitter.net/) - Platform to write visual art in JavaScript limited to 140 characters. ([Code](https://github.com/lionleaf/dwitter))
- [Optimizr](https://github.com/undefinedbuddy/optimizr) - Aggressive compile-time optimizations for modern JavaScript.
- [Tiny Treeshaker](https://github.com/magicmark/tiny-treeshaker) - JavaScript tree shaking in 200 lines of code. ([HN](https://news.ycombinator.com/item?id=27670859))
- [Tips For Using Async/Await in JavaScript (2021)](https://www.youtube.com/watch?v=_9vgd9XKlDQ)
- [`export default thing` is different to `export { thing as default }` (2021)](https://jakearchibald.com/2021/export-default-thing-vs-thing-as-default/) ([Tweet](https://twitter.com/jaffathecake/status/1411976300110258176))
- [Just JavaScript](https://justjavascript.com/) - Discover and rebuild your JavaScript mental models.
- [Iterator Helpers](https://github.com/tc39/proposal-iterator-helpers) - Methods for working with iterators in ECMAScript. ([Web](https://tc39.es/proposal-iterator-helpers/))
- [Loupe](http://latentflip.com/loupe/) - Visualization to help you understand how JavaScript's call stack/event loop/callback queue interact with each other. ([Code](https://github.com/latentflip/loupe))
- [JavaScript on Compute@Edge](https://developer.fastly.com/learning/compute/javascript/) - Compile JavaScript to WebAssembly and run it on Fastly.
- [Txiki.js](https://github.com/saghul/txiki.js) - Tiny JavaScript runtime built with QuickJS and libuv. ([HN](https://news.ycombinator.com/item?id=28020944))
- [JavaScript error logging](https://github.com/cheeaun/javascript-error-logging) - Collection of JavaScript error logging services, and comparisons among them.
- [Regenerated.Dev](https://regenerated.dev/) - Rethinking JavaScript with Generator Functions.
- [JavaScript Array Explorer](https://sdras.github.io/array-explorer/) - Find the array method you need without digging through the docs. ([Code](https://github.com/sdras/array-explorer))
- [Currying in JavaScript](https://javascript.info/currying-partials) ([HN](https://news.ycombinator.com/item?id=28583069))
- [Multithreaded JavaScript Book (2021)](https://www.oreilly.com/library/view/multithreaded-javascript/9781098104429/) - Explores the various features that JavaScript runtimes have at their disposal for implementing multithreaded programming.
- [JS-Interpreter](https://github.com/NeilFraser/JS-Interpreter) - Sandboxed JavaScript interpreter in JavaScript. Execute arbitrary JavaScript code line by line in isolation and safety.
- [EStruct](https://github.com/RayLuxembourg/estruct) - Traverses JavaScript projects and maps all the dependencies and relationships to a JSON.
- [unbuild](https://github.com/unjs/unbuild) - Unified JavaScript build system.
- [The Shocking Immaturity of JavaScript (2021)](https://dev.to/jaredcwhite/the-shocking-immaturity-of-javascript-c70) ([HN](https://news.ycombinator.com/item?id=28816961))
- [Nope](https://github.com/bvego/nope-validator) - Small, simple and fast JS validator.
- [jstime](https://github.com/jstime/jstime) - Another JavaScript Runtime.
- [JavaScript Concrete Syntax Tree](https://github.com/cst/cst)
- [Fastly Compute@Edge JS Runtime](https://github.com/fastly/js-compute-runtime)
- [Responsible JavaScript Book](https://responsiblejs.dev/) ([Twitter](https://twitter.com/responsiblejs))
- [Creating Callable Objects in JavaScript (2019)](https://hackernoon.com/creating-callable-objects-in-javascript-d21l3te1)
- [Reflection at Reflect: The Reflect and Proxy APIs (2021)](https://reflect.run/articles/reflection-at-reflect/)
- [Deep-copying in JavaScript (2018)](https://surma.dev/things/deep-copy/index.html) ([Tweet](https://twitter.com/DasSurma/status/1456999878798807044))
- [Reduce minified code size by property mangling](https://lihautan.com/reduce-minified-code-size-by-property-mangling/)
- [ECMAScript Pattern Matching](https://github.com/tc39/proposal-pattern-matching)
- [JSConf](https://www.jsconf.in/) ([Twitter](https://twitter.com/jsconfin))
- [JS Sockets API Proposal](https://github.com/jasnell/sockets-api) ([Tweet](https://twitter.com/jasnell/status/1460415509459968003))
- [JS Enum Proposal](https://github.com/Jack-Works/proposal-enum) ([Tweet](https://twitter.com/robpalmer2/status/1459834460954890243))
- [JS Records & Tuples Proposal](https://github.com/tc39/proposal-record-tuple)
- [ECMAScript Proposals](https://www.proposals.es/) ([Code](https://github.com/saadq/proposals.es))
- [dum](https://github.com/egoist/dum) - npm scripts runner written in Rust. ([HN](https://news.ycombinator.com/item?id=30778321))
- [Esprima-python](https://github.com/Kronuz/esprima-python) - ECMAScript parsing infrastructure for multipurpose analysis.
- [Software Tools in JavaScript by Greg Wilson Book (2021)](https://leanpub.com/stjs/)
- [ReevaJS](https://github.com/ReevaJS/reeva) - JavaScript engine written from scratch using Kotlin.
- [Web Developer Tools secrets that shouldn’t be secrets (2021)](https://christianheilmann.com/2021/11/01/developer-tools-secrets-that-shouldnt-be-secrets/) ([Reddit](https://www.reddit.com/r/programming/comments/r8wph3/web_developer_tools_secrets_that_shouldnt_be/))
- [The async/await post we promised (2021)](https://dev.to/ablydev/the-asyncawait-post-we-promised-2c50)
- [JavaScript Cross-Site Scripting: How it Happens and Mitigation Steps (2021)](https://www.youtube.com/watch?v=YlMCp_vzbQo)
- [JavaScript Self-Profiling API Proposal](https://github.com/WICG/js-self-profiling) - Proposal for a programmable JS profiling API for collecting JS profiles from real end-user environments.
- [Caterwaul](https://github.com/spencertipping/caterwaul) - JavaScript-to-JavaScript Compiler. ([Web](http://caterwauljs.org/)) ([HN](https://news.ycombinator.com/item?id=29563076))
- [MuJS](https://github.com/ccxvii/mujs) - Embeddable JavaScript interpreter in C. ([Web](https://mujs.com/))
- [JavaScript Structs: Fixed Layout Objects](https://github.com/tc39/proposal-structs)
- [HEY is running its JavaScript off import maps (2022)](https://world.hey.com/dhh/hey-is-running-its-javascript-off-import-maps-2abcf203)
- [A pipe operator for JavaScript: introduction and use cases (2022)](https://2ality.com/2022/01/pipe-operator.html) ([HN](https://news.ycombinator.com/item?id=30097586))
- [Tera](https://github.com/gigamono/tera) - Lean Secure Runtime for JavaScript.
- [Ratel](https://github.com/ratel-rust/ratel-core) - High performance JavaScript to JavaScript compiler with a Rust core.
- [TC39 – Specifying JavaScript](https://tc39.es/) ([Web Code](https://github.com/tc39/tc39.github.io))
- [sane-fmt](https://github.com/sane-fmt/sane-fmt) - Opinionated code formatter for TypeScript and JavaScript.
- [Run JavaScript in WebAssembly](https://github.com/second-state/wasmedge-quickjs) - High-performance, secure, extensible, and OCI-complaint JavaScript runtime for WasmEdge.
- [Insightful JavaScript Q&A Twitter Thread from Dan Abramov (2022)](https://twitter.com/dan_abramov/status/1492880870499360769)
- [The State of JS 2021](https://2021.stateofjs.com/en-US/) ([HN](https://news.ycombinator.com/item?id=30357788))
- [The Story of Asynchronous JavaScript (2022)](https://www.youtube.com/watch?v=rivBfgaEyWQ)
- [JavaScript on Exercism](https://exercism.org/tracks/javascript) ([Code](https://github.com/exercism/javascript))
- [Source Map-Aware Code Generation (2022)](https://gal.hagever.com/posts/source-map-aware-code-generation)
- [Things you don't need JavaScript for (2022)](https://lexoral.com/blog/you-dont-need-js/) ([HN](https://news.ycombinator.com/item?id=30512512))
- [Proposal for Type Syntax in JavaScript](https://devblogs.microsoft.com/typescript/a-proposal-for-type-syntax-in-javascript/) ([Code](https://github.com/giltayar/proposal-types-as-comments)) ([Reddit](https://www.reddit.com/r/javascript/comments/taejmb/a_proposal_for_type_syntax_in_javascript/)) ([Tweet](https://twitter.com/Rich_Harris/status/1501638466857283588)) ([Reddit](https://www.reddit.com/r/programming/comments/taejl7/a_proposal_for_type_syntax_in_javascript/)) ([Tweet](https://twitter.com/alexandereardon/status/1501768797501341699)) ([HN](https://news.ycombinator.com/item?id=30653574))
- [First look: adding type annotations to JavaScript (2022)](https://2ality.com/2022/03/type-annotations-first-look.html) ([HN](https://news.ycombinator.com/item?id=30626458))
- [Escargot](https://github.com/Samsung/escargot) - Memory optimized JavaScript engine for mid-range devices such as mobile phone, tablet and TV.
- [Pipe Operator (|>) for JavaScript](https://github.com/tc39/proposal-pipeline-operator) ([Tweet](https://twitter.com/ianstormtaylor/status/1537837121553178626))
- [Partial Application Syntax for ECMAScript](https://github.com/tc39/proposal-partial-application)
- [The Elephant in The Event Loop (2022)](https://gashamola.com/2022/03/16/the-elephant-in-the-event-loop.html)
- [Moon](https://github.com/milesj/moon) - Rust program for managing JavaScript based monorepo's.
- [Temporal API](https://github.com/tc39/proposal-temporal) - Standard objects and functions for working with dates and times.
- [How to Read the ECMAScript Specification](https://timothygu.me/es-howto/) ([Code](https://github.com/TimothyGu/es-howto))
- [Broken Promises - James Snell, NearForm (2019)](https://www.youtube.com/watch?v=XV-u_Ow47s0) ([Tweet](https://twitter.com/matteocollina/status/1507402939022860291))
- [JS Event Loop Visualizer](https://www.jsv9000.app/) ([Tweet](https://twitter.com/elmd_/status/1508743380859338768))
- [Google Closure Compiler NPM](https://github.com/google/closure-compiler-npm) - Check, compile, optimize and compress JavaScript with Closure-Compiler.
- [unminify](https://github.com/xfthhxk/unminify) - Unminifies JS stacktrace errors.
- [JavaScript on GraalVM](https://github.com/oracle/graaljs) - High performance implementation of the JavaScript programming language. Built on the GraalVM by Oracle Labs.
- [I Avoid Async/Await (2022)](https://uniqname.medium.com/why-i-avoid-async-await-7be98014b73e) ([HN](https://news.ycombinator.com/item?id=31050650))
- [Shift AST](https://shift-ast.org/) - ECMAScript parser that produces a Shift format AST. ([Code](https://github.com/shapesecurity/shift-parser-js)) ([Bending JavaScript with shift-ast](https://loophole-letters.vercel.app/shift-ast))
- [JavaScript function composition: What’s the big deal? (2022)](https://jrsinclair.com/articles/2022/javascript-function-composition-whats-the-big-deal/)
- [How not to write property tests in JavaScript (2021)](https://jrsinclair.com/articles/2021/how-not-to-write-property-tests-in-javascript/)
- [Find what JavaScript variables are leaking into the global scope (2022)](https://mmazzarolo.com/blog/2022-02-14-find-what-javascript-variables-are-leaking-into-the-global-scope/)
- [Building a JavaScript Bundler (2022)](https://cpojer.net/posts/building-a-javascript-bundler) ([HN](https://news.ycombinator.com/item?id=31165509))
- [ECMAScript proposal: do expressions](https://github.com/tc39/proposal-do-expressions)
- [Meticulous](https://meticulous.ai/) - Catch JavaScript errors before they hit prod. ([HN](https://news.ycombinator.com/item?id=31236066))
- [I fell in love with low-JS (2022)](https://edofic.com/posts/2022-01-28-low-js/) ([HN](https://news.ycombinator.com/item?id=31249181)) ([Lobsters](https://lobste.rs/s/859lel/how_i_fell_love_with_low_js))
- [TodoMVC App Written in Vanilla JS in 2022](https://github.com/1Marc/todomvc-vanillajs-2022)
- [Modern JavaScript environments (2022)](https://twitter.com/wesbos/status/1524033215140446211)
- [Advanced JavaScript Objects](https://github.com/carltheperson/advanced-js-objects) - E-book entirely about JavaScript objects.
- [JavaScript Module Fragments Proposal](https://github.com/tc39/proposal-module-fragments) - Named, inline JS modules, which can be used for bundling multiple modules into a single JavaScript file.
- [Design Patters in JS for Humans](https://github.com/sohamkamani/javascript-design-patterns-for-humans)
- [devjar](https://devjar.vercel.app/) - Bundless runtime for your ESM JavaScript project in browser. ([Code](https://github.com/huozhi/devjar))
- [Fast way to do a JS socket API](https://twitter.com/jarredsumner/status/1529049475204849666)
- [Processing Arrays non-destructively: `for-of` vs. `.reduce()` vs. `.flatMap()` (2022)](https://2ality.com/2022/05/processing-arrays-non-destructively.html)
- [Now and .then: Debugging Async JavaScript - Jenn Creighton (2022)](https://www.youtube.com/watch?v=V-lu0YjJYdk)
- [JavaScript tree-sitter](https://github.com/tree-sitter/tree-sitter-javascript) - JavaScript and JSX grammar for tree-sitter.
- [Awesome JavaScript Interviews](https://github.com/rohan-paul/Awesome-JavaScript-Interviews) - Popular JavaScript / React / Node / Mongo stack Interview questions and their answers.
- [Monorepos in JavaScript and TypeScript (2022)](https://www.robinwieruch.de/javascript-monorepos/) ([HN](https://news.ycombinator.com/item?id=31596419))
- [Compiling a Subset of JavaScript to ARM Assembly in Haskell (2022)](https://www.micahcantor.com/blog/js-to-asm-in-hs/)
- [ECMAScript Explicit Resource Management](https://github.com/tc39/proposal-explicit-resource-management)
- [JavaScript: The New Toys Book](https://thenewtoys.dev/)
- [Server Rendering in JavaScript Series](https://dev.to/ryansolid/series/13734)
- [Ecma International approves ECMAScript 2022: What’s new?](https://2ality.com/2022/06/ecmascript-2022.html)
- [JS Functions - The Ultimate Guide](https://jsfunctions.io/)
- [JS Module Blocks](https://github.com/tc39/proposal-js-module-blocks) ([Web](https://tc39.es/proposal-js-module-blocks/))
- [Ask HN: Why is everything in JavaScript changing so fast? (2022)](https://news.ycombinator.com/item?id=31969958)
- [Stop Using JavaScript Objects (2022)](https://www.youtube.com/watch?v=hRSwSAr-gok)
- [Mastering Async/Await Book](https://asyncawait.net/)
- [Roll your own JavaScript runtime (2022)](https://deno.com/blog/roll-your-own-javascript-runtime) ([Code](https://github.com/denoland/roll-your-own-javascript-runtime))
- [Checking if a JavaScript native function is monkey patched (2022)](https://mmazzarolo.com/blog/2022-07-30-checking-if-a-javascript-native-function-was-monkey-patched/)
- [Optimizing for JavaScript is hard (2022)](https://jfmengels.net/optimizing-javascript-is-hard/) ([Lobsters](https://lobste.rs/s/sdoeyq/optimizing_for_javascript_is_hard))
