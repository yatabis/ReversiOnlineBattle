#!/bin/bash

FROM alpine:latest as client

WORKDIR /go/src/ReversiOnlineBattle
RUN apk add --no-cache nodejs && apk add --update npm && npm install -g yarn
COPY ./scripts ./scripts/

WORKDIR ./scripts
RUN yarn install

WORKDIR ../
COPY ./website/src ./website/src/

WORKDIR ./scripts
RUN yarn build


FROM golang:latest as server

WORKDIR /go/src/ReversiOnlineBattle
COPY . .
ENV GO111MODULE on

WORKDIR ./cmd
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main


FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=server /go/src/ReversiOnlineBattle/cmd/main /main
COPY --from=client /go/src/ReversiOnlineBattle/website/static /website/static/
RUN chmod +x /main

ENV PORT 8080

CMD ["/main"]
