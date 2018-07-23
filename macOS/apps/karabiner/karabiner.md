# [Karabiner](https://github.com/tekezo/Karabiner-Elements)
Karabiner is an absolutely amazing app that lets you remap keys at a very low level on macOS.

I have completely remapped my keyboard with it and every key on my keyboard is a custom modifier key that I can program to do what I want.

For example you can make caps lock into an escape key when pressed once but if you hold it, it becomes a [hyper key](http://brettterpstra.com/2017/06/15/a-hyper-key-with-karabiner-elements-full-instructions/). Hyper key means that a key now serves two purposes, once when pressed alone and once when held down. So for example for remapping caps lock, we can remap it to act as escape when pressed alone once but if we hold down on it it becomes ⌘ + ⌃ modifier key. So `caps lock + F` becomes ⌘ + ⌃ + F. And so on.

I take this idea further and define these kind of hyper keys but for __every single key on my keyboard__. I like to call them _sticky keys_.

## My personal Karabiner setup
I generate my Karabiner config with a ruby script [here](https://github.com/nikitavoloboev/karabiner-generator#readme). It contains all the necessary functions needed to make sticky key and simultaneous key definitions work.

## Notes
- I can embed simultaneous key mappings inside sticky key definitions.

## Links
- [KE complex modifications](https://github.com/pqrs-org/KE-complex_modifications) - Has [website](https://pqrs.org/osx/karabiner/complex_modifications/) too.
- [My detailed post on Alfred forum mentioning how I use Karabiner](https://www.alfredforum.com/topic/10673-how-to-make-the-alfred-search-window-a-frontmost-app/?do=findComment&comment=57212)
- [Interesting setup](https://github.com/dunkarooftop/thought/blob/master/keymaps.org)
- [Karabiner Elements documentation](https://qiita.com/s-show/items/a1fd228b04801477729c) (in Japanese)
- [Tekezo's launcher mode](https://github.com/pqrs-org/KE-complex_modifications/blob/2348fb5ae3f0b04cea16b6b07ff6cf18e58885fb/docs/json/personal_tekezo_launcher_mode_v4.json)
- [Karabiner.json reference manual](https://pqrs.org/osx/karabiner/json.html)
- [Karabiner JSON spec](https://pqrs.org/osx/karabiner/json.html)
- [Write Karabiner config in YAML and then convert to JSON](https://github.com/15cm/dotfiles/tree/master/.config/karabiner)
- [jKeyboard](https://github.com/jhelvy/jKeyboard)
- [shell command JSON](https://pqrs.org/osx/karabiner/json.html#typical-complex_modifications-examples-open-alfred-when-escape-is-held-down)
- [Karaconv](https://github.com/durka/karaconv) - Converter from Karabiner to Karabiner-Elements text configuration format.
- [Cursor keys belong at the center of your keyboard](http://tonsky.me/blog/cursor-keys/)
- [Sticky shift key example](https://github.com/rcmdnk/KE-complex_modifications/blob/master/docs/json/sticky.json)
- [Karabiner KeyCodes and Modifiers](https://github.com/tekezo/Karabiner-Elements/issues/925)
- [New Hyper Key](https://josh.blog/2017/07/new-hyper-key)

## Interesting setups
- [A. King](https://github.com/akork/karabiner/blob/master/karabiner.json)