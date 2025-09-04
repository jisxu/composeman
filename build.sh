#!/usr/bin/env bash
set -e

APP_NAME="composeman"
BUILD_DIR="bin"

mkdir -p $BUILD_DIR

# 当前版本号（可从 git tag 获取）
VERSION=$(git describe --tags --always 2>/dev/null || echo "dev")
LDFLAGS="-s -w -X main.Version=$VERSION"

echo "===> 构建 $APP_NAME (version: $VERSION)"

# 本地编译
echo "-> 构建本地平台二进制"
CGO_ENABLED=0 go build -ldflags "$LDFLAGS" -o $BUILD_DIR/$APP_NAME .

## 交叉编译 Linux / macOS / Windows
#echo "-> 交叉编译 Linux amd64"
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$LDFLAGS" -o $BUILD_DIR/${APP_NAME}-linux-amd64 .
#
#echo "-> 交叉编译 macOS amd64"
#CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "$LDFLAGS" -o $BUILD_DIR/${APP_NAME}-darwin-amd64 .
#
#echo "-> 交叉编译 Windows amd64"
#CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "$LDFLAGS" -o $BUILD_DIR/${APP_NAME}-windows-amd64.exe .

echo "✅ 构建完成，文件在 $BUILD_DIR/ 下"
