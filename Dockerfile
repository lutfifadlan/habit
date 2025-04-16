FROM golang:1.24 as build

# Enable CGO and install gcc
ENV CGO_ENABLED=1
RUN apt-get update && apt-get install -y gcc

WORKDIR /app

# Copy the Go modules
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o out

# Use a smaller image for the final container
FROM ubuntu:22.04

WORKDIR /app

# Copy the built binary from the build stage
COPY --from=build /app/out /app/out

# Copy any necessary files and run the binary
CMD ["./out"]