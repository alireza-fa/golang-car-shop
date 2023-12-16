package middlewares

import (
	"bytes"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/pkg/logging"
	"github.com/gin-gonic/gin"
	"io"
	"strings"
	"time"
)

type BodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w BodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func DefaultStructuredLogger(cfg *config.Config) gin.HandlerFunc {
	logger := logging.NewLogger(cfg)
	return structuredLogger(logger)
}

func structuredLogger(logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.FullPath(), "swagger") {
			c.Next()
		} else {
			blw := &BodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
			start := time.Now() // start
			path := c.FullPath()
			raw := c.Request.URL.RawQuery

			bodyBytes, _ := io.ReadAll(c.Request.Body)
			c.Request.Body.Close()
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			c.Writer = blw
			c.Next()

			params := gin.LogFormatterParams{}
			params.TimeStamp = time.Now()
			params.Latency = params.TimeStamp.Sub(start)
			params.ClientIP = c.ClientIP()
			params.Method = c.Request.Method
			params.StatusCode = c.Writer.Status()
			params.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
			params.BodySize = c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}
			params.Path = path

			keys := map[logging.ExtraKey]interface{}{}
			keys[logging.Path] = params.Path
			keys[logging.ClientIp] = params.ClientIP
			keys[logging.Method] = params.Method
			keys[logging.Latency] = params.Latency
			keys[logging.StatusCode] = params.StatusCode
			keys[logging.ErrorMessage] = params.ErrorMessage
			keys[logging.BodySize] = params.BodySize
			keys[logging.RequestBody] = string(bodyBytes)
			keys[logging.ResponseBody] = blw.body.String()

			logger.Info(logging.RequestResponse, logging.Api, "", keys)
		}
	}
}
