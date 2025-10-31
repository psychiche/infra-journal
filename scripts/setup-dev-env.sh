#!/usr/bin/env bash
set -euo pipefail

# Go environment
if command -v go >/dev/null 2>&1; then
  echo "Go version: $(go version)"
else
  echo "请安装 Go 1.22+" >&2
fi

# Python environment
if command -v python3 >/dev/null 2>&1; then
  echo "Python: $(python3 --version)"
  if command -v pip3 >/dev/null 2>&1; then
    pip3 install --upgrade pip >/dev/null || true
    pip3 install redis >/dev/null || true
    echo "已安装 Python 依赖（redis）"
  fi
else
  echo "请安装 Python 3.9+" >&2
fi

# Kafka/Redis hints
echo "请确保本地 Kafka 与 Redis 已就绪（或使用 Docker Compose）。"
