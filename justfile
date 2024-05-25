run: clean install-dev generate gomod dev

clean:
    rm -rf ./gen

install-dev:
    #!/usr/bin/env sh

    # Dev CLI to auto-reload
    go install github.com/cosmtrek/air@latest

    # Dev CLI to concat protobuf files
    go install github.com/syumai/protocat/cmd/protocat

    # Sqlc - Database Code Generation
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

    # GRPC - Protocol Buffers
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

    # ConnectRPC - Codegen & OpenAPI
    go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
    go install github.com/sudorandom/protoc-gen-connect-openapi@main

    # GolangCI-Lint
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.1

gomod:
    go mod tidy

dev:
    air

format:
    golangci-lint run ./... --fix

lint:
    golangci-lint run ./...

generate: generate-sql generate-proto

generate-sql:
    sqlc generate

generate-proto:
    #!/usr/bin/env sh

    mkdir -p ./gen/api
    
    # Merge All proto file in a single api.proto
    PROTO_FILES=$(find ./internal/api -name '*.proto' -printf '%p ')
    protocat $PROTO_FILES > gen/api/api.proto
    echo 'option go_package = "github.com/kefniark/go-web-server/gen/api";' >> gen/api/api.proto
    
    # Auto-generate GRPC/Client code based on api.proto
    protoc --experimental_allow_proto3_optional -I ./gen/api \
        --go_out=./gen/api --go_opt paths=source_relative \
        --go-grpc_out=./gen/api --go-grpc_opt paths=source_relative \
        --connect-go_out=./gen/api --connect-go_opt paths=source_relative \
        --connect-openapi_out=./gen/api --connect-openapi_opt=base=base.openapi.yaml \
        api.proto