# Keyboard Maestro macros I use

## Install

You can install all the macros I use by downloading them [from here](https://gumroad.com/l/kmmacros) for a small price. The folder is always updated with the latest macros I use.

Inside the downloaded `KM` folder, you can double click in Finder on any of the macros that you want to install to KM.

## Explanation

One of my favorite applications that I use to significantly ease my time using my mac is [Keyboard Maestro](https://www.keyboardmaestro.com/main/). This app is essentially an IDE for automation.

It's a wonderful tool that allows you to create very powerful macros that can range from simply pasting some text, to moving the mouse pointer to some location, to doing various calculations with the input you provide and lot lot more.

I wrote about how I manage my huge macro library [here](https://forum.keyboardmaestro.com/t/notation-i-use-to-manage-my-macros/8907).

All global macros will have no trigger as they are called with AppleScript from [Karabiner](../karabiner/karabiner.md).

A lot of these macros were made with great help from the [Keyboard Maestro forum community](https://forum.keyboardmaestro.com/latest).

Just to emphasize how much KM has impacted my life. My most used macro from all these macros I share is macro to open Safari browser which I've ran [some 117,091 times](https://i.imgur.com/09KfToP.png) since I made it.

I also love using [Alfred Maestro](https://github.com/iansinnott/alfred-maestro) workflow that lets me search through my entire KM library of actions to activate.

## KM Plugins

It is advised you install these KM plugins as they make using KM editor more pleasant.

- [KMFAM - Favorite actions and macros](https://forum.keyboardmaestro.com/t/macro-kmfam-favorite-actions-and-macros/4854) - This plugin lets you save complex and non complex actions and macros. You can then search over these actions and macros to quickly add them.

## Macros

Macros are little _KM scripts_ that contain a series of actions. The macros can then be executed from a trigger, usually a hotkey.

In my case all the global macros I share (macros that can run no matter what application is currently active) have no trigger. This is because I can call these macros from Karabiner using this AppleScript code:

```applescript
tell application "Keyboard Maestro Engine"
do script "g: Kill All Macros"
end tell
```

Where `g: Kill All Macros` is a macro name. Where `g:` is a prefix that stands for `global`.

Non global macros that are binded to an app will have a hotkey however. One thing to note is that the hotkey triggers are chosen with respect to my custom keyboard layout.

I have binded my control key to A key with [Karabiner](../karabiner/karabiner.md). My command key is E key and Command + Shift modifier is Q key. So pressing `E` key, holding it and after pressing `F` will trigger `⌘ + F` hotkey. Therefore some bindings will only make sense in context of my own layout. For example `⌃ + W` is easier to reach for me then `⌘ + D`. Because `a + w` is nicer to press then `e + d` on my keyboard. [Karabiner](../karabiner/karabiner.md) is one powerful abstraction that makes managing 1000+ macros easy.

### Global Macro groups

Global macro groups can be found inside `macros/global`.

The macro groups have a prefix like `g:` that is used to give a theme and context to the macros inside the group. All the macros inside the group start with that prefix. The global groups included are:

- `g: General` - All my global macros that have no category.
- `open: Applications` - Big list of macros that open applications. The cool thing about them is that they will open applications when triggered but if you are already in the application and run the macro it will bring you to last opened app.
- `w: Open websites (Safari)` - Open websites in Safari browser and switch to the websites if the URL is already open somewhere. (Credit to [Christopher Stone](https://github.com/ccstone))
- `c: Open websites (Chrome)` - Open websites in Chrome browser and switch to the websites if the URL is already open somewhere. (Credit to [Christopher Stone](https://github.com/ccstone))
- `g: Text manipulation` - Small macros that do text manipulation. Like uppercasing selected text or lowercasing it.

### Application macro groups

Application specific macro groups can be found inside `macros/app`.

There are many macro groups inside the folder. These macros will only run in the application as specified by the name of the macro group.

### Programming language macro groups

Programming language macro groups can be found `macros/programming`.

This folder only contains one macro group `Go: Insert`. Because currently I mostly program in this language. The macro group contains some macros that insert some code I often tend to type in Go that I decided to automate. In my Karabiner, I have binded my `.` key to be a programming language specific modifier. When in `Go` Karabiner mode, `. + ..` will run various macros from this group. An example macro may type `log.println()` when I press `. + g` and move my caret between the ().

## Palettes

I love using [KM Palettes](https://wiki.keyboardmaestro.com/manual/Palettes). A KM palette is a macro group that contains some actions inside. You give the palette a hotkey trigger and when you press the hotkey, a palette will show up, like this one:

![](https://i.imgur.com/b6KB2zM.png)

The letters to the right will then activate the actions inside the palette. Similar to above macro groups, I have many Global palettes that I activate from Karabiner. I binded my `Space` key to be ⇧ + ⌃ + ⌥ modifier. This way I can open various palettes by holding `Space` and pressing different keys on my keyboard.

### Global palettes

Global palettes can be found inside `palettes/global`.

All the palette groups start with `t:` prefix which indicates they are Global palettes. `p:` prefix would be more appropriate but the key is not placed as nicely as `t` on my keyboard and is harder to reach. The macros inside the palette don't start with a prefix.

One of the palettes included in there is `t: KM groups (Palettes)`. This palette when activated looks like this:

![](https://i.imgur.com/z2TWtio.png)

It lists the global palettes I have made. When I activate any of the options, it quickly jumps to this palette group in the KM Editor so I can quickly edit and change it to my will. You can extend the groups to suit your own palettes by simply changing the `macroGroupName` variable in the code of the macros.

### Application specific palettes

Application specific palettes can be found inside `palettes/app`.

These palettes all start with either `tf:` or `ta:`. The reason is that I have binded `Space + a` and `Space + f` to activate application specific palettes. I thus name the palettes appropriately to distinguish them from my `Global` palettes.
