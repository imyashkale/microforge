# Step 1: Start with a base image
FROM golang:1.21.5-alpine3.19

# Step 2: Set the working directory
WORKDIR /app

# Step 3: Copy Go mod files
COPY go.mod .
COPY go.sum .

# Step 4: Download dependencies
RUN go mod download

# Step 5: Copy application files
COPY . .

ENV GIN_MODE=release

# Step 6: Build the application
RUN go build -o microforge-service ./cmd

# Step 7: Specify the entry point
CMD ["./microforge-service"]
