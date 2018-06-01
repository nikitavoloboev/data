# [iTerm](https://www.iterm2.com/)
Use Zsh as my shell together with [Antibody](https://github.com/getantibody/antibody) shell plugin manager and [Pure](https://github.com/sindresorhus/pure) theme.

Assigned w + j with [Karabiner](karabiner/karabiner.md) to open the app from Keyboard Maestro in seconds.

I made my own [Ayu theme](https://github.com/nikitavoloboev/my-mac-os/tree/master/iterm#readme) for it that looks like this:

![](https://i.imgur.com/MFmvp4F.png)

It goes well with [Pure zsh theme](https://github.com/sindresorhus/pure) and [Ayu theme for Vim](https://github.com/ayu-theme/ayu-vim).

![](https://i.imgur.com/m6CK29L.png)

## My workflow
I always have 4 tabs (sessions) always open in iTerm. I open other tabs as the need arises and I start working on other projects.

My always open tabs are:
1. `::` - This tab never has a program open inside. It is my `runner` terminal. I have it assigned to open with `f+j` and I use it to run commands that are not project related. It's my `general` kind of session.
2. `dev` - This tab is where the majority of my work is done. I use Is used to work through projects I cloned in [~/src/clones](../../unix/my-file-system.md) and forked in [~/src/forks](../../unix/my-file-system.md). Access it with [Directory Watches workflow](https://github.com/nikitavoloboev/small-workflows/blob/master/augmentations/Directory%20watches.alfredworkflow?raw=true). Bound to `f+.`.
3. `kar` - Contains my [Karabiner configuration](https://github.com/nikitavoloboev/dotfiles/blob/master/karabiner/private.xml) I edit. Bound to `v+:`.
4. `vim` - Contains my [Vim config](https://github.com/nikitavoloboev/dotfiles/blob/master/nvim/init.vim). Bound to `v+a`.

All other tabs are added on top of this `default` configuration. Usually it looks something like this:
![](https://i.imgur.com/fM9yCCX.png)

Where `la` is some project I am working on.

I then use [this workflow](https://github.com/isometry/alfred-tty) to quickly switch between active tty sessions:
![](https://i.imgur.com/UUDBHcq.png)

## Links
- [Make macOS hotkeys work with iTerm](https://stackoverflow.com/questions/6205157/iterm-2-how-to-set-keyboard-shortcuts-to-jump-to-beginning-end-of-line/29403520#29403520)
- [iTerm metal renderer](https://gitlab.com/gnachman/iterm2/wikis/Metal-Renderer)