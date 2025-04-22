FROM golang:1.24-alpine AS builder

COPY . /auth/
WORKDIR /auth/

RUN go mod download
RUN go build -o ./bin/auth cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /auth/bin/auth .

CMD ["./auth"]
