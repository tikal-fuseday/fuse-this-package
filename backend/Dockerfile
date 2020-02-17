FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY main.go ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
