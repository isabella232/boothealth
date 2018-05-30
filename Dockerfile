FROM golang:1.10-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers

RUN mkdir -p /go/src/github.com/status-im/boothealth
ADD . /go/src/github.com/status-im/boothealth
RUN cd /go/src/github.com/status-im/boothealth && go build -o boothealth .

FROM alpine:latest

RUN apk add --no-cache ca-certificates bash

COPY --from=builder /go/src/github.com/status-im/boothealth/boothealth /usr/local/bin/boothealth
ENTRYPOINT ["/usr/local/bin/boothealth"]