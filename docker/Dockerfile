FROM golang:1.21-rc-alpine3.17
RUN mkdir /app
ADD .. /app
WORKDIR /app
RUN go build -o main ./cmd/app/main.go

CMD ["/app/main"]
