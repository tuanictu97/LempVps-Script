package middleware

import (
	"fmt"

	"io.github.tuanictu97/api/internal/x/contextx"
	"io.github.tuanictu97/api/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type TraceConfig struct {
	SkippedPathPrefixes []string
	RequestHeaderKey    string
	ResponseTraceKey    string
}

var DefaultTraceConfig = TraceConfig{
	RequestHeaderKey: "X-Request-Id",
	ResponseTraceKey: "X-Trace-Id",
}

func Trace() gin.HandlerFunc {
	return TraceWithConfig(DefaultTraceConfig)
}

func TraceWithConfig(config TraceConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkippedPathPrefixes(c, config.SkippedPathPrefixes...) {
			c.Next()
			return
		}

		traceID := c.GetHeader(config.RequestHeaderKey)
		if traceID == "" {
			traceID = fmt.Sprintf("trace-%s", xid.New().String())
		}

		ctx := contextx.NewTraceID(c.Request.Context(), traceID)
		ctx = logger.NewTraceID(ctx, traceID)
		c.Request = c.Request.WithContext(ctx)
		c.Writer.Header().Set(config.ResponseTraceKey, traceID)
		c.Next()
	}
}
