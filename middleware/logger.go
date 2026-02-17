package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 1. 요청 정보 출력
		fmt.Printf("[REQUEST] %s %s\n", c.Request.Method, c.Request.URL.Path)

		c.Next() // 다음 핸들러 실행 (핵심!)

		// 2. 소요 시간 계산 및 상태 코드 출력
		latency := time.Since(t)
		status := c.Writer.Status()
		fmt.Printf("[RESPONSE] 상태코드: %d | 소요시간: %v\n", status, latency)
	}
}
