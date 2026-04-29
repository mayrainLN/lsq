## ADDED Requirements

### Requirement: DeepSeek AI provider
系统 SHALL 实现 DeepSeek AI 接口调用，发送包含单词的 prompt，要求返回结构化 JSON（释义 + 3 条例句）。API Key 通过环境变量配置。

#### Scenario: Successful DeepSeek query
- **WHEN** 系统以 ai_provider=deepseek 调用 DeepSeek API 查询一个有效英文单词
- **THEN** 返回包含中文释义和 3 条英文例句（含中文翻译）的结构化数据

#### Scenario: DeepSeek API failure
- **WHEN** DeepSeek API 调用超时或返回错误
- **THEN** 系统返回友好错误信息，HTTP 状态码 502 或 500

### Requirement: Tongyi Qianwen AI provider
系统 SHALL 实现通义千问 AI 接口调用，功能与 DeepSeek 一致，返回相同结构的数据。API Key 通过环境变量配置。

#### Scenario: Successful Tongyi query
- **WHEN** 系统以 ai_provider=tongyi 调用通义千问 API 查询一个有效英文单词
- **THEN** 返回包含中文释义和 3 条英文例句（含中文翻译）的结构化数据

#### Scenario: Tongyi API failure
- **WHEN** 通义千问 API 调用超时或返回错误
- **THEN** 系统返回友好错误信息，HTTP 状态码 502 或 500

### Requirement: Unified AI provider interface
系统 SHALL 定义统一的 AIProvider 接口，通过 ai_provider 参数字符串动态选择实现。不支持的 provider 名称 MUST 返回错误。

#### Scenario: Select provider by name
- **WHEN** 前端传入 ai_provider 为 "deepseek" 或 "tongyi"
- **THEN** 系统调用对应的 AI 实现

#### Scenario: Unknown provider
- **WHEN** 前端传入不支持的 ai_provider 值
- **THEN** 系统返回 400 错误，提示不支持的 AI 提供商
