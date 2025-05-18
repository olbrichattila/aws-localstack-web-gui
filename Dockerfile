# Stage 1: Build the Go binary
FROM golang:1.24 AS builder

WORKDIR /app

COPY go-api/ ./
RUN go mod tidy && \
    go build -o lsmanager .

# Stage 2: Final image
FROM debian:bookworm-slim

WORKDIR /app

RUN mkdir frontend && mkdir database

COPY ./frontend/build ./frontend
COPY ./go-api/migrations/ ./migrations
COPY ./init.sh .
COPY ./go-api/docker-env.example ./.env
COPY --from=builder /app/lsmanager .

RUN chmod +x init.sh

EXPOSE 80

ENTRYPOINT ["./init.sh"]
