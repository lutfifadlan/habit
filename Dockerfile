FROM golang:alpine

# Install gcc
RUN apk add --no-cache gcc libc-dev

# Enable CGO
ENV CGO_ENABLED=1

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

RUN go build -o binary

EXPOSE 8080

ENTRYPOINT ["/app/binary"]