FROM golang:1.24

WORKDIR /app

COPY . .
RUN go mod tidy

# RUN air init
CMD ["go", "run", "main.go"]
