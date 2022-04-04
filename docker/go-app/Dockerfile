# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY main.go ./

RUN go build -o /go-app


##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go-app /go-app

USER 1000

ENTRYPOINT ["/go-app"]



