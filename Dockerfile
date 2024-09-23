
FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY models .

RUN go build -o main .

FROM ubuntu:latest

LABEL authors="Daniel Kałużny"

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

ENTRYPOINT ["./main"]
