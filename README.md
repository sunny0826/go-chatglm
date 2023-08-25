# Go ChatGLM

[![Go Reference](https://pkg.go.dev/badge/github.com/sunny0826/go-chatglm.svg)](https://pkg.go.dev/github.com/sunny0826/go-chatglm)
[![Go Report Card](https://goreportcard.com/badge/github.com/sunny0826/go-chatglm)](https://goreportcard.com/report/github.com/sunny0826/go-chatglm)

This library provides unofficial Go clients for [ChatGLM API](https://open.bigmodel.cn/dev/api). We support:

- [Get Token](https://open.bigmodel.cn/usercenter/apikeys)
- [ChatGLM](https://open.bigmodel.cn)

## Installation:

```
go get github.com/sunny0826/go-chatglm
```

## ChatGLM example usage:

```go
package main

import (
	"fmt"
	"github.com/sunny0826/go-chatglm"
	"os"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	m := chatglm.ModelAPI{
		Model:       chatglm.ChatGLMLite,
		TopP:        0.7,
		Temperature: 0.9,
		Prompt: []map[string]interface{}{
			{"role": "user", "content": "你好"},
			{"role": "assistant", "content": "我是人工智能助手"},
			{"role": "user", "content": "你叫什么名字"},
			{"role": "assistant", "content": "我叫chatGLM"},
			{"role": "user", "content": "你都可以做些什么事"},
		},
	}
	resp, err := m.Invoke(apiKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
```