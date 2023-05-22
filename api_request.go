package chatglm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	BaseURL = "https://maas.aminer.cn/api/paas"

	EnginesPath   = "/model/v1/open/engines/"
	EnginesPathV2 = "/model/v2/open/engines/"
	// #nosec G101
	TokenPath        = "/passApiToken/createApiToken"
	QueryOrderResult = "/request-task/query-request-task-result"
)

type Params map[string]any

type EngineRequest struct {
	TopP          float64  `json:"top_p"`
	Temperature   float64  `json:"temperature"`
	Prompt        string   `json:"prompt"`
	RequestTaskNo string   `json:"requestTaskNo"`
	History       []string `json:"history"`
}

type EngineResponse struct {
	Code    int          `json:"code"`
	Msg     string       `json:"msg"`
	Data    ResponseData `json:"data"`
	Success bool         `json:"success"`
}

type ResponseData struct {
	Prompt         string `json:"prompt"`
	OutputText     string `json:"outputText"`
	InputTokenNum  *int   `json:"inputTokenNum"`
	OutputTokenNum *int   `json:"outputTokenNum"`
	TotalTokenNum  int    `json:"totalTokenNum"`
	RequestTaskNo  string `json:"requestTaskNo"`
	TaskOrderNo    string `json:"taskOrderNo"`
	TaskStatus     string `json:"taskStatus"`
}

type TokenResponse struct {
	Code    string `json:"code"`
	Data    string `json:"data"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}

func GetToken(apiKey, publicKey string) (string, error) {
	content := time.Now().UnixNano() / int64(time.Millisecond)
	crypto, err := rsaEncode([]byte(strconv.FormatInt(content, 10)), publicKey)
	if err != nil {
		return "", err
	}
	params := Params{
		"apiKey":    apiKey,
		"encrypted": crypto,
	}
	resp, err := sendPost(BaseURL+TokenPath, params)
	if err != nil {
		return "", err
	}
	var data TokenResponse
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return "", err
	}
	if data.Success {
		return data.Data, nil
	} else {
		return "", fmt.Errorf(data.Msg)
	}
}

func Chat(abilityType, engineType, authToken string, params Params) (map[string]any, error) {
	reqEngineAPIURL := fmt.Sprintf("%s/model/v1/open/%s/%s", BaseURL, abilityType, engineType)
	resp, err := sendPost(reqEngineAPIURL, params, authToken)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(resp, &data)
	return data, err
}

func sendPost(url string, params Params, authToken ...string) ([]byte, error) {
	reqBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if len(authToken) > 0 {
		req.Header.Set("Authorization", authToken[0])
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func ExecuteEngine(abilityType, engineType, authToken string, params EngineRequest,
	timeout ...time.Duration) (EngineResponse, error) {
	reqEngineAPIURL := fmt.Sprintf("%s%s%s/%s", BaseURL, EnginesPath, abilityType, engineType)
	resp, err := sendPostWithTimeout(reqEngineAPIURL, params, authToken, timeout...)
	if err != nil {
		return EngineResponse{}, err
	}
	var data EngineResponse
	err = json.Unmarshal(resp, &data)
	return data, err
}

func sendPostWithTimeout(url string, params EngineRequest, authToken string, timeout ...time.Duration) ([]byte, error) {
	reqBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)

	client := &http.Client{}
	if len(timeout) > 0 {
		client.Timeout = timeout[0]
	} else {
		client.Timeout = time.Hour
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
