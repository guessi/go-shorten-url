FROM public.ecr.aws/docker/library/golang:1.19-alpine3.17 as BUILDER
RUN apk add --no-cache git
WORKDIR ${GOPATH}/src/github.com/guessi/go-shorten-url
COPY *.go go.mod go.sum ./
RUN CGO_ENABLED=0 go install

FROM scratch
COPY --from=BUILDER /go/bin/go-shorten-url /opt/
COPY ./config/redirections.json /opt/config/redirections.json
WORKDIR /opt/
VOLUME /opt/config/
EXPOSE 8080
CMD ["/opt/go-shorten-url"]
