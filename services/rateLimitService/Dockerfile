FROM golang:alpine as build

RUN CGO_ENABLED=0 GOOS=linux go get -v github.com/envoyproxy/ratelimit/src/service_cmd

FROM alpine:3.8 AS final
WORKDIR /ratelimit/config

RUN apk --no-cache add ca-certificates
COPY --from=build /go/bin/service_cmd /usr/local/bin/ratelimit