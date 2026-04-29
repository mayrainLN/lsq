## ADDED Requirements

### Requirement: User registration
系统 SHALL 提供用户注册接口 `POST /api/register`，接收 username 和 password，将密码通过 bcrypt 哈希后存入 MySQL users 表。用户名 MUST 唯一。

#### Scenario: Successful registration
- **WHEN** 用户提交有效的用户名和密码（用户名未被占用）
- **THEN** 系统创建用户记录，密码以 bcrypt 哈希存储，返回成功信息

#### Scenario: Duplicate username
- **WHEN** 用户提交的用户名已存在于数据库中
- **THEN** 系统返回错误提示「用户名已存在」，不创建记录

#### Scenario: Invalid input
- **WHEN** 用户提交空用户名或空密码
- **THEN** 系统返回参数校验错误

### Requirement: User login with JWT
系统 SHALL 提供登录接口 `POST /api/login`，验证用户名密码后签发 JWT Token（HS256，24 小时有效期）。

#### Scenario: Successful login
- **WHEN** 用户提交正确的用户名和密码
- **THEN** 系统返回 JWT Token，Token payload 包含 user_id 和过期时间

#### Scenario: Wrong credentials
- **WHEN** 用户提交错误的用户名或密码
- **THEN** 系统返回认证失败错误，不签发 Token

### Requirement: JWT authentication middleware
系统 SHALL 在受保护的 API 路由上应用 JWT 鉴权中间件，从请求头 `Authorization: Bearer <token>` 中解析并验证 Token。

#### Scenario: Valid token
- **WHEN** 请求携带有效且未过期的 JWT Token
- **THEN** 中间件解析出 user_id 并注入请求上下文，请求正常通过

#### Scenario: Missing or invalid token
- **WHEN** 请求未携带 Token 或 Token 格式无效
- **THEN** 中间件返回 401 Unauthorized，请求被拦截

#### Scenario: Expired token
- **WHEN** 请求携带的 JWT Token 已过期
- **THEN** 中间件返回 401 Unauthorized
