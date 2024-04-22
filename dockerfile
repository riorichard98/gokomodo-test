# Start from the official Golang image based on Alpine Linux
FROM golang:latest AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy only the go.mod file
COPY go.mod .

# Run go mod tidy to ensure dependencies are in sync
RUN go mod tidy

# Copy the rest of the source code from the current directory to the Working Directory inside the container
COPY . .

# Copy the .env file
COPY .env .

# Set PSQL_HOST to gokomodo-pg in .env file
RUN sed -i 's/PSQL_HOST=localhost/PSQL_HOST=gokomodo-pg/' .env

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Start a new stage from scratch
FROM alpine:latest

# Copy the executable and .env file from the builder stage
COPY --from=builder /app/main /main
COPY --from=builder /app/.env /.env

# Expose port 3001 to the outside world
EXPOSE 3001

# Command to run the executable
ENTRYPOINT ["./main"] 
