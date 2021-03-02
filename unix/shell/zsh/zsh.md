# [Zsh](http://en.wikipedia.org/wiki/Z_shell)

My shell of choice. My config for it can be seen [here](https://github.com/nikitavoloboev/dotfiles/tree/master/zsh).

## Notes

- Exported variables get passed on to child processes. Non-exported variables do not.
- I don't need to add `function` text before function definition.
- Search `zshall` man page to find all zsh widgets and more things.
- I can use `zsh/zprof` to profile zsh startup time. Docs for it can be seen in `zshmodules`.
- `zsh -x` - See what Zsh executes when it starts new shell.
- [`cmd -` and TAB will list options for command.](https://twitter.com/rsms/status/1304877145743790080)

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
- [Zsh and Fish’s simple but clever trick for highlighting missing linefeeds (2020)](https://www.vidarholen.net/contents/blog/?p=878) ([HN](https://news.ycombinator.com/item?id=23520240))
- [Nice Zsh configs/tips](https://twitter.com/rsms/status/1304837079826747392)
- [tabtab](https://github.com/denosaurs/tabtab) - Generate CLI completions for zsh, bash, and fish.
- [Oh My Zsh](https://github.com/ohmyzsh/ohmyzsh) - Open source, community-driven framework for managing your zsh configuration.
- [zsh-defer](https://github.com/romkatv/zsh-defer) - Deferred execution of zsh commands.
- [Zsh Compinit ... RTFM (2018)](https://www.danielmoch.com/posts/2018/11/zsh-compinit-rtfm/) - Optimizing Zsh startup time.
- [Mastering Zsh](https://github.com/rothgar/mastering-zsh) - Advanced topics to take advantage of zsh.
- [Zsh Tricks to Blow Your Mind (2021)](https://www.twilio.com/blog/zsh-tricks-to-blow-your-mind) ([HN](https://news.ycombinator.com/item?id=26175894))
- [sandboxd](https://github.com/benvan/sandboxd) - Speed up your bashrc / zshrc: avoids running (slow) setup commands until you actually need them.
