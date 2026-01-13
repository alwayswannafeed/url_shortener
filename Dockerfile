FROM golang:1.25-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/alwayswannafeed/url_shortener
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/url_shortener /go/src/github.com/alwayswannafeed/url_shortener

FROM alpine:latest

COPY --from=buildbase /usr/local/bin/url_shortener /usr/local/bin/url_shortener
COPY config.yaml /config.yaml

RUN apk add --no-cache ca-certificates

ENTRYPOINT ["url_shortener"]