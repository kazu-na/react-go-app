FROM golang:1.17.1-alpine3.14

# Goビルドに必要なアプリケーションをインストールする
RUN apk update \
    && apk add --no-cache git \
    && go get github.com/gin-gonic/gin \
    && go get github.com/cosmtrek/air \
    && go build -o /go/bin/air github.com/cosmtrek/air \
    && go get github.com/gorilla/mux \
    && go get github.com/jinzhu/gorm \
    && go get github.com/go-sql-driver/mysql \
    && go get github.com/gin-contrib/cors \
    && go get gopkg.in/ini.v1

ENV GO111MODULE=on

WORKDIR /go/src/api

# go moduleで依存パッケージを取得してダウンロードする
# ローカルの、go.modとgo.sumをコピー
COPY go.mod go.sum ./
# 依存ライブラリをダウンロードする
RUN go mod download

COPY ./api /go/src/api

CMD ["air", "-c", ".air.toml"]