package output

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	ShowErr   bool        `json:"show_err"`
	Timestamp int64       `json:"timestamp"`
	Version   string      `json:"version"`
	Data      interface{} `json:"data"`
}

func NewResp(code int, msg string, showErr bool) *Response {
	return &Response{
		Code:      code,
		Message:   msg,
		ShowErr:   showErr,
		Timestamp: time.Now().Unix(),
		Version:   "1",
		Data:      gin.H{},
	}
}
func Result(data *Response, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, data)
}

func Ok(c *gin.Context, data interface{}) {
	resp := NewResp(0, "success", false)
	resp.Data = data
	Result(resp, c)
}

func Fail(c *gin.Context, err error) {
	var data *Response
	if e, ok := err.(*Error); ok {
		data = NewResp(e.Code, e.Msg, e.ShowErr)
	} else {
		data = NewResp(-5000, "server error", false)
	}
	Result(data, c)
}
