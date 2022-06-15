---
title: Karabiner
---

# [Karabiner](https://github.com/tekezo/Karabiner-Elements)

Karabiner is life changing tool that [lets you remap keys at a low level on macOS](https://medium.com/@nikitavoloboev/karabiner-god-mode-7407a5ddc8f6).

I completely remapped my keyboard with it and every key on my keyboard is a custom modifier key that I can program to do what I want.

For example you can make caps lock into an escape key when pressed once but if you hold it, it becomes a hyper key. Hyper key means that a key now serves two purposes, once when pressed alone and once when held down. So for example for remapping caps lock, we can remap it to act as escape when pressed alone once but if we hold down on it it becomes ⌘ + ⌃ modifier key. So `caps lock + F` becomes ⌘ + ⌃ + F. And so on.

I take this idea further and define these kind of hyper keys but for **every single key on my keyboard**.

![](https://imgs.xkcd.com/comics/borrow_your_laptop_2x.png)

## My personal Karabiner setup

I generate [my Karabiner config](https://github.com/nikitavoloboev/dotfiles/blob/master/karabiner/karabiner.edn) with [Goku](https://github.com/yqrashawn/GokuRakuJoudo).

Will make a web explorer for Karabiner/Goku bindings soon. But briefly, I group keys by theme. `q` = cmd+shift. `w` = open apps. `e` = cmd. `r` = open more apps. `i` = insert symbols like & or |. `o` = Alfred searches. `a` = control. `s` = arrow keys, selecting text & more text helpers. `d` = `d+v` will do left mouse click, use that a lot. `f` = various helpers. `g` = move windows left/right/maximize & more. `hjkl` are no hyper keys as I press on them often for scrolling in vim/browser and need instant response, hyper keys have a little delay after press. `:` = caps lock key, so `:+w` will insert W. `z` = open chat apps, like `z+k` opens telegram. `x` = spotify helpers, like current song, search songs & more. `c` = helpers, use `c+s` often to clone current open GitHub URL & open it in VSCode. `v` = change volume, next song & more. `b` = open 'news' sites like Reddit/HN. `n` = Alfred searches, use `n+s` to trigger [web searches](https://github.com/nikitavoloboev/alfred-web-searches) often. `m` = search [Dash docsets](https://kapeli.com/dash) with Alfred. `.` = insert text fast like `console.log()` & more.

## Interesting configs

-  [Goku In The Wild](https://github.com/yqrashawn/GokuRakuJoudo/blob/master/in-the-wild.md)
-  [George Nica](https://github.com/kiinoda/goku)

## Notes

-  I can embed simultaneous key mappings inside sticky key definitions.

## Code

### 3 finger trackpad actions

```edn
{:des "hold three fingers on trackpad & press keys" :rules [
   [:condi ["multitouch_extension_finger_count_total" 3]]
   [:f :button2] ;
   :v [:button1 :!Cv]]
;]}
```

## Links

-  [Karabiner God Mode (2018)](https://medium.com/@nikitavoloboev/karabiner-god-mode-7407a5ddc8f6) - How to use Karabiner to take your use of mac to the next level.
-  [KE complex modifications](https://github.com/pqrs-org/KE-complex_modifications) - Has [website](https://pqrs.org/osx/karabiner/complex_modifications/) too.
-  [My detailed post on Alfred forum mentioning how I use Karabiner](https://www.alfredforum.com/topic/10673-how-to-make-the-alfred-search-window-a-frontmost-app/?do=findComment&comment=57212)
-  [Interesting setup](https://github.com/dunkarooftop/thought/blob/master/keymaps.org)
-  [Karabiner Elements documentation](https://qiita.com/s-show/items/a1fd228b04801477729c) (in Japanese)
-  [Tekezo's launcher mode](https://github.com/pqrs-org/KE-complex_modifications/blob/2348fb5ae3f0b04cea16b6b07ff6cf18e58885fb/docs/json/personal_tekezo_launcher_mode_v4.json)
-  [Karabiner.json reference manual](https://pqrs.org/osx/karabiner/json.html)
-  [Karabiner JSON spec](https://pqrs.org/osx/karabiner/json.html)
-  [Write Karabiner config in YAML and then convert to JSON](https://github.com/15cm/dotfiles/tree/master/.config/karabiner)
-  [jKeyboard](https://github.com/jhelvy/jKeyboard)
-  [shell command JSON](https://pqrs.org/osx/karabiner/json.html#typical-complex_modifications-examples-open-alfred-when-escape-is-held-down)
-  [Karaconv](https://github.com/durka/karaconv) - Converter from Karabiner to Karabiner-Elements text configuration format.
-  [Cursor keys belong at the center of your keyboard](http://tonsky.me/blog/cursor-keys/)
-  [Sticky shift key example](https://github.com/rcmdnk/KE-complex_modifications/blob/master/docs/json/sticky.json)
-  [Karabiner KeyCodes and Modifiers](https://github.com/tekezo/Karabiner-Elements/issues/925)
-  [New Hyper Key](https://josh.blog/2017/07/new-hyper-key)
-  [Karabiner Elements profile switcher (Alfred Workflow)](https://github.com/awinecki/karabiner-elements-profile-switcher)
-  [All about macOS event observation (2019)](https://docs.google.com/presentation/d/1nEaiPUduh1vjks0rDVRTcJaEULbSWWh1tVdG2HF_XSU/edit#slide=id.g5b38b1767c_0_2)
-  [Customize Karabiner With Goku (2019)](https://johnlindquist.com/customize-karabiner-with-goku)
-  [Karabiner Elements Keycodes](https://github.com/aerobounce/karabiner-elements-keycodes)
-  [Hacking your keyboard with karabiner (2019)](https://blog.kaush.co/2019/12/25/hacking-your-keyboard/) ([HN](https://news.ycombinator.com/item?id=21891082))
-  [Merge karabiner.json with complex modification rules](https://gist.github.com/narze/527ac6321c24cfde71bc3b30b7c078f3)
-  [Karabiner Complex Rules Generator](https://genesy.github.io/karabiner-complex-rules-generator/) ([Code](https://github.com/genesy/karabiner-complex-rules-generator))
-  [Karabiner DriverKit VirtualHIDDevice](https://github.com/pqrs-org/Karabiner-DriverKit-VirtualHIDDevice)
-  [Optimize Your Keyboard (2020)](https://www.pscp.tv/w/1vOxworogovxB)
-  [Custom Key Bindings with Karabiner (2020)](https://zacjones.io/custom-key-bindings)
-  [Master of keyboard is master of automation (Key remapper implementation in Swift)](https://github.com/creasty/Keyboard)
-  [Remapping keys on macOS](https://blog.codefront.net/2020/06/24/remapping-keys-on-macos)
-  [Goku Configs on GitHub](https://github.com/search?l=&o=desc&q=extension%3A.edn+filename%3Akarabiner.edn&s=&type=Code)
-  [Putting your Keyboard on Steroids with Karabiner Elements (2020)](https://dev.to/swyx/notes-on-karabiner-elements-from-john-lindquist-4cmo)
-  [Karabiner Docs](https://karabiner-elements.pqrs.org/docs/)
-  [Compose Key On macOS](https://github.com/Granitosaurus/macos-compose) ([HN](https://news.ycombinator.com/item?id=24553013))
-  [Remap Keyboard Shortcuts Karabiner Elements (2020)](https://www.youtube.com/watch?v=vysHEYTp0H4)
-  [Новый мак, .dotfiles #4, Karabiner, keyboard maestro, Alfred (2022)](https://www.youtube.com/watch?v=wSbZoW0cAww)
