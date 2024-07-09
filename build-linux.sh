#!/bin/sh

echo "打包中..."

set GOOS=linux
go build -o main
set GOOS=windos

echo "打包完成！"
