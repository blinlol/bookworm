# Bookworm

[go](go.dev) + [gin](https://pkg.go.dev/github.com/gin-gonic/gin) + [vugu](https://pkg.go.dev/github.com/vugu/vugu) web project.

## Project structure
- [model](./model/) directory contains struct definitions for used models and API requests and responses.
    - [dao](./model/dao/) directory implements interaction with Postgres and provides an interface for other packages (early versions used Mongo)
- [cmd](./cmd/) directory stores programs for API server and UI variants
    - [api/server.go](./cmd/api/server.go) implements and starts API server. It interacts with the DB via dao.
    - [ui](./cmd/ui/) stores Vugu fronted code. Use [devserver.go](./cmd/ui/devserver.go) to run development server and [dist.go](./cmd/ui/dist.go) to create distribution binaries.
    - [tg](./cmd/tg/) implements a simple Telegram bot, which can show books via markup and show random quotes from them.
    - [codeium_ui](./cmd/codeium_ui/) frontend with JS by llama.

## Run
To start all services simply run `docker compose up` in the project directory

## Goals
The main objectives of the project were to learn
- how to do HTTP REST API with Golang, 
- what is MongoDB and how to use it,
- how to use Docker,
- try to write a webassembly app with go.