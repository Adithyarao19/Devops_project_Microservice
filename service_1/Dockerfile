
FROM golang:1.24-alpine AS builder
WORKDIR /src


RUN apk add --no-cache git ca-certificates


COPY go.mod go.sum ./
RUN if [ -f go.mod ]; then go mod download; fi


COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/service_1 ./...

FROM scratch
COPY --from=builder /out/service_1 /service_1
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
EXPOSE 8080
ENTRYPOINT ["/service_1"]
