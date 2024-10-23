# Use the official Go image as the base image
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download

# Copy the entire project into the container
COPY . .

# Set CGO_ENABLED to 0 to build statically
ENV CGO_ENABLED=0

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest

# Install necessary packages for running the application
RUN apk add --no-cache ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Ensure the binary has execute permissions
RUN chmod +x ./main

# Command to run the executable
CMD ["./main"]

# Expose the port your app runs on
EXPOSE 3000
