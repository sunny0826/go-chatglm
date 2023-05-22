# Go ChatGLM

This library provides unofficial Go clients for [ChatGLM API](https://open.bigmodel.ai/howuse/describesummary). We support:

- [Get Token](https://open.bigmodel.ai/howuse/authentication)
- [ChatGLM](https://open.bigmodel.ai/howuse/chatglm_6b)

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
	// 注意这里仅为了简化编码每一次请求都去获取token， 线上环境token有过期时间， 客户端可自行缓存，过期后重新获取。
	token, err := chatglm.GetToken(os.Getenv("CHATGLM_API_KEY"), os.Getenv("CHATGLM_PUBLIC_KEY"))
	if err != nil {
		fmt.Println("获取token失败，请检查 API_KEY 和 PUBLIC_KEY")
		return
	}
	// 能力类型
	abilityType := "chatGLM"
	// 引擎类型
	engineType := "chatGLM"

	// 请求参数
	data := chatglm.EngineRequest{
		TopP:        0.7,                                    //TopP采样又称核采样（Nucleus Sampling）。topP值会定义候选集在概率分布中的概率密度。若topP值为0.6，则解码器输出的概率密度的前60%会作为候选集。我们建议您不要与topK和temperature同时使用或同时进行调节。取值范围0～1.0，当topP为0时，该参数不起作用
		Temperature: 0.9,                                    //温度系数，取值>0.0，默认为1.0。更大的温度系数表示模型生成的多样性更强。取值范围0.5～1.0
		Prompt:      "摸鱼先生在家里蹲大学担任什么职务？",                    //输入内容，范围1～2048
		History:     []string{"家里蹲大学校长是谁?", "家里蹲大学校长是摸鱼先生"}, // 会话历史,只支持偶数，Q A Q A 的形式传进去
	}
	engineResponse, err := chatglm.ExecuteEngine(abilityType, engineType, token, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(engineResponse)
}

```