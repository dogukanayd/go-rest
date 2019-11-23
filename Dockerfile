FROM golang:latest

LABEL MAINTAINER="dogukanayd <me@dogukanaydogdu.com>"

ENV APP_PATH=/go/src/github.com/dogukanayd/go-rest

RUN apt-get update
RUN go get "github.com/gin-gonic/gin"
RUN go get "github.com/go-sql-driver/mysql"

COPY . ${APP_PATH}
WORKDIR ${APP_PATH}


