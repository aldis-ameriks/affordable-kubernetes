FROM golang:alpine AS builder
ADD ./main.go /go/src/goapp/main.go

RUN set -ex && \
    cd /go/src/goapp && \
    CGO_ENABLED=0 go build -a -v && \
    ls && \
    mv ./goapp /usr/bin/goapp

FROM busybox
COPY --from=builder /usr/bin/goapp /usr/local/bin/goapp
ENTRYPOINT [ "goapp" ]