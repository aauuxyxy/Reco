FROM golang:1.23.0

COPY src /go/src/

WORKDIR /go/src/

RUN apk update \
&& apk add --no-cache-git \
&& go mod tidy

EXPOSE 8080