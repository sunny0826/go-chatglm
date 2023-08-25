package chatglm

import (
	"fmt"
	"os"
	"time"
)

type ModelAPI struct {
	Model       string
	Prompt      []map[string]interface{}
	TopP        float64
	Temperature float64
}

const (
	BaseURL           = "https://open.bigmodel.cn/api/paas/v3/model-api"
	InvokeTypeSync    = "invoke"
	InvokeTypeAsync   = "async-invoke"
	InvokeTypeSSE     = "sse-invoke"
	ApiTimeoutSeconds = 300 * time.Second
	ChatGLMLite       = "chatglm_lite"
	ChatGLMStd        = "chatglm_std"
	ChatGLMPro        = "chatglm_pro"
)

func (m ModelAPI) Invoke(apiKey string) (map[string]interface{}, error) {
	token, err := generateToken(apiKey)
	if err != nil {
		return nil, err
	}
	return post(buildApiUrl(m.Model, InvokeTypeSync), token, m.buildParams(), ApiTimeoutSeconds)
}

func (m ModelAPI) AsyncInvoke(apiKey string) (map[string]interface{}, error) {
	token, err := generateToken(apiKey)
	if err != nil {
		return nil, err
	}
	return post(buildApiUrl(m.Model, InvokeTypeAsync), token, m.buildParams(), ApiTimeoutSeconds)
}

func (m ModelAPI) buildParams() map[string]interface{} {
	params := make(map[string]interface{})
	params["prompt"] = m.Prompt
	params["top_p"] = m.TopP
	params["temperature"] = m.Temperature
	return params
}

func buildApiUrl(module, invokeMethod string) string {
	var url string
	url = os.Getenv("ZHIPUAI_MODEL_API_URL")
	if url == "" {
		url = BaseURL
	}
	return fmt.Sprintf("%s/%s/%s", url, module, invokeMethod)
}
