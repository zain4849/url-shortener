# Use the official Golang image as a base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install Air for live reloading
RUN go install github.com/cosmtrek/air@latest

# Copy the source code into the container
COPY . .

# Build the Go app
# RUN go build -o main cmd/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable with Air
CMD ["air", "-c", ".air.toml"]

