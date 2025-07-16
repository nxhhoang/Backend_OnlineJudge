FROM golang:1.24

WORKDIR /app

COPY . .
RUN go mod tidy

RUN ["mkdir", "/storage"]

# RUN air init
CMD ["go", "run", "cmd/main.go"]
