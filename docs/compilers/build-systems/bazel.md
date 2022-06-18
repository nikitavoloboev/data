---
title: Bazel
---

# [Bazel](https://bazel.build)

## Links

- [Buildtools for bazel](https://github.com/bazelbuild/buildtools) - Contains developer tools for working with Google's bazel buildtool.
- [BazelCon 2019 videos](https://www.youtube.com/playlist?list=PLxNYxgaZ8Rsf-7g43Z8LyXct9ax6egdSj)
- [BazelCon 2020](https://opensourcelive.withgoogle.com/events/bazelcon2020)
- [Bazel Reaches 1.0 Milestone (2019)](https://opensource.googleblog.com/2019/10/bazel-reaches-10-milestone.html) ([HN](https://news.ycombinator.com/item?id=21288185))
- [HN: Bazel 2.0](https://news.ycombinator.com/item?id=21863393)
- [Make local development (with Bazel) great again (2020)](https://www.youtube.com/watch?v=rQv1sjLU4cI)
- [cargo-raze](https://github.com/google/cargo-raze) - Generate Bazel build from Cargo dependencies.
- [bazel-linting-system](https://github.com/thundergolfer/bazel-linting-system) - Experimental system for registering, configuring, and invoking source-code linters in Bazel.
- [Articles on Bazel by Jay Conrod](https://www.jayconrod.com/tags/bazel)
- [Bazelisk](https://github.com/bazelbuild/bazelisk) - User-friendly launcher for Bazel.
- [Building Uber’s Go Monorepo with Bazel (2020)](https://eng.uber.com/go-monorepo-bazel/) ([HN](https://news.ycombinator.com/item?id=23180255))
- [PodToBUILD](https://github.com/pinterest/PodToBUILD) - Easy way to build CocoaPods with Bazel - it integrates pods end to end with an easy to use macro.
- [Bazel Performance in a CI Environment (2020)](https://filipnikolovski.com/posts/bazel-performance-in-a-ci-environment/)
- [Bazel Code](https://github.com/bazelbuild/bazel)
- [Bazel for Open-source C / C++ Libraries Distribution (2020)](https://liuliu.me/eyes/bazel-for-libraries-distribution-an-open-source-library-author-perspective/) ([HN](https://news.ycombinator.com/item?id=24490089))
- [Bazel Blog](https://blog.bazel.build/) ([Code](https://github.com/bazelbuild/bazel-blog))
- [Bazel Rust Rules](https://github.com/bazelbuild/rules_rust) - Rules for building Rust projects with Bazel.
- [apple_rules_lint](https://github.com/apple/apple_rules_lint) - Framework for adding lint checks to Bazel projects.
- [Python Rules for Bazel](https://github.com/bazelbuild/rules_python)
- [LLVM Bazel BUILD files](https://github.com/google/llvm-bazel) - Bazel Build System Support for LLVM. ([HN](https://news.ycombinator.com/item?id=24925368))
- [Bazel's Continuous Integration Setup](https://github.com/bazelbuild/continuous-integration)
- [What happened at Bazelcon 2020](https://www.gasparevitta.com/posts/what-happened-at-bazelcon-2020/)
- [bazel-remote](https://github.com/buchgr/bazel-remote) - Remote build cache for Bazel.
- [Skylib](https://github.com/bazelbuild/bazel-skylib) - Library of Starlark functions for manipulating collections, file paths, and various other data types in the domain of Bazel build rules.
- [Bazel Container Image Rules](https://github.com/bazelbuild/rules_docker) - Rules for building and handling Docker images with Bazel.
- [bazel-cache](https://github.com/znly/bazel-cache) - Minimal cloud oriented Bazel gRPC cache.
- [bazels3cache](https://github.com/Asana/bazels3cache) - Small web server for a Bazel cache, proxies to S3; allows Bazel to work offline; async uploads to make Bazel faster.
- [BuildBuddy](https://github.com/buildbuddy-io/buildbuddy) - Open source Bazel build event viewer, result store, and remote cache. ([Web](https://www.buildbuddy.io/))
- [dbx_build_tools](https://github.com/dropbox/dbx_build_tools) - Dropbox's Bazel rules and tools.
- [Gazelle](https://github.com/bazelbuild/bazel-gazelle) - Bazel build file generator for Bazel projects. It natively supports Go and protobuf, and it may be extended to support new languages and custom rule sets.
- [Bazel Buildfarm](https://github.com/bazelbuild/bazel-buildfarm) - Bazel remote caching and execution service.
- [Example Bazel Monorepo](https://github.com/thundergolfer/example-bazel-monorepo) - Example Bazel-ified monorepo, supporting Golang, Java, Python, Scala, and Typescript.
- [Apple Rules for Bazel](https://github.com/bazelbuild/rules_apple) - Bazel rules to build apps for Apple platforms.
- [BazelCon](https://conf.bazel.build/) ([Supporting files for BazelCon](https://github.com/bazelbuild/bazelcon))
- [JavaScript rules for Bazel](https://github.com/bazelbuild/rules_nodejs)
- [Scoot](https://github.com/twitter/scoot) - Distributed task runner, supporting both a proprietary API and Bazel's Remote Execution.
- [Aspect Build Systems](https://www.aspect.dev/) - Expert Bazel consulting. ([GitHub](https://github.com/aspect-build))
- [Bazel Central Registry](https://github.com/bazelbuild/bazel-central-registry) - Central registry of Bazel modules for the Bzlmod external dependency system.
- [Jazelle](https://github.com/uber-web/jazelle) - Incremental, cacheable builds for large JavaScript monorepos using Bazel.
- [Bazel watcher](https://github.com/bazelbuild/bazel-watcher) - Tools for building Bazel targets when source files change.
- [Turbo Cache](https://github.com/allada/turbo-cache) - Extremely fast and efficient bazel cache service (CAS) written in rust.
- [bazel-differ](https://github.com/ewhauser/bazel-differ) - Command line interface for Bazel that helps you do incremental builds across different Git versions.
- [Starlark Language](https://docs.bazel.build/versions/main/skylark/language.html) - Python-inspired configuration language. Used by Bazel. ([Starlark LSP](https://github.com/tilt-dev/starlark-lsp))
- [Awesome Bazel](https://github.com/jin/awesome-bazel)
- [bazel-diff](https://github.com/Tinder/bazel-diff) - Performs Bazel Target Diffing between two revisions in Git, allowing for Test Target Selection and Selective Building.
- [Target Determinator](https://github.com/bazel-contrib/target-determinator) - Determines which Bazel targets were affected between two git commits.
- [rules_multirun](https://github.com/keith/rules_multirun) - Bazel rules for running multiple commands in parallel in a single bazel invocation.
- [Example repository for how to setup a bazel build that uses rust on iOS and Android](https://github.com/keith/bazel-rust-mobile-demo)
- [Bazel rules for JavaScript](https://github.com/aspect-build/rules_js) - High-performance Bazel rules for running NodeJS tools and building JavaScript projects.
