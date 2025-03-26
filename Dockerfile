# Start with the official Golang image
FROM golang:1.22.5 as base

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod .

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Use a minimal base image for the final container
FROM gcr.io/distroless/base

# Set the working directory inside the container
WORKDIR /root/

# Copy the built Go binary from the base stage
COPY --from=base /app/main .
COPY . ./static

# Expose the application port
EXPOSE 3000

# Command to run the application
CMD ["./main"]