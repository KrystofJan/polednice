{
  description = "Rusalka nix flake and build system";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";

    alejandra.url = "github:kamadorueda/alejandra/3.1.0";
    alejandra.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs = inputs @ {
    self,
    nixpkgs,
    alejandra,
    ...
  }: let
    systems = ["x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin"];
    forAllSystems = f:
      nixpkgs.lib.genAttrs systems (
        system: let
          pkgs = import nixpkgs {inherit system;};
        in
          f {inherit system pkgs;}
      );
  in {
    devShells = forAllSystems ({
      system,
      pkgs,
    }: {
      default = import ./shell.nix {
        inherit pkgs system alejandra;
      };
    });

    formatter = nixpkgs.lib.genAttrs systems (
      system:
        alejandra.packages.${system}.default
    );
  };
}
