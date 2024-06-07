{ pkgs, lib, config, inputs, ... }:

{
  dotenv.enable = true;
  
  # https://devenv.sh/packages/
  packages = [
    # Golang
    pkgs.go_1_22
    pkgs.golangci-lint

    # Golang Dev Tools
    pkgs.air
    pkgs.sqlc
    pkgs.templ

    # Protobuf & Code Generation
    pkgs.protobuf
    pkgs.buf
    pkgs.protoc-gen-go
    pkgs.protoc-gen-go-grpc
    pkgs.protoc-gen-connect-go

    # Linters
    pkgs.nodePackages.prettier
  ];

  # https://devenv.sh/scripts/
  scripts.prepare.exec = ''
    go mod download
    go get "github.com/sudorandom/protoc-gen-connect-openapi@v0.7.2"
    go install "github.com/sudorandom/protoc-gen-connect-openapi@v0.7.2"
    go get "github.com/otiai10/copy"
  '';

  # scripts.up.exec = "docker compose -f ./docker/example.yaml up";
  # scripts.down.exec = "docker compose -f ./docker/example.yaml down -v";
  scripts.mango.exec = "go run ./pkg/mango-cli $*";

  enterShell = ''
    prepare
    mango generate

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
  # services.postgres.enable = true;

  # https://devenv.sh/languages/
  # languages.nix.enable = true;
  # languages.go.enable = true;

  # https://devenv.sh/pre-commit-hooks/
  # pre-commit.hooks.shellcheck.enable = true;

  # https://devenv.sh/processes/
  # processes.ping.exec = "ping example.com";
  # containers.example.name = "example";
  # containers.example.copyToRoot = ./dist/example;
  # containers.example.startupCommand = "example-linux-amd64";

  # See full reference at https://devenv.sh/reference/options/
}
