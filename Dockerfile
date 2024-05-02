# syntax=docker/dockerfile:1

FROM golang:1.20

WORKDIR /tmp/app

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /tmp/hello-docker

EXPOSE 4000

ENTRYPOINT ["/tmp/hello-docker"]