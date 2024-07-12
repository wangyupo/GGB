#!/bin/sh

# 等待 MySQL 启动
dockerize -wait tcp://10.1.0.2:3306 -timeout 60s

# 启动服务（使用 config.docker.yaml 配置文件来启动应用程序）
./server -c config.docker.yaml
