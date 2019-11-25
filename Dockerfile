FROM golang:latest

LABEL MAINTAINER="dogukanayd <me@dogukanaydogdu.com>"

ENV APP_PATH=/go/src/github.com/dogukanayd/go-rest

RUN apt-get update

COPY . ${APP_PATH}
WORKDIR ${APP_PATH}


