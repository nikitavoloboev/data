# [Nix](https://github.com/NixOS/nix)

[nix.dev](https://nix.dev/) is great place to start learning/using Nix.

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
- Derivations are variables + build script.
- Docs are in `.xml` files in `/docs`
- To build derivation in nixpkgs, at root of `nixpkgs`, run `nix build -f . <pkg-name>` (ie `nix build -f . wifi-password`)
- To find out the SHA256, run `nix-prefetch-url -A <pkg-name>.src` (ie `nix-prefetch-url -A wifi-password.src`). `nix-prefetch-url` works with GitHub. This gives you the SHA256 you can copy.
- [I think Nix's approach is the right way to build container images--build the image layers declaratively, reproducibly, incrementally (no, Dockerfile builds aren't incremental because dependency trees aren't linear), and without a container runtime dependency.](https://twitter.com/weberc2/status/1334927112550174721)
- [Nix doesn't solve dependency resolution problems. Only pinning. There's ground to break there. Dependency management not being part of flakes is my biggest gripe with it. It would be our chance to be the once size fits all solution but we failed to deliver.](https://twitter.com/ProgrammerDude/status/1375451276234928132)
- [There's a lot of unexplored potential of Nix in granular build systems and displacing systems like Bazel. If applied correctly, it lets smaller organisations get much of the benefit of Google-style monorepos but without as much maintenance overhead.](https://news.ycombinator.com/item?id=26748696)

## Code

```bash
# Build nix package locally.

# cd into cloned https://github.com/NixOS/nixpkgs

# Build package from default.nix inside nixpkgs. Will put result as ./result if succeeds
# i.e. nix-build -A watchexec -> will build watchexec package
nix-build -A <package>

# Install the build and put it `~/.nix-profile/bin`
nix-env -i ./result
```

```bash
# Garbage collect
sudo nix-collect-garbage --delete-older-than 30d
```

## Links

- [Nix Manual](https://nixos.org/manual/nix/stable/)
- [Nix Pills](https://nixos.org/guides/nix-pills/) ([Code](https://github.com/NixOS/nix-pills))
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
- [example-nix](https://github.com/shajra/example-nix) - A way to develop software with Nix.
- [Hercules CI](https://hercules-ci.com/) - Hosted CI for building Nix projects on your infrastructure.
- [Dysnomia](https://github.com/svanderburg/dysnomia) - Tool and plug-in system that can be used to automatically deploy mutable components.
- [Disnix](https://github.com/svanderburg/disnix) - Nix-based distributed service deployment tool.
- [NUR](https://github.com/nix-community/NUR) - Nix User Repository: User contributed nix packages.
- [Eris](https://github.com/thoughtpolice/eris) - Binary cache for Nix.
- [pypi2nix](https://github.com/garbas/pypi2nix) - Generate Nix expressions for Python packages.
- [hnix-store](https://github.com/haskell-nix/hnix-store) - Haskell implementation of the nix store API.
- [nix-cheatsheet](https://github.com/knedlsepp/nix-cheatsheet)
- [Nix RFCs](https://github.com/NixOS/rfcs)
- [nix-linter](https://github.com/Synthetica9/nix-linter) - Linter for the Nix expression language.
- [Install Nix docs by Mozilla](https://docs.mozilla-releng.net/develop/install-nix.html) - Pretty good.
- [Nix scripts shared across IOHK projects](https://github.com/input-output-hk/iohk-nix)
- [niv](https://github.com/nmattia/niv) - Painless dependencies for Nix projects.
- [Cachix](https://cachix.org/) - Build Nix packages once and share them for good.
- [Alternative Haskell Infrastructure for Nixpkgs](https://github.com/input-output-hk/haskell.nix) - Works by automatically translating your Cabal or Stack project and its dependencies into Nix code.
- [nix-bundle](https://github.com/matthewbauer/nix-bundle) - Bundle Nix derivations to run anywhere.
- [crate2nix](https://github.com/kolloch/crate2nix) - Nix build file generator for rust crates. ([Lobsters](https://lobste.rs/s/26xnzy/crate2nix_nix_build_file_generator_for))
- [lorri](https://github.com/nix-community/lorri) - nix-shell replacement for project development.
- [Awesome Nix](https://github.com/nix-community/awesome-nix)
- [nixfmt](https://github.com/serokell/nixfmt) - Formatter for Nix code.
- [Nix for devs](https://github.com/uniphil/nix-for-devs) - Collection of recipes focused on nix-shell to make setting up project environments easy.
- [nixpkgs-fmt](https://github.com/nix-community/nixpkgs-fmt) - Nix code formatter for nixpkgs.
- [Nix builder for Kubernetes](https://gist.github.com/tazjin/08f3d37073b3590aacac424303e6f745)
- [Nixery](https://nixery.dev/) - Provides the ability to pull ad-hoc container images from a Docker-compatible registry server. ([Code](https://github.com/tazjin/nixery)) ([Talk](https://www.youtube.com/watch?v=pOI9H4oeXqA))
- [Nixery: Improved Layering Design (2019)](https://tazj.in/blog/nixery-layers)
- [hnix-lsp](https://github.com/domenkozar/hnix-lsp) - Language Server Protocol for Nix.
- [Make Nix precisely emulate gitignore](https://github.com/hercules-ci/gitignore.nix)
- [Nixery](https://github.com/google/nixery) - Container registry which transparently builds images using the Nix package manager.
- [wharfix](https://github.com/google/nixery) - Minimal stateless+readonly docker registry based on nix expressions. Heavily inspired by Nixery.
- [Nix - A One Pager](https://github.com/tazjin/nix-1p) - A (more or less) one page introduction to Nix, the language.
- [yants](https://github.com/tazjin/yants) - Tiny type-checker for data in Nix, written in Nix.
- [nix-shorts](https://github.com/justinwoo/nix-shorts) - Collection of short notes about Nix, down to what is immediately needed for users.
- [rnix-parser](https://github.com/nix-community/rnix-parser) - Nix parser written in Rust.
- [Naersk](https://github.com/nmattia/naersk) - Nix support for building cargo crates.
- [nix-diff](https://github.com/Gabriel439/nix-diff) - Explain why two Nix derivations differ.
- [Optimising Docker Layers for Better Caching with Nix (2018)](https://grahamc.com/blog/nix-and-layered-docker-images)
- [nix-du](https://github.com/symphorien/nix-du) - Visualise which gc-roots to delete to free some space in your nix store.
- [NixCon 2019 - Main Track](https://www.youtube.com/watch?v=aUG9aGYYCY8)
- [Nix flakes (2019)](https://www.youtube.com/watch?v=UeBX7Ide5a0)
- [Nix: How and Why it Works (2019)](https://www.youtube.com/watch?v=lxtHH838yko)
- [Nix recipes for Haskellers](https://www.srid.ca/haskell-nix.html)
- [format-nix](https://github.com/justinwoo/format-nix) - Simple formatter for Nix using tree-sitter-nix.
- [go-nix](https://github.com/orivej/go-nix) - Nix language parser and Nix-compatible file hasher in Go.
- [nix-dns](https://github.com/kirelagin/nix-dns) - Nix DSL for DNS zone files.
- [Nix-based app VMs](https://github.com/jollheef/appvm)
- [nix-index](https://github.com/bennofs/nix-index) - Quickly locate nix packages with specific files.
- [Nix-bisect](https://github.com/timokau/nix-bisect) - Bisect Nix Builds.
- [Thoughts on Nix (2020)](https://christine.website/blog/thoughts-on-nix-2020-01-28) ([Lobsters](https://lobste.rs/s/tp6o0q/thoughts_on_nix))
- [I was Wrong about Nix (2020)](https://christine.website/blog/i-was-wrong-about-nix-2020-02-10) ([HN](https://news.ycombinator.com/item?id=22295102)) ([Lobsters](https://lobste.rs/s/4otqzp/i_was_wrong_about_nix))
- [Grafanix](https://github.com/stolyaroleh/grafanix) - Visualize your Nix dependencies.
- [What's your configuration.nix like?](https://www.reddit.com/r/NixOS/comments/9aa08b/whats_your_configurationnix_like/)
- [Built with Nix](https://builtwithnix.org/) - Build software only once. ([Code](https://github.com/nix-community/builtwithnix.org))
- [How I Start: Nix (2020)](https://christine.website/blog/how-i-start-nix-2020-03-08) ([Lobsters](https://lobste.rs/s/lktf6u/how_i_start_nix))
- [rnix-lsp](https://github.com/nix-community/rnix-lsp) - WIP Language Server for Nix.
- [Language server for Nix language](https://github.com/ebkalderon/nix-language-server)
- [Eelco Dolstra's talks/papers](https://edolstra.github.io/) ([Code](https://github.com/edolstra/edolstra.github.io))
- [Nix Haskell Monorepo Tutorial](https://github.com/fghibellini/nix-haskell-monorepo)
- [The journey of packaging a .NET app on Nix (2020)](https://sgt.hootr.club/molten-matter/dotnet-on-nix/)
- [nixpkgs](https://github.com/NixOS/nixpkgs) - Nix Packages collection.
- [Nix IRC logs](https://logs.nix.samueldr.com/)
- [Nixology (2020)](https://www.youtube.com/playlist?list=PLRGI9KQ3_HP_OFRG6R-p4iFgMSK1t5BHs) - Series of videos I've been releasing within Shopify to help promote and educate about Nix.
- [Nix function to easily create derivations (packages) to install binaries from location](https://twitter.com/mitchellh/status/1259521715211657216)
- [What Is Nix (2020)](https://engineering.shopify.com/blogs/engineering/what-is-nix) ([HN](https://news.ycombinator.com/item?id=23251754)) ([Lobsters](https://lobste.rs/s/bgwsd8/what_is_nix))
- [What Is Nix and Why You Should Use It (2020)](https://serokell.io/blog/what-is-nix)
- [comma](https://github.com/Shopify/comma) - Runs software without installing it. Wraps together nix run and nix-index.
- [nix-derivation](https://github.com/Gabriel439/Haskell-Nix-Derivation-Library) - Parse and render \*.drv files.
- [nix-build-uncached](https://github.com/Mic92/nix-build-uncached) - CI friendly wrapper around nix-build.
- [nix-tests](https://github.com/cleverca22/nix-tests) - Scratchpad for small experimental things I am doing with Nix.
- [Nix Flakes, Part 1: An introduction and tutorial (2020)](https://www.tweag.io/posts/2020-05-25-flakes.html)
- [Nix Flakes, Part 1: An introduction and tutorial](https://www.tweag.io/blog/2020-05-25-flakes/) ([Lobsters](https://lobste.rs/s/h99uo8/nix_flakes_part_1_introduction_tutorial)) ([Lobsters 2](https://lobste.rs/s/iqoegm/nix_flakes_part_1_introduction_tutorial))
- [nix-overlays of Antonio Monteiro](https://github.com/anmonteiro/nix-overlays)
- [A Nix terminology primer by a newcomer (2020)](https://stephank.nl/p/2020-06-01-a-nix-primer-by-a-newcomer.html)
- [Statistical Rethinking and Nix (2020)](https://rgoswami.me/posts/rethinking-r-nix/)
- [flake-utils](https://github.com/numtide/flake-utils) - Pure Nix flake utility functions.
- [nix.dev](https://nix.dev/) - Opinionated guide for developers wanting to get things done using the Nix ecosystem. ([Code](https://github.com/nix-dot-dev/nix.dev))
- [Nix language antipatterns](https://nix.dev/anti-patterns/language.html)
- [So, tell me about Nix (2020)](https://ghedam.at/15490/so-tell-me-about-nix)
- [nixdu](https://github.com/utdemir/nixdu) - Interactively browse the dependency graph of your Nix derivations.
- [Nix Package Versions](https://lazamar.co.uk/nix-versions/) - Search for old versions of Nix packages. ([Code](https://github.com/lazamar/nix-package-versions)) ([Reddit](https://www.reddit.com/r/NixOS/comments/gc68ec/find_older_versions_of_a_package_in_the_nix/))
- [Opinionated Nix repository template](https://github.com/nix-dot-dev/getting-started-nix-template) - Based on nix.dev tutorials, repository template to get you started with Nix.
- [Tools to manage a Nix-based project](https://github.com/shajra/nix-project)
- [Building static Haskell binary with Nix on Linux (2020)](https://blog.patchgirl.io/haskell/2020/07/13/static-haskell-binary.html)
- [Template for NUR repositories](https://github.com/rvolosatovs/nur-packages)
- [Bramble](https://github.com/maxmcd/bramble) - Functional build system inspired by nix. ([Article](https://maxmcd.com/posts/bramble/)) ([Lobsters](https://lobste.rs/s/g1tqfe/bramble_purely_functional_build_system))
- [The easiest way (I've found) to create your own Nix channel (2020)](https://lucperkins.dev/blog/nix-channel/)
- [Nix Monorepo](https://github.com/lucperkins/nix-monorepo) - How you might use Nix in a larger, multi-language project.
- [A Tutorial Introduction to Nix (2020)](https://rgoswami.me/posts/ccon-tut-nix/)
- [nixbuild.net](https://nixbuild.net/) - Cloud service that runs your Nix builds. It takes away the effort of maintaining and scaling build clusters, and integrates easily with any setup that uses Nix. ([Docs](https://docs.nixbuild.net/)) ([Tweet](https://twitter.com/utdemir/status/1299850167546380288))
- [Manix](https://github.com/mlvzk/manix) - Fast CLI documentation searcher for Nix.
- [Review of home manager (2020)](https://magnusson.io/post/home-manager-review/) ([Lobsters](https://lobste.rs/s/pys2ta/review_home_manager))
- [Nix Quick Install Action](https://github.com/nixbuild/nix-quick-install-action) - GitHub Action installs Nix in single-user mode, and adds almost no time at all to your workflow's running time. ([Web](https://github.com/marketplace/actions/nix-quick-install))
- [Setting up a Nix S3 binary cache (2020)](https://fzakaria.com/2020/07/15/setting-up-a-nix-s3-binary-cache.html) ([Lobsters](https://lobste.rs/s/myps2p/setting_up_nix_s3_binary_cache))
- [swift2nix: Run Swift inside Nix builds (2020)](https://euandre.org/2020/10/05/swift2nix-run-swift-inside-nix-builds.html)
- [sorri](https://github.com/nmattia/sorri) - Simple, lightweight implementation of Tweag's lorri.
- [Nix and the nix-shell for easily redistributable scripts (2020)](https://knezevic.ch/posts/nix-shell-redistributable-scripts/)
- [nix-buildkite](https://github.com/circuithub/nix-buildkite) - Take a Nix job description and turn it into separate Buildkite steps with dependencies. ([Tweet](https://twitter.com/acid2/status/1314507071224676352))
- [pre-commit-hooks.nix](https://github.com/cachix/pre-commit-hooks.nix) - Seamless integration of pre-commit git hooks with Nix.
- [Nix + Haskell setup](https://twitter.com/acid2/status/1314507569541640192)
- [caching your nix-shell (2020)](https://fzakaria.com/2020/08/11/caching-your-nix-shell.html)
- [nix-direnv](https://github.com/nix-community/nix-direnv) - Fast, persistent use_nix implementation for direnv.
- [NixCon 2020](https://2020.nixcon.org/) ([Stream](https://www.youtube.com/watch?v=7sQa04olUA0)) ([Code](https://github.com/nixcon/2020.nixcon.org)) ([HN](https://news.ycombinator.com/item?id=24799659))
- [Nix UX improvements](https://github.com/tweag/nix-ux)
- [NixCon 2020 talk about Nix flakes](https://github.com/serokell/nixcon2020-talk)
- [Nickel](https://github.com/tweag/nickel) - Lightweight configuration language. Its purpose is to automate the generation of static configuration files. ([Nickel: better configuration for less](https://www.tweag.io/blog/2020-10-22-nickel-open-sourcing/))
- [Nix-based process management framework](https://github.com/svanderburg/nix-processmgmt)
- [Local Nix Cache](https://github.com/andir/local-nix-cache) - Poor and hacky attempt at re-serving local nix packages that came from trusted sources.
- [Nix parallelism & Import From Derivation (2020)](https://fzakaria.com/2020/10/20/nix-parallelism-import-from-derivation.html) ([Reddit](https://www.reddit.com/r/NixOS/comments/jf79zy/nix_parallelism_import_from_derivation/))
- [macOS Nix setup: an alternative to Homebrew (2020)](https://wickedchicken.github.io/post/macos-nix-setup/) ([Lobsters](https://lobste.rs/s/rtluby/macos_nix_setup_alternative_homebrew)) ([HN](https://news.ycombinator.com/item?id=27825240))
- [update-nix-fetchgit](https://github.com/expipiplus1/update-nix-fetchgit) - Program to automatically update fetchgit values in Nix expressions.
- [Cache install Nix packages](https://github.com/nix-actions/cache-install) - Use the GitHub Actions cache for Nix packages.
- [nixbuild.net Action](https://github.com/nixbuild/nixbuild-action) - GitHub Action for using the nixbuild.net service.
- [fromElisp](https://github.com/talyz/fromElisp) - Emacs Lisp reader in Nix.
- [On-demand linked libraries for Nix (2020)](https://fzakaria.com/2020/11/17/on-demand-linked-libraries-for-nix.html) ([Lobsters](https://lobste.rs/s/nawo6d/on_demand_linked_libraries_for_nix))
- [Towards a Content-Addressed Model for Nix (2020)](https://www.tweag.io/blog/2020-09-10-nix-cas/) ([HN](https://news.ycombinator.com/item?id=27070858))
- [deploy-rs](https://github.com/serokell/deploy-rs) - Simple, multi-profile Nix-flake deploy tool.
- [TodoMVC-Nix](https://github.com/nix-community/todomvc-nix) - One-stop reference to build TodoMVC application inside the Nix world.
- [Trustix: Distributed trust and reproducibility tracking for binary caches (2020)](https://www.tweag.io/blog/2020-12-16-trustix-announcement/) ([Code](https://github.com/tweag/trustix))
- [Binary Verification with Trustix starring Adam Höse (2021)](https://overcast.fm/+i6QEhiFYA)
- [Building with Nix Flakes for, eh .. reasons (2021)](https://blog.ysndr.de/posts/internals/2021-01-01-flake-ification/)
- [What is the right approach to handling binary, non-executable data dependencies of packages? (2021)](https://www.reddit.com/r/NixOS/comments/kqe57g/what_is_the_right_approach_to_handling_binary/)
- [Scrive Nix Workshop](https://scrive.github.io/nix-workshop/) ([Code](https://github.com/scrive/nix-workshop))
- [nix-npm-buildpackage](https://github.com/serokell/nix-npm-buildpackage) - Build nix packages that use npm/yarn.
- [dev-env](https://github.com/digital-asset/dev-env) - Nix with training wheels.
- [Basinix](https://github.com/jonringer/basinix) - Pull request reviewing tool for nixpkgs.
- [Nix-template](https://github.com/jonringer/nix-template) - Make creating nix expressions easy. Provide a nice way to create largely boilerplate nix-expressions.
- [nixpkgs-hammering](https://github.com/jtojnar/nixpkgs-hammering) - Set of nit-picky rules that aim to point out and explain common mistakes in nixpkgs package pull requests.
- [nix-script](https://github.com/BrianHicks/nix-script) - Write scripts in compiled languages that run in the nix ecosystem, with no separate build step. ([Article](https://bytes.zone/posts/nix-script/))
- [nix-update](https://github.com/Mic92/nix-update) - Updates versions/source hashes of nix packages. It is designed to work with nixpkgs but also other package sets.
- [Nix Portable](https://github.com/DavHau/nix-portable) - Static, Permissionless, Installation-free, Pre-configured.
- [nix-optics](https://github.com/masaeedu/nix-optics) - Using profunctor optics to focus modifications in Nix.
- [Use Nix flakes without any fluff](https://github.com/gytis-ivaskevicius/flake-utils-plus)
- [Nix-environments](https://github.com/nix-community/nix-environments) - Repository to maintain out-of-tree shell.nix files.
- [Determinate Systems](https://determinate.systems/) - Confidently build and deploy to the cloud, stadium, or stock exchange. Expert help with the Nix ecosystem from Graham. ([Twitter](https://twitter.com/DeterminateSys)) ([GitHub](https://github.com/DeterminateSystems))
- [How to Learn Nix](https://ianthehenry.com/posts/how-to-learn-nix/)
- [Nix is the ultimate DevOps toolkit (2021)](https://tech.channable.com/posts/2021-04-09-nix-is-the-ultimate-devops-toolkit.html) ([HN](https://news.ycombinator.com/item?id=26748696)) ([Lobsters](https://lobste.rs/s/5gbbp2/nix_is_ultimate_devops_toolkit))
- [devshell](https://github.com/numtide/devshell) - Per project developer environments.
- [Nix Data](https://github.com/nix-community/nix-data) - Set of packages for data-scientists with batteries-included.
- [Nomia](https://github.com/scarf-sh/nomia) - System for precise, efficient resource management across every domain and resource type.
- [Chimera](https://github.com/craftweg/chimera) - Nix-based package manager with the focus on developer experience.
- [Data Science with Nix: Parameter Sweeps (2021)](https://blog.nixbuild.net/posts/2021-04-26-data-science-with-nix-parameter-sweeps.html)
- [Nixpkgs rules for Bazel](https://github.com/tweag/rules_nixpkgs) - Rules for importing Nixpkgs packages into Bazel.
- [Practical Nix Flakes (2021)](https://serokell.io/blog/practical-nix-flakes)
- [npmlock2nix](https://github.com/tweag/npmlock2nix) - Utilizing npm lockfiles to create Nix expressions for NPM based projects.
- [Laurn](https://github.com/baloo/laurn) - Run a dev-environment in a pure-ish nix environment.
- [nix-filter](https://github.com/numtide/nix-filter) - Small self-container source filtering lib.
- [nix-eval-lsp](https://github.com/aaronjanse/nix-eval-lsp) - Nix language server that evaluates code.
- [Flakes are such an obviously good thing: ...but the design and development process should be better (2021)](https://grahamc.com/blog/flakes-are-an-obviously-good-thing) ([Lobsters](https://lobste.rs/s/ipvgon/flakes_are_such_obviously_good_thing))
- [scratchix](https://github.com/dramforever/scratchix) - Linux From Scratch, but it's Nix.
- [Understanding Nix's String Context (2018)](https://shealevy.com/blog/2018/08/05/understanding-nixs-string-context/)
- [Replit now supports every programming language in Nix (2021)](https://blog.replit.com/nix) ([HN](https://news.ycombinator.com/item?id=27269292))
- [nix-graph](https://github.com/awakesecurity/nix-graph) - Reify the Nix build graph into a Haskell graph data structure.
- [Digga](https://github.com/divnix/digga) - Feature rich and configurable framework for harnessing Nix Flakes.
- [Nix solves the package manager ejection problem (2021)](https://zeroindexed.com/nix-ejection-problem) ([HN](https://news.ycombinator.com/item?id=27344677))
- [Cross compilation using Nix](https://nix.dev/tutorials/cross-compilation) ([Tweet](https://twitter.com/burdiyan/status/1448778093188030464))
- [nix-std](https://github.com/chessai/nix-std) - no-nixpkgs standard library for the nix expression language.
- [Nix unstable installer](https://github.com/numtide/nix-unstable-installer) - Place to host Nix unstable releases.
- [Nix Learning](https://github.com/humancalico/nix-learning) - Links to blog posts, articles, videos, etc for learning Nix.
- [nix-plugins](https://github.com/shlevy/nix-plugins) - Collection of useful Nix native plugins.
- [Nix.Ci](https://nix.ci/) - Provides the CI integration infrastructure for Nixpkgs and NixOS. More commonly known as OfBorg. ([Code](https://github.com/NixOS/ofborg))
- [nix-user-chroot](https://github.com/nix-community/nix-user-chroot) - Install & Run nix without root permissions.
- [Another Nix Success Story (2021)](https://maxdeviant.com/shards/2021/another-nix-success-story/) ([Lobsters](https://lobste.rs/s/sfacek/another_nix_success_story))
- [nix bundle](https://nixos.org/manual/nix/unstable/command-ref/new-cli/nix3-bundle.html) - Bundle an application so that it works outside of the Nix store.
- [Building Container Images with Nix (2021)](https://thewagner.net/blog/2021/02/25/building-container-images-with-nix/) ([HN](https://news.ycombinator.com/item?id=28240748))
- [Debug symbols for binaries with Nix (2021)](http://rski.github.io/2021/09/05/nix-debugging.html)
- [Ditch your version manager and use Nix (2021)](https://juliu.is/ditch-your-version-manager/) ([Lobsters](https://lobste.rs/s/emyfhx/ditch_your_version_manager)) ([HN](https://news.ycombinator.com/item?id=28565072))
- [Advanced shell packaging: resholve YADM's nixpkg (2021)](https://t-ravis.com/post/nix/advanced_shell_packaging_resholve_yadm/)
- [nix-parsec](https://github.com/nprindle/nix-parsec) - Parser combinators in Nix.
- [nix-prefetch-github](https://github.com/seppeljordan/nix-prefetch-github) - Prefetch sources from github for nix build tool.
- [Building reproducible Development environment with nix-shell (2021)](https://mudrii.medium.com/building-reproducible-development-environment-b1d4fb51a302)
- [Building with Nix on Replit](https://docs.replit.com/tutorials/30-build-with-nix) ([HN](https://news.ycombinator.com/item?id=28733156))
- [dream2nix](https://github.com/DavHau/dream2nix) - Generic framework for 2nix converters (converting from other build systems to nix).
- [Novice Nix: Flake Templates (2021)](https://peppe.rs/posts/novice_nix:_flake_templates/)
- [Personal provisioning machines with Nix](https://github.com/shajra/shajra-provisioning)
- [Issues you encountered learning Nix? (2021)](https://twitter.com/theprincessxena/status/1448984266071752721)
- [Declarative Dev Environments using Nix (2021)](https://marcopolo.io/code/declarative-dev-environments/)
- [Peerix](https://github.com/cid-chan/peerix) - Peer-to-peer binary cache for nix derivations.
- [update-flake-lock](https://github.com/DeterminateSystems/update-flake-lock) - GitHub Action that will update your flake.lock file whenever it is run.
- [Replit - Faster Nix Repl Startup (2021)](https://blog.replit.com/nix-perf-improvements)
- [Statix](https://github.com/nerdypepper/statix) - Lints and suggestions for the nix programming language. ([Reddit](https://www.reddit.com/r/NixOS/comments/qh47fz/statix_lints_and_suggestions_for_the_nix/))
- [Cicero](https://github.com/input-output-hk/cicero) - Workflow execution engine. Workflow is a description of (dependent) steps using the Nix configuration language.
- [Cross-compiling and Static-linking with Nix (2021)](https://functor.tokyo/blog/2021-10-20-nix-cross-static)
- [Nix 2.4 (2021)](https://discourse.nixos.org/t/nix-2-4-released/15822) ([Lobsters](https://lobste.rs/s/cumwee/nix_2_4_released)) ([HN](https://news.ycombinator.com/item?id=29073301))
- [nixcrpkgs](https://github.com/pololu/nixcrpkgs) - Tools for cross-compiling standalone applications using Nix.
- [Flake structure](https://github.com/freezeboy/flake-example)
- [nix-patchtools](https://github.com/svanderburg/nix-patchtools) - Autopatching binary packages to make them work with Nix.
- [A Critique of Nix Package Manager](https://www.iohannes.us/en/commentary/nix-critique/) ([Reddit](https://www.reddit.com/r/NixOS/comments/qs529l/a_critique_of_nix_package_manager/))
- [rusty-nix](https://github.com/Kloenk/rusty-nix) - Nix written in rust.
- [Some Nix questions (2021)](https://www.reddit.com/r/NixOS/comments/qsfkow/new_to_nix_some_questions/)
- [Nix - Flakes for out-of-tree code (2021)](https://www.youtube.com/watch?v=90P-Ml1318U)
- [Is it worth learning and migrating to Flakes as of November 2021?](https://www.reddit.com/r/NixOS/comments/qv9ag1/is_it_worth_learning_and_migrating_to_flakes_as/)
- [dreampkgs](https://github.com/nix-community/dreampkgs) - Collection of software packages managed with dream2nix, a framework for automated packaging.
- [Flake templates](https://github.com/NixOS/templates)
