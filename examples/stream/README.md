# Stream

```go
package main

import (
	"fmt"
	"os"

	"github.com/sunny0826/go-chatglm"
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
	sseClient, err := m.SSEInvoke(apiKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	eventChan := sseClient.ReadEvents()

	for event := range eventChan {
		switch event.Event {
		case "add":
			fmt.Println(event.Data)
		case "error", "interrupted":
			fmt.Println(event.Data)
		case "finish":
			fmt.Println(event.Data)
			for key, value := range event.Meta {
				fmt.Printf("%s: %s\n", key, value)
			}
		default:
			fmt.Println(event.Data)
		}
	}
}
```