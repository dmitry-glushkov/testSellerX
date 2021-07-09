FROM golang:1.16.5-alpine

WORKDIR /task/sellerx/
ADD ./  /task/sellerx/

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN go mod download

RUN go build -o bin/task cmd/main.go
ENTRYPOINT [ "bin/task" ]