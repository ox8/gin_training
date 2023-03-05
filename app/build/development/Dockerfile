FROM golang:1.20.1-alpine

WORKDIR /go/src
COPY ./app .

RUN apk upgrade --update && apk --no-cache add git

RUN go install github.com/cosmtrek/air@latest && go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]