---
title: Chrome DevTools
---

# [Chrome DevTools](https://developer.chrome.com/devtools)

## Notes

- `$_` holds the previous result in the console. So you can do `$_.something(..)` to access the returned object items.
- [In DevTools hold shift while hovering over a request and it will highlight the initiator in green and dependencies in red](https://twitter.com/addyosmani/status/1260479896888975362)
- Pressing Ctrl + Space in Dev Tools Network filter search will give useful filters for requests. In Console it will give autosuggestions.
- [Can use Chrome DevTools in iOS Safari 15 via the JSBox app](https://twitter.com/Baconbrix/status/1441840354563530756)
- [Chrome Dev Tools can visualize backend performance with Server Timing](https://twitter.com/addyosmani/status/1445644998477815808)
- [Can add custom headers to requests in the Network tab. Right-click the columns bar > Response Headers > Manage Header Columns.](https://twitter.com/dhh/status/1445036316023005195)
- [ChromeDevTools has CSS Shadow Editor](https://twitter.com/addyosmani/status/1447456939466706946)
- [ChromeDevTools has CSS Animations inspector](https://twitter.com/addyosmani/status/1447079452987387905)
- [Can hide scripts/assets loaded by extensions from your network tab, put "-scheme:chrome-extension" in the filter input box](https://twitter.com/threepointone/status/1446064032407080966)
- [You can dynamically insert console.logs without touching your code. Just right-click on the line and choose "logpoint", enter the variables you want to log and hit Enter.](https://twitter.com/marvinhagemeist/status/1527356830757933058)
- [Holding the shift key while hovering over a request on Chrome DevTools will highlight the initiator in green and dependencies in red.](https://twitter.com/guaca/status/1544967782206431234)
- [To avoid seeing errors from extensions in Chrome DevTools, check 'Selected context only'](https://twitter.com/sw12/status/1545475163201978368)

## Links

- [What's new in Chrome DevTools (Google I/O '18)](https://www.youtube.com/watch?v=mfuE53x4b3k)
- [Chrome DevTools Masterclass (2016)](https://www.youtube.com/watch?v=KykP5Z5E4kA)
- [chromedp](https://github.com/chromedp/chromedp) - Faster, simpler way to drive browsers supporting the Chrome DevTools Protocol. ([Examples](https://github.com/chromedp/examples))
- [Chrome DevTools Protocol](https://github.com/ChromeDevTools/devtools-protocol)
- [Make your own DevTools](https://kentcdodds.com/blog/make-your-own-dev-tools)
- [Rod](https://github.com/ysmood/rod) - High-level Chrome Devtools controller directly based on Chrome DevTools Protocol.
- [Chrome DevTools UI](https://github.com/ChromeDevTools/devtools-frontend)
- [DevTools tips (2020)](https://twitter.com/brian_d_vaughn/status/1250659369496145921)
- [Hidden Features of Chrome DevTools (2020)](https://martinheinz.dev/blog/33)
- [How to Debug JavaScript Apps with Chrome DevTools (2020)](https://blog.asayer.io/how-to-debug-javascript-apps-with-chrome-devtools)
- [Emulate vision deficiencies in DevTools (2020)](https://addyosmani.com/blog/emulate-vision-deficiencies-devtools/)
- [Awesome Chrome DevTools](https://github.com/ChromeDevTools/awesome-chrome-devtools)
- [Introduction to Dev Tools](https://github.com/jkup/mastering-chrome-devtools)
- [Abusing the Chromium Devtools Scope Pane (2021)](https://medium.com/@weizmangal/javascript-anti-debugging-some-next-level-sh-t-part-2-abusing-chromium-devtools-scope-pane-b2796c00331d) ([HN](https://news.ycombinator.com/item?id=28422781))
- [How to debug event listeners with your browser's developer tools (2021)](https://gomakethings.com/how-to-debug-event-listeners-with-your-browsers-developer-tools/)
- [Chrome DevTools Protocol Docs](https://chromedevtools.github.io/devtools-protocol/)
- [Wayang](https://github.com/go-rod/wayang) - High-level Devtools driver directly based on DevTools Protocol.
- [chobitsu](https://github.com/liriliri/chobitsu) - Chrome devtools protocol JavaScript implementation.
- [DevTools Tips](https://devtoolstips.org/) - Collection of useful cross-browser DevTools tips. ([Code](https://github.com/captainbrosset/devtools-tips))
- [Replay](https://github.com/puppeteer/replay) - Library that provides an API to replay and stringify recordings created using Chrome DevTools Recorder.
- [Detecting memory leaks in chrome dev tools](https://twitter.com/mgechev/status/1519534448227803138)
- [Puppeteer Heap Snapshot](https://github.com/adriancooney/puppeteer-heap-snapshot) - API and CLI tool to fetch and query Chrome DevTools heap snapshots.
- [Chrome DevTools Protocol Proxy](https://github.com/wendigo/chrome-protocol-proxy) - Intelligent proxy for debugging purposes.
- [Investigate Animation Performance with DevTools (2019)](https://calibreapp.com/blog/investigate-animation-performance-with-devtools)
- [Can I DevTools?](https://www.canidev.tools/)
