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

## 项目目录结构

```
GGB/
├── api                  # API 控制器
│   └── v1               # 版本
├── config               # 全局配置
├── core                 # 核心功能
│   ├── viper.go         # viper初始化
│   └── zap.go           # zap日志初始化
├── enums                # 枚举
├── global               # 全局常量
├── initialize           # 项目初始化
│   ├── router.go        # 路由初始化
│   └── gorm.go          # 数据库初始化
├── log                  # 日志
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
```
