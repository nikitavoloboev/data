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
- [Benefits of using nix](https://www.reddit.com/r/haskell/comments/7wmhyi/an_opinionated_guide_to_haskell_in_2018/du2506q/)
- [Nix, the purely functional build system](http://www.boronine.com/2018/02/02/Nix/) - Great intro article.
- [A Gentle Introduction to the Nix Family](https://ebzzry.io/en/nix/)
- [Nix: A Reproducible Setup for Linux and macOS](http://nmattia.com/posts/2018-03-21-nix-reproducible-setup-linux-macos.html)
- [hnix](https://github.com/jwiegley/hnix) - Haskell re-implementation of the Nix expression language.
- [Haskell & Nix](https://github.com/Gabriel439/haskell-nix)
- [Brian McKenna - Nix for Functional Systems](https://www.youtube.com/watch?v=mIxtBVKo7JE)
- [Learning Nix by Example: Building FFmpeg 4.0](https://blog.kiloreux.me/2018/05/24/learning-nix-by-example-building-ffmpeg-4-dot-0/)
- [nix-shell and Shebang Lines](http://iam.travishartwell.net/2015/06/17/nix-shell-shebang/)
- [Domen Kožar, Lead DevOps, Nix workshop](https://www.youtube.com/watch?v=BjRGlKNHeEc)
- [Cheap Docker images with Nix](http://lethalman.blogspot.com/2016/04/cheap-docker-images-with-nix_15.html)
- [When to use declarative approach and when not](https://www.reddit.com/r/NixOS/comments/95vczu/when_to_use_declarative_approach_and_when_not/)
- [example-nix](https://github.com/shajra/example-nix#readme) - A way to develop software with Nix.
- [Hercules](https://hercules-ci.com/) - Hosted CI for building Nix projects on your infrastructure.
- [Dysnomia](https://github.com/svanderburg/dysnomia) - Tool and plug-in system that can be used to automatically deploy mutable components.
- [Disnix](https://github.com/svanderburg/disnix) - Nix-based distributed service deployment tool.
- [NUR](https://github.com/nix-community/NUR) - Nix User Repository: User contributed nix packages.
- [Eris](https://github.com/thoughtpolice/eris) - Binary cache for Nix.
- [pypi2nix](https://github.com/garbas/pypi2nix) - Generate Nix expressions for Python packages.
- [hnix-store](https://github.com/haskell-nix/hnix-store) - Haskell implementation of the nix store API.
- [nix-cheatsheet](https://github.com/knedlsepp/nix-cheatsheet)
- [Nix RFCs](https://github.com/NixOS/rfcs#readme)
- [nix-linter](https://github.com/Synthetica9/nix-linter) - Linter for the Nix expression language.
- [Install Nix docs by Mozilla](https://docs.mozilla-releng.net/develop/install-nix.html) - Pretty good.
- [Nix scripts shared across IOHK projects](https://github.com/input-output-hk/iohk-nix)
- [niv](https://github.com/nmattia/niv) - Painless dependencies for Nix projects.
- [Cachix](https://cachix.org/) - Build Nix packages once and share them for good.
- [Alternative Haskell Infrastructure for Nixpkgs](https://github.com/input-output-hk/haskell.nix) - Works by automatically translating your Cabal or Stack project and its dependencies into Nix code.
- [nix-bundle](https://github.com/matthewbauer/nix-bundle) - Bundle Nix derivations to run anywhere.
- [crate2nix](https://github.com/kolloch/crate2nix) - Nix build file generator for rust crates. ([Lobsters](https://lobste.rs/s/26xnzy/crate2nix_nix_build_file_generator_for))
- [lorri](https://github.com/target/lorri) - nix-shell replacement for project development.
- [Awesome Nix](https://github.com/nix-community/awesome-nix#readme)
- [nixfmt](https://github.com/serokell/nixfmt) - Formatter for Nix code.
- [Nix for devs](https://github.com/uniphil/nix-for-devs) - Collection of recipes focused on nix-shell to make setting up project environments easy.
- [nixpkgs-fmt](https://github.com/nix-community/nixpkgs-fmt) - Nix code formatter for nixpkgs.
- [Nix builder for Kubernetes](https://gist.github.com/tazjin/08f3d37073b3590aacac424303e6f745)
- [Nixery](http://nixery.dev/nixery.html) - Provides the ability to pull ad-hoc container images from a Docker-compatible registry server. ([Code](https://github.com/google/nixery))
- [hnix-lsp](https://github.com/domenkozar/hnix-lsp) - Language Server Protocol for Nix.
- [Make Nix precisely emulate gitignore](https://github.com/hercules-ci/gitignore)
- [Nixery](https://github.com/google/nixery) - Container registry which transparently builds images using the Nix package manager.
- [Nix - A One Pager](https://github.com/tazjin/nix-1p#readme) - A (more or less) one page introduction to Nix, the language.
