# ベースイメージとしてGolangを使用
FROM golang:latest

# アプリケーションのディレクトリを作成
RUN mkdir /app

# 作業ディレクトリを設定
WORKDIR /app

# ローカルのファイルをコンテナにコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .

# アプリケーションを起動
CMD ["./main"]