# Build stage
FROM golang:1.22-alpine AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create a directory for the app
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN go build -o pixelQR ./cmd/server/main.go

# Final stage
FROM alpine:latest

# Set environment variables
ENV GIN_MODE=release

# Set up a non-root user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Create a directory for the app and set permissions
RUN mkdir /app && chown appuser:appgroup /app

# Switch to the non-root user
USER appuser

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/pixelQR /app/

# Expose the port the app runs on
EXPOSE 3000

# Command to run the binary
CMD ["/app/pixelQR"]
