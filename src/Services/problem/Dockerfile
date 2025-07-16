FROM golang:1.24

WORKDIR /app

COPY . .
RUN go mod tidy

RUN ["mkdir", "/storage"]
EXPOSE 3000
# RUN air init
CMD ["go", "run", "cmd/main.go"]
