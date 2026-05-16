FROM golang:1.24-alpine

WORKDIR /app

COPY . .

RUN go mod init mtproto-ui-next || true
RUN go mod tidy || true

EXPOSE 8080

CMD ["go", "run", "./cmd/server"]
