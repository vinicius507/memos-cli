{
  inputs = {
    nixpkgs.url = "github:cachix/devenv-nixpkgs/rolling";
    devenv = {
      url = "github:cachix/devenv";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  nixConfig = {
    extra-trusted-public-keys = "devenv.cachix.org-1:w1cLUi8dv3hnoSPGAuibQv+f9TZLr6cv/Hm9XgU50cw=";
    extra-substituters = "https://devenv.cachix.org";
  };

  outputs = {
    self,
    nixpkgs,
    devenv,
    systems,
    ...
  } @ inputs: let
    system = "x86_64-linux";
    pkgs = nixpkgs.legacyPackages.${system};
  in {
    packages.${system}.memos-cli = with pkgs;
      buildGoModule {
        pname = "memos-cli";
        version = "0.1.0";
        src = ./.;
        vendorHash = "sha256-du7r9qNu0pNZgQfpMp+YQCl7iayUkxm5FdSAsGZ0DPI=";
        installPhase = ''
          install -Dm755 $GOPATH/bin/memos-cli $out/bin/memos
        '';
        meta = with lib; {
          mainProgram = "memos";
          description = "A CLI for managing Memos";
          homepage = "https://github.com/vinicius507/memos-cli";
          license = licenses.mit;
        };
      };
    devShells.${system}.default = devenv.lib.mkShell {
      inherit inputs pkgs;
      modules = [
        {
          env.CGO_ENABLED = 0;
          packages = with pkgs; [
            cobra-cli
          ];
          languages.go.enable = true;
        }
      ];
    };
  };
}
