package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.RouterGroup) {
	r.GET("/health",healthHandler())
}

func healthHandler() func(c *gin.Context) {
		return func(c *gin.Context) {
			c.String(http.StatusOK, "Ok")
		}
	}