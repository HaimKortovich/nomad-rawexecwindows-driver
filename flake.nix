{
  description = "Nomad windows raw exec driver";

  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
    nixpkgs-unstable.url = "nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs, nixpkgs-unstable, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        unstable = nixpkgs-unstable.legacyPackages.${system};
        pkgs = nixpkgs.legacyPackages.${system};
      in with pkgs; {
        devShells.default = mkShell {
          buildInputs = [ nixfmt unstable.gopls operator-sdk unstable.go ];
        };
      });
}
