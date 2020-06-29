#!/bin/bash

FROM golang:latest as builder

WORKDIR /go/src/ReversiOnlineBattle
COPY . .

ENV GO111MODULE on

WORKDIR ./cmd
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main

FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/ReversiOnlineBattle/cmd/main /main
RUN chmod +x /main

ENV PORT 8080

CMD ["/main"]
