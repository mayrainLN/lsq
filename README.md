# AI 智能单词本

一个前后端分离的英语学习 Web 应用。用户查询单词后，AI 大模型生成释义和例句，支持保存到个人单词本。

## 架构说明

```
浏览器 ──80──▸ Nginx (frontend)
                ├── /         → Vue3 SPA 静态资源
                └── /api/     → proxy_pass ──▸ Go Gin (backend:8080)
                                                   │
                                                   ├── JWT 认证
                                                   ├── DeepSeek / 通义千问 API
                                                   └── MySQL 8.0 (db:3306)
```

**技术栈：**
- 后端：Go 1.21 + Gin + GORM + JWT + Viper
- 前端：Vue3 + TypeScript + Vite + Vue Router + Axios
- 数据库：MySQL 8.0
- 部署：Docker + Docker Compose + Nginx

## 前置依赖

- [Docker](https://docs.docker.com/get-docker/) & Docker Compose
- AI API Key（至少配置一个）：
  - [DeepSeek API Key](https://platform.deepseek.com/)
  - [通义千问 API Key](https://dashscope.console.aliyun.com/)

## 快速启动

### 1. 配置 API Key

在项目根目录创建 `.env` 文件：

```bash
DEEPSEEK_API_KEY=your-deepseek-api-key
TONGYI_API_KEY=your-tongyi-api-key
JWT_SECRET=your-jwt-secret
```

### 2. 一键启动

```bash
docker-compose up -d
```

首次启动会自动：
- 拉取 MySQL 8.0 镜像并初始化数据库（执行 `docs/init.sql`）
- 编译构建后端 Go 应用
- 构建前端 Vue3 应用并部署到 Nginx

### 3. 访问应用

打开浏览器访问：**http://localhost**

### 4. 停止服务

```bash
docker-compose down
```

## 开发环境

### 后端

```bash
cd backend
cp .env .env.local   # 修改配置
go run main.go       # 启动后端 :8080
```

### 前端

```bash
cd frontend
npm install
npm run dev          # 启动前端 :5173（自动代理 /api → :8080）
```

## 功能清单

- [x] 用户注册 / 登录（JWT 鉴权）
- [x] 智能查词（DeepSeek / 通义千问）
- [x] 手动保存单词到单词本
- [x] 单词本分页展示
- [x] 单词软删除
- [x] 跨域处理（开发: Vite Proxy / 生产: Nginx 反向代理）
- [x] Docker Compose 一键部署

## 跨域处理

本项目**严格禁止**后端 CORS 中间件：

- **开发环境**：Vite `server.proxy` 将 `/api` 请求代理到 `http://localhost:8080`
- **生产环境**：Nginx `proxy_pass` 将 `/api/` 转发到内部 `backend:8080` 容器

## 项目文档

- [API 接口文档](docs/api.md)
- [数据库设计文档](docs/db.md)
- [数据库初始化脚本](docs/init.sql)
