# Use an official Go image as a build image
FROM golang:1.24 as build

# Enable CGO and install gcc
ENV CGO_ENABLED=1
RUN apt-get update && apt-get install -y gcc

WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod . 
COPY go.sum .
RUN go mod download

# Copy the source files and build the application
COPY . .
RUN go build -o out

# Use a smaller base image for the final output
FROM ubuntu:22.04

WORKDIR /app

# Copy the built binary from the build stage
COPY --from=build /app/out /app/out

# Copy any necessary files and run the binary
CMD ["./out"]
