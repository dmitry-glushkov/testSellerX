package model

import "time"

// Структура, описывающая сообщения
type Message struct {
	ID        int
	ChatID    int
	AuthorID  int
	Text      string
	CreatedAt time.Time
}
