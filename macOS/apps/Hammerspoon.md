# Hammerspoon

[Hammerspoon](https://github.com/Hammerspoon/hammerspoon) is an amazing open source app that lets you write some Lua code to interact with macOS on a kernel level. 

# Notes

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