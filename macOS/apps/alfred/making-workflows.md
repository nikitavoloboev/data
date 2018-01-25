# Making Workflows
There is a [detailed article](https://medium.com/@NikitaVoloboev/writing-alfred-workflows-in-go-2a44f62dc432) available on how you can write workflows in Go language.

You can also take a look through [Alfred's help](https://www.alfredapp.com/help/) and get help from the [forum](https://www.alfredforum.com/).

## My workflow in developing workflows
I like to write all my workflows in Go language unless the workflow can be done using only Alfred builtin objects and some python/bash code.

### Workflows in Go language
In cases where I am planning to use Go to write the workflow. Here are the steps I follow. This assumes you have installed [alfred cli tool](https://godoc.org/github.com/jason0x43/go-alfred/alfred).

1. Create a folder in my GOPATH inside `~/go/src/github.com/nikitavoloboev/`.
2. Prefix the folder name with `alfred-`.
	- i.e. `alfred-twitter`.
3. Initialise the folder with git as any project.
4. Inside my `alfred-` folder I create a `workflow` directory.
	- this is where the actual workflow will live in as well as the binary file compiled from my Go code.
5. Create the workflow I want to make in alfred preferences.
	- add the icon, name and unique bundle id.
6. Carry the info.plist and .png file of the logo to this workflow directory I created.
	- I use Alfred's _Move To..._ file action for quickly transferring files.
	- You can find these files by going here.

![](https://i.imgur.com/rVlcl9y.png)

- After transferring the files, I delete the workflow from Alfred preferences
- Inside my `alfred-` directory I run `alfred link`.
	- That creates a symlink between the workflow directory and where the actual workflow lies.
 - Write my Go code inside `alfred-` folder and when I want to compile it to run from Alfred, I run `alfred build`.

