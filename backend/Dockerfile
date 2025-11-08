# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o migrate ./cmd/migrate
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o seed ./cmd/seed

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy binaries from builder
COPY --from=builder /app/server .
COPY --from=builder /app/migrate .
COPY --from=builder /app/seed .
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./server"]

