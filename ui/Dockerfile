FROM golang:1.23-alpine AS build
WORKDIR /src
COPY dist.go generate.go main_wasm.go server.go \
     root.vugu \
     go.mod go.sum ./
RUN go run dist.go
# WORKDIR /src/dist
# ENTRYPOINT [ "./server" ]

FROM alpine AS release
COPY --from=build /src/dist/ /
WORKDIR /
ENTRYPOINT [ "/server" ]
