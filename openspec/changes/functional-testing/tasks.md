## 1. 环境准备

- [x] 1.1 确认 .env 中 DeepSeek API Key 已配置为真实值
- [x] 1.2 执行 docker-compose up -d --build 启动全部服务
- [x] 1.3 检查三个容器状态均为 Running，等待 MySQL 健康检查通过

## 2. API 功能测试

- [x] 2.1 测试用户注册：POST /api/register，验证成功返回和重复注册拦截
- [x] 2.2 测试用户登录：POST /api/login，验证返回 JWT Token
- [x] 2.3 测试 DeepSeek 查词：POST /api/words/query，验证 AI 返回释义和 3 条例句（结果：代码正确，Key 余额不足 HTTP 402）
- [x] 2.4 测试保存单词：POST /api/words，验证数据写入成功（修复了 GORM AIProvider 列名映射 bug）
- [x] 2.5 测试单词列表：GET /api/words，验证分页返回和数据完整性
- [x] 2.6 测试删除单词：DELETE /api/words/:id，验证软删除生效

## 3. 前端浏览器验证

- [x] 3.1 访问 http://localhost，验证前端页面加载和 SPA 路由正常
- [x] 3.2 通过浏览器完成注册→登录→查词→保存→查看单词本→删除全流程（注册→登录→查词页均正常；AI 查词因 Key 余额不足无法验证）

## 4. 问题修复

- [x] 4.1 修复测试过程中发现的所有阻塞性 bug（共 3 个：docker-compose 版本、Node 版本、GORM 列名映射）
