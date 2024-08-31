# ベースイメージを指定
FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

# goplantuml のインストール
RUN go install github.com/jfeliu007/goplantuml/cmd/goplantuml@latest

# pip3 のインストールと、requirements.txt のインストール
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends python3-pip graphviz \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/* \
    && pip3 install --no-cache-dir -U pip setuptools wheel \
    && pip3 install --no-cache-dir -U -r https://raw.githubusercontent.com/jfeliu007/goplantuml/main/requirements.txt

# ワークディレクトリを設定
WORKDIR /workspace
