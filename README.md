# gin-demo

#### 介绍
基于golang web框架Gin搭建的项目demo

#### 软件架构
golang  
Gin  
gorm  
jwt token认证go-jwt  
zap 日志  
Redis

#### 项目目录

|  文件/目录名称   | 说明  |
|  ----  | ----  |
| app/common | 公共模块（请求、响应结构体等） |
|app/controllers | 业务调度器 |
|app/middleware| 中间件 |
|app/models| 数据库结构体 |
|app/services| 业务层 |
|bootstrap| 项目启动初始化 |
|config| 配置文件 |
|config| 配置结构体 |
|global| 全局变量 |
|routes| 路由定义 |
|static| 静态资源（允许外部访问） |
|storage| 系统日志、文件等静态资源） |
|config.yaml| 配置文件 |
|utils| 工具函数 |
|main.go| 项目启动文件 |
|README.md   |        |
