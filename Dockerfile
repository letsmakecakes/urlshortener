# Build stage
FROM golang:1.23-alpine AS builder

# Install git and SS certificates
RUN apk add --no-cache git ca-certificates tzdata

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code to the working directory
COPY . .

# Build the Go application for linux with CGO disabled
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/urlshortener ./cmd/api

# Final stage
FROM alpine:latest

# Import SSL certificates and timezone data from the builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Create a non-root user
RUN adduser -D -g '' appuser

# Copy the built binary from the builder stage
COPY --from=builder /app/urlshortener /app/urlshortener

# Copy the .env file to the working directory
COPY .env /app/.env

# Set the working directory inside the container
WORKDIR /app

# Change ownership of the /app directory to the non-root user
RUN chown -R appuser:appuser /app

# Switch to the non-root user
USER appuser

# Expose the port the application will run on
EXPOSE 8080

# Set the entrypoint to the built binary
ENTRYPOINT ["/app/urlshortener"]