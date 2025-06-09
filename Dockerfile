FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o todo-api .

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/todo-api .

EXPOSE 8080

CMD ["./todo-api"]