FROM golang:1.23

WORKDIR /app

COPY . .

RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/vehicle-registration-manager


CMD ["./main"]


# docker build -t vehicle-ap .
# docker run -p 8080:8080 vehicle-ap