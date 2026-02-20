package middleware

import (
	"net/http"
	"os" // 🆕 추가
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtKey := []byte(os.Getenv("JWT_SECRET"))
		if len(jwtKey) == 0 {
			jwtKey = []byte("your_secret_key")
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "인증 헤더가 없습니다."})
			c.Abort()
			return
		}

		// ... (이하 동일)
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "인증 형식이 잘못되었습니다."})
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "유효하지 않은 토큰입니다."})
			c.Abort()
			return
		}
		c.Next()
	}
}
