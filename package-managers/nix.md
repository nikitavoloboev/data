# [Nix](https://github.com/NixOS/nix)
Nix package manager is really interesting. I am in the process of transitioning to using it over homebrew.

## Notes
- Nix lets you roll back changes atomically.
- nix-shell lets you make build environments that are totally reproducible across machines, and don’t interfere with each other. You can freely mix any number of libraries of versions or software on the same machine and they don’t conflict.
- With Ubuntu, every time you want to fix something with your car, you roll it into the garage, pop open the hood and get to work. It's intensive
labour, results will vary, and undoing a change can be really difficult.
	- With NixOS, it's like 3D printing a new car every time. You'll design a model, press a button, and the car gets built from scratch. If you don't like it, tweak the design a bit, and print a new car. If the new car breaks, just go back to the previous known-good one, which is already in your garage. You can even take the design documents to your friend and generate an exactly identical model.
- `sudo` command sets the wrong `$HOME`, have to use `sudo -i` for nix commands that need sudo.
- Nix is Turing complete configuration.

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