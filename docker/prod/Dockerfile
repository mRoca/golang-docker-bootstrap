FROM golang:1.12-alpine AS builder

ARG PACKAGE_NAME

RUN apk add --update --no-cache \
    git \
  && rm -rf /var/cache/apk/* && rm -rf /tmp/*

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/${PACKAGE_NAME}

COPY src/${PACKAGE_NAME}/Gopkg.toml src/${PACKAGE_NAME}/Gopkg.lock ./

RUN dep ensure --vendor-only

COPY src/${PACKAGE_NAME}/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/cli ./cmd/cli/main.go

# ===================================

FROM alpine:3.9

ARG PACKAGE_NAME

WORKDIR /srv

COPY --from=builder /go/src/${PACKAGE_NAME}/bin/cli /srv/cli

ENTRYPOINT ["/srv/cli"]

STOPSIGNAL SIGQUIT
