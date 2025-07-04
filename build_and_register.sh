#!/bin/bash

set -e

# 빌드할 바이너리 이름
BIN_NAME="sorting"

# 현재 디렉토리에서 main.go 빌드
go build -o "$BIN_NAME" main.go

# 바이너리를 /usr/local/bin 으로 이동 (sudo 필요할 수 있음)
if [ -w /usr/local/bin ]; then
    mv -f "$BIN_NAME" /usr/local/bin/
else
    sudo mv -f "$BIN_NAME" /usr/local/bin/
fi

echo "설치 완료: /usr/local/bin/$BIN_NAME"
echo "터미널에서 '$BIN_NAME' 명령어로 실행할 수 있습니다."
