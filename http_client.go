package chatglm

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

var headers = map[string]string{
	"Accept":       "application/json",
	"Content-Type": "application/json; charset=UTF-8",
}

func post(api_url, token string, params map[string]interface{}, timeout time.Duration) (map[string]interface{}, error) {
	client := resty.New()
	client.SetTimeout(timeout)
	resp, err := client.R().
		SetHeaders(headers).
		SetHeader("Authorization", token).
		SetBody(params).
		Post(api_url)

	if err != nil {
		return nil, fmt.Errorf("请求异常：%v", err)
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("响应异常：%s", resp.Body())
	}
	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("JSON 解析异常：%v", err)
	}

	return result, nil

}

func stream(api_url, token string, params map[string]interface{}, timeout time.Duration) (*resty.Response, error) {
	client := resty.New()
	client.SetTimeout(timeout)
	resp, err := client.R().
		SetHeaders(headers).
		SetHeader("Authorization", token).
		SetQueryParam("stream", "true").
		SetBody(params).
		Post(api_url)

	if err != nil {
		return nil, fmt.Errorf("请求异常：%v", err)
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("响应异常：%s", resp.Body())
	}

	return resp, nil
}

func get(api_url, token string, timeout time.Duration) (map[string]interface{}, error) {
	client := resty.New()
	client.SetTimeout(timeout)
	resp, err := client.R().
		SetHeaders(headers).
		SetHeader("Authorization", token).
		Get(api_url)

	if err != nil {
		return nil, fmt.Errorf("请求异常：%v", err)
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("响应异常：%s", resp.Body())
	}
	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("JSON 解析异常：%v", err)
	}

	return result, nil
}
