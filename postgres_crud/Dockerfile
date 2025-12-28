# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Install git for fetching modules
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server ./

# Final stage
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /server /server
EXPOSE 6000
ENTRYPOINT ["/server"]
