FROM golang:1.11-alpine as builder
COPY . /go/src/brfutebol/hostdetect
ENV GO111MODULE=on
WORKDIR /go/src/brfutebol/hostdetect
RUN apk -U add git build-base && \
    mkdir -p /build && \
    go build  -ldflags '-extldflags "-static"' -o /build/hostdetect

FROM alpine:3.7
RUN apk -U add ca-certificates curl && rm -rf /var/cache/apk/*
WORKDIR /opt
COPY --from=builder /build/hostdetect .
HEALTHCHECK --interval=2m --timeout=3s \
  CMD curl -f http://localhost:8080/metrics || exit 1
CMD ["./opt/hostdetect"]