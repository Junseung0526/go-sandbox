package middleware

import (
	"github.com/gin-gonic/gin"
)

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
