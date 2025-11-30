FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o d20app

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/d20app .
EXPOSE 8080
CMD ["./d20app"]