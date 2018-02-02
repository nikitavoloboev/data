# NixOS
Is interesting. [This](https://nixos.org/~eelco/pubs/phd-thesis.pdf) is a PhD thesis on this OS.

Whilst I don't use NixOS as my primary OS. I am trying to use its [nix package manager](../package-managers/nix.md) on macOS where possible.

## Notes
- Nix never uses host dependencies, it always builds with exactly precise dependencies every time, and will always refer to them from then on.
- [Even if you curate your system, it gathers dust: configuration files left to rot, manually installed packages that didn't get uninstalled properly, orphaned packages difficult to track down... You could argue that it shouldn't happen in the first place, but that actually takes discipline. In NixOS, this is managed for you. Once you do nixos collect-garbage -d, you know that your system is only left with what it needs. Nothing more, nothing less.](https://www.reddit.com/r/NixOS/comments/441ymh/nixos_users_tell_me_what_are_the_cons/czmu9lo/)

## Links
- [Notes on nixOS package manager](https://yoshuawuyts.gitbooks.io/knowledge/content/bin/nix.html)