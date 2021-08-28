FROM golang:1.17

RUN go install github.com/canthefason/go-watcher/cmd/watcher@latest

WORKDIR /app

CMD watcher
