# 数据库设计文档

## 概述

本项目使用 MySQL 8.0，共设计 3 张业务表：`users`、`words`、`sentences`。采用外键约束保证数据一致性，支持 GORM 软删除。

## ER 关系

```
users (1) ──< (N) words (1) ──< (N) sentences
```

- 一个用户拥有多个单词记录
- 一个单词记录拥有多条例句

## 表结构

### users 表 — 用户信息

| 字段 | 类型 | 约束 | 说明 |
|------|------|------|------|
| id | BIGINT UNSIGNED | PRIMARY KEY, AUTO_INCREMENT | 用户 ID |
| username | VARCHAR(64) | NOT NULL, UNIQUE INDEX | 用户名，唯一 |
| password | VARCHAR(255) | NOT NULL | bcrypt 哈希后的密码 |
| created_at | DATETIME(3) | — | 创建时间 |
| updated_at | DATETIME(3) | — | 更新时间 |
| deleted_at | DATETIME(3) | INDEX | 软删除标记 |

**索引：**
- `idx_username` — UNIQUE，用户名唯一约束
- `idx_deleted_at` — 软删除查询优化

### words 表 — 单词记录

| 字段 | 类型 | 约束 | 说明 |
|------|------|------|------|
| id | BIGINT UNSIGNED | PRIMARY KEY, AUTO_INCREMENT | 单词记录 ID |
| user_id | BIGINT UNSIGNED | NOT NULL, FOREIGN KEY → users(id) | 所属用户 ID |
| word | VARCHAR(128) | NOT NULL | 英文单词 |
| definition | TEXT | NOT NULL | AI 生成的中文释义 |
| ai_provider | VARCHAR(32) | NOT NULL | AI 提供商（deepseek/tongyi） |
| created_at | DATETIME(3) | — | 创建时间 |
| updated_at | DATETIME(3) | — | 更新时间 |
| deleted_at | DATETIME(3) | INDEX | 软删除标记 |

**索引：**
- `idx_user_id` — 按用户查询单词列表
- `idx_user_word` — UNIQUE，同一用户下单词唯一（防重复保存）
- `idx_deleted_at` — 软删除查询优化

**外键：**
- `fk_words_user` → `users(id)` ON DELETE CASCADE

### sentences 表 — 例句

| 字段 | 类型 | 约束 | 说明 |
|------|------|------|------|
| id | BIGINT UNSIGNED | PRIMARY KEY, AUTO_INCREMENT | 例句 ID |
| word_id | BIGINT UNSIGNED | NOT NULL, FOREIGN KEY → words(id) | 所属单词 ID |
| english | TEXT | NOT NULL | 英文例句 |
| chinese | TEXT | NOT NULL | 中文翻译 |
| created_at | DATETIME(3) | — | 创建时间 |
| updated_at | DATETIME(3) | — | 更新时间 |
| deleted_at | DATETIME(3) | INDEX | 软删除标记 |

**索引：**
- `idx_word_id` — 按单词查询例句
- `idx_deleted_at` — 软删除查询优化

**外键：**
- `fk_sentences_word` → `words(id)` ON DELETE CASCADE

## 设计说明

1. **软删除**：所有表都包含 `deleted_at` 字段，GORM 执行 DELETE 时自动设置该字段而非物理删除。
2. **例句独立成表**：每个单词固定 3 条例句，采用一对多关系而非 JSON 存储，便于后续扩展和单独查询。
3. **联合唯一索引**：`words` 表的 `(user_id, word)` 联合唯一索引确保同一用户不会重复保存同一单词。
4. **级联删除**：外键设置 `ON DELETE CASCADE`，删除用户时自动清理其单词和例句数据。
