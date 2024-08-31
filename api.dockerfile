FROM golang:1.23-alpine AS build
WORKDIR /src
COPY . . 
RUN go build -o /api_server github.com/blinlol/bookworm/cmd/api

FROM alpine AS release
COPY --from=build /api_server /api_server
ENTRYPOINT [ "/api_server" ]
