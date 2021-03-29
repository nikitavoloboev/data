# [Alfred](https://www.alfredapp.com)

Alfred is a powerful launcher that you can program to show anything you want. It saved me a lot of time as I use it to search through anything instantly.

![](https://i.imgur.com/PtXa6By.png)

It has a great [community](http://www.alfredforum.com/) and [amazing workflows](https://github.com/learn-anything/alfred-workflows) that you can use.

I wrote [an article](https://medium.com/@nikitavoloboev/writing-alfred-workflows-in-go-2a44f62dc432) on how anyone can start developing workflows of their own using Go language and [AwGo](https://github.com/deanishe/awgo) library.

## Workflows I use

My most used workflows at current time are:

- [Searchio](https://github.com/deanishe/alfred-searchio)
- [GitHub Jump](https://github.com/nikitavoloboev/small-workflows#workflow-augmentations)
- [Web Searches](https://github.com/nikitavoloboev/alfred-web-searches)
- [Learn Anything](https://github.com/nikitavoloboev/alfred-learn-anything)
- [Safari Assistant](https://github.com/deanishe/alfred-safari-assistant)
- [Ask, Create Share](https://github.com/nikitavoloboev/alfred-ask-create-share)
- [My Mind](https://github.com/nikitavoloboev/alfred-my-mind)
- [Maestro](https://github.com/iansinnott/alfred-maestro)
- [Fast Menu Bar Search](https://github.com/ascandroli/menudump/releases/download/1.8.0/Menu.Bar.Search-v1_8.alfredworkflow)

Coupled with many [small workflows](https://github.com/nikitavoloboev/small-workflows) I made. I hope to add more workflows of my own [to the mix](https://github.com/learn-anything/alfred-workflows) with time.

## Alfred theme

I use [Mono Light theme](https://www.alfredapp.com/extras/theme/yyoqZV6XGS/) with macOS light appearance (during day).

![](https://i.imgur.com/d5is1ao.png)

I use [Mono theme](https://www.alfredapp.com/extras/theme/xzcLtcIIDe/) with macOS dark appearance (evenings & no light areas).

![](https://i.imgur.com/Y4oKBoT.png)

## Snippets

I use Alfred together with [Keyboard Maestro](../keyboard-maestro/keyboard-maestro.md) to automate text expansions. Alfred for simple expansions, KM for complex or per app expansions.

![](https://i.imgur.com/vWP4lkz.png)

### Symlink workflows

It is a great idea to symlink workflows you are working on so you can work on them in the comfort of your file system and not inside Alfred UUID named directories.

I use [this script](https://gist.github.com/deanishe/35faae3e7f89f629a94e) to achieve this. In example of [this workflow](https://github.com/nikitavoloboev/small-workflows/tree/master/folder-search) you can use it as follows. Assuming the `workflow` directory contains Alfred workflow contents (`info.plist` file). You can run `workflow-install.py -s workflow`.

I can also then use [this script](https://gist.github.com/deanishe/b16f018119ef3fe951af) and build my workflow with `workflow-build.py -v source`.

## Notes

- [Package workflows from CLI with version numbers](https://www.alfredforum.com/topic/10838-how-to-package-workflows-from-the-command-line/?tab=comments#comment-55677)
- Alfred can preserve a user's Hotkey shortcuts, keywords and workflow variable values. Everything else in the workflow directory gets overwritten.
- I prefix my own workflows that I have not published out to the world with **a.** prefix. This way I know off hand that this workflow is either private or unreleased yet.
- Sometimes when passing things from one object into another, you want to clear your {query} being passed to get a clean prompt. To achieve this you can set argument utility and set it as empty like here:

![](https://i.imgur.com/seduWW7.png)

- You can exclude certain folders or files from being searched. For example, if you have a file filter that searches through some folders and its contents but you don't wish for it to search some big `node_modules` folder. You can simply add this folder to Spotlight preferences under _privacy_ tab like here:

![](https://i.imgur.com/D0NP2s3.png)

- Adding `alfred:ignore` tag to files/folders will have Alfred ignore the file/folder in its searches.
- [If you copy the selected text, then open Alfred's clipboard manager, you can use cmd+s on the selected clipboard entry, and this will open Alfred's preferences and pre-fill a snippet for you to save.](https://www.alfredforum.com/topic/15852-how-to-save-selected-text-as-alfred-snippet/)
- File note found? Go to Alfred's preferences > Help > Troubleshooting and run the File Search Troubleshooting, dragging in some of the files you are expecting to see in your File Filter. Check the `Content Type:` field. It could be something like `com.unknown.md` instead of `net.daringfireball.markdown` you might expect.

## Links

- [Reporting Problems with Workflows](https://www.alfredforum.com/topic/10224-reporting-problems-with-workflows/)
- [How to: workflow / environment variables](https://www.alfredforum.com/topic/9070-how-to-workflowenvironment-variables/?tab=comments#comment-45177)
- [fuzzy-list](https://github.com/derickfay/fuzzylist) - Self-updating list filter workflow template.
- [fuzzy-search](https://github.com/deanishe/alfred-fuzzy) - Fuzzy search helper.
- [Deanishe's approach in building Go workflows](https://github.com/deanishe/awgo/commit/5f0051950af39371385f2dfda96483eb1423e565#r29572675)
- [Pacmax](https://pacmax.org/) - Explore & Share Great Packages for Alfred.
- [Fix for running compiled workflows in Catalina](https://github.com/deanishe/alfred-sublime-text/issues/20#issuecomment-539578934)
- [Alfred Workflows Store](https://www.alfredworkflows.store/)
- [AlfredExtraPane](https://github.com/mr-pennyworth/alfred-extra-pane) - Spotight-like rich previews for Alfred workflows. ([Forum post](https://www.alfredforum.com/topic/16111-wip-poc-spotlight-like-rich-preview-pane-for-alfred-workflows/))
- [Workflow Icon Generator](https://icons.deanishe.net/)
- [How to Programmatically Add Folders to the Spotlight Ignore List (2020)](https://mattprice.me/2020/programmatically-modify-spotlight-ignore/)
- [Alfred Workflow ScriptFilter in Swift](https://github.com/godbout/AlfredWorkflowScriptFilter) ([Discussion](https://www.alfredforum.com/topic/16719-alfredworkflowscriptfilter-swift/))
