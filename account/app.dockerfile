# Use latest Go Alpine base image for build stage
FROM golang:1.22-alpine3.20 AS build

# Install necessary build tools
RUN apk --no-cache add gcc g++ make ca-certificates

# Set working directory for build
WORKDIR /go/src/dev.murali.go-microservice

# Copy module files and vendor directory
COPY go.mod go.sum ./
COPY vendor ./vendor
COPY account ./account

# Build the Go app with vendoring
RUN go build -mod=vendor -o /go/bin/app ./account/cmd/account

# Use minimal Alpine image for final stage
FROM alpine:3.20

# Set working directory for final stage
WORKDIR /usr/bin

# Copy built binary from build stage
COPY --from=build /go/bin/app .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./app"]
