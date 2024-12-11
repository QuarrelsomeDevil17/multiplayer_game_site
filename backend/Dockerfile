# Stage 1: Build the Go application
FROM golang:1.21 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/main.go

# Stage 2: Create a lightweight runtime image
FROM gcr.io/distroless/base-debian10 AS final

# Set the Current Working Directory inside the container for runtime
WORKDIR /root/

# Copy the pre-built binary from the builder image
COPY --from=builder /app/main .

# Expose port 8080 (or whichever port your application uses)
EXPOSE 8080

# Set environment variables (optional: change based on your needs)
ENV PORT=8080

# Define the command to run your application
CMD ["./main"]
