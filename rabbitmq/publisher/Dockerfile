FROM golang:1.17-buster as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /publisher

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /publisher /publisher

USER nonroot:nonroot

ENTRYPOINT ["/publisher"]