# ---------- Stage 1: Build ----------
FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o goflow ./cmd/api


# ---------- Stage 2: Run ----------
FROM gcr.io/distroless/static-debian13:nonroot

WORKDIR /app

COPY --from=builder /app/goflow .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./goflow"]