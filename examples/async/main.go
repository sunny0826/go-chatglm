package main

import (
	"fmt"
	"os"

	"github.com/sunny0826/go-chatglm"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	//m := chatglm.ModelAPI{
	//	Model:       chatglm.ChatGLMLite,
	//	TopP:        0.7,
	//	Temperature: 1,
	//	Prompt: []map[string]interface{}{
	//		{"role": "user", "content": "你好"},
	//		{"role": "assistant", "content": "我是人工智能助手"},
	//		{"role": "user", "content": "你叫什么名字"},
	//		{"role": "assistant", "content": "我叫chatGLM"},
	//		{"role": "user", "content": "你都可以做些什么事"},
	//	},
	//}
	//resp, err := m.AsyncInvoke(apiKey)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(resp)

	taskID := "75753966931688458417874899228181488752"
	m := chatglm.ModelAPI{Model: chatglm.ChatGLMLite}
	resp, err := m.QueryAsyncInvokeResult(apiKey, taskID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
