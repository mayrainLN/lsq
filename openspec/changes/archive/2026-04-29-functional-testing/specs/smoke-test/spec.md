## ADDED Requirements

### Requirement: Docker Compose startup
docker-compose up -d SHALL 成功启动 db、backend、frontend 三个容器，且全部处于 healthy/running 状态。

#### Scenario: All services start
- **WHEN** 在项目根目录执行 `docker-compose up -d --build`
- **THEN** 三个容器全部启动，`docker-compose ps` 显示 Up 状态

### Requirement: User registration API
POST /api/register SHALL 成功创建用户。

#### Scenario: Register new user
- **WHEN** 发送 `{"username":"testuser","password":"test123456"}` 到 `/api/register`
- **THEN** 返回 200 和 `{"message":"注册成功"}`

#### Scenario: Duplicate registration
- **WHEN** 使用已注册的用户名再次注册
- **THEN** 返回 409 和 `用户名已存在`

### Requirement: User login API
POST /api/login SHALL 返回有效 JWT Token。

#### Scenario: Login with correct credentials
- **WHEN** 发送正确的用户名密码到 `/api/login`
- **THEN** 返回 200 和包含 `token` 字段的 JSON

### Requirement: DeepSeek word query
POST /api/words/query SHALL 调用 DeepSeek API 返回释义和 3 条例句。

#### Scenario: Query word via DeepSeek
- **WHEN** 已登录用户发送 `{"word":"hello","ai_provider":"deepseek"}` 到 `/api/words/query`
- **THEN** 返回 200，包含 definition 和 3 条 sentences，saved 为 false

### Requirement: Save and list words
POST /api/words SHALL 保存单词，GET /api/words SHALL 返回分页列表。

#### Scenario: Save then list
- **WHEN** 保存查词结果后请求 GET /api/words
- **THEN** 列表中包含刚保存的单词，total >= 1

### Requirement: Delete word
DELETE /api/words/:id SHALL 软删除单词。

#### Scenario: Delete saved word
- **WHEN** 删除已保存的单词后再次请求列表
- **THEN** 该单词不再出现在列表中
