# goバージョン
FROM golang:1.18.3-alpine
# アップデートとgitのインストール！！
RUN apk add --update &&  apk add git
# ワーキングディレクトリの設定
WORKDIR /go/src/app

# 基本docker-composeのvoluneでマウントしているが、dockerfileでgo mod tidyを実行するためにCOPYする。Copyしないとエラーになる。
COPY go.mod .
#go modのダウンロード
RUN go mod tidy