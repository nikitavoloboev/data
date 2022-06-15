---
title: Hammerspoon
---

# [Hammerspoon](https://github.com/Hammerspoon/hammerspoon)

## Code

```lua
-- Call Hammerspoon functions from AppleScript

-- First activate applescript support:
hs.allowAppleScript(true)

-- Then you can call functions from your lua code like this:
tell application "Hammerspoon"
execute lua code "showClipBoardContent()"
end tell

-- In the above case it would call showClipBoardContent() function
```

```lua
-- Bind function to hotkey
hs.hotkey.bind("ctrl", "return", function()
  hs.notify.new({title="Hammerspoon", informativeText="Hello World"}):send()
end)
```

```lua
-- URL handler alert

hs.urlevent.bind("someAlert", function(eventName, params)
    hs.alert.show("Hey there alert")
end)

-- After having this line in init.lua
-- you can then call it from the shell like so
open -g hammerspoon://someAlert
```

## Links

- [Spoons](http://www.hammerspoon.org/Spoons/) - Pure-Lua plugins for users to use in their Hammerspoon configs. ([Doc](https://github.com/Hammerspoon/hammerspoon/blob/master/SPOONS.md)) ([Code](https://github.com/Hammerspoon/Spoons))
- [Interesting HS config](https://github.com/S1ngS1ng/HammerSpoon)
- [Accessing Accessibility Objects with Hammerspoon](https://github.com/asmagill/hs._asm.axuielement)
- [HN: Hammerspoon (2019)](https://news.ycombinator.com/item?id=21801178)
- [Spacehammer](https://github.com/agzam/spacehammer) - Hammerspoon config inspired by Spacemacs.
- [Using Fennel to configure Hammerspon](https://lobste.rs/s/tkm8nh/using_fennel_configure_hammerspon)
- [Akinori's HS modules](https://github.com/knu/hs-knu) - Mainly for keyboard customization.
- [VimMode.spoon](https://github.com/dbalatero/VimMode.spoon) - Adds vim keybindings to all macOS inputs.
- [Lobsters: Hammerspoon (2021)](https://lobste.rs/s/ltu8du/hammerspoon)
- [HN: Hammerspoon (2021)](https://news.ycombinator.com/item?id=29533495)
