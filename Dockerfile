FROM golang:1.11-alpine3.8 as BUILDER
RUN apk add --no-cache git
RUN go get -u github.com/guessi/go-shorten-url
WORKDIR ${GOPATH}/src/github.com/guessi/go-shorten-url
RUN go build

FROM alpine:3.8
COPY --from=BUILDER /go/bin/go-shorten-url /opt/
WORKDIR /opt/
VOLUME /opt/config/
EXPOSE 8080
CMD ["/opt/go-shorten-url"]
