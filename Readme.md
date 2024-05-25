# Go Web ServerTemplate

## Description

**WIP**: Try to build a modern and full fledged Web Server in Golang.
The goal is to get a versatile base that can be used for a variety of projects.

Features:
* Heavily use code generation (schema first approach) to reduce boilerplate
* API Endpoints based on [Connect](https://connectrpc.com/) (Support GRPC, Rest, OpenAPI, Subscriptions, ...)
* HTML Endpoints based on [Templ](https://github.com/a-h/templ)
* Database Access based on [SQLC](https://github.com/sqlc-dev/sqlc) (Support sqlite, postgres, mysql)
* Dev Tools:
  * Live-Reload with [Air](https://github.com/cosmtrek/air)
  * Linting & Golang Good practice with [Golangci-lint](https://github.com/golangci/golangci-lint)

## TODO
* [ ] Support .env
* [ ] Web UI (HTMX/Templ)
* [ ] Background Task (queue)
* [ ] Handle assets/public folder
* [ ] Doc generation
* [ ] Tests / 2e2
* [ ] Distribution (auto build docker images / binaries)

## Folder Structure

* `internal/api`: API Schema + Handlers
* `internal/db`: SQL Schema for [SQLC](https://github.com/sqlc-dev/sqlc) (`.sql`)
* `internal/middlewares`: Manipulate Requests and Responses (Auth, Logging, Cors)

## Dev Commands

* `just`: Bootstrap project
* `just dev`: Start the server in development mode
* `just format`: Use Golangci-lint to auto-fix as much code as possible
* `just lint`: Run Golangci-lint
* `just generate`: Run all the code generators