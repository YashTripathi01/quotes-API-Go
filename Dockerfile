# Build stage
FROM golang:1.20-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/main ./

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set environment variable
ENV DATABASE_URL=mongodb://localhost:27017/quotes

USER appuser

EXPOSE 1323

CMD ["./main"]
