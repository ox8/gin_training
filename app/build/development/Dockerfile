FROM golang:1.20.1-alpine

WORKDIR /go/src
COPY . .

RUN apk upgrade --update && apk --no-cache add git

RUN go install github.com/cosmtrek/air@latest
RUN go install entgo.io/ent/cmd/ent@v0.11.9

CMD ["air", "-c", "build/development/.air.toml"]