# Making Workflows
There is a [detailed article](https://medium.com/@NikitaVoloboev/writing-alfred-workflows-in-go-2a44f62dc432) available on how you can write workflows in Go language.

You can also take a look through [Alfred's help](https://www.alfredapp.com/help/) and get help from the [forum](https://www.alfredforum.com/).

## How I develop workflows
I like to write all my workflows in Go language unless the workflow can be done using only Alfred builtin objects and some python/bash code. Here are steps I follow to create my workflows, both Go workflows and non Go.

### Using Go and [AwGo](https://github.com/deanishe/awgo)
1. Create a folder in `~/dev/alfred`. Prefix the folder name with `alfred-`. i.e. `alfred-web-searches`.
2. Initialise the folder with git as any project.
3. Inside the `alfred-..` folder I create a `workflow` directory. This is where the actual workflow will live in as well as the binary file compiled from my Go code.
4. Create the workflow I want to make in Alfred Preferences. Add the icon, name and unique bundle ID. And add the workflow to my custom `Mine` category so I can quickly filter out all the workflows I've ever made inside Alfred.
5. Carry the info.plist and .png file of the logo to this workflow directory I created. I use Alfred's _Move To..._ file action for quickly transferring files. You can find these files by going here.

    ![](https://i.imgur.com/5UBCGbc.png)

6. Delete the workflow I made from Alfred preferences.
7. Inside my `alfred-..` directory I run `alfred link`. That creates a symlink between the workflow directory and where the actual workflow lies. This uses [alfred](https://godoc.org/github.com/jason0x43/go-alfred/alfred) CLI tool.
8. Write my Go code inside `alfred-..` folder and when I want to compile it to run from Alfred, I run `alfred build`.
9. Write the code and make a MVP. Not forgetting to add/change version number of my workflow.
10. If I want to share the workflow with the world (which is usually always), I create a GitHub repository and push my workflow there, making sure that AwGo updating works so all the users can receieve future updates.

### Using Alfred builtin objects + some code
1. Create a folder in `~/dev/alfred/small-workflows`. Name the folder by workflow name, lowercased and separated by dashes without `alfred` prefix.
2. Inside the folder I create a `workflow` directory.
3. Create the workflow I want to make in Alfred Preferences. Add the icon, name and unique bundle ID. Add the workflow to `Mine` category.
4. Carry the info.plist and .png file of the logo to `workflow` directory I just made.
5. Delete the workflow I made from Alfred preferences.
6. Inside my workflow's folder in _small workflows_, I run `workflow-install -s workflow`. `workflow-install` command is placed inside my [~/.dotfiles/bin](https://github.com/nikitavoloboev/dotfiles/tree/master/bin) which is added to my $PATH. The [script](https://gist.github.com/deanishe/35faae3e7f89f629a94e) symlinks the `workflow` directory to Alfred.
7. If the workflow is private to me, I prefix it's name with `a. ` before the workflow name to indicate that I have not released the workflow yet or it contains sensitive information.
8. Create the Alfred objects I need to make the workflow does what I want. If the workflow needs something more complex, I sometimes create Python files inside the workflow dir to do what I want and call them from the Script Filters.
9. In cases where I want to release the workflow, I strip the `a. ` prefix from the name. Then go to to my [Objects library](https://github.com/nikitavoloboev/small-workflows/tree/master/objects-library) workflow and transfer a [OneUpdater](https://github.com/vitorgalvao/alfred-workflows/tree/master/OneUpdater) object to my workflow.
10. Change `readonly remote_info_plist` and `readonly workflow_url` variables in OneUpdater script to point to correct links. Then add a version number to the workflow if I haven't yet. Export the workflow to `~/Desktop` with [DirectoryWatches](https://github.com/nikitavoloboev/small-workflows/blob/master/augmentations/Workflow%20directory.alfredworkflow?raw=true) workflow and then viewing `~/Desktop` with [Directory Watches](https://github.com/nikitavoloboev/small-workflows/blob/master/augmentations/Directory%20watches.alfredworkflow?raw=true) workflow, I file action on the workflow in `~/Desktop` to move it quickly to the directory in my [small workflows](https://github.com/nikitavoloboev/small-workflows).
11. `git commit` and `push` the directory contents to GitHub.

## Updating existing workflows
### Go workflows
I first commit and push the changes to my repo.

I then use [Workflow Directory](https://github.com/nikitavoloboev/small-workflows/blob/master/augmentations/Workflow%20directory.alfredworkflow?raw=true) workflow to search the workflow I want to export and with a modifier press, export that workflow to `~/Desktop`. This action also copies the version number of the workflow.

I then use modified by me wonderful [Repos](https://github.com/deanishe/alfred-repos) workflow which searches Git projects in directories I specify to open a __new release__ page in my browser on a modifier keypress. You can grab my hacked version of Repos workflow [here](https://github.com/nikitavoloboev/small-workflows/blob/master/augmentations/Git%20Repos.alfredworkflow?raw=true).

Using the workflow I search the `alfred-..` workflow I want to release an update for and with a modifier press quickly open the __new release__ page. I then add the version number I have copied in my clipboard and add the name to the release and I attach the workflow sitting in my `~/Desktop` to the release.

### Small workflows
I use [Workflow Directory](https://github.com/nikitavoloboev/small-workflows/blob/master/augmentations/Workflow%20directory.alfredworkflow?raw=true) workflow to search the workflow I want to export and with a modifier press, export that workflow to `~/Desktop`.

I delete the original workflow sitting in the directory of the workflow and then with Alfred's `Move to..` transfer the workflow to the directory of the workflow. Commit and push.