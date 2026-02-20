package handlers

import (
	"go-study/database"
	"go-study/models"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan Message)
	mutex     = sync.Mutex{}
)

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	defer func() {
		mutex.Lock()
		delete(clients, conn)
		mutex.Unlock()
	}()

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			break
		}
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast

		// 🆕 DB에 채팅 내역 저장
		chatEntry := models.ChatMessage{
			Username: msg.Username,
			Content:  msg.Content,
		}
		database.DB.Create(&chatEntry)

		mutex.Lock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}

func GetChatHistory(c *gin.Context) {
	var messages []models.ChatMessage
	// 최근 50개의 메시지만 가져오기
	database.DB.Order("created_at desc").Limit(50).Find(&messages)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   messages,
	})
}
