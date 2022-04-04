FROM golang:1.17-alpine as build

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY ./main.go ./

RUN go build -o /docker-app



FROM alpine:3.15.0

WORKDIR /

COPY --from=build /docker-app ./

EXPOSE 3000

USER 1000

CMD [ "/docker-app" ]
