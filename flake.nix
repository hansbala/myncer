{
  description = "Myncer dev environment";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-23.11"; # ← known to include grpc-web
  outputs = { self, nixpkgs }: {
    devShells.default = nixpkgs.legacyPackages.${builtins.currentSystem}.mkShell {
      buildInputs = [
        nixpkgs.legacyPackages.${builtins.currentSystem}.go
        nixpkgs.legacyPackages.${builtins.currentSystem}.protobuf
        nixpkgs.legacyPackages.${builtins.currentSystem}.protoc-gen-grpc-web
        nixpkgs.legacyPackages.${builtins.currentSystem}.nodejs
        nixpkgs.legacyPackages.${builtins.currentSystem}.yarn
        nixpkgs.legacyPackages.${builtins.currentSystem}.buf
      ];
    };
  };
}
