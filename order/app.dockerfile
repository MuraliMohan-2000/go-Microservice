# Build stage using latest Go Alpine image
FROM golang:1.22-alpine3.20 AS build

# Install build dependencies
RUN apk --no-cache add gcc g++ make ca-certificates

# Set the working directory for build context
WORKDIR /go/src/dev.murali.go-microservice

# Copy module files
COPY go.mod go.sum ./

# Copy vendor and service directories
COPY vendor ./vendor
COPY account ./account
COPY catalog ./catalog
COPY order ./order

# Build the 'order' service binary
RUN go build -mod=vendor -o /go/bin/app ./order/cmd/order

# Runtime stage using latest Alpine base image
FROM alpine:3.20

# Set the working directory for runtime
WORKDIR /usr/bin

# Copy the compiled binary from build stage
COPY --from=build /go/bin/app .

# Expose the application's port
EXPOSE 8080

# Command to run the application
CMD ["./app"]