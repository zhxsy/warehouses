package output

import (
	"context"
	"github.com/cfx/warehouses/app"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

type ErrorLog struct {
	Msg string
}

func NewErrLog(msg string) *ErrorLog {
	err := &ErrorLog{
		Msg: msg,
	}
	return err
}

func (e *ErrorLog) Error() string {
	return e.Msg
}

// 错误记录
func (e *ErrorLog) Log(err error) *ErrorLog {
	app.Log().WithField("msg", e.Msg).Warn(err)
	return e
}

// 日志扩展数据
func LogCtx(ctx context.Context) (l *logrus.Entry) {
	le, ok := ctx.(*gin.Context)
	if ok {
		l = app.Log().WithField("x-tracing-id", le.Request.Context().Value("trace_id"))
	} else {
		l = app.Log().WithField("x-tracing-id", ctx.Value("trace_id"))
	}
	return
}

// 脚本，新增context 记录
func NewContext() context.Context {
	traceId := xid.New().String()
	newCtx := context.WithValue(context.Background(), "trace_id", traceId)

	md := metadata.Pairs("trace_id", traceId)
	newCtx = metadata.NewOutgoingContext(newCtx, md)
	return newCtx
}
