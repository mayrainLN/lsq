## ADDED Requirements

### Requirement: Smart word query
系统 SHALL 提供单词查询接口 `POST /api/words/query`，接收 word 和 ai_provider 参数。鉴权通过后先查数据库，已保存则直接返回；未保存则调用 AI 接口获取释义和 3 条例句后返回（不自动保存）。

#### Scenario: Query saved word
- **WHEN** 已登录用户查询一个已保存到单词本中的单词
- **THEN** 系统直接从数据库读取该单词的释义和例句，标记 `saved: true` 返回

#### Scenario: Query new word via AI
- **WHEN** 已登录用户查询一个未保存的单词，指定 ai_provider
- **THEN** 系统调用对应 AI 接口，返回释义和 3 条例句，标记 `saved: false`，不写入数据库

#### Scenario: Empty word input
- **WHEN** 用户提交空的 word 参数
- **THEN** 系统返回参数校验错误

### Requirement: Save word manually
系统 SHALL 提供单词保存接口 `POST /api/words`，接收完整的单词数据（word、definition、sentences、ai_provider），将其与当前用户绑定后写入数据库。

#### Scenario: Save new word
- **WHEN** 用户提交一个未保存过的单词完整数据
- **THEN** 系统在 words 表创建记录并在 sentences 表插入对应例句，绑定当前 user_id

#### Scenario: Save duplicate word
- **WHEN** 用户尝试保存一个已存在于自己单词本中的单词
- **THEN** 系统返回提示「该单词已保存」，不重复创建记录

### Requirement: Word list with pagination
系统 SHALL 提供分页查询接口 `GET /api/words`，接收 page 和 page_size 参数，返回当前用户的单词列表及分页信息。

#### Scenario: Get first page
- **WHEN** 用户请求第 1 页，page_size=10
- **THEN** 系统返回最多 10 条单词记录（含例句），以及 total 总数、当前页码等分页元数据

#### Scenario: Empty word list
- **WHEN** 用户尚未保存任何单词
- **THEN** 系统返回空列表，total=0

### Requirement: Delete word (soft delete)
系统 SHALL 提供单词删除接口 `DELETE /api/words/:id`，根据 word ID 对记录进行软删除。只能删除当前用户自己的单词。

#### Scenario: Delete own word
- **WHEN** 用户删除自己保存的某个单词
- **THEN** 系统对该 word 记录及其关联 sentences 进行软删除（设置 deleted_at），API 返回成功

#### Scenario: Delete non-existent or others' word
- **WHEN** 用户尝试删除不存在的单词或其他用户的单词
- **THEN** 系统返回 404 或权限错误
