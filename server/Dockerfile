FROM golang:1.8.0
MAINTAINER nsoushi

RUN mkdir /src

COPY ./server /src/server

EXPOSE 8080

ENTRYPOINT ["/src/server"]