# Stage 1: Build the Go binary
FROM golang:1.24 AS builder

# Set working directory inside container
WORKDIR /app

# Copy Go modules and download dependencies
COPY go-api/ ./

# Build the Go application
RUN go mod tidy && \
    go build -o lsmanager .

# Stage 2: Final image
FROM debian:bookworm-slim

# Set working directory in final image
WORKDIR /app

COPY ./frontend/build /app/
COPY --from=builder /app/lsmanager .


# Expose any needed port (e.g., 8080)
EXPOSE 8080

# Run the binary
ENTRYPOINT ["./lsmanager"]