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
```
