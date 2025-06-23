{
  description = "Myncer Dev environment";

  inputs = {
    nixpkgs = { url = "github:NixOS/nixpkgs/nixos-24.05"; };
    nixpkgs-unstable = { url = "github:NixOS/nixpkgs/nixpkgs-unstable"; };
    flake-utils = { url = "github:numtide/flake-utils"; };
  };

  outputs = { self, nixpkgs, nixpkgs-unstable, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        stable = import nixpkgs { inherit system; };
        unstable = import nixpkgs-unstable { inherit system; };
      in {
        devShells.default = stable.mkShell {
          buildInputs = [
            stable.go_1_22
            stable.nodejs_20
            stable.pnpm
            stable.docker
            unstable.buf
          ];

          shellHook = ''
            export PATH=$PWD/node_modules/.bin:$PATH
            echo "🧪 Myncer flake shell ready"
          '';
        };
      });
}
