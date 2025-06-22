{
  description = "Myncer Dev environment";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-24.05";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        devShells.default = pkgs.mkShell {
          buildInputs = [
            pkgs.go_1_22
            pkgs.nodejs_20
            pkgs.pnpm
            pkgs.protobuf
            pkgs.docker
          ];

          shellHook = ''
            export PATH=$PWD/node_modules/.bin:$PATH
            echo "🧪 Myncer flake shell ready"
          '';
        };
      });
}
