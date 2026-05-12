FROM golang:1.26.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/app ./cmd/main.go

FROM ubuntu:24.04

WORKDIR /app

COPY --from=builder /app/app /app

RUN ls -la .

RUN pwd

EXPOSE 8000

CMD ["./app"]