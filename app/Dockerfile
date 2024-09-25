# Stage 1: ビルドステージ (Go のバージョンは適宜変更してください)
FROM golang:1.19 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o main .

# Stage 2: 実行環境 (軽量な Alpine Linux を使用)
FROM alpine:latest

WORKDIR /app

# マルチステージビルドでビルドしたバイナリをコピー
COPY --from=builder /app/main /app/main

# ポートを公開
EXPOSE 8080

# アプリケーションの実行
CMD ["/app/main"]