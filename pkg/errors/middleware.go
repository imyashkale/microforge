package errors

import (
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/imyashkale/microforge/pkg/log"
)

func Handler(logger log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			l := logger.With(ctx.Request.Context())
			if err := recover(); err != nil {
				l.Errorf("recovered from panic (%v): %s", err, debug.Stack())
				if err != nil {
					rs := NewResponse(err)
					ctx.JSON(rs.StatusCode(), rs)
					ctx.Abort()
					err = nil
				}
			}
		}()
		ctx.Next()
	}
}
