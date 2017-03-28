FROM golang:1.8.0
MAINTAINER nsoushi

RUN mkdir /src

COPY ./client /src/client

EXPOSE 8081

ENTRYPOINT ["/src/client"]