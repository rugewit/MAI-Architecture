# Use the official Golang image as the base image
FROM golang:1.22

# Set the current working directory inside the container
WORKDIR /basket_service

# Copy the Go module files and download dependencies
COPY ./go.mod ./
RUN go mod download

# Copy the rest of the application source code
COPY ./ ./

# Copy wait-for-it.sh into the container
COPY wait-for-it.sh /usr/local/bin/wait-for-it.sh
RUN chmod +x /usr/local/bin/wait-for-it.sh

# Build the Go application
RUN go build -o basket_service ./cmd/main.go

EXPOSE 8083

# Update entrypoint to wait for PostgreSQL to be ready
CMD ["sh", "-c", "/usr/local/bin/wait-for-it.sh db_mongo:27017 -- ./basket_service"]