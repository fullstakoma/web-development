# ビルド用のGo公式イメージを指定
FROM golang:1.20 AS builder

# 作業ディレクトリの作成
WORKDIR /app

# クロスコンパイル設定
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# ソースコードをコンテナにコピー
COPY go.mod ./
RUN go mod download
# アプリケーションのビルド
COPY . .
RUN go build -o main .

# 実行用の軽量イメージとしてAlpine Linuxを利用
FROM alpine:latest

WORKDIR /app

# ビルド済みバイナリをコピー
COPY --from=builder /app/main .

# セキュリティのためnonrootユーザーを使用（オプション）
RUN chmod +x /app/main

# アプリケーションがリッスンするポート（この例では8080）を公開
EXPOSE 8080

# コンテナ起動時に実行するコマンド
CMD ["/app/main"]