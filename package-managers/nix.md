# [Nix](https://github.com/NixOS/nix)
## Notes
- Nix never uses host dependencies, it always builds with exactly precise dependencies every time, and will always refer to them from then on.
- Nix lets you roll back changes atomically.
- nix-shell lets you make build environments that are totally reproducible across machines, and don’t interfere with each other. You can freely mix any number of libraries of versions or software on the same machine and they don’t conflict.
- With Ubuntu, every time you want to fix something with your car, you roll it into the garage, pop open the hood and get to work. It's intensive labour, results will vary, and undoing a change can be difficult.
	- With NixOS, it's like 3D printing a new car every time. You'll design a model, press a button, and the car gets built from scratch. If you don't like it, tweak the design a bit, and print a new car. If the new car breaks, just go back to the previous known-good one, which is already in your garage. You can even take the design documents to your friend and generate an exactly identical model.
- `sudo` command sets the wrong `$HOME`, have to use `sudo -i` for nix commands that need sudo.
- Nix is Turing complete language used for configuration and building packages.
- Can use nox, nix search, nix-repl, [nixOS packages](https://nixos.org/nixos/packages.html#) to search for packages.
- Think of Nix (the language) as an expression-based programming language where every program evaluates to a single (possibly complex) value; that resulting value is what is used as eg. the configuration of your system or a package, but it means that you can generate that value based on arbitrary logic and abstractions like you would with a regular programming language.
- As for domain-specific package managers, It Depends; it's possible with varying degrees of hackiness (and I definitely use eg. npm for development), but for a 'real' deployment - whether as a service on a server or as a local application - you'd want to convert your project's metadata to a Nix expression and let Nix handle the dependency management.
- Overlay adds/overrides something in the global package set.
- In general, you should only install things with nix, and not use any other package managers.
- The main idea of the Nix approach is to store software components in isolation from each other in a central component store, under path names that contain cryptographic hashes of all inputs involved in building the component, such as `/nix/store/rwmfbhb2znwp...-ﬁrefox1.0.4`.
- Don't install libraries with Nix.

## Links
- [Nix manual](https://nixos.org/nix/manual/)
- [Nix pills](https://nixos.org/nixos/nix-pills/index.html)
- [Nix on Darwin – History, challenges, and where it's going by Dan Peebles](https://www.youtube.com/watch?v=73mnPBLL_20)
- [Interesting nix config](https://github.com/jwiegley/nix-config) - macOS based.
- [Benefits of using nix](https://www.reddit.com/r/haskell/comments/7wmhyi/an_opinionated_guide_to_haskell_in_2018/du2506q/)
- [Nix, the purely functional build system](http://www.boronine.com/2018/02/02/Nix/) - Great intro article.
- [A Gentle Introduction to the Nix Family](https://ebzzry.io/en/nix/)
- [Nix: A Reproducible Setup for Linux and macOS](http://nmattia.com/posts/2018-03-21-nix-reproducible-setup-linux-macos.html)
- [hnix](https://github.com/jwiegley/hnix) - Haskell re-implementation of the Nix expression language.
- [Haskell & Nix](https://github.com/Gabriel439/haskell-nix)
- [neat macOS nix darwin config.nix](https://github.com/LnL7/nix-darwin/blob/master/modules/examples/lnl.nix)
- [Brian McKenna - Nix for Functional Systems](https://www.youtube.com/watch?v=mIxtBVKo7JE)
- [Learning Nix by Example: Building FFmpeg 4.0](https://blog.kiloreux.me/2018/05/24/learning-nix-by-example-building-ffmpeg-4-dot-0/)
- [nix-shell and Shebang Lines](http://iam.travishartwell.net/2015/06/17/nix-shell-shebang/)