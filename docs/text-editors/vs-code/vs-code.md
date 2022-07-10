---
title: VS Code
---

# [VS Code](https://github.com/Microsoft/vscode)

My favorite editor that I use to write code in. I use [many extensions](vs-code-extensions.md) for it.

My config for it can be found [here](https://github.com/nikitavoloboev/dotfiles/blob/master/vscode/settings.json).

I use VS Code [GitHub theme](https://marketplace.visualstudio.com/items?itemName=GitHub.github-vscode-theme) with [JetBrains Mono](https://www.jetbrains.com/lp/mono/) font. Here is how it looks:

I switch between GitHub Light and GitHub Dark themes as I change between macOS appearances.

![](https://i.imgur.com/ZZTGK1A.png)

> GitHub Light

![](https://i.imgur.com/rd4Ve3X.png)

> GitHub Dark

## Notes

- [VS Code is architected in a way where extensions are not eagerly activated by default. Each extension can declare a list of activation events, such as e.g. opening a file of a certain language, invoking a specific command, starting debugging, etc.](https://news.ycombinator.com/item?id=16940419)
- `Developer: Show Running Extensions` command -> Shows currently running extensions. Is good for profiling.
- [VS Code has excellent integrated node debugging. It integrates seamlessly with the entire tool ecosystem (eg I use ts-node-dev for autoreloading + typescript support, and the VS Code debugger Just Works). And because it's inside the editor, ndb features like the ability to put breakpoints in a file before it is required are irrelevant. It's all at your fingertips, just put a breakpoint right where you're coding, hit F5 to attach the debugger to your devserver and step through the code.](https://news.ycombinator.com/item?id=17581521)
- To hide non useful `Outline` section in sidebar, just open the context menu on the outline section's header and select the "Hide" action.
- [Can use "\*" in the advanced search options in @code to search for similar subfolders across multiple parent folders.](https://twitter.com/SteveGodderidge/status/1285978384049348609)
- [Can Keyboard map Emmet commands: inward and outward to select HTML/JSX blocks](https://twitter.com/jaredpalmer/status/1385938591323414529)
- [When you're writing markdown in VSCode you can just highlight text and ctrl+v to automatically make it a link.](https://twitter.com/hollylawly/status/1398375511366115328)
- [VSCode can let you run Jupyter Notebooks line by line and see variables](https://www.youtube.com/shorts/zTxmCVGaYic)

## Links

- [VS Code Docs](https://code.visualstudio.com/docs)
- [Introductory Videos](https://code.visualstudio.com/docs/getstarted/introvideos)
- [Visual Studio Live Share Docs & Feedback](https://github.com/MicrosoftDocs/live-share)
- [code-server](https://github.com/codercom/code-server) - Run VS Code on a remote server.
- [sshcode](https://github.com/codercom/sshcode) - CLI to automatically install and run code-server over SSH.
- [Vim to Code](https://github.com/asantos00/vim-to-code) - Comprehensive (almost) list of resources, tutorials, and inspiration for migrating to Visual Studio code from Vim.
- [VSCode Keyboard Shortcuts For Productivity (2019)](https://www.youtube.com/watch?v=Xa5EU-qAv-I)
- [VS Code Recipes](https://github.com/microsoft/vscode-recipes) - Collection of recipes for using VS Code with particular technologies.
- [Snippet Generator](https://snippet-generator.app/) - Generate snippets for VSCode/SublimeText. ([Code](https://github.com/pawelgrzybek/snippet-generator))
- [VS Code Tips](https://www.youtube.com/channel/UCyYh-eAr74avLwOyPa1dDNg/videos)
- [code-server](https://github.com/cdr/code-server) - Run VS Code on a remote server.
- [sshcode](https://github.com/cdr/sshcode) - CLI to automatically install and run code-server over SSH.
- [sail](https://github.com/cdr/sail) - Instant, pre-configured VS Code development environments.
- [StackBlitz](https://stackblitz.com/) - Online IDE powered by Visual Studio Code. ([Code](https://github.com/stackblitz/core)) ([GitHub](https://github.com/stackblitz))
- [The best Parts of Visual Studio Code are proprietary (2020)](https://underjord.io/the-best-parts-of-visual-studio-code-are-proprietary.html) ([HN](https://news.ycombinator.com/item?id=24047638))
- [VS Code Rocks](https://vscode.rocks/) - Weekly news on the newest features and updates to VSCode.
- [VSCode to Prism Themes](https://prism.dotenv.dev/)
- [The Era of Visual Studio Code (2020)](https://blog.robenkleene.com/2020/09/21/the-era-of-visual-studio-code/) ([HN](https://news.ycombinator.com/item?id=24558788)) ([Lobsters](https://lobste.rs/s/d2uhbm/era_visual_studio_code))
- [My new coding workflow: VS Code + Remote-SSH extension (2020)](https://jlelse.blog/dev/code-using-vps/) ([Lobsters](https://lobste.rs/s/mqclc7/my_new_coding_workflow_vs_code_remote_ssh))
- [Cloud Code for Visual Studio Code](https://github.com/GoogleCloudPlatform/cloud-code-vscode) - Brings the power and convenience of IDEs to developing cloud-native Kubernetes and Cloud Run applications.
- [GitHub and VS Code](https://vscode.github.com/)
- [VS Code Tips](https://twitter.com/vscodetips)
- [VS Code Wiki](https://github.com/Microsoft/vscode/wiki) ([Code](https://github.com/microsoft/vscode-wiki))
- [Learning all VSCode shortcuts evolved my developing habits (2020)](https://tkainrad.dev/posts/learning-all-vscode-shortcuts-evolved-my-developing-habits/)
- [VSCode Dev Containers](https://github.com/microsoft/vscode-dev-containers) - VS Code Remote / GitHub Codespaces Container Definitions.
- [VS Code Day Event](https://code.visualstudio.com/vscode-day) ([2021 Talks](https://www.youtube.com/playlist?list=PLj6YeMhvp2S6uB23beQaffszlavLq3lNq))
- [VS Code JavaScript snippets](https://github.com/xabikos/vscode-javascript)
- [VSCode Learn Docs](https://code.visualstudio.com/learn)
- [VS Code can do that?](https://vscodecandothat.com/) - All the best things about Visual Studio Code that nobody ever bothered to tell you.
- [Hacking VSCode - fun side projects that boost productivity (2021)](https://www.youtube.com/watch?v=XY9MaaR1dRI)
- [Make VS Code Awesome](https://makevscodeawesome.com/)
- [snp](https://github.com/djyde/snp) - VS Code code snippet generator.
- [HN: Reflections on IDEA vs VS Code (2021)](https://news.ycombinator.com/item?id=26367410)
- [Productive VS Code](https://productivevscode.com/) - Change the way you VS Code in 1 hour.
- [OpenVSCode Server](https://github.com/gitpod-io/openvscode-server) - Run upstream VS Code on a remote machine with access through a modern web browser from any device, anywhere. ([Article](https://www.gitpod.io/blog/openvscode-server-launch))
- [How OpenVSCode Server turns VS Code into a web IDE](https://sourcegraph.com/github.com/gitpod-io/openvscode-server/-/blob/docs/sourcedive.snb.md) ([HN](https://news.ycombinator.com/item?id=28685978))
- [How We Made Bracket Pair Colorization 10,000x Faster (2021)](https://code.visualstudio.com/blogs/2021/09/29/bracket-pair-colorization) ([HN](https://news.ycombinator.com/item?id=28692470)) ([Lobsters](https://lobste.rs/s/ujj1cu/how_we_made_bracket_pair_colorization_10)) ([Tweet](https://twitter.com/kmett/status/1443213748840325132))
- [VSCode.dev](https://code.visualstudio.com/blogs/2021/10/20/vscode-dev) - Online VSCode Editor. ([Article](https://code.visualstudio.com/blogs/2021/10/20/vscode-dev)) ([HN](https://news.ycombinator.com/item?id=28932206))
- [Ask HN: Could VSCode be the new Emacs? (2021)](https://news.ycombinator.com/item?id=29159192)
- [VSCode Web](https://github.com/Felx-B/vscode-web) - Visual Studio Code for browser.
- [nix-devcontainer](https://github.com/xtruder/nix-devcontainer) - Swiss army knife container for vscode development environments.
- [File Nesting Config for VS Code](https://github.com/antfu/vscode-file-nesting-config)
- [CodeTerminal](https://github.com/xcodebuild/CodeTerminal) - Cross platform terminal app from Visual Studio Code.
- [VS Code's Issue Triage GitHub Actions](https://github.com/microsoft/vscode-github-triage-actions)
- [VSCodium](https://vscodium.com/) - Open Source Binaries of VSCode. ([Code](https://github.com/VSCodium/vscodium)) ([HN](https://news.ycombinator.com/item?id=31604932))
- [Visual Studio Code Community Discussions](https://github.com/microsoft/vscode-discussions)
- [Monaco VSCode API](https://github.com/CodinGame/monaco-vscode-api) - NPM module that implements the VSCode api and redirects calls to Monaco editor.
- [Vim for VSCode](https://vimforvscode.com/) - Learn to use Vim inside your favorite editor.
- [The Visual Studio Code Server](https://code.visualstudio.com/blogs/2022/07/07/vscode-server) ([HN](https://news.ycombinator.com/item?id=32024882))
