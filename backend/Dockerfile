# Use the Go base image
FROM golang:1.23-alpine

# Set the working directory
WORKDIR /usr/src/app

# Install dependencies
RUN apk add --no-cache git bash

# Install Air for live reloading during development
RUN go install github.com/cosmtrek/air@v1.41.0

# Add Go binaries to PATH
ENV PATH="/go/bin:${PATH}"

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Add .env file to the container
COPY .env ./

# Add .air.toml file
COPY .air.toml ./

# Clean up any unnecessary dependencies
RUN go mod tidy

# Run the application with Air (assuming .air.toml exists in the project)
CMD ["air", "-c", ".air.toml"]
