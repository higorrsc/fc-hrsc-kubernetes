FROM golang:1.17-alpine
COPY . .
RUN go build -o server server.go
CMD ["./server"]
