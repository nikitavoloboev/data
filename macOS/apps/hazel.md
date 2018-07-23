# [Hazel](https://www.noodlesoft.com)
I currently use Hazel to instantly commit any changes I make to all the [curated lists](https://github.com/learn-anything/curated-lists#readme) I maintain. As well as CSV files of [Web Searches](https://github.com/nikitavoloboev/alfred-web-searches#readme) and [Ask Create Share](https://github.com/nikitavoloboev/alfred-ask-create-share#readme).

I also automate committing the README of [my macOS](https://github.com/nikitavoloboev/my-mac-os#readme) & [iOS](https://github.com/nikitavoloboev/my-ios#readme) repo as I want to instantly push any changes I make to the repo.

Since I want to keep my macOS repo always updated, I made a macro to open the README file in Sublime Text so I can quickly make a change, save and the change will instantly be committed with `update readme` message.

The Hazel rule for this is simple and looks like this:
![](https://i.imgur.com/EF3elcv.png)

With this as the shell script:
![](https://i.imgur.com/9FgVmxm.png)

Here it is in code:
```bash
git add README.md
git commit -m "update readme"
git push
```

To have Hazel know about commands I installed with Nix. I need to run this code first in the shell script:
```bash
export USER=nikitavoloboev
if [ -e '/nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh' ]; then
   . '/nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh'
fi

```

## Notes
- If you want to have your rules to be applied onto subfolders as well as the directory chosen, add this rule.
![](https://i.imgur.com/yPfhkBo.png)

- If extensions are hidden, don't add extensions in the rule editor too.

## Links
- [Reference hazel file path](https://forum.keyboardmaestro.com/t/reference-hazels-file-path/9138)
- [Hazel Debug Mode](https://www.noodlesoft.com/kb/hazel-debug-mode/)