{
  description = "Shanepee";

  # Nixpkgs / NixOS version to use.
  inputs.nixpkgs.url = "nixpkgs/nixos-24.11";

  outputs = { self, nixpkgs }:
    let
      # System types to support.
      supportedSystems = [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];

      # Helper function to generate an attrset '{ x86_64-linux = f "x86_64-linux"; ... }'.
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;

      # Nixpkgs instantiated for supported system types.
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });
    in
    {
      packages = forAllSystems (system:
      let
        pkgs = nixpkgsFor.${system};
      in
        {
          api = pkgs.buildGoModule {
            pname = "api";
            version = "0.0.0";

            vendorHash = "sha256-PnMZcYXHLgmitVapjWSgT6m1GR3ohrWJmQ1hhBOfX0Q=";

            src = ./api;

            subPackages = [ "cmd/api" ];
          };
        }
      );

      devShells = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [
              go
              gopls
              gotools
              go-tools

              go-swag
              wire
            ];
          };
        });
    };
}
