# [Karabiner](https://github.com/tekezo/Karabiner-Elements)
Karabiner is an absolutely amazing app that lets you remap keys at a very low level on macOS. 

For example you can make caps lock into an escape key when pressed once but if you do hold it, it becomes a [hyper key](http://brettterpstra.com/2017/06/15/a-hyper-key-with-karabiner-elements-full-instructions/). Hyper key means that a key now serves two purposes, once when pressed alone and once when held down. So for example for remapping caps lock, we can remap it to act as escape when pressed alone once but if we hold down on it it becomes ⌘ + ⌃ modifier key. So `caps lock + F` becomes ⌘ + ⌃ + F. And so on.

I however take this idea further and define these kind of hyper keys but for __every single key on my keyboard__. I like to call them _sticky keys_ and you can read more about it [here](./sticky-keys.md).

## My personal Karabiner setup
I share my Karabiner private.xml [here](https://github.com/nikitavoloboev/dotfiles/blob/master/karabiner/private.xml). It is approximately 12,000 lines long so might take a bit of time to load. However I edit it in vim and vim handles this big file very easily. I also use `_` to indicate jumping points in the config. So I can jump between various XML blocks by searching `_km def` for defining KM actions to map, `_alfred def` for mapping alfred triggers to run and `sticky ..` to jump between sticky key definitions. 

I document all the bindings I have made on my laptop as a mind map [here](https://my.mindnode.com/c7EmmKvaxCyCEuTzcpkGB4MGeLpWdR8nsJK4rjDh). It is pretty large but if you have MindNode, you can download it locally and it would be quite fast to use.

It documents both simultaneous key presses that is if I press _j and k together_ and my custom modifier keys (_sticky keys_).

## Mind Map key
Mind Maps are pretty awesome because I can literally map my modifier mappings and bindings as a QWERTY map :
![](https://i.imgur.com/4yk23Be.png)
I hope you get the idea for how powerful both Karabiner and this way of visualising bindings is. And you will make and share your own private.xml config file or .json file if you are using [Karabiner Elements](https://github.com/tekezo/Karabiner-Elements).

## Notes
- [interesting setup](https://github.com/dunkarooftop/thought/blob/master/keymaps.org)