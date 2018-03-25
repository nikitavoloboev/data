# [Sublime text](https://www.sublimetext.com)
Use this editor in addition to VS Code and Neovim for its blazing fast speed of opening files.

I use it primarily to edit markdown files like [this wiki](../other/wiki-workflow.md). I also edit config files and open large and small files for quick edits.

I use [Ayu theme](https://github.com/dempfi/ayu).

## Tricks I use
- I also made a km macro that I binded with `v + r` to open `temp-notes.md` file I have saved on my system in Sublime. It acts as a kind of quick dump buffer for writing a thing that will be deleted afterwards. I can use my vim mode to fast edit text and the paste the text I wrote to another app like Google docs.
- I [binded jj](https://github.com/nikitavoloboev/dotfiles/blob/master/sublime/Default%20(OSX).sublime-keymap) to go to normal mode from insert. This way when I load a file in sublime, it doesn't sometimes immediately go to normal mode. So I can instantly open file and safely spam j to go down a page without writing the j's.

## Packages I use
[Here](https://github.com/nikitavoloboev/dotfiles/blob/master/sublime/Package%20Control.sublime-settings) is my up-to-date list of Sublime packages I use.

## Interesting packages
- [Linter flake8](https://github.com/SublimeLinter/SublimeLinter-flake8)

## Notes
- Entering into sublime console (View -> Show Console):
	- `sublime.log_commands(True)`
	- `sublime.log_input(True)` - Allows you to then see what commands you type as well as actions you make map to as sublime bindings.
	- You can then turn all logging by running commands above with `False` or restart Sublime.
- `New view into file` will split current file into two tabs.