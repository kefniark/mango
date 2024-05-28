{ pkgs, lib, config, inputs, ... }:

{
  dotenv.enable = true;
  
  # https://devenv.sh/packages/
  packages = [
    # Golang
    pkgs.go_1_22
    pkgs.gotools
    pkgs.golangci-lint

    # Protobuf & Code Generation
    pkgs.protobuf
    # pkgs.buf need 1.32
    pkgs.protoc-gen-go
    pkgs.protoc-gen-go-grpc
    pkgs.protoc-gen-connect-go

    # SQLC
    pkgs.sqlc
    pkgs.air

    # Templ
    pkgs.templ

    # TailwindCSS & DaisyUI
    pkgs.nodejs_22

    # Linters
    pkgs.nodePackages.prettier
  ];

  # https://devenv.sh/scripts/
  scripts.prepare.exec = ''
    go mod download
    go get "github.com/sudorandom/protoc-gen-connect-openapi@v0.7.2"
    go install "github.com/sudorandom/protoc-gen-connect-openapi@v0.7.2"
    
    mango prepare
    mango generate
  '';

  scripts.mango.exec = "go run ./cli $*";

  enterShell = ''
    prepare

    echo ""
    echo "----- 🚀 Mango Development Shell -----"
    echo " 💻 mango dev : Start dev server"
    echo " 💻 mango build : Compile server go into binary"
    echo " 💻 mango format : Format code"
    echo " 💻 mango lint : Lint code"
    echo " 💻 mango generate : Code generation (Sql queries, Protobuf, Openapi, ...)"
    echo "------"
  '';

  # https://devenv.sh/tests/
  # enterTest = ''
  #   echo "Running tests"
  #   git --version | grep "2.42.0"
  # '';

  # https://devenv.sh/services/
  services.postgres.enable = true;


  # https://devenv.sh/languages/
  languages.nix.enable = true;
  languages.go.enable = true;

  # https://devenv.sh/pre-commit-hooks/
  # pre-commit.hooks.shellcheck.enable = true;

  # https://devenv.sh/processes/
  # processes.ping.exec = "ping example.com";

  # See full reference at https://devenv.sh/reference/options/
}
