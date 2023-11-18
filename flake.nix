{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    templ.url = "github:a-h/templ";
  };

  outputs = { self, flake-utils, nixpkgs, templ }:
    flake-utils.lib.eachDefaultSystem (system: 
      let
        pkgs = import nixpkgs { inherit system; };
        templ-pkg = templ.packages.${system}.templ;
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [ 
              go
              air
              templ-pkg
              tailwindcss
          ];
        };
      }
    );
}
