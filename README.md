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

## Todo

- [x] 发送、验证邮件
- [ ] 集成 Elasticsearch

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

# 使用 mysql 镜像创建容器（容器命名为 mysql；使用自定义网络，绑定IP为 10.1.0.2；将 docker 宿主机的 3307 端口映射到容器的 3306 端口；初始化 root 用户的密码为 123456；初始化数据库 ggb；挂载 mysql 数据和配置卷到本地，以持久化数据；设置字符集为 utf8mb4；排序规则为 utf8mb4_general_ci；）
docker run -itd --name mysql --network my-net --ip 10.1.0.2 -p 3307:3306 -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=ggb -v C:/dockerVolumes/mysql/data:/var/lib/mysql -v C:/dockerVolumes/mysql/mysql.conf.d:/etc/mysql/conf.d:ro --restart unless-stopped mysql:8.0 --character-set-server=utf8mb4 --collation-server=utf8mb4_general_ci

# 拉取 redis 镜像
docker pull redis:latest

# 使用 redis 镜像创建容器（容器命名为 redis；使用自定义网络，绑定IP为 10.1.0.3；将 docker 宿主机的 6380 端口映射到容器的 6379 端口；挂载 redis 数据和配置卷到本地）
docker run -itd --name redis --network my-net --ip 10.1.0.3 -p 6380:6379 -v C:/dockerVolumes/redis/data:/data -v C:/dockerVolumes/redis/redis.conf:/usr/local/etc/redis/redis.conf:ro --restart unless-stopped redis:latest

# 拉取 nginx 镜像
docker pull nginx:latest

# 使用 nginx 镜像创建容器（容器命名为 nginx；使用自定义网络，绑定IP为 10.1.0.4；将 docker 宿主机的 81 端口映射到容器的 80 端口；挂载 nginx 配置卷到本地）
docker run -itd --name nginx --network my-net --ip 10.1.0.4 -p 81:80 -v C:/dockerVolumes/nginx/nginx.conf:/etc/nginx/nginx.conf:ro -v C:/dockerVolumes/nginx/conf.d:/etc/nginx/conf.d:ro -v C:/dockerVolumes/nginx/html:/usr/share/nginx/html -v C:/dockerVolumes/nginx/log:/var/log/nginx --restart unless-stopped nginx:latest
```

3）修改 config.docker.yaml 配置

```bash
# 修改 mysql 配置
mysql:
  host: 10.1.0.2          # 这里填mysql容器的IP
  password: 123456        # 这里填mysql的root密码

# 修改 redis 配置
redis:
  addr: 10.1.0.3:6379     # 这里填redis容器的IP:端口
  
# （无需执行，仅作命令展示）查看容器的 IP
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' containerName
```

4）创建本项目的 docker 镜像（docker image），并创建容器

```bash
# 创建项目的 docker 镜像（镜像名为 ggb，tag 默认为 latest，）
docker build -t ggb .     # 也可指定 tag，如：docker build -t ggb:v0.0.1 .

# 创建项目容器（server 服务会在容器启动时自动运行）
docker run --name ggb_server --network my-net --ip 10.1.0.113 -p 5313:5312 --restart unless-stopped ggb

# （无需执行，仅作命令展示）启动已有容器，并附加到其控制台输出（-a），同时保持交互模式（-i）
docker start -a -i my-container
```

### 2、如何将 docker 镜像移动到另一个环境中加载并使用？

```bash
# 导出 docker 镜像
# eg：docker save -o <path_to_tar_file> <image_name>:<tag>
docker save -o ggb.tar ggb:latest

# 上传镜像文件
scp ggb.tar user@remote_host:/path/to/destination

# 加载 docker 镜像
# eg：docker load -i <path_to_tar_file>
docker load -i ggb.tar

# 确认镜像加载成功（你应该能看到 ggb:latest 镜像在列表中）
docker images
```

### 3、如何使用 docker-compose 部署该项目？

1）安装 docker-compose

```bash
# 更新包管理器和安装依赖
sudo apt update
sudo apt install -y curl

# 下载最新版本的 Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

# 赋予执行权限
sudo chmod +x /usr/local/bin/docker-compose

# 创建符号链接（可选，但推荐）
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose

# 验证安装
docker-compose --version      # 这应该输出 Docker Compose 的版本号，表示安装成功
```

2）将本项目根目录下的 docker-compose.yml 复制到你的服务器上项目所处目录中去

3）在你的项目所处目录中，运行以下指令

```bash
# 启动所有服务
docker-compose up -d

# （需执行，仅作命令展示）查看运行中的服务
docker-compose ps

# （需执行，仅作命令展示）停止所有服务但不删除容器
docker-compose stop

# （需执行，仅作命令展示）停止所有服务并删除所有容器
docker-compose down
```

_注意：如果发现 mysql 没有新建数据库 ggb，可能是因为已存在的 mysql 映射卷的缘故，请执行以下命令：_

```bash
# 停止所有服务并删除所有容器
docker-compose down

# 删除 mysql 的映射卷
sudo rm -r C:/dockerVolumes/mysql/

# 重启启动所有服务
docker-compose up -d
```

### 4、如何使用数据库工具管理 docker 容器中的 mysql？

使用 Navicat 等数据库工具，连接 docker 容器中的 mysql 即可，具体配置如下：

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
```

### 5、如何查看 docker 容器中的服务的实时日志？

```bash
# 确认 docker-compose 中 ggb_server 的 GIN_MODE=debug
ggb_server:
  environment:
    - GIN_MODE=debug

# 执行问题注解 3 中的步骤，导出镜像->加载镜像

# 使用 docker-compose 启动容器
docker-compose up -d
    
# 查看服务实时日志
docker logs -f ggb_server
```

### 6、如何访问 OpenAPI（Swagger）？

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

### 7、如何集成 GGB 的前端项目？

```bash
# 下载前端项目到本地
git clone https://github.com/wangyupo/GGB_FE

# 进入项目目录
cd GGB_FE

# 安装依赖（node >= v20.16.0）
npm install

# 运行前端项目
npm run dev
```

## 其它项目推荐

1、**[v3s](https://github.com/wangyupo/v3s)** 基于 Vue3、Vite5、Vue Router、Pinia 和 Element Plus 构建的高效后台管理模板。结合 VSCode 插件 v3s snippets，助力快速开发业务应用。

2、**[vue3-cookbook](https://github.com/wangyupo/vue3-cookbook)** 为您提供了一些组件化的范例和资源，以帮助您在 Vue 3 中起步。包括组合式 API (Composition API)、Pinia、Vue Router 使用示例，以及 axios 和 TailwindCSS 集成示例。

3、**[FE-Guide](https://github.com/wangyupo/FE-Guide)** 本文档旨在为前端团队搭建一个标准化的技术栈和代码风格指南，帮助团队成员在开发过程中保持一致，提升整体开发体验。

## License

[MIT © Richard McRichface.](https://github.com/wangyupo/GGB/blob/main/LICENSE)