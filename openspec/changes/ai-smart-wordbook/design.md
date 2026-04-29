## Context

本项目是一个课程作业，要求从零构建「AI 智能单词本」全栈应用。技术栈已明确限定：Go+Gin 后端、Vue3+Vite 前端、MySQL 8.0 数据库、Docker Compose 编排。项目对跨域处理、数据库初始化方式、目录结构有严格约束。

当前状态：空项目，需要从零搭建全部代码和基础设施。

## Goals / Non-Goals

**Goals:**
- 完整实现用户认证 + 单词学习两大业务模块
- 严格遵守跨域处理规范（Vite proxy + Nginx 反向代理，禁止后端 CORS）
- 通过 `docker-compose up -d` 一键启动全部服务
- 产出 README、API 文档、DB 文档三份交付文档

**Non-Goals:**
- 不做高可用/集群部署
- 不做 CI/CD 流水线
- 不做国际化/多语言
- 不做复杂的前端状态管理（Pinia 仅在需要时引入）

## Decisions

### 1. 后端分层架构

采用经典三层架构：`api`（路由+Handler）→ `service`（业务逻辑）→ `model`（数据模型+GORM）。

**理由**: 作业体量适中，三层架构足够清晰；过度抽象（如 DDD）在此场景下增加不必要复杂度。额外增加 `middleware` 层放置 JWT 鉴权中间件，`config` 层管理 Viper 配置。

### 2. 数据库设计：三表结构

| 表 | 用途 |
|---|---|
| `users` | 用户信息，存储哈希密码 |
| `words` | 单词主记录，关联 user_id |
| `sentences` | 例句，关联 word_id（一词多句） |

**理由**: 例句独立成表（而非 JSON 字段存储）便于后续扩展查询；words 和 sentences 一对多关系清晰。通过 user_id 外键实现用户与单词的绑定。软删除使用 GORM 的 `deleted_at` 字段。

**替代方案**: 将例句 JSON 序列化存入 words 表的 text 字段。放弃原因：不利于按例句检索，且不符合关系型设计规范。

### 3. AI 接口对接：策略模式

定义统一的 `AIProvider` 接口，DeepSeek 和通义千问各自实现。通过前端传入的 `ai_provider` 参数动态选择。

```
type AIProvider interface {
    QueryWord(word string) (*WordResult, error)
}
```

**理由**: 遵循开闭原则，新增 AI 提供商只需新增实现，不改动已有代码。每个 provider 内部处理各自 API 的 prompt 构造和响应解析。

### 4. JWT 鉴权方案

- 使用 `golang-jwt/jwt/v5` 签发 HS256 Token
- Token 有效期 24 小时
- 前端存储在 localStorage，请求时通过 `Authorization: Bearer <token>` 携带
- 后端 Gin 中间件统一校验，解析 UserID 注入 Context

### 5. 跨域处理（硬性约束）

- **开发环境**: Vite `server.proxy` 将 `/api` 请求代理到 `http://localhost:8080`
- **生产环境**: Nginx `location /api/` 通过 `proxy_pass` 转发到 `http://backend:8080`
- **后端零 CORS 配置**: 不引入任何 CORS 中间件

### 6. Docker 编排策略

- 后端 Dockerfile：两阶段构建（golang:alpine 编译 → scratch/alpine 运行）
- 前端 Dockerfile：两阶段构建（node:alpine 编译 → nginx:alpine 运行）
- docker-compose 定义三个 service：db、backend、frontend
- 仅 frontend(Nginx) 暴露 80 端口到宿主机
- 数据库通过挂载 `docs/init.sql` 到 `/docker-entrypoint-initdb.d/` 初始化

## Risks / Trade-offs

- **AI API Key 安全**: 通过 `.env` 文件管理，docker-compose 中以环境变量注入 → 不提交真实 Key 到仓库，仅提交模板
- **AI 接口超时/限流**: 设置 HTTP Client 30 秒超时 → 前端显示友好错误提示
- **JWT 无状态缺陷**: Token 一旦签发无法服务端撤销 → 本项目为作业级别，可接受；生产环境需引入 Token 黑名单或 Redis
- **MySQL 初次启动延迟**: backend 容器启动可能早于 db 就绪 → docker-compose 使用 `depends_on` + 后端代码增加数据库连接重试逻辑
