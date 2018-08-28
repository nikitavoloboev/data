# [iTerm](https://www.iterm2.com/)
Use Zsh as my shell together with [Antibody](https://github.com/getantibody/antibody) plugin manager and [Pure](https://github.com/sindresorhus/pure) theme.

Assigned w + j with [Karabiner](karabiner/karabiner.md) to open the app from Keyboard Maestro in seconds.

I made my own [Night Blue theme](https://github.com/nikitavoloboev/my-mac-os/tree/master/iterm#readme) for it that looks like this:

![](https://i.imgur.com/Emw7NAj.png)

It goes well with [Pure Zsh plugin](https://github.com/sindresorhus/pure) and [Night Blue theme for Vim](https://github.com/nikitavoloboev/night-blue-vim#readme).

![](https://i.imgur.com/DR2E94n.png)

## My workflow
I always have at least 3 tabs (sessions) always open in iTerm. I open other tabs as the need arises and I start working on other projects.

My always open tabs are:
1. `dev` - This tab is where the majority of my work is done or when I need to quickly jump around my file system. I use it to work through projects I cloned in [~/src/clones](../../unix/my-file-system.md). Access it with [Directory Watches workflow](https://github.com/nikitavoloboev/small-workflows/blob/master/augmentations/Directory%20watches.alfredworkflow?raw=true). Bound to `f+j`.
2. `kar` - Contains my [Karabiner configuration](karabiner/karabiner.md) I edit. Bound to `v+:`.
3. `vim` - Contains my [Vim config](https://github.com/nikitavoloboev/dotfiles/blob/master/nvim/init.vim). Bound to `v+a`.

All other tabs are added on top of this `default` configuration. Usually it looks something like this:
![](https://i.imgur.com/3JvddNy.png)

Where `learn-anynthing` in 4th tab is another project I am working on.

I then use [this workflow](https://github.com/isometry/alfred-tty) to quickly switch between active tty sessions:
![](https://i.imgur.com/qO8wAaN.png)

## Links
- [Make macOS hotkeys work with iTerm](https://stackoverflow.com/questions/6205157/iterm-2-how-to-set-keyboard-shortcuts-to-jump-to-beginning-end-of-line/29403520#29403520)
- [iTerm metal renderer](https://gitlab.com/gnachman/iterm2/wikis/Metal-Renderer)