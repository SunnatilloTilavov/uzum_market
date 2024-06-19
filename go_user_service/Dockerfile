# Use a Go version 1.19 or later
FROM golang:1.19-alpine3.16 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main cmd/main.go

# Start a new stage from scratch
FROM alpine:3.16

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /app/main

# Expose port 8082 to the outside world
EXPOSE 8082

# Command to run the executable
ENTRYPOINT ["/app/main"]
