FROM golang:1.7-alpine

RUN apk --update add \
    ca-certificates

ADD . /go/src/github.com/denderello/scrapy
WORKDIR /go/src/github.com/denderello/scrapy

RUN go install -v

ENTRYPOINT ["scrapy"]
