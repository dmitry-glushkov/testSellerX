package model

import "time"

type Message struct {
	ID        int
	ChatID    int
	AuthorID  int
	Text      string
	CreatedAt time.Time
}
