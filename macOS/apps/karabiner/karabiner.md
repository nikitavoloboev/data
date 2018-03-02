# [Karabiner](https://github.com/tekezo/Karabiner-Elements)
Karabiner is an absolutely amazing app that lets you remap keys at a very low level on macOS.

I have completely remapped my keyboard with it and every key on my keyboard is a custom modifier key that I can program to do what I want.

For example you can make caps lock into an escape key when pressed once but if you do hold it, it becomes a [hyper key](http://brettterpstra.com/2017/06/15/a-hyper-key-with-karabiner-elements-full-instructions/). Hyper key means that a key now serves two purposes, once when pressed alone and once when held down. So for example for remapping caps lock, we can remap it to act as escape when pressed alone once but if we hold down on it it becomes ⌘ + ⌃ modifier key. So `caps lock + F` becomes ⌘ + ⌃ + F. And so on.

I however take this idea further and define these kind of hyper keys but for __every single key on my keyboard__. I like to call them _sticky keys_ and you can read more about it [here](sticky-keys.md).

## My personal Karabiner setup
I share my Karabiner private.xml [here](https://github.com/nikitavoloboev/dotfiles/blob/master/karabiner/private.xml). It is approximately 12,000 lines long so might take a bit of time to load. However I edit it in vim and vim handles this big file very easily. I also use `_` to indicate jumping points in the config. So I can jump between various XML blocks by searching `_km def` for defining KM actions to map, `_alfred def` for mapping alfred triggers to run and `sticky ..` to jump between sticky key definitions.

## Notes
- I can embed simultaneous key mappings inside sticky key definitions.

## Links
- [My detailed post on Alfred forum mentioning how I use Karabiner](https://www.alfredforum.com/topic/10673-how-to-make-the-alfred-search-window-a-frontmost-app/?do=findComment&comment=57212)
- [Interesting setup](https://github.com/dunkarooftop/thought/blob/master/keymaps.org)
- [Karabiner Elements documentation](https://qiita.com/s-show/items/a1fd228b04801477729c) (in Japanese)
- [Tekezo's launcher mode](https://github.com/pqrs-org/KE-complex_modifications/pull/206/files)
- [Karabiner.json reference manual](https://pqrs.org/osx/karabiner/json.html)
- [Karabiner JSON spec](https://pqrs.org/osx/karabiner/json.html)
- [Write Karabiner config in YAML and then convert to JSON](https://github.com/15cm/dotfiles/tree/master/.config/karabiner)