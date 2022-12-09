package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 加一个打印请求的中间件，但是gin.Default()里面包含了日志打印
func NetLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		now := time.Now()
		log.Printf("[%s] %s %s %s", now.Format(time.RFC3339), c.Request.Method, c.Request.URL.Path, clientIP)
		c.Next()
	}
}
