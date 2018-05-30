# [Karabiner](https://github.com/tekezo/Karabiner-Elements)
Karabiner is an absolutely amazing app that lets you remap keys at a very low level on macOS.

I have completely remapped my keyboard with it and every key on my keyboard is a custom modifier key that I can program to do what I want.

For example you can make caps lock into an escape key when pressed once but if you hold it, it becomes a [hyper key](http://brettterpstra.com/2017/06/15/a-hyper-key-with-karabiner-elements-full-instructions/). Hyper key means that a key now serves two purposes, once when pressed alone and once when held down. So for example for remapping caps lock, we can remap it to act as escape when pressed alone once but if we hold down on it it becomes ⌘ + ⌃ modifier key. So `caps lock + F` becomes ⌘ + ⌃ + F. And so on.

I take this idea further and define these kind of hyper keys but for __every single key on my keyboard__. I like to call them _sticky keys_.

## Sticky keys
[This snippet](https://gist.github.com/b591b290c6a55ac47b19158c721415a4) makes `p` key into a modifier key. This only works for old Karabiner on El Captain. I am still not certain how it will look like in Sierra.

You can probably use [this](https://github.com/tekezo/Karabiner-Elements/issues/926) to achieve this behaviour on Sierra but I have not tested it.

## My personal Karabiner setup
I share my Karabiner private.xml [here](https://github.com/nikitavoloboev/dotfiles/blob/master/karabiner/private.xml). It is approximately 14,000 lines long so might take a bit of time to load. However I edit it in vim and vim handles this big file very easily. I comment out the necessary sections and jump to them with searches like `km def` for defining KM actions to map, `alfred def` for mapping alfred triggers to run and `sticky ..` to jump between sticky key definitions.

## Notes
- I can embed simultaneous key mappings inside sticky key definitions.

## Links
- [KE complex modifications](https://github.com/pqrs-org/KE-complex_modifications) - Has [website](https://pqrs.org/osx/karabiner/complex_modifications/) too.
- [My detailed post on Alfred forum mentioning how I use Karabiner](https://www.alfredforum.com/topic/10673-how-to-make-the-alfred-search-window-a-frontmost-app/?do=findComment&comment=57212)
- [Interesting setup](https://github.com/dunkarooftop/thought/blob/master/keymaps.org)
- [Karabiner Elements documentation](https://qiita.com/s-show/items/a1fd228b04801477729c) (in Japanese)
- [Tekezo's launcher mode](https://github.com/pqrs-org/KE-complex_modifications/pull/206/files)
- [Karabiner.json reference manual](https://pqrs.org/osx/karabiner/json.html)
- [Karabiner JSON spec](https://pqrs.org/osx/karabiner/json.html)
- [Write Karabiner config in YAML and then convert to JSON](https://github.com/15cm/dotfiles/tree/master/.config/karabiner)
- [jKeyboard](https://github.com/jhelvy/jKeyboard)
- [shell command JSON](https://pqrs.org/osx/karabiner/json.html#typical-complex_modifications-examples-open-alfred-when-escape-is-held-down)
- [Karaconv](https://github.com/durka/karaconv) - Converter from Karabiner to Karabiner-Elements text configuration format.
- [Cursor keys belong at the center of your keyboard](http://tonsky.me/blog/cursor-keys/)

## Interesting setups
- [A. King](https://github.com/akork/karabiner/blob/master/karabiner.json)