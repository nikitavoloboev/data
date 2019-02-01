# [Sublime Text](https://www.sublimetext.com)

Use this editor in addition to [VS Code](../vs-code/vs-code.md) and [Neovim](../vim/vim.md) for its blazing fast speed of opening files.

I use it primarily to edit markdown files like [this wiki](../../other/wiki-workflow.md). I also edit config files and open large and small files for quick edits.

I use [many plugins](sublime-text-plugins.md) together with [Ayu theme](https://github.com/dempfi/ayu).

## Notes

- Entering into sublime console (View -> Show Console):
  - `sublime.log_commands(True)`
  - `sublime.log_input(True)` - Allows you to then see what commands you type as well as actions you make map to as sublime bindings.
  - You can then turn all logging by running commands above with `False` or restart Sublime.
- `New view into file` will split current file into two tabs.
- I binded jj to go to normal mode from insert. This way when I load a file in sublime, it doesn't sometimes immediately go to normal mode. So I can instantly open file and safely spam j to go down a page without writing the j's.
