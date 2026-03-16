FROM golang:1.24-alpine AS builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod ./

RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o main ./cmd/api/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/main .
RUN mkdir -p /app/data
CMD ["./main"]
