# ベースイメージを指定
FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

# goplantuml のインストール
RUN go install github.com/jfeliu007/goplantuml/cmd/goplantuml@latest

# ワークディレクトリを設定
WORKDIR /workspace
