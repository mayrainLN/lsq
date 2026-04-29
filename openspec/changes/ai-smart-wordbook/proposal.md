## Why

完成「全栈开发实战」课程作业：构建一个 AI 智能单词本 Web 应用。项目重点考察前后端分离架构、跨域代理处理、第三方 AI 接口对接、关系型数据库设计，以及基于 Docker 的全栈容器编排部署能力。

## What Changes

- 搭建前后端分离项目骨架，后端 Go+Gin、前端 Vite+Vue3
- 实现用户注册/登录模块，使用 JWT 鉴权，密码 bcrypt 哈希存储
- 实现单词智能查询功能，对接 DeepSeek / 通义千问 AI 大模型，返回释义 + 3 条例句
- 实现单词手动保存、分页列表、软删除功能
- 设计 MySQL 数据库表结构（users、words、sentences），通过 init.sql 初始化
- 配置 Vite proxy（开发环境）和 Nginx 反向代理（生产环境）处理跨域，**禁止后端 CORS 中间件**
- 编写 Docker 多阶段构建（后端）和 Nginx 镜像（前端），通过 docker-compose 一键编排
- 编写 README.md、API 接口文档、数据库设计文档

## Capabilities

### New Capabilities
- `user-auth`: 用户注册与登录，JWT Token 签发与验证，密码哈希加密
- `word-learning`: 单词智能查询（AI 调用）、手动保存、分页列表、软删除
- `ai-integration`: 对接 DeepSeek 和通义千问 API，返回结构化 JSON（释义 + 例句）
- `frontend-ui`: Vue3 前端页面，包含登录/注册、单词查询、单词本列表、分页器
- `deployment`: Docker 多阶段构建、Nginx 反向代理、docker-compose 三服务编排
- `project-docs`: README.md、API 文档、数据库设计文档、init.sql 初始化脚本

### Modified Capabilities

（无已有能力需要修改，这是全新项目）

## Impact

- **目录结构**: 创建 `backend/`、`frontend/`、`docs/` 完整项目目录
- **依赖引入**: Go 模块（gin, gorm, jwt-go, viper, bcrypt）、Node 模块（vue3, axios, vue-router）
- **外部服务**: MySQL 8.0 数据库、DeepSeek API、通义千问 API
- **部署产物**: 3 个 Docker 容器（db、backend、frontend），对外仅暴露 Nginx 80 端口
