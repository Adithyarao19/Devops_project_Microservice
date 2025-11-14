FROM golang:1.22-alpine

WORKDIR /app

COPY . .

# Enable Go modules
RUN go mod init service_1 || true

# Add Prometheus dependencies
RUN go get github.com/prometheus/client_golang/prometheus \
    && go get github.com/prometheus/client_golang/prometheus/promhttp

RUN go build -o main

CMD ["./main"]
