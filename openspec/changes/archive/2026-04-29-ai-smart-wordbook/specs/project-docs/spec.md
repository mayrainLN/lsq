## ADDED Requirements

### Requirement: README with run guide
项目根目录 SHALL 包含 README.md，内含项目简介、架构说明、前置依赖说明、API Key 配置说明、一键启动命令 `docker-compose up -d`、访问方式说明。

#### Scenario: New user follows README
- **WHEN** 新用户按照 README.md 中的步骤操作
- **THEN** 用户能成功配置 `.env` 文件中的 AI API Key，执行 `docker-compose up -d` 启动全部服务，并通过浏览器访问应用

### Requirement: API documentation
docs/api.md SHALL 记录每个 API 接口的路径、方法、鉴权要求、请求参数、成功返回示例、错误码及含义。

#### Scenario: API doc completeness
- **WHEN** 开发者查阅 docs/api.md
- **THEN** 可找到所有业务接口的完整调用说明，包括注册、登录、查词、保存、列表、删除

### Requirement: Database design documentation
docs/db.md SHALL 阐述所有数据库表的字段名、数据类型、主/外键、索引、业务含义，以及表间关联关系。

#### Scenario: DB doc completeness
- **WHEN** 开发者查阅 docs/db.md
- **THEN** 可了解 users、words、sentences 三表的完整字段设计和关联关系

### Requirement: Database init script
docs/init.sql SHALL 包含所有建表语句，通过 Docker 挂载到 MySQL `/docker-entrypoint-initdb.d/` 自动执行。MUST NOT 依赖 GORM AutoMigrate。

#### Scenario: Database initialization
- **WHEN** MySQL 容器首次启动并执行 init.sql
- **THEN** users、words、sentences 三张表被正确创建，包含所有字段、索引和外键约束
