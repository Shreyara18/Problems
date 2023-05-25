# Use the official Golang image as the base image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod ./

# Download the Go dependencies
RUN go mod download

# Copy the remaining source code to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port that the application listens on
EXPOSE 8080

# Set the command to run the binary executable
CMD ["./main"]
