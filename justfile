run: clean install-dev generate gomod dev

clean:
    rm -rf ./gen

install-dev:
    #!/usr/bin/env sh

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
    go run cmd/server/main.go

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
    protoc --experimental_allow_proto3_optional -I ./internal/api \
        --go_out=./gen/api --go_opt paths=source_relative \
        --go-grpc_out=./gen/api --go-grpc_opt paths=source_relative \
        --connect-go_out=./gen/api --connect-go_opt paths=source_relative \
        --connect-openapi_out=./gen/api --connect-openapi_opt=base=base.openapi.yaml \
        ./internal/**/*.proto

    # --openapiv2_out ./gen/api --openapiv2_opt generate_unbound_methods=true \
    # --grpc-gateway_out=./gen/gw --grpc-gateway_opt paths=source_relative \
    # --grpc-gateway_opt generate_unbound_methods=true \
    # --grpc-gateway_opt grpc_api_configuration=./internal/api/config.yaml \
    # --grpc-gateway_opt standalone=true \
        