# My file system
I have an intricate system where I put files according to their purpose and context thus I always know where to find them. And I access everything from Alfred.

## ~/Desktop
My Desktop is nearly always empty. It acts as a kind of `temp` folder where every file that is put there needs to be acted upon and either moved to some other place in my system or deleted.

I use [this workflow](https://github.com/nikitavoloboev/small-workflows/tree/master/clean-folders) to completely trash everything inside `~/Desktop` with one hotkey. I also use [this workflow](https://github.com/nikitavoloboev/small-workflows/blob/master/augmentations/Directory%20watches.alfredworkflow?raw=true) to quickly scan the contents of `~/Desktop` from Alfred.

## ~/Downloads
I try to keep this folder like `Desktop` always empty. This is the folder where I download things to from the browser as well as other places.

I use [this workflow](https://github.com/nikitavoloboev/small-workflows/blob/master/augmentations/Recent%20Downloads.alfredworkflow?raw=true) to scan through contents of it from Alfred.

## ~/dev
Everything that is code is put into this folder.

```Bash
~/dev
❯ ls
alfred  applescript  bots  clones  curated  forks  orgs  ideas  learn  ml  personal  practice  python  rust  safari  test  uni  web  Xcode
```

## ~/dev/clones
I often love checking out various GitHub repos. Everything that I clone, I clone into this folder. I [use km macro](https://medium.com/@NikitaVoloboev/insta-cloning-ff5f38eb1d32) that will clone the repo that is currently open in my Safari tab. It will put the repo in `~/dev/clones` and then open it in VS Code. I also have similar macros that will only clone the repo or clone the repo and open it in Sublime.

I then filter contents of the folder with [this workfow](https://github.com/nikitavoloboev/small-workflows/blob/master/augmentations/Directory%20watches.alfredworkflow?raw=true).

## ~/dev/forks
If I forked something and I want to keep working on, I keep it inside `forks` folder.

## ~/dev/orgs
Inside this folder I keep all the things that are open source on GitHub sorted by organisation. In my case it's only [learn-anything](https://github.com/learn-anything/):

```Bash
~/dev/orgs
❯ ls
learn-anything
```

## ~/go/src/github.com/nikitavoloboev
I keep all my Go libraries inside my GOPATH.

## ~/dev/test
Inside this folder I have a bunch of language specific folders that usually just have one file inside them with the extension of that language. Here is how that folder looks for me:

```Bash
~/dev/test
❯ ls
bash-test	c-test		clojure-test	go-test		haskell-test	js-test		lisp-test	python-test	web-test
```

I then have KM macros to quickly open these files:
![](https://i.imgur.com/Ygqwukd.png)

## ~/dev/ideas
Most of my projects I start, start out in this folder. If something works out and I like the idea and want to develop it further, I move the project away from `ideas`.

## ~/dev/Xcode
Contains all my Swift iOS and macOS projects. It is further divided into macOS, iOS and playgrounds folders.

## ~/dev/learn
I use the folder to learn new technologies, languages and things. Perhaps I am completing some course or going through some book that has exercises. I put it there.

```Bash
~/dev/learn
❯ ls
go-learn  graphics-learn  ml-learn  react-learn
```

## ~/dev/alfred
All my Alfred workflows are placed there. And each one is symlinked either with [worklow install](https://gist.github.com/deanishe/35faae3e7f89f629a94e) or [alfred CLI tool](https://godoc.org/github.com/jason0x43/go-alfred/alfred) for Go workflows.

```Bash
~/dev/alfred
❯ ls
alfred-ask-create-share  alfred-birthday      alfred-inline-searches  alfred-my-mind  alfred-twitter
alfred-awesome-lists     alfred-breathe       alfred-lastfm           alfred-pixave   alfred-web-searches
alfred-awgo-practice     alfred-github-users  alfred-learn-anything   alfred-trello   small-workflows
```

## ~/app
I put various app configuration and app specific files in there. Here is how the folder looks for me:

```Bash
~/app
❯ ls
dash  focus  hammerspoon  iterm  km  paw  safari
```

## ~/Documents
I use Documents to store things like books, research papers, uni work, various app related things and files, audio books and more. Here is how my Documents folder looks like:

```Bash
~/Documents
❯ ls
Audio Books  Books  Design  History  LaTeX  Personal docs  Pixave  Podcasts  Uni  Watch
```