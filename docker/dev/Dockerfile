FROM golang:1.12-alpine

RUN apk add --update --no-cache \
    git \
    inotify-tools \
    ffmpeg \
  && rm -rf /var/cache/apk/* && rm -rf /tmp/*

RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u golang.org/x/lint/golint

STOPSIGNAL SIGQUIT
