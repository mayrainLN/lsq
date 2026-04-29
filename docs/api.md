# API 接口文档

## 基础信息

- **Base URL**: `/api`
- **认证方式**: JWT Bearer Token（在 Header 中携带 `Authorization: Bearer <token>`）
- **请求格式**: JSON（`Content-Type: application/json`）

---

## 1. 用户注册

**POST** `/api/register`

**鉴权**: 无需

**请求参数 (Body)**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| username | string | 是 | 用户名，2-64 位 |
| password | string | 是 | 密码，至少 6 位 |

**成功返回** (200):

```json
{
  "message": "注册成功"
}
```

**错误码**:

| 状态码 | error | 说明 |
|--------|-------|------|
| 400 | 参数校验失败: ... | 缺少字段或格式错误 |
| 409 | 用户名已存在 | 用户名重复 |
| 500 | 注册失败 | 服务器内部错误 |

---

## 2. 用户登录

**POST** `/api/login`

**鉴权**: 无需

**请求参数 (Body)**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| username | string | 是 | 用户名 |
| password | string | 是 | 密码 |

**成功返回** (200):

```json
{
  "message": "登录成功",
  "token": "eyJhbGciOiJIUzI1NiIs..."
}
```

**错误码**:

| 状态码 | error | 说明 |
|--------|-------|------|
| 400 | 参数校验失败: ... | 缺少字段 |
| 401 | 用户名或密码错误 | 认证失败 |

---

## 3. 智能查询单词

**POST** `/api/words/query`

**鉴权**: 需要 JWT Token

**请求参数 (Body)**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| word | string | 是 | 要查询的英文单词 |
| ai_provider | string | 是 | AI 提供商（`deepseek` 或 `tongyi`） |

**成功返回** (200) — 已保存的单词:

```json
{
  "saved": true,
  "word": "hello",
  "definition": "int. 你好；喂（用于问候或引起注意）",
  "ai_provider": "deepseek",
  "sentences": [
    { "english": "Hello, how are you?", "chinese": "你好，你怎么样？" },
    { "english": "She said hello to everyone.", "chinese": "她向每个人打了招呼。" },
    { "english": "Hello, is anyone there?", "chinese": "喂，有人在吗？" }
  ]
}
```

**成功返回** (200) — 新查询（AI 生成）:

```json
{
  "saved": false,
  "word": "serendipity",
  "definition": "n. 意外发现美好事物的能力；机缘巧合",
  "ai_provider": "tongyi",
  "sentences": [...]
}
```

**错误码**:

| 状态码 | error | 说明 |
|--------|-------|------|
| 400 | 参数校验失败: ... | 缺少 word 或 ai_provider |
| 400 | 不支持的 AI 提供商: ... | ai_provider 值无效 |
| 401 | 未提供认证令牌 | 未登录 |
| 502 | AI 查询失败: ... | AI 接口调用出错 |

---

## 4. 保存单词

**POST** `/api/words`

**鉴权**: 需要 JWT Token

**请求参数 (Body)**:

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| word | string | 是 | 英文单词 |
| definition | string | 是 | 中文释义 |
| ai_provider | string | 是 | AI 来源 |
| sentences | array | 是 | 例句数组 |
| sentences[].english | string | 是 | 英文例句 |
| sentences[].chinese | string | 是 | 中文翻译 |

**成功返回** (200):

```json
{
  "message": "保存成功",
  "id": 1
}
```

**错误码**:

| 状态码 | error | 说明 |
|--------|-------|------|
| 400 | 参数校验失败: ... | 缺少必填字段 |
| 401 | 未提供认证令牌 | 未登录 |
| 409 | 该单词已保存 | 重复保存 |
| 500 | 保存失败 | 服务器错误 |

---

## 5. 获取单词列表（分页）

**GET** `/api/words`

**鉴权**: 需要 JWT Token

**请求参数 (Query)**:

| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| page | int | 否 | 1 | 页码 |
| page_size | int | 否 | 10 | 每页条数（最大 100） |

**成功返回** (200):

```json
{
  "total": 25,
  "page": 1,
  "page_size": 10,
  "words": [
    {
      "ID": 1,
      "word": "hello",
      "definition": "int. 你好",
      "ai_provider": "deepseek",
      "sentences": [
        { "english": "Hello, world!", "chinese": "你好，世界！" }
      ],
      "CreatedAt": "2026-04-29T10:00:00Z"
    }
  ]
}
```

**错误码**:

| 状态码 | error | 说明 |
|--------|-------|------|
| 401 | 未提供认证令牌 | 未登录 |

---

## 6. 删除单词

**DELETE** `/api/words/:id`

**鉴权**: 需要 JWT Token

**路径参数**:

| 字段 | 类型 | 说明 |
|------|------|------|
| id | int | 单词记录 ID |

**成功返回** (200):

```json
{
  "message": "删除成功"
}
```

**错误码**:

| 状态码 | error | 说明 |
|--------|-------|------|
| 400 | 无效的单词 ID | ID 格式错误 |
| 401 | 未提供认证令牌 | 未登录 |
| 404 | 单词不存在或无权操作 | 找不到或非本人的单词 |
