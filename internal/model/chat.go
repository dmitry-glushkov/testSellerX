package model

import "time"

// Структура, описывающая чат
type Chat struct {
	ID        int    `json:"user_id"`
	Name      string `json:"username"`
	UsersID   []int  `json:"users_ids"`
	CreatedAt time.Time
}
