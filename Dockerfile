FROM golang:alpine

RUN apk update && apk add --no-cache git
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main ./cmd/api/main.go
CMD ["/app/main"]