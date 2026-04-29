## ADDED Requirements

### Requirement: Login and register pages
前端 SHALL 提供登录和注册页面，表单包含用户名和密码输入框。登录成功后将 JWT Token 存入 localStorage，并跳转到主页。

#### Scenario: Login flow
- **WHEN** 用户在登录页输入正确的用户名密码并提交
- **THEN** 前端调用 `/api/login`，将返回的 Token 存入 localStorage，跳转到单词查询页

#### Scenario: Register flow
- **WHEN** 用户在注册页输入用户名密码并提交
- **THEN** 前端调用 `/api/register`，注册成功后提示用户前往登录

#### Scenario: Unauthenticated access
- **WHEN** 未登录用户访问受保护页面
- **THEN** 前端自动跳转到登录页

### Requirement: Word query page
前端 SHALL 提供单词查询页面，包含单词输入框和 AI 模型下拉选择器（DeepSeek / 通义千问）。查询结果展示释义和例句，并显示「保存到单词本」按钮。

#### Scenario: Query and display result
- **WHEN** 用户输入单词、选择 AI 模型后点击查询
- **THEN** 前端调用 `/api/words/query`，展示返回的释义和 3 条例句

#### Scenario: Save queried word
- **WHEN** 用户在查询结果页面点击「保存到单词本」按钮
- **THEN** 前端调用 `/api/words` 保存数据，按钮变为已保存状态

#### Scenario: Already saved word
- **WHEN** 查询的单词已存在于用户单词本中
- **THEN** 返回结果标记为已保存，不显示保存按钮（或显示「已保存」状态）

### Requirement: Word list page with pagination
前端 SHALL 提供单词本列表页面，分页展示用户保存的所有单词。MUST 包含分页器组件。

#### Scenario: View word list
- **WHEN** 用户访问单词本页面
- **THEN** 前端调用 `/api/words?page=1&page_size=10`，以卡片或列表形式展示单词及其释义、例句

#### Scenario: Pagination navigation
- **WHEN** 用户点击分页器切换页码
- **THEN** 前端请求对应页数据并更新展示

#### Scenario: Delete word from list
- **WHEN** 用户在列表中点击某个单词的删除按钮并确认
- **THEN** 前端调用 `DELETE /api/words/:id`，从列表中移除该单词

### Requirement: Request interceptor with JWT
前端 SHALL 在 HTTP 请求拦截器中自动附加 `Authorization: Bearer <token>` 请求头。Token 过期或缺失时自动跳转登录页。

#### Scenario: Auto-attach token
- **WHEN** 前端发起任何 API 请求
- **THEN** 请求拦截器自动从 localStorage 读取 Token 并附加到 Authorization 头

#### Scenario: Handle 401 response
- **WHEN** 后端返回 401 状态码
- **THEN** 前端清除本地 Token，跳转到登录页
