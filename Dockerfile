FROM golang:1.25

WORKDIR /app

# 依存関係のファイルをコピー
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# 開発用: go runで直接実行
CMD ["go", "run", "./cmd/api/main.go"]

