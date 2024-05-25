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

    # TailwindCSS
    pkgs.tailwindcss

    # Linters
    pkgs.nodePackages.prettier
  ];

  # https://devenv.sh/scripts/
  scripts.clean.exec = ''
  rm -rf ./gen
  rm -rf ./dist
  rm -rf ./dev.db
  '';

  scripts.install-deps.exec = ''
    go mod download
    go get "github.com/sudorandom/protoc-gen-connect-openapi@v0.7.2"
    go install "github.com/sudorandom/protoc-gen-connect-openapi@v0.7.2"
  '';

  scripts.start.exec = ''
    air -c ./config/air.toml
  '';

  scripts.build.exec = ''
    rm -rf ./dist
    mkdir -p ./dist
    go build -ldflags="-s -w" -o ./dist/server ./main.go
    cp -r ./static ./dist/static
  '';

  scripts.format.exec = ''
    golangci-lint run ./... --fix --config config/golangci.yaml
    sqlc vet -f config/sqlc.yaml
    prettier "**/*.{json,yaml,md}" --write
  '';

  scripts.lint.exec = ''
    golangci-lint run ./...  --config config/golangci.yaml
    prettier "**/*.{json,yaml,md}" --check
  '';

  scripts.generate.exec = ''
    generate-sql &
    generate-proto &
    generate-templ &
    generate-tailwind &
    wait
  '';

  scripts.generate-sql.exec = ''
    sqlc generate -f config/sqlc.yaml
  '';

  scripts.generate-proto.exec = ''
    mkdir -p ./gen/api
    BUF_CMD=$(which buf)

    # Auto-generate GRPC/Client code based on api.proto
    $BUF_CMD dep update
    $BUF_CMD generate
  '';

  scripts.generate-templ.exec = ''
    templ generate
  '';

  scripts.generate-tailwind.exec = ''
    tailwindcss -i ./config/tailwind.css -o ./assets/css/index.css
  '';

  enterShell = ''
    if [ ! -d directory ]; then
      install-deps
      generate
    fi
    echo "----- 🚀 Server Devenv -----"
    echo ""
    echo "💻 Scripts:"
    echo " > start : Start dev server"
    echo " > build : Compile server go into binary"
    echo " > format : Format code"
    echo " > lint : Lint code"
    echo " > generate : Code generation (Sql queries, Protobuf, Openapi, ...)"
    echo "------"
  '';

  # https://devenv.sh/tests/
  # enterTest = ''
  #   echo "Running tests"
  #   git --version | grep "2.42.0"
  # '';

  # https://devenv.sh/services/

  # https://devenv.sh/languages/
  languages.nix.enable = true;
  languages.go.enable = true;

  # https://devenv.sh/pre-commit-hooks/
  # pre-commit.hooks.shellcheck.enable = true;

  # https://devenv.sh/processes/
  # processes.ping.exec = "ping example.com";

  # See full reference at https://devenv.sh/reference/options/
}
