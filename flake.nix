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
            pkgs.openjdk17
          ];

          shellHook = ''
            export PATH=$PWD/node_modules/.bin:$PATH
            export OPENAPI_GENERATOR_CLI_JAR="$PWD/thirdparty/openapi-generator-cli.jar"

            if [ ! -f "$OPENAPI_GENERATOR_CLI_JAR" ]; then
              echo "⬇️  Downloading OpenAPI Generator CLI JAR..."
              mkdir -p thirdparty
              curl -L -o "$OPENAPI_GENERATOR_CLI_JAR" https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/7.4.0/openapi-generator-cli-7.4.0.jar
            fi

            alias openapi-gen="java -jar $OPENAPI_GENERATOR_CLI_JAR"
            echo "🧪 Myncer flake shell ready"
          '';
        };
      });
}
