# Use the official Go image as a base image
FROM golang:latest

# Set the current working directory inside the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Download dependencies if you have any

# Build the Go application
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]