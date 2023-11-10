# Step 1: Build the Go binary
FROM golang:1.19.5-alpine AS builder

# Set the working directory inside the container
WORKDIR /build

# Copy Go files into the container
COPY . .

# Download dependencies and build the binary
RUN go build -o server .

# Step 2: Create the runtime image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage to the current stage
COPY --from=builder /build/zeroconf-server .

# Set the binary as the entry point of the container
ENTRYPOINT ["/app/zeroconf-server"]
