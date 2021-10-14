# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-alpine AS build

COPY yagows/ /yagows/
COPY examples/ /examples/

ENV CGO_ENABLED=0
RUN cd /examples/app && go build -ldflags "-s -w" -o /httpserver

##
## Deploy
##
FROM scratch

WORKDIR /

COPY --from=build /httpserver /httpserver

EXPOSE 8090

ENV VERSION=1.0

ENTRYPOINT [ "/httpserver"]