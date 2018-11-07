FROM golang:alpine AS builder
ADD ./ /go/src/smallgo/

RUN set -ex && \
    cd /go/src/smallgo && \
    CGO_ENABLED=0 go build -a -v && \
    ls && \
    mv ./smallgo /usr/bin/smallgo

FROM busybox
COPY --from=builder /usr/bin/smallgo /usr/local/bin/smallgo
ENTRYPOINT [ "smallgo" ]