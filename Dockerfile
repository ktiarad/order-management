# Dockerfile
FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# RUN apk add --no-cache postgresql-client

CMD ["go", "run", "main.go"]
