# Dockerfile
# Stage 1: Build the application
FROM golang:1.24.2-alpine AS builder

# Set working directory
WORKDIR /app

# Install dependencies for Node.js and pnpm
RUN apk add --no-cache nodejs npm curl
RUN npm install -g pnpm

# Copy package.json and pnpm-lock.yaml for dependency caching
COPY package.json pnpm-lock.yaml ./
RUN pnpm install

# Copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Generate Templ files
RUN go tool templ generate

# Generate CSS from tailwind and daisy 
RUN pnpm build:css 

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

# Stage 2: Create minimal runtime image
FROM alpine:3.20.7

# Install ca-certificates for HTTPS (needed for CDN requests)
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy static assets (if any)
COPY --from=builder /app/static ./static

# Expose port 3000
EXPOSE 3000

# Run the binary
CMD ["/app/main"]
