# [VS Code](https://github.com/Microsoft/vscode)
My favorite editor that I use to write code in. I use [many extensions](vs-code-extensions.md) for it.

My config for it can be found [here](https://github.com/nikitavoloboev/dotfiles/blob/master/vscode/settings.json).

I use [Monokai Octagon](https://www.monokai.pro/vscode/) theme with [Fira Code](https://github.com/tonsky/FiraCode) font. Here is how it looks:

![](https://i.imgur.com/hss9CH8.png)

## Notes
- [VS Code is architected in a way where extensions are not eagerly activated by default. Each extension can declare a list of activation events, such as e.g. opening a file of a certain language, invoking a specific command, starting debugging, etc.](https://news.ycombinator.com/item?id=16940419)
- `Developer: Show Running Extensions` command -> Shows currently running extensions. Is good for profiling.
- [VS Code has excellent integrated node debugging. It integrates seamlessly with the entire tool ecosystem (eg I use ts-node-dev for autoreloading + typescript support, and the VS Code debugger Just Works). And because it's inside the editor, ndb features like the ability to put breakpoints in a file before it is required are irrelevant. It's all at your fingertips, just put a breakpoint right where you're coding, hit F5 to attach the debugger to your devserver and step through the code.](https://news.ycombinator.com/item?id=17581521)

## Links
- [VS Code Docs](https://code.visualstudio.com/docs)
- [Introductory Videos](https://code.visualstudio.com/docs/getstarted/introvideos)