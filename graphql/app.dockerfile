# Stage 1: Build
FROM golang:1.22-alpine3.20 AS build

# Install build dependencies
RUN apk add --no-cache build-base ca-certificates

# Set working directory inside the container
WORKDIR /go/src/dev.murali.go-microservice

# Cache and install Go dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY vendor vendor
COPY account account
COPY catalog catalog
COPY order order
COPY graphql graphql

# Build the Go application
RUN go build -mod=vendor -o /app/bin/app ./graphql

# Stage 2: Minimal Runtime Image
FROM alpine:3.20

# Set working directory in the runtime container
WORKDIR /usr/bin

# Copy built binary from builder stage
COPY --from=build /app/bin/app .

# Ensure CA certificates are available
RUN apk add --no-cache ca-certificates

# Expose application port
EXPOSE 8080

# Run the application
CMD ["./app"]
