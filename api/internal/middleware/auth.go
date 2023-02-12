package middleware

import (
	"github.com/gin-gonic/gin"
	"io.github.tuanictu97/api/internal/x/contextx"
	"io.github.tuanictu97/api/internal/x/utilx"
	"io.github.tuanictu97/api/pkg/logger"
)

type AuthConfig struct {
	Disable             bool
	SkippedPathPrefixes []string
	Skipper             func(c *gin.Context) bool
	DefaultUserID       func(c *gin.Context) string
	ParseUserID         func(c *gin.Context) (string, error)
}

func AuthWithConfig(config AuthConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkippedPathPrefixes(c, config.SkippedPathPrefixes...) ||
			(config.Skipper != nil && config.Skipper(c)) {
			c.Next()
			return
		}

		var userID string
		if config.Disable {
			userID = config.DefaultUserID(c)
		} else {
			if v, err := config.ParseUserID(c); err != nil {
				utilx.ResError(c, err)
				return
			} else {
				userID = v
			}
		}

		ctx := contextx.NewUserID(c.Request.Context(), userID)
		ctx = logger.NewUserID(ctx, userID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
