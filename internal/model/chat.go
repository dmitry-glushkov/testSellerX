package model

import "time"

// Структура, описывающая чат
type Chat struct {
	ID        int
	Name      string
	UsersID   []int
	CreatedAt time.Time
}
