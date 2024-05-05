# syntax=docker/dockerfile:1

FROM golang:1.20 as build

WORKDIR /app

COPY go.mod ./

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /app

FROM scratch

COPY --from=build /app /

EXPOSE 4000

ENTRYPOINT ["/hello-docker"]