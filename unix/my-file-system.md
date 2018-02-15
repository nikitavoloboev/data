# My file system
I have a very intricate system where I put files according to their purpose and context thus I always know where to find them. And I access everything from Alfred.

## ~/Desktop
My Desktop is nearly always empty. It acts as a kind of `temp` folder where every file that is put there needs to be acted upon and either moved to some other place in my system or deleted.

I use [this workflow](https://github.com/nikitavoloboev/small-workflows/tree/master/clean-folders) to completely trash everything inside `~/Desktop` with one hotkey. I also use [this workflow](https://github.com/nikitavoloboev/small-workflows/blob/master/augmentations/Directory%20watches.alfredworkflow?raw=true) to quickly scan the contents of `~/Desktop` from Alfred.

## ~/Downloads
I try to keep this folder like `Desktop` always empty. This is the folder where I download things to from the browser as well as other places.

I use [this workflow](https://github.com/nikitavoloboev/small-workflows/blob/master/augmentations/Recent%20Downloads.alfredworkflow?raw=true) to scan through contents of it from Alfred.

## ~/dev
Everything that is code is put into this folder.
![](https://i.imgur.com/ezM5esL.png)

## ~/dev/clones
I often love checking out various GitHub repos. Everything that I clone, I clone into this folder. I use km macro that will clone the repo that is currently open in my Safari tab. It will put the repo in `~/dev/clones` and then open it in VS Code. I also have similar macros that will only clone the repo or clone the repo and open it in Sublime.

I can also filter contents of the folder with [this workfow](https://github.com/nikitavoloboev/small-workflows/blob/master/augmentations/Directory%20watches.alfredworkflow?raw=true).

## ~/dev/github
Inside this folder I keep all the things that are open source on GitHub sorted by organisation. In my case it just contains two folders:
![](https://i.imgur.com/QRqdrl7.png)

And inside I sort the projects according to their purpose too:
![](https://i.imgur.com/seM6wJQ.png)

## ~/go/src/github.com/nikitavoloboev
I keep all my Go projects inside my GOPATH.

## ~/dev/testing
Inside this folder I have a bunch of language specific folders that usually just have one file inside them with the extension of that language. Here is how that folder looks for me:
![](https://i.imgur.com/0hmPty7.png)

Only `go-test` is put inside my GOPATH. I then have km macros that open these files so I can quickly prototype or test out some code quickly. Here is how such macro looks:
![](https://i.imgur.com/ZNL31og.png)

## ~/dev/ideas
Most of my projects I start, start out in this folder. If something works out and I like the idea and want to develop it further, I move the project away from `ideas`.

## ~/dev/forks
If I forked something and I want to keep working on the repo later, I keep it inside `forks` folder.

## ~/dev/Xcode
Contains all my Swift iOS and macOS projects. It is further divided into macOS, iOS and playgrounds folders.

## ~/dev/learning
I use the folder to learn new technologies, languages and things. Perhaps I am completing some course or going through some book that has exercises. I usually put them there.

## ~/app
I put various app configuration and app specific files in there. For example my `Alfred preferences` is synced to `~/app/alfred/Alfred.alfredpreferences`. Here is how the folder looks for me:
![](https://i.imgur.com/E27yBe5.png)

## ~/Documents
I use Documents to store things like books, research papers, uni work, various app related things and files, audio books and more. Here is how my Documents folder looks like:
![](https://i.imgur.com/sMPGInd.png)

I use a custom icon for Documents folder just so it's easy to distinguish them from the rest using my [folder search workflow](https://github.com/nikitavoloboev/small-workflows/tree/master/folder-search).

And since I use Karabiner and it gives me such a freedom of infinite binding of keys. I dedicate my `4` key for quickly opening various folders in Alfred file viewer.