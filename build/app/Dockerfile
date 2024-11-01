FROM golang:1.23.2-alpine as builder

WORKDIR /usr/local/src

RUN apk --no-cache add gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN swag init -g app.go -d ./cmd/main/,./internal/ --quiet   
RUN go build -o ./bin/app ./cmd/main/app.go


FROM alpine as runner

COPY --from=builder /usr/local/src/bin/app/ /
COPY config/config.yaml /config.yaml

CMD ["/app"]
