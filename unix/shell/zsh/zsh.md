# [Zsh](http://en.wikipedia.org/wiki/Z_shell)

My shell of choice. My config for it can be seen [here](https://github.com/nikitavoloboev/dotfiles/tree/master/zsh).

## Notes

- Exported variables get passed on to child processes. Non-exported variables do not.
- I don't need to add `function` text before function definition.
- Search `zshall` man page to find all zsh widgets and more things.
- I can use `zsh/zprof` to profile zsh startup time. Docs for it can be seen in `zshmodules`.
- `zsh -x` - See what Zsh executes when it starts new shell.

## Links

- [Zsh custom widgets](https://sgeb.io/posts/2014/04/zsh-zle-custom-widgets/)
- [Prezto](https://github.com/sorin-ionescu/prezto) - Instantly Awesome Zsh.
- [zsh-dirnav](https://github.com/gparker42/zsh-dirnav) - Directory navigation functions for zsh.
- [Zsh prompt with asynchronous Git status](https://vincent.bernat.ch/en/blog/2019-zsh-async-vcs-info)
- [ZSH History Database](https://github.com/larkery/zsh-histdb)
- [Building portable Zsh](https://github.com/xxh/zsh-portable)
- [Speeding up Zsh](https://blog.jonlu.ca/posts/speeding-up-zsh)
- [Some zshrc tricks (2020)](https://www.arp242.net/zshrc.html) ([Lobsters](https://lobste.rs/s/tgmzke/some_zshrc_tricks))
- [5 Types Of ZSH Aliases You Should Know](https://thorsten-hans.com/5-types-of-zsh-aliases) ([HN](https://news.ycombinator.com/item?id=23309310))