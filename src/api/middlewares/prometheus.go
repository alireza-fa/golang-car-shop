package middlewares

import (
	"github.com/alireza-fa/golang-car-shop/pkg/metrics"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Prometheus() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.FullPath()
		method := c.Request.Method
		c.Next()
		status := c.Writer.Status()
		metrics.HttpDuration.WithLabelValues(path, method, strconv.Itoa(status)).
			Observe(float64(time.Since(start) / time.Millisecond))
	}
}
