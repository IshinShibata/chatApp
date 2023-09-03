# 使用するベースイメージ
FROM golang:1.19.4

# アプリケーションのディレクトリを設定
WORKDIR /go/src/chat-app

# 依存関係をコピーしてinstall
COPY go.mod .
COPY go.sum .
RUN go mod download && \
    go get github.com/cosmtrek/air && \
    go install github.com/cosmtrek/air@latest

# アプリケーションのソースをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .

# アプリケーションを実行
CMD ["./main"]
