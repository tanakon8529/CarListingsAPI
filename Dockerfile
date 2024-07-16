# Dockerfile
FROM golang:1.21

# Install git and other dependencies
RUN apt-get update && apt-get install -y git

# Install swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Generate Swagger documentation
RUN swag init -g api/main.go

# Ensure .env file is copied to the root of the application
COPY .env .env

WORKDIR /app/api

RUN go build -o /daveslist-emdpcv

EXPOSE 8080

CMD ["/daveslist-emdpcv"]
