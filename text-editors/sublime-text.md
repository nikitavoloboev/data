# Sublime text
Use this editor in addition to VS Code and Neovim for its blazing fast speed of opening files.

I use it primarily to edit markdown files like [this wiki](https://wiki.nikitavoloboev.xyz/other/wiki-workflow.html). I also edit config files and open large and small files for quick edits.

[Six](https://github.com/guillermooo/Six) vim plugin for this editor is genuinely amazing. Alongside [Ayu theme](https://github.com/dempfi/ayu).

## Notes
- Entering into sublime console (View -> Show Console):
	- `sublime.log_commands(True)`
	- `sublime.log_input(True)` - Allows you to then see what commands you type as well as actions you make map to as sublime bindings.
	- You can then turn all logging by running commands above with `False` or restart Sublime.
- `New view into file` will split current file into two tabs.