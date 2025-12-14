FROM golang:1.22-alpine

WORKDIR /app
COPY . .

RUN go build -o nova-qrcode ./cmd/server

EXPOSE 8080
CMD ["./nova-qrcode"]