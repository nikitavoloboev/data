# My file system

I have an intricate system where I put files according to their purpose and context thus I always know where to find them. And I access everything from Alfred.

## ~/Desktop

My Desktop is nearly always empty. It acts as a kind of `temp` folder where every file that is put there needs to be acted upon and either moved to some other place in my system or deleted.

I use [this workflow](https://github.com/nikitavoloboev/small-workflows/tree/master/clean-folders#readme) to completely trash everything inside `~/Desktop` with one hotkey. I also use [this workflow](https://github.com/nikitavoloboev/small-workflows/blob/master/augmentations/Directory%20watches.alfredworkflow?raw=true) to quickly scan the contents of `~/Desktop` from Alfred.

## ~/Documents

All my actual documents (books, papers..) are put there. The folder synced with iCloud and I often access it from my phone via Spotlight search.

```Bash
~/Documents
❯ exa
Backup  Books  Design  PDFs  Pixave  Research
```

## ~/Downloads

I try to keep this folder like `Desktop` always empty. This is the folder where I download things to from the browser as well as other places.

I use [this workflow](https://github.com/nikitavoloboev/small-workflows/blob/master/augmentations/Recent%20Downloads.alfredworkflow?raw=true) to scan through contents of it from Alfred.

## ~/src

Everything that is code is put into this folder.

```Bash
~/src
❯ exa
alfred  apps  bots  build-to-learn  clones  curated  games  ideas  languages  libs  models  nix  orgs  personal  puzzles  scripts  vim-plugins  vscode-extensions  web  Xcode
```

## ~/src/clones

I often love checking out various GitHub repos. Everything that I clone, I clone into this folder. I use [km macro](https://medium.com/@nikitavoloboev/insta-cloning-ff5f38eb1d32) that will clone the repo that is currently open in my Safari tab. It will put the repo in `~/src/clones` and then open it in VS Code. I also have similar macros that will only clone the repo or clone the repo and open it in Sublime Text.

I then filter contents of the folder with [this workflow](https://github.com/nikitavoloboev/small-workflows/blob/master/augmentations/Directory%20watches.alfredworkflow?raw=true).

## ~/src/orgs

Inside this folder I keep all code from orgs I am working or worked on.

```Bash
~/src/orgs
❯ exa
deedmob  enpicom  learn-anything
```

## ~/src/build-to-learn

The folder is OSS on [GitHub](https://github.com/nikitavoloboev/build-to-learn). Inside this folder I have a bunch of language/framework specific folders. I use it to test out new tools/ideas in there.

## ~/src/Xcode

Contains all my Swift iOS and macOS projects.

```Bash
~/src/Xcode
❯ a
iOS  macOS  playgrounds  safari-extensions
```

## ~/src/alfred

All my Alfred workflows are placed there. And each one is symlinked with [workflow-install](https://gist.github.com/deanishe/35faae3e7f89f629a94e).

```Bash
~/src/alfred
❯ exa
alfred-ask-create-share  alfred-awesome-lists  alfred-github  alfred-learn-anything  alfred-my-mind  alfred-npm  alfred-timer  alfred-trello  alfred-web-searches  small-workflows
```

## ~/src/curated

Keep all the GitHub [curated lists](https://github.com/learn-anything/curated-lists#readme) there where all edits to the `README.md` files is automatically committed with [Hazel](../macOS/apps/hazel.md).

```bash
~/src/curated
❯ exa
alfred-workflows  cheat-sheets        courses        find-work           games   ios-apps    movies       privacy-respecting     quotes        research-papers    slack-groups    talks      websites
blogs             chrome-extensions   curated-lists  firefox-extensions  humans  macos-apps  newsletters  programming-languages  reddit        safari-extensions  spectrum        telegram   youtube
books             command-line-tools  documentaries  forums              images  mindmaps    podcasts     quora                  reddit-multi  series             stack-exchange  tv-series
```

## ~/Dropbox

Use Dropbox to sync configuration files. I also keep this wiki in `Write` dir in Dropbox so I can edit it on the phone with [Ulysses](../macOS/apps/ulysses.md).

```bash
~/Dropbox
❯ exa
Alfred  Apps  Config  IFTTT  Shared  Shares  Temp  upload.sh  Write
```

## Links

- [Ask HN: How do you organise your hard drive? (2018)](https://news.ycombinator.com/item?id=18836472)
- [Ask HN: How do you keep your files organized on macOS? (2019)](https://news.ycombinator.com/item?id=19327264)
