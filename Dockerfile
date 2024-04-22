# Stage 1: Build the Go binary
FROM golang:1.22.1-alpine

WORKDIR /customer-data-api

COPY . .

# Download dependencies
RUN go mod download

#build
RUN go build -v -o /customer-data-api/customer-data-api ./cmd/main.go

# Expose port 8000
EXPOSE 8000

#run app
ENTRYPOINT [ "/customer-data-api/customer-data-api" ]