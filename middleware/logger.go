package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		fmt.Printf("[REQUEST] %s %s\n", c.Request.Method, c.Request.URL.Path)

		c.Next()

		latency := time.Since(t)
		status := c.Writer.Status()
		fmt.Printf("[RESPONSE] 상태코드: %d | 소요시간: %v\n", status, latency)
	}
}
