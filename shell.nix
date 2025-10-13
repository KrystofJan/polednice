{
  pkgs ? import <nixpkgs> {},
  system ? builtins.currentSystem,
  alejandra,
}:
pkgs.mkShell {
  name = "default";

  buildInputs = with pkgs; [
    go
    gopls
    gotools
    sqlite
    sqlc

    # nix stuff
    nixd
    (alejandra.packages.${system}.default)
  ];

  shellHook = ''
    echo "Let's get crackin'"
    zsh
  '';
}
