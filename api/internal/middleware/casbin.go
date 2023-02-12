package middleware

import (
	"github.com/gin-gonic/gin"
	"io.github.tuanictu97/api/internal/x/utilx"
)

type CasbinConfig struct {
	SkippedPathPrefixes []string
	Skipper             func(c *gin.Context) bool
	GetEnforcer         func(c *gin.Context) *casbin.Enforcer
	GetRoleIDs          func(c *gin.Context) []string
}

func CasbinWithConfig(config CasbinConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkippedPathPrefixes(c, config.SkippedPathPrefixes...) ||
			(config.Skipper != nil && config.Skipper(c)) {
			c.Next()
			return
		}

		roleIDs := config.GetRoleIDs(c)
		enforcer := config.GetEnforcer(c)
		for _, roleID := range roleIDs {
			if b, err := enforcer.Enforce(roleID, c.Request.URL.Path, c.Request.Method); err != nil {
				utilx.ResError(c, err)
				return
			} else if b {
				c.Next()
				return
			}
		}

		utilx.ResError(c, utilx.ErrPermissionDenied)
	}
}
