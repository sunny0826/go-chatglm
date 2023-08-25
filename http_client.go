package chatglm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

var headers = map[string]string{
	"Accept":       "application/json",
	"Content-Type": "application/json; charset=UTF-8",
}

func post(apiURL, token string, params map[string]interface{}, timeout time.Duration) (map[string]interface{}, error) {
	client := resty.New()
	client.SetTimeout(timeout)
	resp, err := client.R().
		SetHeaders(headers).
		SetHeader("Authorization", token).
		SetBody(params).
		Post(apiURL)

	if err != nil {
		return nil, fmt.Errorf("请求异常：%w", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("响应异常：%s", resp.Body())
	}
	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("JSON 解析异常：%w", err)
	}

	return result, nil
}

//noinspection ALL
func stream(apiURL, token string, params map[string]interface{}, timeout time.Duration) (*resty.Response, error) {
	client := resty.New()
	client.SetTimeout(timeout)
	resp, err := client.R().
		SetHeaders(headers).
		SetHeader("Authorization", token).
		SetQueryParam("stream", "true").
		SetBody(params).
		Post(apiURL)

	if err != nil {
		return nil, fmt.Errorf("请求异常：%w", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("响应异常：%s", resp.Body())
	}

	return resp, nil
}

//noinspection ALL
func get(apiURL, token string, timeout time.Duration) (map[string]interface{}, error) {
	client := resty.New()
	client.SetTimeout(timeout)
	resp, err := client.R().
		SetHeaders(headers).
		SetHeader("Authorization", token).
		Get(apiURL)

	if err != nil {
		return nil, fmt.Errorf("请求异常：%w", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("响应异常：%s", resp.Body())
	}
	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("JSON 解析异常：%w", err)
	}

	return result, nil
}
