FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/vehicle-registration-manager


CMD ["./main"]

