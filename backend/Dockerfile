FROM golang:1.22-alpine

# Install gcc for SQLite
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy all code
COPY . .

# Build binary
RUN CGO_ENABLED=1 go build -o main ./cmd/server/main.go

EXPOSE 8080

CMD ["./main"]