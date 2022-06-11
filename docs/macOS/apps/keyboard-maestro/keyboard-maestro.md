---
title: Keyboard Maestro
---

# [Keyboard Maestro](https://www.keyboardmaestro.com/main/)

Keyboard Maestro is essentially an IDE for automation. You create macros of actions that you can then easily call from Karabiner.

It has a [wonderful community](https://forum.keyboardmaestro.com/) that is happy to help with whatever you are trying to achieve.

I share [all the macros I use](km-macros.md) with the app.

## Notes

- It is always better to run compiled AppleScripts.
- [Assert action](https://forum.keyboardmaestro.com/t/assert-action/8374). "If not condition, then fail". Useful for checking if you got right kind of data.
- You need to [set ENV_PATH variable](https://forum.keyboardmaestro.com/t/create-a-path-environment-variable-for-keyboard-maestro-and-add-usr-local-bin-to-the-default-path/10064) to make sure all your shell tools work in KM.
- You can read the full text of any error message in the Engine.log file (Help ➤ Open Logs Folder), and make sure (at least when testing or having problems) that the Execute Shell Script is configured to display the result in a window so you can see any error message.
- [Keyboard Maestro converts any variables that start with ENV\_ into environment variables when it launches sub-processes.](https://forum.keyboardmaestro.com/t/how-to-sent-env-in-keyboard-maestro/11947/2)
- `defaults write com.stairways.keyboardmaestro.editor MouseGetCountdown -int 3` will change the [timer for getting coordinates from 5 secs to 3](https://forum.keyboardmaestro.com/t/is-there-a-way-to-set-default-waiting-time-to-get-coordinates-from-5-seconds-to-3-seconds/14980).
- Can dismiss palettes with `return` key.
- Useful to sort macros by `Date modified`. Applies to most other tools too like Finder.
- [The trick is to work incrementally. Find out how to successfully evaluate and display something atomic. And then gradually elaborate, layer upon layer.](https://forum.keyboardmaestro.com/t/having-trouble-with-execute-shell-script/19628/9)

## Code

### System Clipboard variable

`%SystemClipboard%`

### Var as token

`%Variable%%{left}{key:⌃L}`

### Var for shell

`"$KMVAR_"{left}{key:⌃L}`

### Variable

`%Variable%{^}%`

## Links

- [How I manage my huge KM library](https://forum.keyboardmaestro.com/t/notation-i-use-to-manage-my-macros/8907)
- [KM forum](https://forum.keyboardmaestro.com/latest) - Great place to discuss all things related to KM.
- [Best Macro List](https://forum.keyboardmaestro.com/t/best-macro-list/4118)
- [Automating KM editor](https://forum.keyboardmaestro.com/t/automating-the-keyboard-maestro-editor/4184/31)
- [6 months in, what I wish I knew on day 1 with KM](https://forum.keyboardmaestro.com/t/6-months-in-what-i-wish-i-knew-on-day-1-with-keyboard-maestro/4949)
- [Exporting KM Macros with KM Macro](http://chauncey.io/projects/keyboard-maestro-export-macros/)
- [Little Approach I use to quickly prototype new macros](https://forum.keyboardmaestro.com/t/little-approach-i-use-to-quickly-prototype-new-macros/8091)
- [Simplicity vs Complexity](https://forum.keyboardmaestro.com/t/simplicity-vs-complexity/11259)
- [KM Wiki](https://wiki.keyboardmaestro.com/Home_Page) ([Quick Start](https://wiki.keyboardmaestro.com/manual/Quick_Start))
- [Keyboard Maestro Course](https://www.youtube.com/playlist?list=PLsFZ0c2Wsoy8U8NA1C6Bii5ylV4sFyIzt)
- [From TextExpander to Keyboard Maestro… again (2021)](https://leancrew.com/all-this/2021/07/from-textexpander-to-keyboard-maestro-again/)
- [Archiving and version control for Keyboard Maestro (2021)](https://robservatory.com/archiving-and-version-control-for-keyboard-maestro/)
- [Design Maestro](https://dmaestro.webflow.io/) - Automate your design workflow.
- [Alfred - Conductor for Keyboard Maestro](https://www.thoughtasylum.com/alfred/alfred_conductor_for_keyboard_maestro/)
