# Use Go 1.24.0 with ARM64 support
FROM --platform=linux/arm64 golang:1.24.0-alpine

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first (for better caching)
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy all files into the container
COPY . .

# Expose the application port
EXPOSE 8001

# Build the Go application
RUN go build -o main .

# Command to run the service
CMD ["/app/main"]