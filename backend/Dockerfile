# Start with the base Go image
FROM golang:1.21rc2-bookworm

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port on which the application will run
EXPOSE 3030

# Set the entrypoint for the Docker container
CMD ["./main"]