# Build stage using latest Go Alpine image
FROM golang:1.22-alpine3.20 AS build

# Install build dependencies
RUN apk --no-cache add gcc g++ make ca-certificates

# Set working directory
WORKDIR /go/src/dev.murali.go-microservice

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Copy vendor and catalog directories
COPY vendor ./vendor
COPY catalog ./catalog

# Build the catalog service binary
RUN go build -mod=vendor -o /go/bin/app ./catalog/cmd/catalog

# Runtime stage using minimal Alpine image
FROM alpine:3.20

# Set working directory
WORKDIR /usr/bin

# Copy built binary from build stage
COPY --from=build /go/bin/app .

# Expose the service port
EXPOSE 8080

# Set the startup command
CMD ["./app"]
