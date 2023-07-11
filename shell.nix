{ pkgs ? import (builtins.fetchTarball {
        url = "https://github.com/NixOS/nixpkgs/archive/55070e598e0e03d1d116c49b9eff322ef07c6ac6.tar.gz";
    }) {} }:
pkgs.mkShell {
  packages = [
      pkgs.darwin.apple_sdk.frameworks.Security 
      pkgs.pkgconfig 
      pkgs.openssl

      # python
      pkgs.python3
      pkgs.python3Packages.pip
      pkgs.python3Packages.pytest

      # nodeJs
      pkgs.nodejs-18_x
      pkgs.nodePackages.pnpm
      pkgs.nodePackages.typescript

      # rust
      # issue with rust-anaylzer and for neovim and vscode
      # pkgs.cargo
      # pkgs.rustc

      # golang
      # don't know how to make workspace works
      # pkgs.go
  ];
}
