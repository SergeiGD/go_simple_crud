FROM golang:1.23.2-alpine

WORKDIR /usr/local/src

RUN apk --no-cache add bash gcc musl-dev 

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .
COPY ./build/migrations/entrypoint.sh .

RUN chmod +x entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
