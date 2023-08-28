# Go ChatGLM

[![Go Reference](https://pkg.go.dev/badge/github.com/sunny0826/go-chatglm.svg)](https://pkg.go.dev/github.com/sunny0826/go-chatglm)
[![Go Report Card](https://goreportcard.com/badge/github.com/sunny0826/go-chatglm)](https://goreportcard.com/report/github.com/sunny0826/go-chatglm)
[![GitHub release (with filter)](https://img.shields.io/github/v/release/sunny0826/go-chatglm?logo=go&link=https%3A%2F%2Fgithub.com%2Fsunny0826%2Fgo-chatglm%2Freleases)](https://github.com/sunny0826/go-chatglm/releases)

> **Synchronization with [official Python SDK](https://pypi.org/project/zhipuai/) feature & version numbers.**

This library provides unofficial Go clients for [ChatGLM API](https://open.bigmodel.cn/dev/api). We support:

- [Get Token](https://open.bigmodel.cn/usercenter/apikeys)
- [ChatGLM](https://open.bigmodel.cn)

## Installation:

```
go get github.com/sunny0826/go-chatglm
```

## Select Model

- ChatGLMLite
- ChatGLMStd
- ChatGLMPro

```go
m := chatglm.ModelAPI{
    Model:       chatglm.ChatGLMLite,
    // Model:       chatglm.ChatGLMStd,
    // Model:       chatglm.ChatGLMPro,
}
```

## Examples

- [sync](https://github.com/sunny0826/go-chatglm/tree/main/examples/sync)
- [async](https://github.com/sunny0826/go-chatglm/tree/main/examples/async)
- [stream](https://github.com/sunny0826/go-chatglm/tree/main/examples/stream)

## license

[Apache License 2.0](./LICENSE)