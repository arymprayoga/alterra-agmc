FROM golang:1.18.6-alpine3.16 AS builder
LABEL maintainer="Ary Muhammad Prayoga<arymprayoga@gmail.com>"
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /day10
EXPOSE 8080
CMD ["./day10"]