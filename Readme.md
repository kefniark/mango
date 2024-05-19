# Go Web ServerTemplate

## Description

**WIP**: Try to build a modern Web Server Structure in Golang:
* Use Code generation to reduce boilerplate and focus on what matters
* API Endpoints based on [Connect](https://connectrpc.com/) (Support GRPC, Rest, OpenAPI, Subscriptions, ...)
* HTML Endpoints based on [Templ](https://github.com/a-h/templ)
* Database Access based on [SQLC](https://github.com/sqlc-dev/sqlc) (Support sqlite, postgres, mysql)
* Dev Tools:
  * Live-Reload
  * Linting & Golang Good practice with [Golangci-lint](https://github.com/golangci/golangci-lint)

## TODO
* [ ] Support .env
* [ ] Web UI (HTMX/Templ)
* [ ] Background Task (queue)
* [ ] Handle assets/public folder
* [ ] Doc generation
* [ ] Distribution (auto build docker images / binaries)

## Folder Structure

* `internal/api`: API Schema + Handlers
* `internal/db`: SQL Schema for [SQLC](https://github.com/sqlc-dev/sqlc) (`.sql`)
* `internal/middlewares`: Manipulate Requests and Responses (Auth, Logging, Cors)

## Dev Commands

* `just`: Dev Start Server
* `just dev`
* `just format`
* `just lint`
* `just generate`: 