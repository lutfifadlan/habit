# Start from the official Golang image for building the binary
FROM golang:1.24 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. These layers are cached if go.mod and go.sum remain unchanged
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application binary
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o main .

# Now create a small image for the application using Alpine Linux
FROM alpine:3.18

# Set the working directory
WORKDIR /app

# Copy the binary from the builder image
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

# Command to run the binary
CMD ["./main"]
