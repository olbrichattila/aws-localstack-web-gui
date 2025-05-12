# Stage 1: Build the Go binary
FROM golang:1.24 AS builder

WORKDIR /app

COPY go-api/ ./
RUN go mod tidy && \
    go build -o lsmanager .

# Stage 2: Final image
FROM debian:bookworm-slim

WORKDIR /app

COPY ./frontend/build /app/
COPY --from=builder /app/lsmanager .

EXPOSE 80

ENTRYPOINT ["./lsmanager"]
