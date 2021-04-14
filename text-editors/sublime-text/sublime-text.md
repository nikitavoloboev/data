# [Sublime Text](https://www.sublimetext.com)

Use this editor in addition to [VS Code](../vs-code/vs-code.md) and [Neovim](../vim/vim.md) for its blazing fast speed of opening files.

I use it primarily to edit markdown files like [this wiki](../../other/wiki-workflow.md). I also edit config files and open large and small files for quick edits.

I use [many plugins](sublime-text-plugins.md) together with [Ayu theme](https://github.com/dempfi/ayu).

I switch between Ayu Light and Ayu Mirage themes as I change between macOS appearances.

![](https://i.imgur.com/vdTDYe1.png)

> Ayu Light

![](https://i.imgur.com/sdIqSvT.png)

> Ayu Mirage

## Notes

- Entering into sublime console (View -> Show Console):
  - `sublime.log_commands(True)`
  - `sublime.log_input(True)` - Allows you to then see what commands you type as well as actions you make map to as sublime bindings.
  - You can then turn all logging by running commands above with `False` or restart Sublime.
- `New view into file` will split current file into two tabs.
- I binded jj to go to normal mode from insert. This way when I load a file in sublime, it doesn't sometimes immediately go to normal mode. So I can instantly open file and safely spam j to go down a page without writing the j's.

## Links

- [OdatNurd - Sublime Text Tutorials](https://www.youtube.com/user/nurdz/playlists)
- [Sublime Text Community Documentation](https://docs.sublimetext.io/guide/) ([Code](https://github.com/sublimetext-io/docs.sublimetext.io))
- [PackageDev](https://github.com/SublimeText/PackageDev) - Provides syntax highlighting and other helpful utility for Sublime Text resource files.
- [Sublime Text 4 change log](https://gist.github.com/jfcherng/7bf4103ea486d1f67b7970e846b3a619)
- [HN: Sublime Text 4 (2021)](https://news.ycombinator.com/item?id=26646142)
- [Building an ultimate writing machine from Sublime Text](https://tonsky.me/blog/sublime-writer/)
