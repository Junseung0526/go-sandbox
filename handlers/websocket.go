package handlers

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 접속한 클라이언트들을 관리할 맵과 뮤텍스
var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan Message)
	mutex     = sync.Mutex{}
)

// 메시지 구조체
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

	// 1. 새로운 클라이언트 등록
	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	// 연결 종료 시 클라이언트 제거
	defer func() {
		mutex.Lock()
		delete(clients, conn)
		mutex.Unlock()
	}()

	for {
		var msg Message
		// JSON 형태의 메시지 읽기
		err := conn.ReadJSON(&msg)
		if err != nil {
			break
		}
		// 브로드캐스트 채널로 메시지 전송
		broadcast <- msg
	}
}

// 2. 모든 클라이언트에게 메시지를 전달하는 루틴 (main에서 실행)
func HandleMessages() {
	for {
		msg := <-broadcast
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
