package models

import "gorm.io/gorm"

type ChatMessage struct {
	gorm.Model
	Username string `json:"username"`
	Content  string `json:"content"`
}
