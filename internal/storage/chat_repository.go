package storage

import "testSellerX/internal/model"

// Для работы с таблицей Chat
type ChatRepository struct {
	storage *Storage
}

func (cr *ChatRepository) Create(c *model.Chat) error {
	return cr.storage.db.QueryRow(
		"INSERT INTO Chat chat_name VALUES $1 RETURNING (id, created_at)",
		c.Name,
	).Scan(&c.ID, &c.CreatedAt)
}
