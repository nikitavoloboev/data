# [iTerm](https://www.iterm2.com/)
Moved completely to using the command line for all my development needs. Assigned [w + j](karabiner/karabiner.md) with Karabiner to open the app from Keyboard Maestro in seconds.

I also made my own [Ayu theme](https://github.com/nikitavoloboev/my-mac-os/tree/master/iterm#readme) for it that looks like this:
![](https://i.imgur.com/KZYHoa9.png)

[Here](https://gist.github.com/nikitavoloboev/3fbe13ce427132d0297f411b62f49034) are all the [Homebrew](http://brew.sh/index.html) packages I like and use.

## My workflow
I always have 5 tabs (sessions) always open in iTerm. I open other tabs as the need arises and I start working on other projects.

My always open tabs are:
1. `::` - This tab never has a program open inside. It is my `runner` terminal. I have it assigned to open with `f + j` and I use it to run commands that are not project related. It's my `general` kind of session.
2. `now` - This tab is where the majority of my work is done. If I am working on any kind of project, I usually work on it through this tab.
3. `clones` - Is used to run various projects I have saved in [~/dev/clones](../../unix/my-file-system.md).
4. `kar` - Contains my [Karabiner configuration file](https://github.com/nikitavoloboev/dotfiles/blob/master/karabiner/private.xml) I edit. I have assigned `v + :` to always switch to this tab from no matter what app I am on my system.
5. `vim` - Contains my [vim config file](https://github.com/nikitavoloboev/dotfiles/blob/master/nvim/init.vim).

All other tabs are added on top of this `default` configuration. Usually it looks something like this:
![](https://i.imgur.com/cMY26z2.png)

Where `la` is some project I am working on.

I then use [this workflow](https://github.com/isometry/alfred-tty) to quickly switch between many tty sessions:
![](https://i.imgur.com/KMvqvzF.png)

## Links
- [Make macOS hotkeys work with iTerm](https://stackoverflow.com/questions/6205157/iterm-2-how-to-set-keyboard-shortcuts-to-jump-to-beginning-end-of-line/29403520#29403520)
- [iTerm metal renderer](https://gitlab.com/gnachman/iterm2/wikis/Metal-Renderer)