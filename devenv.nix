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
  '';

  scripts.build.exec = ''mango build'';
  scripts.clean.exec = "mango clean";
  scripts.generate.exec = "mango generate";
  scripts.mango.exec = "go run ./cli $*";

  scripts.format.exec = ''
    golangci-lint run ./... --fix --config config/golangci.yaml
    sqlc vet -f config/sqlc.yaml
    prettier "**/*.{js,css,json,yaml,md}" --write
  '';

  scripts.lint.exec = ''
    golangci-lint run ./...  --config config/golangci.yaml
    prettier "**/*.{js,css,json,yaml,md}" --check
  '';

  scripts.generate-tailwind.exec = ''
    cd config/tailwind
    npx tailwindcss -i ./tailwind.css -o ../../assets/css/index.css
  '';

  enterShell = ''
    if [ ! -d example/.mango ]; then
      prepare
      generate
    fi

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
