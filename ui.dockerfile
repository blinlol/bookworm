FROM golang:1.23-alpine AS build
WORKDIR /src
COPY . .
WORKDIR /src/cmd/ui
RUN go run dist.go

FROM alpine AS release
COPY --from=build /src/cmd/ui/dist /dist
WORKDIR /dist
ENTRYPOINT [ "./server" ]
