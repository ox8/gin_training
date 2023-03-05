FROM golang:1.20.1-alpine

WORKDIR /go/src
COPY . .

RUN apk upgrade --update && apk --no-cache add git

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", "build/development/.air.toml"]