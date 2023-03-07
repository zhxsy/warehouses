package utils

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
)

// Printf 打印输出
func Printf(t *testing.T, err error, r interface{}) {
	if err != nil {
		t.Error("\n错误: ", err)
		return
	}
	if r != nil {
		jsIndent, err := json.MarshalIndent(r, "", "\t")
		if err != nil {
			fmt.Printf("%+v", r)
		}
		fmt.Println("\n结果: ", string(jsIndent))
	}
}

// 封装get 请求参数 /api?age=11&name=vic
func UrlPathWithParams(path string, params map[string]string) string {
	parseURL, _ := url.Parse(path)

	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}

	parseURL.RawQuery = values.Encode()
	return parseURL.String()
}
