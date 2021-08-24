FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# build-time variable with default port
ARG SERVE_PORT=8090

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app. CGO is needed for sqlite.
# As per doc sqlite is a CGO enabled package so required to set the environment variable CGO_ENABLED=1 and have a gcc compile present within path.
RUN CGO_ENABLED=1 GOOS=linux go build --tags="json1" -a -installsuffix cgo -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
