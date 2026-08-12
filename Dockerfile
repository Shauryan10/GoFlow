# ---------- Stage 1: Build ----------
FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o goflow ./cmd/api


# ---------- Stage 2: Run ----------
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/goflow .

EXPOSE 8080

CMD ["./goflow"]