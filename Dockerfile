# 在 Docker 的多阶段构建中，最终镜像的内容完全取决于最后一个阶段中复制的文件和目录。前面的构建阶段可以包含任意
# 数量的文件和指令，但只有那些在最后一个阶段明确复制的文件和目录会出现在最终的镜像中（只有在最后一个阶段明确复制
# 的文件会出现在最终的镜像中。构建阶段的所有其他文件和目录都不会出现在最终镜像中，除非它们被显式复制）。
# -> 多阶段构建优点：可以确保最终镜像包含最少的必要文件，从而减小镜像大小，减少潜在的安全风险，并优化部署效率。

# 第一阶段：构建（在第一阶段进行构建和编译，生成最终需要的文件。）
# 使用基于 Alpine Linux 的 Go 官方镜像作为构建阶段的基础镜像
FROM golang:1.22.5-alpine3.20 as builder

# 设置工作目录为
WORKDIR /go/src/github.com/wangyupo/ggb/server

# 将当前目录下的所有文件复制到构建阶段的工作目录中
COPY . .

# 设置 Go 环境变量，清理依赖，并构建项目
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

# 第二阶段：运行（第二阶段使用精简的基础镜像，仅复制运行时需要的文件）
# 使用最新的 Alpine Linux 官方镜像作为运行阶段的基础镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /go/src/github.com/wangyupo/ggb/server

# 从构建阶段复制文件到运行阶段（最终只保留这些被复制的文件，构建阶段的其它文件则不会被保留）
COPY --from=0 /go/src/github.com/wangyupo/ggb/server/server ./
COPY --from=0 /go/src/github.com/wangyupo/ggb/server/resource ./resource/
COPY --from=0 /go/src/github.com/wangyupo/ggb/server/config.docker.yaml ./

# 声明容器在运行时监听 5312 端口
EXPOSE 5312

# 启动容器（执行 ./server -c config.docker.yaml，使用 config.docker.yaml 配置文件来启动应用程序）
ENTRYPOINT ["./server", "-c", "config.docker.yaml"]
