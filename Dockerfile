FROM golang:alpine

LABEL maintainer="Agus Wibawantara"


RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base
RUN mkdir /app
RUN go install github.com/githubnemo/CompileDaemon@latest
RUN go install -v golang.org/x/tools/gopls@latest

WORKDIR /app

COPY . .
COPY .env .

RUN go get -d -v ./...

RUN go install -v ./...

ENTRYPOINT CompileDaemon --build="go build -buildvcs=false -o ./build/build ." --command="./build/build" -build-dir=/app