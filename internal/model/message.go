package model

import "time"

// Структура, описывающая сообщения
type Message struct {
	ID        int    `json:"message_id"`
	ChatID    int    `json:"chat_id"`
	AuthorID  int    `json:"author_id"`
	Text      string `json:"text"`
	CreatedAt time.Time
}
