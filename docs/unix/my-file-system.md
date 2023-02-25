# My file system

I try put files according to their purpose and context thus I always know where to find them. And I access everything from [Alfred](../macOS/apps/alfred/alfred.md)'s [native file search](https://www.alfredapp.com/help/features/file-search/) or [folder search workflow](https://github.com/nikitavoloboev/alfred/tree/master/folder-search).

I try to keep all code I write [open source](../open-source/open-source.md). And also share [all of my `~/src` folder](https://github.com/nikitavoloboev/#src) as it is laid out on my mac.

## ~/Desktop

Think of Desktop as a kind of `todo` pile. Just folders/files that either I am actively working on.

![](https://images.nikiv.dev/desktop-23.png)

> One example of how it looks

The `move` folder is when things get out of control with too many files ~/Desktop, I 'archive' them for later into `move` folder to sort later.

## ~/Documents

Contains pretty much all my non code files/videos/.. The folder synced with iCloud and I can access it from my phone too.

```Bash
~/Documents
❯ exa
AudioBooks Books Design Docs Dropcode Images LA Learn Music Org Other Papers PDFs Personal Screen Studio Tax Use Video WebArchives
```

## ~/Downloads

A huge folder of downloads, both [torrent](../networking/peer-to-peer/bittorrent.md) and from a browser.

I sort the folder by `recently modified` in Finder and act on each file as need arises.

Once in a while I try sort through contents of it so the folder doesn't get too large without items being acted upon. Ideally useful things are neatly organized into `~/Documents`.

I use [this workflow](https://github.com/nikitavoloboev/small-workflows/blob/master/augmentations/Recent%20Downloads.alfredworkflow?raw=true) to scan through contents of the folder from Alfred quickly. Workflow will show most recently added item as first item so I can instantly press `return` to open it.

## ~/src

Everything that is code is put into this folder. Share some contents of it [here](https://github.com/nikitavoloboev/#src).

```Bash
~/src
❯ exa
cli config curated docs extensions games go.work lib new orgs other personal run scripts web
```

## ~/clones

I often love checking out various GitHub repos. Everything that I clone, I clone into this folder. I use [km macro](https://medium.com/@nikitavoloboev/insta-cloning-ff5f38eb1d32) that will clone the repo that is currently open in my Safari tab. It will put the repo in `~/clones` and then open it in VS Code. I also have similar macros that will only clone the repo or clone the repo and open it in Sublime Text.

I then filter contents of the folder with [this workflow](https://github.com/nikitavoloboev/small-workflows/blob/master/augmentations/Directory%20watches.alfredworkflow?raw=true).

## ~/src/extensions

Contains different extensions to various apps in form of workflows/plugins.

```Bash
~/src/extensions
❯ exa
alfred  raycast  safari  vim  vscode  xcode
```

## ~/src/extensions/alfred

All my Alfred workflows are placed there. And each one is symlinked with [workflow-install](https://gist.github.com/deanishe/35faae3e7f89f629a94e).

```Bash
~/src/extensions/alfred
❯ exa
alfred-ask-create-share  alfred-awesome-lists  alfred-github  alfred-learn-anything  alfred-my-mind  alfred-npm  alfred-pocket  alfred-timer  alfred-trello  alfred-web-searches  small-workflows
```

## ~/src/curated

Keep all the GitHub [curated lists](https://github.com/learn-anything/curated-lists) there where all edits to the `README.md` files is automatically committed with [Hazel](../macOS/apps/hazel.md).

```bash
~/src/curated
❯ exa
alfred-workflows  cheat-sheets        courses        events              forums        humans      movies       privacy-respecting     quotes           safari-extensions  stack-exchange  tv-series
blogs             chrome-extensions   curated-lists  find-work           games         ios-apps    newsletters  programming-languages  reddit           slack-groups       talks           websites
books             command-line-tools  documentaries  firefox-extensions  github-stars  macos-apps  podcasts     quora                  research-papers  spectrum           telegram        youtube
```

## ~/Dropbox

No longer use Dropbox.

```bash
~/Dropbox
❯ exa
Alfred  Apps  Config  IFTTT  Public  Shared  Temp  upload.sh  Write
```

## Links

- [Ask HN: How do you organise your hard drive? (2018)](https://news.ycombinator.com/item?id=18836472)
- [Ask HN: How do you keep your files organized on macOS? (2019)](https://news.ycombinator.com/item?id=19327264)
- [How do you organize your \$HOME directory? (2019)](https://lobste.rs/s/zpw6py/how_do_you_organize_your_home_directory)
- [Ask HN: How do you organise your files and folders? (2020)](https://news.ycombinator.com/item?id=23404900)
- [Ask HN: How do you manage your personal documents? (2021)](https://news.ycombinator.com/item?id=29161110)
