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

## 语言环境

golang版本 >= v1.22

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

1）新建自定义网络

```bash
# 查看已存在的 docker 网络，确认没有重名
docker network ls

# 新建网络，IP 地址范围从 10.1.0.0 到 10.1.255.255，网络名称为 my-net
docker network create --subnet=10.1.0.0/16 my-net

# （无需执行，仅作命令展示）删除自定义网络
docker network rm my-net
```

2）拉取镜像，并在自定义网络上创建容器

```bash
# 拉取 mysql 镜像
docker pull mysql:8.0

# 使用 mysql 镜像创建容器（容器命名为 mysql；使用自定义网络，绑定IP为 10.1.0.2；将 docker 宿主机的 3307 端口映射到容器的 3306 端口；初始化 root 用户的密码为 123456；挂载 mysql 数据和配置卷到本地，以持久化数据）
docker run -itd --name mysql --network my-net --ip 10.1.0.2 -p 3307:3306 -e MYSQL_ROOT_PASSWORD=123456 -v C:/dockerVolumes/mysql/data:/var/lib/mysql -v C:/dockerVolumes/mysql/mysql.conf.d:/etc/mysql/conf.d:ro mysql:8.0

# 拉取 redis 镜像
docker pull redis:latest

# 使用 redis 镜像创建容器（容器命名为 redis；使用自定义网络，绑定IP为 10.1.0.3；将 docker 宿主机的 6380 端口映射到容器的 6379 端口；挂载 redis 数据和配置卷到本地）
docker run -itd --name redis --network my-net --ip 10.1.0.3 -p 6380:6379 -v C:/dockerVolumes/redis/data:/data -v C:/dockerVolumes/redis/redis.conf:/usr/local/etc/redis/redis.conf:ro redis:latest

# 拉取 nginx 镜像
docker pull nginx:latest

# 使用 nginx 镜像创建容器（容器命名为 nginx；使用自定义网络，绑定IP为 10.1.0.3；将 docker 宿主机的 81 端口映射到容器的 80 端口；挂载 nginx 配置卷到本地）
docker run -itd --name nginx --network my-net --ip 10.1.0.4 -p 81:80 -v C:/dockerVolumes/nginx/nginx.conf:/etc/nginx/nginx.conf:ro -v C:/dockerVolumes/nginx/conf.d:/etc/nginx/conf.d:ro -v C:/dockerVolumes/nginx/html:/usr/share/nginx/html -v C:/dockerVolumes/nginx/log:/var/log/nginx nginx:latest
```

3）修改 config.docker.yaml 配置

```bash
# 修改 mysql 配置
mysql:
  host: 10.1.0.2          # 这里填mysql容器的IP
  password: 123456        # 这里填mysql的root密码

# 修改 redis 配置
redis:
  addr: 10.1.0.3:6378     # 这里填redis容器的IP:端口
  
# （无需执行，仅作命令展示）查看容器的 IP
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' containerName
```

4）使用 Navicat 等数据库工具，连接 docker 容器中的 mysql，并创建数据库 ggb

```bash
# 主机
localhost

# 端口
3307

# 用户名
root

# 密码
123456

# 数据库名称
ggb

# 数据库字符集
utf8mb4

# 数据库排序规则
utf8mb4_general_ci
```

5）创建本项目的 docker 镜像（docker image），并创建容器

```bash
# 创建项目的 docker 镜像（镜像名为 ggb，tag默认为 latest，）
docker build -t ggb .     # 也可指定tag，如：docker build -t ggb:v0.0.1 .

# 创建项目容器（server 服务会在容器启动时自动运行）
docker run --name ggb_server --network my-net --ip 10.1.0.113 -p 5313:5312 ggb

# （无需执行，仅作命令展示）启动已有容器，并附加到其控制台输出（-a），同时保持交互模式（-i）
docker start -a -i my-container
```

### 2、如何访问 OpenAPI（Swagger）？

```bash
# 生成/更新 API 文档
swag init

# 本地启动项目
go run main.go

# 访问本地 OpenAPI 地址
http://localhost:5312/swagger/index.html

# docker 启动项目（启动项目容器时，已经把 docker 宿主机的 5313 端口已经映射到容器的 5312 端口）
docker run -p 5313:5312 --name=ggb_server ggb

# 访问本地映射的 OpenAPI 地址
http://localhost:5313/swagger/index.html
```

## License

[MIT © Richard McRichface.](https://github.com/wangyupo/GGB/blob/main/LICENSE)