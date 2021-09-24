FROM golang:1.17.1-alpine3.14

COPY api /go/src/api

WORKDIR /go/src/api

# Install dependency module
RUN apk update \
    && apk add --no-cache git \
    && go get github.com/gin-gonic/gin \
    && go get github.com/gorilla/mux \
    && go get github.com/jinzhu/gorm \
    && go get github.com/go-sql-driver/mysql \
    && go get github.com/gin-contrib/cors \
    && go get gopkg.in/ini.v1