FROM golang:1.17-alpine

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY ./go.mod ./
RUN go mod download
EXPOSE 80

# FIX: cmd階層にあるためコマンドは実行させないように変更する
# CMD ["go", "run", "main.go"]