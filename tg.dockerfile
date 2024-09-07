FROM golang:1.23-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /tg_server github.com/blinlol/bookworm/cmd/tg

FROM alpine AS release
COPY --from=build /tg_server /tg_server
ENTRYPOINT [ "/tg_server" ]
