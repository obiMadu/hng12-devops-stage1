# Stage 1: Build
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -o /go/bin/number-classifier ./cmd/server

# Stage 2: Runtime
FROM scratch

# Copy SSL certificates for HTTPS requests to Numbers API
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary from builder
COPY --from=builder /go/bin/number-classifier /number-classifier

# Expose port
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/number-classifier"]
