# Start from the latest golang base image
FROM golang:latest AS builder

# Add Maintainer Info
LABEL maintainer="Fabien Bellanger <valentil@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN GOOS=linux go build -o main .

RUN make update


######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8888 to the outside world
EXPOSE 8888

# Command to run the executable
CMD ["./main", "serve"] 
