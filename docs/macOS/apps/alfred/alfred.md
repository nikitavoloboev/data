---
title: Alfred
---

# [Alfred](https://www.alfredapp.com)

Alfred is a powerful launcher that you can program to show anything you want. It saved me a lot of time as I use it to search through anything instantly.

![](https://i.imgur.com/PtXa6By.png)

It has a great [community](http://www.alfredforum.com/) and [amazing workflows](https://github.com/learn-anything/alfred-workflows) that you can use.

I wrote [an article](https://medium.com/@nikitavoloboev/writing-alfred-workflows-in-go-2a44f62dc432) on how anyone can start developing workflows of their own using Go language and [AwGo](https://github.com/deanishe/awgo) library.

[Raycast](../../../tools/raycast.md) & [Script Kit](https://www.scriptkit.com/) are nice alternatives.

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

I most often use [Mono Dim theme](https://www.alfredapp.com/extras/theme/jzs9j2Kqmu/) when in macOS dark appearance.

![](https://i.imgur.com/0EIwRT7.png)

I use [Mono Light theme](https://www.alfredapp.com/extras/theme/yyoqZV6XGS/) with macOS light appearance (when light is hitting the screen).

![](https://i.imgur.com/d5is1ao.png)

I also made [Mono theme](https://www.alfredapp.com/extras/theme/xzcLtcIIDe/) but stopped using it.

![](https://i.imgur.com/Y4oKBoT.png)

The themes work best with GitHub [VSCode](https://marketplace.visualstudio.com/items?itemName=GitHub.github-vscode-theme) & [iTerm](https://github.com/cdalvaro/github-vscode-theme-iterm) dim themes.

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

## Code

Strip title and subtext from input. Put inside [JSON utility](https://www.alfredapp.com/help/workflows/utilities/json/). Useful to get clean prompts in Alfred (i.e. it's used in [web-searches](https://github.com/nikitavoloboev/alfred-web-searches) workflow)

```json
{
  "alfredworkflow": {
    "arg": "{query}",
    "config": {
      "title": "",
      "runningsubtext": "",
      "subtext": ""
    },
    "variables": {}
  }
}
```

Export workflow to a location with version number in clipboard.

```bash
readonly workflow_dir="{query}"

if [[ ! "${workflow_dir}" == *'Alfred.alfredpreferences/workflows/user.workflow.'* ]]
then
  echo "You need to be inside the workflow’s directory in Alfred’s preferences directory." >&2
  exit 1
fi

readonly workflow_name="$(/usr/libexec/PlistBuddy -c 'print name' "${workflow_dir}/info.plist")"
readonly workflow_version="$(/usr/libexec/PlistBuddy -c 'print version' "${workflow_dir}/info.plist")"
readonly workflow_file="${HOME}/Desktop/${workflow_name}-${workflow_version}.alfredworkflow"

find "${workflow_dir}" -iname '.DS_Store' -delete

if /usr/libexec/PlistBuddy -c 'Print variablesdontexport' "${workflow_dir}/info.plist" &> /dev/null
then
  readonly workflow_dir_to_package="$(mktemp -d)"
  cp -R "${workflow_dir}/"* "${workflow_dir_to_package}"
  readonly tmp_info_plist="${workflow_dir_to_package}/info.plist"
  /usr/libexec/PlistBuddy -c 'print variablesdontexport' "${tmp_info_plist}" | grep '    ' | sed -E 's/ {4}//' | xargs -I {} /usr/libexec/PlistBuddy -c "set variables:'{}' ''" "${tmp_info_plist}"
else
  readonly workflow_dir_to_package="${workflow_dir}"
fi

if ditto -ck "${workflow_dir_to_package}" "${workflow_file}"
then
  echo "Created ${workflow_file}"
  exit 0
else
  echo "There was and error creating ${workflow_file}"
  exit 1
fi
```

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
- [AlfredSwift](https://github.com/mr-pennyworth/AlfredSwift) - Swift package to access and manipulate metadata and preferences for Alfred and workflows.
- [Workflow Development](https://intersect.rknight.me/macos/alfred/workflow-development/)
- [Alfred Workflows (Rust)](https://github.com/rust-playground/alfred-workflows-rs)
- [Updated Third-Party Python 2 Workflows](https://github.com/alfredapp/updated-third-party-python2-workflows)
- [Alfred-Workflow](https://github.com/NorthIsUp/alfred-workflow-py3) - Full-featured library for writing Alfred 3 & 4 workflows.
- [go-alfred](https://github.com/jason0x43/go-alfred) - Alfred Workflow utility library for Go. Useful CLI too.
- [powerpack](https://github.com/rossmacarthur/powerpack) - Supercharge your Alfred workflows by building them in Rust.
- [Alfred Automation Tasks](https://github.com/alfredapp/automation-tasks)
