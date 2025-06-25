FROM golang:1.24

WORKDIR /app

COPY . .
RUN go mod tidy

RUN go install github.com/air-verse/air@latest
RUN go mod download

# RUN air init
CMD ["air", "-c", ".air.toml"]
