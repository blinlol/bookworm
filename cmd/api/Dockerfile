FROM golang:1.23-alpine AS build
RUN go install github.com/blinlol/bookworm/cmd/api@latest
ENTRYPOINT [ "${GOBIN}/api" ]














# FROM golang:1.23-alpine AS build
# WORKDIR /src
# RUN go mod download
# RUN GOOS=linux CGO_ENABLED=0 go build -o /api_server api/server.go


# FROM build AS run-test
# WORKDIR /src
# RUN go vet
# RUN go test -v ./...


# FROM alpine AS release
# WORKDIR /
# COPY --from=build /api_server .
# ENTRYPOINT [ "/api_server" ]











# COPY . /app/src
# WORKDIR /app/src
# RUN go mod download
# RUN GOOS=linux CGO_ENABLED=0 go build -o /api_server

# FROM build AS run-test
# WORKDIR /app/src
# RUN go vet
# RUN go test -v ./...

# FROM alpine AS release
# WORKDIR /
# COPY --from=build /api_server /api_server
# ENTRYPOINT [ "/api_server" ]
