## ADDED Requirements

### Requirement: Backend multi-stage Docker build

后端 Dockerfile SHALL 使用多阶段构建：第一阶段 golang:alpine 编译二进制，第二阶段使用精简基础镜像运行。暴露 8080 端口。

#### Scenario: Build backend image

- **WHEN** 执行 `docker build` 构建后端镜像
- **THEN** 产出精简镜像，仅包含编译后的 Go 二进制和必要的运行时文件

### Requirement: Frontend Nginx Docker build

前端 Dockerfile SHALL 使用两阶段构建：第一阶段 node:alpine 执行 `npm run build`，第二阶段将产物拷贝到 nginx:alpine 的发布目录，并替换自定义 nginx.conf。

#### Scenario: Build frontend image

- **WHEN** 执行 `docker build` 构建前端镜像
- **THEN** 产出基于 nginx:alpine 的镜像，包含 Vite 构建产物和自定义 Nginx 配置

### Requirement: Nginx reverse proxy configuration

nginx.conf SHALL 将 `/api/` 路径通过 `proxy_pass` 反向代理到 `http://backend:8080`，其余路径返回前端静态资源。MUST 支持 SPA history 模式的 `try_files`。

#### Scenario: API request proxy

- **WHEN** 浏览器请求 `/api/words`
- **THEN** Nginx 将请求转发到 backend 容器的 8080 端口

#### Scenario: Frontend SPA routing

- **WHEN** 浏览器直接访问 `/login` 或 `/wordbook` 等前端路由
- **THEN** Nginx 返回 `index.html`，由前端路由接管

### Requirement: Docker Compose orchestration

docker-compose.yml SHALL 定义 db（MySQL 8.0）、backend、frontend 三个服务。仅 frontend 暴露 80 端口到宿主机。backend depends_on db。数据库通过挂载 `docs/init.sql` 初始化。

#### Scenario: One-command startup

- **WHEN** 在项目根目录执行 `docker-compose up -d`
- **THEN** 三个容器依次启动，MySQL 自动执行 init.sql 建表，访问 `http://localhost` 即可使用

#### Scenario: Internal network isolation

- **WHEN** 三个容器运行在同一 Docker 网络中
- **THEN** frontend 可通过容器名 `backend` 访问后端，backend 可通过容器名 `db` 访问数据库；db 和 backend 端口不暴露到宿主机

### Requirement: Vite dev proxy

Vite 配置 SHALL 在开发环境下将 `/api` 前缀的请求代理到 `http://localhost:8080`，实现开发环境下的同源访问。

#### Scenario: Dev mode API proxy

- **WHEN** 前端开发服务器运行时，浏览器请求 `/api/login`
- **THEN** Vite dev server 将请求代理到 `http://localhost:8080/api/login`

