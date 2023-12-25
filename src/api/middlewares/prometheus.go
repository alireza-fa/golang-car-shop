package middlewares

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Prometheus() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.FullPath()
		method := c.Request.Method
		c.Next()
		status := c.Writer.Status()
	}
}
