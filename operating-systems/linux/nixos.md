# [NixOS](https://nixos.org)

NixOS is a Linux distro built around the Nix package system. Nix is built around the idea of immutability. It makes all packages immutable by giving them their own directory identified by a hash that is derived from ALL of that package's dependencies. This has a number of desirable properties:

- It makes it trivial to have multiple versions of the same package installed at the same time and allows you to switch between them at will.
- It is trivial to roll back your system after a failed upgrade. Difficult system recovers after you upgrade to a new unstable version are a thing of the past.
- Non-privileged users can install software completely securely.
- Projects packaged with nix have the best possible build reproducibility because nix accounts for ALL of your dependencies all the way down to the lowest level system libraries, compilers, etc.

Whilst I don't use NixOS as my primary OS. I use [nix package manager](../../package-managers/nix/nix.md) on macOS. And I am exploring using NixOS for servers I use.

## Nix configs (NixOS)

- [Infrastructure](https://github.com/rvolosatovs/infrastructure) ([How to use it](https://github.com/rvolosatovs/infrastructure/issues/3))
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
- [Roman Gonzalez](https://github.com/roman/nix-dots)
- [Kim Zick](https://github.com/rummik/nixos-config)
- [Eric Bailey](https://github.com/yurrriq/dotfiles)
- [Martin Baillie](https://github.com/martinbaillie/dotfiles)
- [Alex Ermolov](https://github.com/wiedzmin/nixos-config)
- [Bastian Kocher](https://github.com/bkchr/nixos-config)
- [Structured flake-based NixOS configuration by Tony O](https://github.com/bqv/nixrc)
- [chessai](https://github.com/chessai/nixos-configs)
- [Bruno Bigras](https://github.com/bbigras/nix-config)
- [Thomas Honeyman](https://github.com/thomashoneyman/.dotfiles)
- [Jorg Thalheim](https://github.com/Mic92/dotfiles)

## Nix configs (macOS)

- [John Wiegley](https://github.com/jwiegley/nix-config)
- [LnL7](https://github.com/LnL7/dotfiles)
- [cmacrae](https://github.com/cmacrae/.nixpkgs/blob/master/darwin-configuration.nix)
- [Tom's nix-configs](https://github.com/nocoolnametom/nix-configs)
- [Phil Pluckthun](https://github.com/kitten/nix-system)

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
- [Getting started with NixOS on Raspberry Pi 3 Model B+](https://github.com/zupo/nix)
- [Collection of NixOS image builders](https://github.com/nix-community/nixos-generators) - Allows to take the same NixOS configuration, and generate outputs for different target formats.
- [HN: Guix An advanced operating system (2019)](https://news.ycombinator.com/item?id=18902823)
- [NixOS on ARM](https://github.com/illegalprime/nixos-on-arm) - WIP to cross compile NixOS to run on ARM targets.
- [Arion](https://github.com/hercules-ci/arion) - Run docker-compose with help from Nix/NixOS.
- [Morph](https://github.com/DBCDK/morph) - Tool for managing existing NixOS hosts.
- [Mobile NixOS](https://github.com/samueldr/mobile-nixos) - Goal is to get a nix-built operating system, preferably NixOS running on mobile devices, e.g. Android phones.
- [Anyone using NixOS as main desktop (2019)](https://www.reddit.com/r/NixOS/comments/eb5nxv/anyone_using_nixos_as_main_destkop/)
- [NixOS: For developers (2020)](https://myme.no/posts/2020-01-26-nixos-for-development.html) ([Lobsters](https://lobste.rs/s/jevfaf/nixos_for_developers))
- [nixos-shell](https://github.com/Mic92/nixos-shell) - Spawns lightweight nixos vms in a shell.
- [Erase your darlings (2020)](https://grahamc.com/blog/erase-your-darlings) ([HN](https://news.ycombinator.com/item?id=22856199)) ([Lobsters](https://lobste.rs/s/2ayklq/erase_your_darlings_immutable))
- [Building a web app with functional programming - NixOS (2020)](https://blog.patchgirl.io/nixos/2020/03/31/nixos.html) ([HN](https://news.ycombinator.com/item?id=22877355))
- [nixos-manager](https://github.com/pmiddend/nixos-manager) - Manage your NixOS graphically.
- [Search NixOS packages and options](https://search.nixos.org/) ([Code](https://github.com/NixOS/nixos-search))
- [My NixOS Desktop Flow (2020)](https://christine.website/blog/nixos-desktop-flow-2020-04-25) ([HN](https://news.ycombinator.com/item?id=22984639)) ([Lobsters](https://lobste.rs/s/yb1oqg/my_nixos_desktop_flow))
- [Is NixOS Reproducible?](https://r13y.com/)
- [Impermanence](https://github.com/nix-community/impermanence) - Modules to help you handle persistent state on systems with ephemeral root storage.
- [NixOS Weekly Newsletter](https://weekly.nixos.org/) - Stay up to date with events, learning resources, and recent developments in NixOS community. ([Code](https://github.com/NixOS/nixos-weekly))
- [Nix Community Infrastructure](https://github.com/nix-community/infra)
- [NixOS: How it works and how to install it (2020)](https://www.youtube.com/watch?v=oPymb2-IXbg)
- [Nix(OS) Thoughts (2020)](https://blog.qtp2t.club/posts/2020-06-20-nix-nixos-thoughts/) ([Lobsters](https://lobste.rs/s/iy17mo/nix_os_thoughts))
- [Lightweight Linux VMs on NixOS (2020)](https://www.srid.ca/2012301.html)
- [adhoc executable patching on nixos (2020)](https://notes.neeasade.net/adhoc-executable-patching-on-nix.html)
- [Nix Flakes: Managing NixOS systems (2020)](https://www.tweag.io/blog/2020-07-31-nixos-flakes/)
- [NixOS Channels (2020)](https://nixos.online/posts/NixOS_channels/)
- [Building and Importing NixOS AMIs on EC2 (2020)](http://jackkelly.name/blog/archives/2020/08/30/building_and_importing_nixos_amis_on_ec2/)
- [Tailscale is magic; even more so with NixOS (2020)](https://fzakaria.com/2020/09/17/tailscale-is-magic-even-more-so-with-nixos.html)
- [Secure, Declarative Key Management with NixOps, Pass, and nix-plugins (2018)](https://elvishjerricco.github.io/2018/06/24/secure-declarative-key-management.html)
- [nixos-install-scripts](https://github.com/nix-community/nixos-install-scripts) - Collection of one-shot scripts to install NixOS on various server hosters and other hardware.
- [NixOS on prgmr and Failing to Learn Nix (2018)](https://push.cx/2018/nixos) ([Lobsters](https://lobste.rs/s/qpbohs/nixos_on_prgmr_failing_learn_nix_2018))
- [Eight Months of NixOS (2020)](https://catgirl.ai/log/nixos-experience/) ([Lobsters](https://lobste.rs/s/7eq5qv/eight_months_nixos))
- [One Week of NixOS (2020)](https://jae.moe/blog/2020/11/one-week-of-nixos/) ([HN](https://news.ycombinator.com/item?id=25024639)) ([Lobsters](https://lobste.rs/s/b7hjcy/one_week_nixos))
- [Nixops Services on Your Home Network (2020)](https://christine.website/blog/nixops-services-2020-11-09) ([Lobsters](https://lobste.rs/s/kzforn/nixops_services_on_your_home_network))
- [sops-nix](https://github.com/Mic92/sops-nix) - Atomic secret provisioning for NixOS based on sops.
- [NixOS infrastructure configurations](https://github.com/NixOS/nixos-org-configurations)
- [nix-ld](https://github.com/Mic92/nix-ld) - Run unpatched dynamic binaries on NixOS.
- [Nix(OS) Thoughts (2020)](https://blog.knightsofthelambdacalcul.us/posts/2020-06-20-nix-nixos-thoughts/) ([Lobsters](https://lobste.rs/s/m3j4yn/nix_os_thoughts))
- [NixOps AWS Plugin](https://github.com/NixOS/nixops-aws) - Tool for deploying NixOS machines in a network or cloud.
- [nixflk](https://github.com/nrdxp/nixflk) - Highly structured NixOS configuration database.
- [NixOS Pre Installer](https://github.com/alexandergall/nixos-pxe-installer) - Set of modules to perform a fully automated installation of a customized NixOS system.
- [krops](https://github.com/krebs/krops) - Lightweight toolkit to deploy NixOS systems, remotely or locally.
- [TangeriNixOS](https://github.com/Pamplemousse/tangerinixos) - NixOS tailored for pentesting.
- [Methods for building custom NixOS AMIs](https://github.com/nh2/nixos-ami-building)
- [Colmena](https://github.com/zhaofengli/colmena) - Simple, stateless NixOS deployment tool modeled after NixOps and Morph, written in Rust.
- [NixOS in the Cloud, step-by-step (2020)](https://justinas.org/nixos-in-the-cloud-step-by-step-part-1)
- [Mayflower Nix Consulting](https://nixos.mayflower.consulting/)
- [agenix](https://github.com/ryantm/agenix) - age-encrypted secrets for NixOS.
- [nix-autobahn](https://github.com/Lassulus/nix-autobahn) - Allows you to download ELF binaries and use them right away in NixOS.
- [Why isn't NixOS more popular (2021)](https://www.reddit.com/r/NixOS/comments/kpntby/why_isnt_nixos_more_popular/)
- [Server-optimized NixOS](https://github.com/arianvp/server-optimised-nixos) - Distribution inspired by NixOS, ChromeOS, Container Optimised Linux and Container Linux. Opinionated, server-first distribution.
- [HN: NixOS Linux (2021)](https://news.ycombinator.com/item?id=25718098)
- [NixOps backend for Google Cloud](https://github.com/nix-community/nixops-gce) - Tool for deploying NixOS machines in a network or cloud.
- [Nixus](https://github.com/Infinisil/nixus) - Experimental deployment tool for multiple NixOS systems.
- [nixos.org website code](https://github.com/NixOS/nixos-homepage)
- [Encrypted Secrets with NixOS (2021)](https://christine.website/blog/nixos-encrypted-secrets-2021-01-20) ([Lobsters](https://lobste.rs/s/gur8yy/encrypted_secrets_with_nixos))
- [Offloading NixOS builds to a faster machine (2021)](https://sgt.hootr.club/molten-matter/nix-distributed-builds/)
- [Why you should never ever use NixOS (2021)](https://hands-on.cloud/why-you-should-never-ever-use-nixos/) ([Lobsters](https://lobste.rs/s/f6i7g0/why_you_should_never_ever_use_nixos)) ([HN](https://news.ycombinator.com/item?id=25881654))
- [NixOS on Raspberry Pi](https://github.com/lucernae/nixos-pi)
- [Learning nix workshop](https://github.com/spacekookie/nixos-workshops)
- [Using NixOS as a router (2021)](https://francis.begyn.be/blog/nixos-home-router)
- [Immutable Systems Infrastructure, or how to mashup Kubernetes and Nix](https://tevps.net/blog/2021/2/20/immutable-systems-infrastructure-or-how-mashup-kub/) ([Lobsters](https://lobste.rs/s/h8ejms/immutable_systems_infrastructure_how))
- [NixOS Flake Example](https://github.com/colemickens/nixos-flake-example) - Demo NixOS config, with optional flakes support. Along with notes on why flakes is useful and worth adopting.
- [Bitte](https://github.com/input-output-hk/bitte) - Nix Ops for Terraform, Consul, Vault, Nomad.
- [nixos-up](https://github.com/samuela/nixos-up) - Fastest NixOS install there is.
- [Declarative Docker Container Service in NixOS (2020)](https://www.breakds.org/post/declarative-docker-in-nixos/)
