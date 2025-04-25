# -------- Stage 1: Build --------
  FROM golang:1.23 AS builder

  WORKDIR /app
  
  # Copy module files and get Go dependencies
  COPY go.mod go.sum ./
  RUN go mod download
  
  # Copy all project files
  COPY . .

  # Build binary
  RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/main ./cmd/api/*.go
  
# -------- Stage 2: Minimal Runtime --------
  FROM alpine:latest
  
  # Copy only the binary and .env if necessary
  COPY --from=builder /app/bin/main /app/main
  COPY --from=builder /app/.env /app/.env
  
  # Optional: Enable port
  EXPOSE 3000
  
  CMD ["/app/main"]