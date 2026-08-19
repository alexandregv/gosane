{
  description = "Go back to sanity with these Go linters";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-26.05";
  };

  outputs = { nixpkgs, ... }:
    let
      forAllSystems = f: nixpkgs.lib.genAttrs nixpkgs.lib.systems.flakeExposed (system: f (import nixpkgs { inherit system; }));
    in
    {
      devShells = forAllSystems (pkgs: {
        default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gofumpt
            golangci-lint
            gopls
            gotools
            tree-sitter
            tree-sitter-grammars.tree-sitter-go
          ];
        };
      });

      packages = forAllSystems (pkgs:
        let
          name = "gosane";
          version = "0.1.0";
        in
        {
          default = pkgs.buildGoModule {
            inherit name version;
            pname = name;
            src = ./.;
            goSum = ./go.sum;
            vendorHash = "sha256-ztgD9in1B8l5FkuC/E5zNvyum6My586ZpBAaeXRBz8Y=";

            env.CGO_ENABLED = 0;

            meta = {
              description = "Go back to sanity with these Go linters";
              homepage = "https://github.com/alexandregv/gosane";
              license = pkgs.lib.licenses.mit;
              maintainers = with pkgs.lib.maintainers; [ alexandregv ];
            };
          };
        }
      );
    };
}
