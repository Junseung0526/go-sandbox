package middleware

import (
	"github.com/gin-gonic/gin"
)

// 함수 이름이 반드시 대문자 Auth여야 합니다.
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "secret-key" {
			c.JSON(401, gin.H{"error": "인증 실패!"})
			c.Abort()
			return
		}
		c.Next()
	}
}
