FROM golang:1.24.4 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./src/cmd/main.go

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/main .

RUN chmod +x main

EXPOSE 8000

CMD ["./main"]
