package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cfx/warehouses/library/third_api"
	"github.com/cfx/warehouses/output"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"

	"testing"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func NewEngine(f func() *gin.Engine) {
	engine = f()
}

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
func PrintfJson(data interface{}) {
	jsIndent, err := json.Marshal(data)
	if err != nil {
		fmt.Println(string(jsIndent))
	} else {
		fmt.Errorf("json 解析失败：%v", err)
	}
}

type Result output.Response

// Printf 格式化输出
func (r *Result) Printf() {
	jsIndent, err := json.MarshalIndent(r, "", "\t")
	if err != nil {
		fmt.Printf("%+v", r)
	}
	fmt.Println("\n结果: ", string(jsIndent))
}

// ParseResult 解析最后结果
//  @Description:
//  @params body
//  @Author vic
//  @return Result
func ParseResult(body []byte) Result {
	var m = Result{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		log.Fatal("json 转换失败: " + err.Error())
	}
	return m
}

// GetHttp get请求
//  @Description:
//  @params uri
//  @Author vic
//  @return Result
func GetHttp(uri string, options ...third_api.Option) Result {
	req := httptest.NewRequest("GET", uri, nil)
	for _, o := range options {
		o.ApplyOption(req)
	}
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	result := w.Result()
	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatal("Get读取结果失败：" + err.Error())
	}

	return ParseResult(body)
}

// PostForm post form表单 请求
//  @Description:
//  @params uri
//  @params param
//  @Author vic
//  @return Result
func PostForm(uri string, param map[string]string) Result {

	p := url.Values{}
	for k, v := range param {
		p.Set(k, v)
	}
	var buf io.Reader
	buf = strings.NewReader(p.Encode())

	req := httptest.NewRequest("POST", uri, buf)
	w := httptest.NewRecorder() // 初始化响应
	engine.ServeHTTP(w, req)

	result := w.Result()
	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatal("Get读取结果失败：" + err.Error())
	}

	return ParseResult(body)
}

//
// post json 请求
//  @Description:
//  @params uri
//  @params param isStg: false develop , true 星巴克stg
//  @Author vic
//  @return Result
func PostJson(uri string, param map[string]interface{}, options ...third_api.Option) Result {
	body := make([]byte, 0)
	p, err := json.Marshal(param)
	if err != nil {
		log.Fatal("map 转换 json 字符串失败" + err.Error())
	}

	var req *http.Request
	req = httptest.NewRequest("POST", uri, bytes.NewReader(p))
	req.Header.Add("Content-Type", "application/json")
	for _, o := range options {
		o.ApplyOption(req)
	}

	w := httptest.NewRecorder() // 初始化响应
	engine.ServeHTTP(w, req)

	result := w.Result()
	defer result.Body.Close()

	body, err = ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatal("Get读取结果失败：" + err.Error())
	}

	return ParseResult(body)
}

func Dump(r interface{}) {
	jsIndent, err := json.MarshalIndent(r, "", "\t")
	if err != nil {
		fmt.Printf("%+v", r)
	}
	fmt.Println("\n结果: ", string(jsIndent))
}

func GetHttpLogin(uri string) Result {
	//token := login()
	token := ""
	if token == "" {
		return Result{
			Message: "token 失败",
		}
	}
	head := &third_api.HeaderOption{
		Key: "token",
		Val: token,
	}

	return GetHttp(uri, head)
}
func PostJsonLogin(uri string, param map[string]interface{}, isStg bool) Result {
	//token := login()
	token := ""
	if token == "" {
		return Result{
			Message: "token 失败",
		}
	}
	head := &third_api.HeaderOption{
		Key: "token",
		Val: token,
	}

	return PostJson(uri, param, head)
}

func login(uri string, data map[string]interface{}) string {
	body := PostJson(
		uri,
		data,
	)
	if body.Code == 0 {
		return getToken(body.Data)
	}
	return ""
}

func getToken(data interface{}) string {
	v := reflect.ValueOf(data)
	mm := v.Interface().(map[string]interface{})
	return mm["Token"].(string)
}
