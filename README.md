<br />
<div align="center">
  <a href="https://github.com/wangyupo/GGB">
    <img src="./logo.png" alt="Logo" width="180" height="160">
  </a>

  <h3 align="center">GGB（猪猪侠）后端框架</h3>

  <p>
    基于 gin 搭建的后端框架。
    <br />
    Beauty and standards ✨
    <br />
    <br />
    ·
    <a href="https://github.com/wangyupo/GGB/issues">报告 Bug</a>
    ·
    <a href="https://github.com/wangyupo/GGB/issues">建议需求</a>
  </p>
</div>

## 项目介绍

GGB（猪猪侠）通过清晰的目录结构和模块化设计，为开发者提供了一套高效、可维护的后端服务架构。无论是初学者还是经验丰富的开发者，都能通过 GGB 快速上手并构建高质量的 Web 应用。

## 项目运行

```bash
# 克隆项目
git clone https://github.com/wangyupo/GGB.git

# 进入项目文件夹
cd GGB

# 使用生成指令，执行配置环境、安装依赖包等一系列操作
go generate

# 编译 
go build main.go

# 运行编译好的包
./main (windows运行命令为 ./main.exe)
```

### 指定配置文件运行

```bash
# 本地启动服务
go run main.go

# 根据指定yaml配置启动服务
go run main.go -c ./config.docker.yaml
```

### windows编译并生成可执行文件

```bash
set GOOS=linux
go build -o main
set GOOS=windos

或直接执行以下命令：

./build-linux.sh
```

## 项目目录结构

```
GGB/
├── api                  # API 控制器
│   └── v1               # v1版本
├── config               # 全局配置
├── core                 # 核心功能
│   ├── viper.go         # viper初始化
│   └── zap.go           # zap日志初始化
├── enums                # 枚举
├── global               # 全局变量
│   └── global.go        # 全局实例
├── initialize           # 项目初始化
│   ├── router.go        # 路由初始化
│   ├── gorm.go          # 数据库初始化
│   └── timer.go         # 定时器初始化
├── log                  # 日志
│   ├── 2024-10-01       # 按日期分类储存
│   └── ...
├── middleware           # 中间件
│   ├── jwt.go           # 鉴权
│   └── operation.go     # 操作日志
├── model                # 数据模型
│   ├── system           # 数据库实体
│   ├── request          # 请求参数
│   └── response         # 响应参数
├── resource             # 资源
│   └── excel            # Excel模板
├── router               # 路由
├── service              # 服务层
│   ├── log              # 日志相关服务
│   └── system           # 系统相关服务
├── uploads              # 上传文件本地存储目录
└── utils                # 工具函数
    ├── timer            # 定时器
    └── upload           # oss
```

## 问题注解

### 1、如何使用docker部署该项目？

1）创建必要容器

```bash
# 拉取 mysql 镜像
docker pull mysql5.7

# 使用 mysql 镜像创建容器（将 docker 宿主机的 3307 端口映射到容器的 3306 端口；容器命名为 mysql；初始化 root 用户的密码为 123456）
docker run -itd -p 3307:3306 --name=mysql -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7

# 拉取 redis 镜像
docker pull redis:latest

# 使用 redis 镜像创建容器（将 docker 宿主机的 6379 端口映射到容器的 6378 端口；容器命名为 redis）
docker run -itd -p 6379:6378 --name=redis redis:latest

# 拉取 nginx 镜像
docker pull nginx:latest

# 使用 nginx 镜像创建容器（将 docker 宿主机的 81 端口映射到容器的 80 端口；容器命名为 nginx）
docker run -itd -p 81:80 --name=nginx nginx:latest
```

2）查看容器 IP，配置 config.docker.yaml

```bash
# 查看 mysql 容器的 IP
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mysql

# 修改 config.docker.yaml 中的 mysql 配置
mysql:
  host: 这里填你查到的mysql容器的IP
  password: 这里填你设置的mysql的root密码

# 查看 redis 容器的 IP
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' redis

# 修改 config.docker.yaml 中的 redis 配置
redis:
  addr: 这里填你查到的redis容器的IP
```

3）创建本项目的docker镜像（docker image），并创建容器

```bash
# 创建项目的 docker 镜像（镜像名为 ggb）
docker build -t ggb .

# 创建项目容器（程序会自动运行）
docker run -p 5313:5312 --name=ggb_server ggb
```

### 2、如何访问OpenAPI（Swagger）？

```bash
# 本地启动项目
http://localhost:5312/swagger/index.html

# docker 启动项目（启动项目容器时，已经把 docker 宿主机的 5313 端口已经映射到容器的 5312 端口）
http://localhost:5313/swagger/index.html
```

## License

[MIT © Richard McRichface.](https://github.com/wangyupo/GGB/blob/main/LICENSE)