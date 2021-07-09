package storage

import "testSellerX/internal/model"

// Для работы с таблицей Messages
type MessageRepository struct {
	storage *Storage
}

func (mr *MessageRepository) Create(m *model.Message) error {
	return mr.storage.db.QueryRow(
		"INSERT INTO Messages (chat_id, author_id, message_text) VALUES ($1, $2, $3) RETURNING id",
		m.ChatID,
		m.AuthorID,
		m.Text,
	).Scan(&m.ID)
}

func (mr *MessageRepository) GetChatMessages(chatId int) ([]model.Message, error) {
	var messages []model.Message

	rows, err := mr.storage.db.Query(
		"SELECT id, chat_id, author_id, message_text, created_at FROM Messages WHERE chat_id = $1",
		chatId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tmp model.Message
		err := rows.Scan(&tmp.ID, &tmp.ChatID, &tmp.AuthorID, &tmp.Text, &tmp.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, tmp)
	}

	return messages, nil
}
