## Why

ai-smart-wordbook 的代码开发已完成，但未经过任何实际验证。需要启动服务并进行端到端功能测试，确认注册→登录→查词→保存→列表→删除全流程可跑通，特别是 DeepSeek AI 接口对接是否正常工作。

## What Changes

- 配置真实的 DeepSeek API Key 到 `.env`
- 启动 docker-compose 全套服务（db + backend + frontend）
- 通过 API 请求逐一验证 6 个接口的功能正确性
- 通过浏览器验证前端页面渲染和交互流程
- 发现并修复测试过程中暴露的 bug

## Capabilities

### New Capabilities
- `smoke-test`: 端到端冒烟测试，验证全链路可用性

### Modified Capabilities

（无）

## Impact

- **配置变更**: `.env` 中 DeepSeek API Key 更新为真实值
- **Bug 修复**: 测试过程中发现的问题将直接修复
