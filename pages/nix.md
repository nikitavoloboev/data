## Commands
```
nix profile list # list profiles
```
## install latest nix version
```
nix profile install github:nixos/nix/2.24.3
```

```
# will find binary `protoc` in nixpkgs
nix-locate -w /bin/protoc
```

```
# repl
nix repl

packages.x86_64-linux.rust-hello-x86_64-apple-darwin # to see package to build

nix build .#rust-hello-x86_64-apple-darwin # builds it

# in repl, can do, to load flake in current repl
:lf .
```

## build docker for mac
```
nix build .#packages.aarch64-darwin.rust-hello-aarch64-apple-darwin-oci
docker load < result
docker image ls
```