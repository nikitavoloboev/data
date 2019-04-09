# Keyboard Maestro macros I use

## Explanation

One of my favorite applications that I use to significantly ease my time using my mac is [Keyboard Maestro](https://www.keyboardmaestro.com/main/). This app is essentially an IDE for automation.

It's a wonderful tool that allows you to create very powerful macros that can range from simply pasting some text, to moving the mouse pointer to some location, to doing various calculations with the input you provide and lot lot more.

I wrote about how I manage my huge macro library [here](https://forum.keyboardmaestro.com/t/notation-i-use-to-manage-my-macros/8907).

All global macros will have no trigger as they are called with AppleScript from [Karabiner](../karabiner/karabiner.md).

A lot of these macros were made with great help from the [Keyboard Maestro forum community](https://forum.keyboardmaestro.com/latest).

Just to emphasize how much KM has impacted my life. My most used macro from all these macros I share is macro to open Safari browser which I've ran some 55,851 times since I made it.

![](https://i.imgur.com/48OtscL.png)

I also love using [Alfred Maestro](https://github.com/iansinnott/alfred-maestro) workflow that lets me search through my entire KM library of actions to activate.

## Install all the macros I use

I share all the macros I use as [macro sync](https://www.keyboardmaestro.com/documentation/6/macrosyncing.html) file [you can download](https://www.dropbox.com/s/mh62g5jrtu70161/Keyboard%20Maestro%20Macros.kmsync?dl=1).

It includes over 1400 macros I use for various purposes to automate my workflow.

![](https://i.imgur.com/a3kUsxS.png)

App specific macros are called with a hotkey bound to the app the macro is running in. And global macros are called with [Karabiner](../karabiner/karabiner.md) directly. Technically even app specific macros are called with Karabiner as my `control` key is `a` key and `command` key is my `e` key.

You can double click the downloaded macro sync file to upload all the 1400+ macros into KM. Be aware that it will overwrite your existing macros so back those up beforehand!

Which macros I call the most can be gleaned from by looking at my [Karabiner config](https://github.com/nikitavoloboev/dotfiles/blob/master/karabiner/karabiner.edn).

Previously I shared these KM macros on [Gumroad as separate exported folders](https://gumroad.com/l/kmmacros).

## KM Plugins

It is advised you install these KM plugins as they make using KM editor more pleasant.

- [KMFAM - Favorite actions and macros](https://forum.keyboardmaestro.com/t/macro-kmfam-favorite-actions-and-macros/4854) - This plugin lets you save complex and non complex actions and macros. You can then search over these actions and macros to quickly add them.

## Macros

Macros are little _KM scripts_ that contain a series of actions. The macros can then be executed from a trigger, usually a hotkey.

In my case all the global macros I share (macros that can run no matter what application is currently active) have no trigger. This is because I can call these macros from Karabiner using this AppleScript code:

```applescript
tell application "Keyboard Maestro Engine"
do script "open: Safari"
end tell
```

Where `open: Safari` is a macro name.

Non global macros that are binded to an app will have a hotkey however. One thing to note is that the hotkey triggers are chosen with respect to my custom keyboard layout.

I have binded my control key to `a` key with [Karabiner](../karabiner/karabiner.md). My command key is `e` key and Command + Shift modifier is `q` key. So pressing `e` key, holding it briefly and after pressing `f` will trigger `⌘ + f` hotkey. Therefore some bindings will only make sense in context of my own layout. For example `⌃ + w` is easier to reach for me then `⌘ + d`. Because `a + w` is nicer to press then `e + d` on my keyboard. [Karabiner](../karabiner/karabiner.md) is one powerful abstraction that makes managing 1000+ macros easy.

## Palettes

I love using [KM Palettes](https://wiki.keyboardmaestro.com/manual/Palettes). A KM palette is a macro group that contains some actions inside. You give the palette a hotkey trigger and when you press the hotkey, a palette will show up, like this one:

![](https://i.imgur.com/b6KB2zM.png)

The letters to the right will then activate the actions inside the palette. Similar to above macro groups, I have many Global palettes that I activate from Karabiner. I binded my `Space` key to be ⇧ + ⌃ + ⌥ modifier. This way I can open various palettes by holding `Space` and pressing different keys on my keyboard.
