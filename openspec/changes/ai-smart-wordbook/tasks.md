## 1. 项目骨架与基础设施

- [ ] 1.1 创建项目根目录结构：`backend/`、`frontend/`、`docs/`
- [ ] 1.2 初始化 Go 模块（`go mod init`），添加 gin、gorm、jwt、viper、bcrypt 等依赖
- [ ] 1.3 创建后端入口 `backend/main.go`，配置 Gin 引擎、路由分组、启动监听 8080
- [ ] 1.4 创建 `backend/.env` 配置模板（DB 连接、JWT 密钥、AI API Key 等）
- [ ] 1.5 创建 `backend/config/` 目录，使用 Viper 加载 .env 配置
- [ ] 1.6 使用 Vite 初始化 Vue3+TypeScript 前端项目到 `frontend/` 目录

## 2. 数据库设计与初始化

- [ ] 2.1 编写 `docs/init.sql`：创建 users 表（id, username, password, created_at, updated_at, deleted_at）
- [ ] 2.2 编写 `docs/init.sql`：创建 words 表（id, user_id, word, definition, ai_provider, created_at, updated_at, deleted_at）
- [ ] 2.3 编写 `docs/init.sql`：创建 sentences 表（id, word_id, english, chinese, created_at, updated_at, deleted_at）
- [ ] 2.4 创建 `backend/model/` 目录，定义 User、Word、Sentence 的 GORM 模型结构体
- [ ] 2.5 创建 `backend/model/database.go`，实现 GORM 数据库连接初始化（含重试逻辑）

## 3. 用户认证模块（后端）

- [ ] 3.1 实现 `POST /api/register`：参数校验、密码 bcrypt 哈希、用户创建、用户名唯一检查
- [ ] 3.2 实现 `POST /api/login`：用户名密码验证、JWT Token 签发（HS256，24h 有效期）
- [ ] 3.3 创建 `backend/middleware/auth.go`：JWT 鉴权中间件，解析 Token 并注入 user_id 到 Context
- [ ] 3.4 在 Gin 路由中配置公开路由组（register/login）和受保护路由组（需 JWT 中间件）

## 4. 单词学习模块（后端）

- [ ] 4.1 实现 `POST /api/words/query`：鉴权后先查库，已保存返回数据库记录，未保存调用 AI 接口
- [ ] 4.2 实现 `POST /api/words`：手动保存单词，写入 words 和 sentences 表，绑定 user_id，防重复保存
- [ ] 4.3 实现 `GET /api/words`：分页查询当前用户单词列表（page, page_size），预加载 sentences 关联
- [ ] 4.4 实现 `DELETE /api/words/:id`：软删除单词及关联例句，校验单词归属当前用户

## 5. AI 接口对接（后端）

- [ ] 5.1 定义 `AIProvider` 接口：`QueryWord(word string) (*WordResult, error)`
- [ ] 5.2 实现 DeepSeek provider：构造 prompt、调用 API、解析返回的释义和 3 条例句
- [ ] 5.3 实现通义千问 provider：构造 prompt、调用 API、解析返回的释义和 3 条例句
- [ ] 5.4 实现 provider 工厂函数：根据 ai_provider 字符串返回对应实现，不支持的返回错误
- [ ] 5.5 统一 AI 返回数据结构 `WordResult`：definition(string)、sentences([]Sentence{english, chinese})

## 6. 前端页面实现

- [ ] 6.1 安装 vue-router、axios 依赖，配置路由（登录、注册、查词、单词本）
- [ ] 6.2 创建 axios 实例：请求拦截器附加 JWT Token，响应拦截器处理 401 跳转登录
- [ ] 6.3 实现登录页面：用户名+密码表单，调用 `/api/login`，Token 存入 localStorage
- [ ] 6.4 实现注册页面：用户名+密码表单，调用 `/api/register`，成功后跳转登录
- [ ] 6.5 实现单词查询页：输入框 + AI 模型下拉选择 + 查询按钮，展示释义和例句，保存按钮
- [ ] 6.6 实现单词本列表页：分页展示单词卡片（含释义和例句），分页器组件，删除按钮
- [ ] 6.7 实现路由守卫：未登录用户访问受保护页面时重定向到登录页

## 7. 跨域代理配置

- [ ] 7.1 配置 `frontend/vite.config.ts`：开发环境 server.proxy 将 `/api` 代理到 `http://localhost:8080`
- [ ] 7.2 编写 `frontend/nginx.conf`：`location /api/` proxy_pass 到 `http://backend:8080`，其余路径 try_files 返回 index.html

## 8. Docker 容器化与编排

- [ ] 8.1 编写 `backend/Dockerfile`：多阶段构建（golang:alpine 编译 → alpine 运行），暴露 8080
- [ ] 8.2 编写 `frontend/Dockerfile`：两阶段构建（node:alpine 编译 → nginx:alpine），复制 nginx.conf
- [ ] 8.3 编写 `docker-compose.yml`：定义 db/backend/frontend 三个 service，配置网络、depends_on、环境变量、init.sql 挂载，仅暴露 frontend 80 端口

## 9. 项目文档

- [ ] 9.1 编写 `docs/db.md`：三表字段说明、数据类型、主/外键、索引、关联关系
- [ ] 9.2 编写 `docs/api.md`：所有接口的路径、方法、鉴权、参数、返回示例、错误码
- [ ] 9.3 编写 `README.md`：项目简介、架构说明、前置依赖、.env 配置指引、一键启动命令、访问方式
