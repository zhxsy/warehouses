package wrapper

import (
	"context"
	"fmt"
	"github.com/rs/xid"
	"time"

	"google.golang.org/grpc/metadata"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

func CheckTraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		span := opentracing.SpanFromContext(c.Request.Context())
		if span != nil {
			if _, ok := span.Context().(jaeger.SpanContext); ok {
				return
			}
		}

		updateCtx(c)
	}
}

// 在ctx中添加trace_id
func updateCtx(c *gin.Context) {
	traceId := c.GetHeader("x-tracing-id")
	if traceId == "" {
		traceId = fmt.Sprintf("autoGen_%s_%d", xid.New().String(), time.Now().Unix())
	}

	newCtx := context.WithValue(c.Request.Context(), "trace_id", traceId)

	md := metadata.Pairs("trace_id", traceId)
	newCtx = metadata.NewOutgoingContext(newCtx, md)

	c.Request = c.Request.WithContext(newCtx)
}
