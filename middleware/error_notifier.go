package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ErrorNotifier() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 || c.Writer.Status() >= 400 {
			sendToDiscord(c)
		}
	}
}

func sendToDiscord(c *gin.Context) {
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if webhookURL == "" {
		return
	}

	content := fmt.Sprintf("🚨 **서버 에러 발생!**\n- **경로**: %s\n- **상태**: %d\n- **에러**: %v",
		c.Request.URL.Path, c.Writer.Status(), c.Errors.String())

	payload := map[string]string{"content": content}
	jsonBody, _ := json.Marshal(payload)

	http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonBody))
}
