package model

import "time"

// Структура, описывающая пользователей
type User struct {
	ID        int
	Username  string
	CreatedAt time.Time
}
