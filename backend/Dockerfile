# Goのベースイメージを取得
FROM golang:alpine

RUN apk update
RUN apk add alpine-sdk

# ターゲットディレクトリを作成
RUN mkdir -p /go/src/app

# コンテナ内の作業ディレクトリを設定
WORKDIR /go/src/app

# go.mod ファイルをコピー
COPY go.mod .

# .go ファイルをコピー
COPY . .

# バイナリをビルド
RUN go build -o main .

# コンテナ起動時に実行されるコマンド
CMD ["./main"]