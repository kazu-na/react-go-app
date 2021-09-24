FROM golang:1.17.1-buster

# Goビルドに必要なアプリケーションをインストールする
RUN go get -u github.com/gin-gonic/gin \
    && go install github.com/cosmtrek/air@latest \
    && go get github.com/gorilla/mux \
    && go get github.com/jinzhu/gorm \
    && go get github.com/go-sql-driver/mysql \
    && go get github.com/gin-contrib/cors \
    && go get gopkg.in/ini.v1

ENV GO111MODULE=on

WORKDIR /go/src/api

# go moduleで依存パッケージを取得してダウンロードする
# ローカルの、go.modとgo.sumをコピー
COPY api/go.mod api/go.sum ./
# 依存ライブラリをダウンロードする
RUN go mod download

COPY ./api .

CMD ["air", "-c", ".air.toml"]