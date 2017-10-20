# Hammerspoon

Hammerspoon is an amazing open source app that lets you write some Lua code to interact with macOS on a kernel level. 

My config for it can be seen [here](https://github.com/nikitavoloboev/dotfiles/blob/master/hammerspoon/init.lua).

## call hammerspoon functions from Applescript
First activate applescript support : 

`hs.allowAppleScript(true)`

Then you can call functions from your lua code like this :

```Lua
tell application "Hammerspoon"
execute lua code "showClipBoardContent()"
end tell
```


In the above case it would call `showClipBoardContent()` function.