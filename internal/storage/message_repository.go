package storage

import "testSellerX/internal/model"

// Для работы с таблицей Messages
type MessageRepository struct {
	storage *Storage
}

func (mr *MessageRepository) Create(m *model.Message) error {
	return mr.storage.db.QueryRow(
		"INSERT INTO Messages (chat_id, author_id, message_text) VALUES ($1, $2, $3)",
		m.ChatID,
		m.AuthorID,
		m.Text,
	).Scan(&m.ID)
}
