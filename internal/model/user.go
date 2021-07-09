package model

import "time"

// Структура, описывающая пользователей
type User struct {
	ID        int    `json:"user_id"`
	Username  string `json:"username"`
	CreatedAt time.Time
}
