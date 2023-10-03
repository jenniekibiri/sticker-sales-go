# Use an official Golang runtime as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code to the container
COPY . .

# Build the Go application
RUN go build -o main

# Expose the port your Go application will listen on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
