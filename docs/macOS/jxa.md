---
title: JavaScript for Automation
---

# [JavaScript for Automation](https://developer.apple.com/library/content/releasenotes/InterapplicationCommunication/RN-JavaScriptForAutomation/Articles/Introduction.html)

## Code

```javascript
// Get URL of front most tab
Application("Google Chrome Canary").windows[0].activeTab.url();
```

```javascript
// Activate app
it = Application("Safari");
it.activate();
```

```javascript
// Minimize all Safari windows but the currently focused one
safari = Application("Safari");

// Front most window has index 0
for (var i = 1; i < safari.windows.length; i++) {
  var win = safari.windows[i];
  if (win.miniaturized() === false) win.miniaturized = true;
}
```

```javascript
// Get selected text in Chrome Canary
Application("Google Chrome Canary").windows[0].activeTab.execute({
  javascript: "window.getSelection().toString()",
});
```

```javascript
// Return current active line in TaskPaper.
(() => {
  "use strict";

  const main = () => {
    const tp3Context = (editor, options) => {
      const main = () =>
        unlines(
          concatMap((x) => {
            const txt = x[options.textProperty];
            return options.skipBlankLines && 0 === txt.length ? [] : [txt];
          }, editor.selection.selectedItems)
        );

      // GENERIC FUNCTIONS FOR TP3 CONTEXT ----------
      // https://github.com/RobTrew/prelude-jxa

      // concatMap :: (a -> [b]) -> [a] -> [b]
      const concatMap = (f, xs) => xs.reduce((a, x) => a.concat(f(x)), []);

      // unlines :: [String] -> String
      const unlines = (xs) => xs.join("\n");

      // TP3 MAIN ---
      return main();
    };

    const ds = Application("TaskPaper").documents,
      lrResult = bindLR(
        ds.length > 0 ? Right(ds.at(0)) : Left("No TaskPaper documents open"),
        (d) =>
          Right(
            d.evaluate({
              script: tp3Context.toString(),
              withOptions: {
                skipBlankLines: true,
                textProperty: "bodyContentString", // or 'bodyContentString'
              },
            })
          )
      );
    return lrResult.Left || lrResult.Right;
  };

  // GENERIC FUNCTIONS FOR JXA CONTEXT ------------------
  // https://github.com/RobTrew/prelude-jxa

  // Left :: a -> Either a b
  const Left = (x) => ({
    type: "Either",
    Left: x,
  });

  // Right :: b -> Either a b
  const Right = (x) => ({
    type: "Either",
    Right: x,
  });

  // bindLR (>>=) :: Either a -> (a -> Either b) -> Either b
  const bindLR = (m, mf) => (undefined !== m.Left ? m : mf(m.Right));

  // MAIN ----
  return main();
})();
```

## Links

- [JavaScript for Automation (JXA) Book](https://bru6.de/jxa/)
- [JXA Cookbook](https://github.com/JXA-Cookbook/JXA-Cookbook/wiki/Foreword)
- [JXA Resources](https://apple-dev.groups.io/g/jxa/wiki/JXA-Resources)
- [Get frontmost tab’s URL and title of various browsers](https://www.alfredforum.com/topic/2013-how-to-get-frontmost-tab%E2%80%99s-url-and-title-of-various-browsers/)
- [JXA Prelude](https://github.com/RobTrew/prelude-jxa) - Generic functions for macOS scripting with JavaScript for Automation.
- [Getting Started with JavaScript for Automation on Yosemite](https://www.macstories.net/tutorials/getting-started-with-javascript-for-automation-on-yosemite/)
- [How To Use “VS Code” for JXA, with Keyboard Maestro (2019)](https://forum.keyboardmaestro.com/t/how-to-use-vs-code-for-jxa-with-keyboard-maestro/15281)
- [How to find JXA actions for apps](https://forum.keyboardmaestro.com/t/how-to-use-vs-code-for-jxa-with-keyboard-maestro/15281/26)
- [The unexpected return of JavaScript for Automation (2021)](https://scriptingosx.com/2021/11/the-unexpected-return-of-javascript-for-automation/) ([HN](https://news.ycombinator.com/item?id=29327349))
- [JXA TS/Node packages](https://github.com/JXA-userland/JXA)
- [JXA’s Parenthesis Paradox (2022)](https://leancrew.com/all-this/2022/05/jxas-parenthesis-paradox/) ([HN](https://news.ycombinator.com/item?id=31579435))
