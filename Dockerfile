# Stage 1: Build the Go binary
FROM golang:1.23.0-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o api-server .

# Stage 2: Run the built binary
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/api-server .

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./api-server"]
