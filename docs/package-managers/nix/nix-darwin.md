# Nix on macOS

There is [nix-darwin](https://github.com/LnL7/nix-darwin) project that brings Nix configuration to describing many of macOS services.

## Notes

- To not enter password on every `darwin-rebuild switch` with [nix-darwin](https://github.com/LnL7/nix-darwin), you can create `/etc/sudoers.d/nix-darwin` file with this content:
  - `<home-user> ALL=(ALL:ALL) NOPASSWD: /run/current-system/sw/bin/darwin-rebuild` (where `home-user` is name of home directory)
- Nix darwin generates files to `etc/static`

## Links

- [Nix on Darwin – History, challenges, and where it's going by Dan Peebles](https://www.youtube.com/watch?v=73mnPBLL_20)
- [Neat macOS nix Darwin config.nix](https://github.com/LnL7/nix-darwin/blob/master/modules/examples/lnl.nix)
- [Changing the configuration.nix location](https://github.com/LnL7/nix-darwin/wiki/Changing-the-configuration.nix-location)
- [Moving from Homebrew to Nix Package Manager (2019)](https://www.softinio.com/post/moving-from-homebrew-to-nix-package-manager/)
- [Nix on Catalina installer](https://static.domenkozar.com/install-catalina) ([GitHub PR](https://github.com/NixOS/nix/pull/3212)) ([Another Gist solution](https://gist.github.com/rabbitonweb/7a12d61294669f0958ebcf1eac347304))
