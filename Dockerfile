FROM golang:1.8-alpine AS compile
COPY hellomkes-web.go /go
RUN go build hellomkes-web.go

FROM alpine:latest
COPY --from=compile /go/hellomkes-web /
COPY  server.* /
#probably should bind to a higher port and run as non root inside the container..
#USER nobody:nobody
EXPOSE 443
ENTRYPOINT ["/hellomkes-web"]
