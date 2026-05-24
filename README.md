# go-nanobot

用 Golang 参考 [nanobot](https://github.com/HKUDS/nanobot) 框架重新实现一个通用 AI Agent。通过这个过程掌握 Agent 开发的原理和方法。

## 项目目标

- 学习并实践 AI Agent 核心架构：消息总线、Agent 循环、工具执行、会话管理
- 用 Go 的并发模型（goroutine / channel）重构 nanobot 的异步消息流
- 逐步实现一个可运行的 Agent，支持 LLM 对话、工具调用、MCP 协议

## 参考架构

原版 nanobot (Python) 的核心数据流：

```
Channel → MessageBus → AgentLoop → AgentRunner → LLM Provider
                ↑                                      |
                └──────── OutboundMessage ←────────────┘
```

### 关键子系统

| 子系统 | 职责 |
|--------|------|
| **Channel** | 接收外部平台消息（Telegram、Discord、Slack 等），发布到消息总线 |
| **MessageBus** | 解耦 Channel 与 Agent 核心的异步消息队列 |
| **AgentLoop** | 管理会话 key、hooks、上下文构建，协调一次对话 turn |
| **AgentRunner** | 执行多轮 LLM 对话：发送消息、接收 tool call、执行工具、流式返回 |
| **Provider** | LLM 提供商适配层（OpenAI、Anthropic、Azure 等） |
| **Tool** | Agent 能力：文件操作、Shell 执行、Web 搜索、MCP 等 |
| **Session** | 会话历史持久化、上下文压缩 |
| **Config** | 配置管理 |

## 技术栈

- **语言**: Go 1.24+
- **并发模型**: goroutine + channel（对应 Python asyncio）
- **配置**: Viper / YAML
- **日志**: slog
- **测试**: testing + testify

## 项目结构（规划）

```
go-nanobot/
├── cmd/              # 入口
├── internal/
│   ├── agent/        # AgentLoop + AgentRunner
│   ├── bus/          # 消息总线
│   ├── channel/      # Channel 接口与实现
│   ├── provider/     # LLM Provider 适配
│   ├── tool/         # 工具注册与执行
│   ├── session/      # 会话管理
│   ├── config/       # 配置加载
│   └── mcp/          # MCP 协议支持
├── reference/        # nanobot 原版源码（仅供参考，不提交）
└── docs/             # 学习笔记与设计文档
```

## 开发计划

1. **Phase 1**: 项目骨架 — Config、Logger、基础消息结构
2. **Phase 2**: 消息总线 — Go channel 实现的异步消息队列
3. **Phase 3**: LLM Provider — OpenAI-compatible 接口
4. **Phase 4**: Agent Loop — 核心对话循环 + 工具调用
5. **Phase 5**: Tool 系统 — 文件操作、Shell 执行
6. **Phase 6**: Session 管理 — 历史持久化、上下文压缩
7. **Phase 7**: MCP 支持 — Model Context Protocol 客户端
8. **Phase 8**: Channel 集成 — 至少一个聊天平台

## License

MIT
