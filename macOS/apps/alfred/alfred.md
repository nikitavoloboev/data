# [Alfred](https://www.alfredapp.com)
Alfred is a very powerful launcher that you can program to show anything you want. It saved me a lot of time in my life.

![](https://i.imgur.com/uzq1w5g.png)

It has a great [community](http://www.alfredforum.com/) and [amazing workflows](https://github.com/learn-anything/alfred-workflows#readme) that you can use.

I wrote [an article](https://medium.com/@nikitavoloboev/writing-alfred-workflows-in-go-2a44f62dc432) on how anyone can start developing workflows of their own using Go language and [AwGo](https://github.com/deanishe/awgo) library.

## Workflows I use
My most used workflows at current time are:
- [Searchio](https://github.com/deanishe/alfred-searchio#readme)
- [GitHub Jump](https://github.com/nikitavoloboev/small-workflows#workflow-augmentations)
- [Web Searches](https://github.com/nikitavoloboev/alfred-web-searches#readme)
- [Learn Anything](https://github.com/nikitavoloboev/alfred-learn-anything#readme)
- [Safari Assistant](https://github.com/deanishe/alfred-safari-assistant#readme)
- [Ask, Create Share](https://github.com/nikitavoloboev/alfred-ask-create-share#readme)
- [My Mind](https://github.com/nikitavoloboev/alfred-my-mind#readme)
- [Maestro](https://github.com/iansinnott/alfred-maestro#readme)
- [Fast Menu Bar Search](https://github.com/ascandroli/menudump/releases/download/1.8.0/Menu.Bar.Search-v1_8.alfredworkflow)

Coupled with many [small workflows](https://github.com/nikitavoloboev/small-workflows) I made. I hope to add more workflows of my own [to the mix](https://github.com/learn-anything/alfred-workflows#readme) with time.

## Alfred theme
I use [Night Blue theme](https://www.alfredapp.com/extras/theme/24fhXfBld7/). Here is how it looks:
![](https://i.imgur.com/OPd7GGA.png)

### Symlink workflows
It is a great idea to symlink workflows you are working on so you can work on them in the comfort of your file system and not inside Alfred UUID named directories.

I use [this script](https://gist.github.com/deanishe/35faae3e7f89f629a94e) to achieve this. In example of [this workflow](https://github.com/nikitavoloboev/small-workflows/tree/master/folder-search#readme) you can use it as follows. Assuming the `workflow` directory contains Alfred workflow contents (`info.plist` file). You can run `workflow-install.py -s workflow`.

I can also then use [this script](https://gist.github.com/deanishe/b16f018119ef3fe951af) and build my workflow with `workflow-build.py -v source`.

## Notes
- [Package workflows from CLI with version numbers](https://www.alfredforum.com/topic/10838-how-to-package-workflows-from-the-command-line/?tab=comments#comment-55677)
- Alfred can preserve a user's Hotkey shortcuts, keywords and workflow variable values. Everything else in the workflow directory gets overwritten.
- I prefix my own workflows that I have not published out to the world with __a.__ prefix. This way I know off hand that this workflow is either private or unreleased yet.
- Sometimes when passing things from one object into another, you want to clear your {query} being passed to get a clean prompt. To achieve this you can set argument utility and set it as empty like here:

![](https://i.imgur.com/seduWW7.png)

- You can exclude certain folders or files from being searched. For example, if you have a file filter that searches through some folders and its contents but you don't wish for it to search some big `node_modules` folder. You can simply add this folder to Spotlight preferences under _privacy_ tab like here:

![](https://i.imgur.com/D0NP2s3.png)

## Links
- [Reporting Problems with Workflows](https://www.alfredforum.com/topic/10224-reporting-problems-with-workflows/)
- [How to: workflow / environment variables](https://www.alfredforum.com/topic/9070-how-to-workflowenvironment-variables/?tab=comments#comment-45177)
- [fuzzy-list](https://github.com/derickfay/fuzzylist) - Self-updating list filter workflow template.
- [fuzzy-search](https://github.com/deanishe/alfred-fuzzy) - Fuzzy search helper.
- [Deanishe's approach in building Go workflows](https://github.com/deanishe/awgo/commit/5f0051950af39371385f2dfda96483eb1423e565)