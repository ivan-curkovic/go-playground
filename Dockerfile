# Use the official Go image as a parent image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Update dependencies
RUN go mod tidy && go get -u

# Build your application
RUN go build -o myapp

# Run the binary program produced by `go install`
CMD ["./myapp"]
