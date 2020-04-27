# [NixOS](https://nixos.org)

NixOS is a Linux distro built around the Nix package system. Nix is built around the idea of immutability. It makes all packages immutable by giving them their own directory identified by a hash that is derived from ALL of that package's dependencies. This has a number of desirable properties:

- It makes it trivial to have multiple versions of the same package installed at the same time and allows you to switch between them at will.
- It is trivial to roll back your system after a failed upgrade. Difficult system recovers after you upgrade to a new unstable version are a thing of the past.
- Non-privileged users can install software completely securely.
- Projects packaged with nix have the best possible build reproducibility because nix accounts for ALL of your dependencies all the way down to the lowest level system libraries, compilers, etc.

Whilst I don't use NixOS as my primary OS. I use [nix package manager](../../package-managers/nix/nix.md) on macOS. And I am exploring using NixOS for servers I use.

## Nix configs (NixOS)

- [Infrastructure](https://github.com/rvolosatovs/infrastructure#readme) ([How to use it](https://github.com/rvolosatovs/infrastructure/issues/3))
- [Bob nix-home](https://github.com/bobvanderlinden/nix-home)
- [Brian McKenna Nix Files](https://github.com/puffnfresh/nix-files) - NixOS configuration and custom Nix derivations.
- [William A. Kennington III](https://github.com/wkennington/nixos) - NixOS configurations for my local cluster of machines.
- [Arian van Putten](https://github.com/arianvp/nixos-stuff)
- [Michael Peyton Jones](https://github.com/michaelpj/nixos-config)
- [Grégoire Martinache's Infrastructure](https://github.com/M-Gregoire/infrastructure)
- [Silvan Mosberger](https://github.com/Infinisil/system)
- [Vincent Ambo's depot](https://github.com/tazjin/depot) - Personal monorepo of my services & tools.
- [Aaron Janse](https://github.com/aaronjanse/dotfiles)
- [Sridhar Ratnakumar](https://github.com/srid/nix-config)
- [homelab](https://github.com/danderson/homelab) - NixOS configurations for home servers. ([Tweet](https://twitter.com/dave_universetf/status/1236634753765269512))
- [Henrik Lissner](https://github.com/hlissner/dotfiles)
- [Edmund Wu](https://github.com/eadwu/nixos-configuration)

## Nix configs (macOS)

- [John Wiegley](https://github.com/jwiegley/nix-config)
- [LnL7](https://github.com/LnL7/dotfiles#readme)
- [cmacrae](https://github.com/cmacrae/.nixpkgs/blob/master/darwin-configuration.nix)
- [Tom's nix-configs](https://github.com/nocoolnametom/nix-configs)

## Notes

- [Even if you curate your system, it gathers dust: configuration files left to rot, manually installed packages that didn't get uninstalled properly, orphaned packages difficult to track down... You could argue that it shouldn't happen in the first place, but that actually takes discipline. In NixOS, this is managed for you. Once you do nixos collect-garbage -d, you know that your system is only left with what it needs. Nothing more, nothing less.](https://www.reddit.com/r/NixOS/comments/441ymh/nixos_users_tell_me_what_are_the_cons/czmu9lo/)

## Links

- [PhD thesis on nixOS](https://nixos.org/~eelco/pubs/phd-thesis.pdf)
- [Search NixOS options](https://nixos.org/nixos/options.html#)
- [Notes on nixOS package manager](https://yoshuawuyts.gitbooks.io/knowledge/content/bin/nix.html)
- [Why nixOS?](https://www.reddit.com/r/NixOS/comments/8bxdyu/why_nixos/)
- [Not OS](https://github.com/cleverca22/not-os) - Operating system generator, based on NixOS, that, given a config, outputs a small (47 MB), read-only squashfs for a runit-based operating system, with support for iPXE and signed boot.
- [NixOS 💜 Chromebook?](https://sphalerite.org/ghotl/posts/2017-11-10-chromebook.html)
- [NixOS Wiki](https://nixos.wiki/wiki/Main_Page)
- [NixOps](https://github.com/NixOS/nixops) - NixOS-based cloud deployment tool.
- [NixOS Discourse forum](https://discourse.nixos.org/)
- [Getting started with NixOS on Raspberry Pi 3 Model B+](https://github.com/zupo/nix#readme)
- [Collection of NixOS image builders](https://github.com/nix-community/nixos-generators) - Allows to take the same NixOS configuration, and generate outputs for different target formats.
- [HN: Guix An advanced operating system (2019)](https://news.ycombinator.com/item?id=18902823)
- [NixOS on ARM](https://github.com/illegalprime/nixos-on-arm) - WIP to cross compile NixOS to run on ARM targets.
- [Arion](https://github.com/hercules-ci/arion) - Run docker-compose with help from Nix/NixOS.
- [Morph](https://github.com/DBCDK/morph) - Tool for managing existing NixOS hosts.
- [Mobile NixOS](https://github.com/samueldr/mobile-nixos) - Goal is to get a nix-built operating system, preferably NixOS running on mobile devices, e.g. Android phones.
- [Anyone using NixOS as main destkop (2019)](https://www.reddit.com/r/NixOS/comments/eb5nxv/anyone_using_nixos_as_main_destkop/)
- [NixOS: For developers (2020)](https://myme.no/posts/2020-01-26-nixos-for-development.html) ([Lobsters](https://lobste.rs/s/jevfaf/nixos_for_developers))
- [nixos-shell](https://github.com/Mic92/nixos-shell) - Spawns lightweight nixos vms in a shell.
- [Erase your darlings (2020)](https://grahamc.com/blog/erase-your-darlings) ([HN](https://news.ycombinator.com/item?id=22856199)) ([Lobsters](https://lobste.rs/s/2ayklq/erase_your_darlings_immutable))
- [Building a web app with functional programming - NixOS (2020)](https://blog.patchgirl.io/nixos/2020/03/31/nixos.html) ([HN](https://news.ycombinator.com/item?id=22877355))
- [nixos-manager](https://github.com/pmiddend/nixos-manager) - Manage your NixOS graphically.
- [Search NixOS packages and options](https://search.nixos.org/) ([Code](https://github.com/NixOS/nixos-search))
- [My NixOS Desktop Flow (2020)](https://christine.website/blog/nixos-desktop-flow-2020-04-25) ([HN](https://news.ycombinator.com/item?id=22984639)) ([Lobsters](https://lobste.rs/s/yb1oqg/my_nixos_desktop_flow))
