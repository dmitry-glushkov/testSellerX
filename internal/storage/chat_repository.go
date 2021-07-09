package storage

import "testSellerX/internal/model"

// Для работы с таблицей Chat
type ChatRepository struct {
	storage *Storage
}

func (cr *ChatRepository) Create(c *model.Chat) error {
	return cr.storage.db.QueryRow(
		"INSERT INTO Chat (chat_name) VALUES ($1) RETURNING id",
		c.Name,
	).Scan(&c.ID)
}

func (cr *ChatRepository) GetChatById(id int) (*model.Chat, error) {
	chat := &model.Chat{}
	if err := cr.storage.db.QueryRow(
		"SELECT * FROM Chat WHERE id = $1",
		id,
	).Scan(&chat.ID, &chat.Name, &chat.CreatedAt); err != nil {
		return nil, err
	}

	chat.UsersID, _ = cr.FindChatUsers(chat.ID)

	return chat, nil
}

func (cr *ChatRepository) FindChatUsers(chatId int) ([]int, error) {
	var users_id []int

	rows, err := cr.storage.db.Query(
		"SELECT user_id FROM User_to_chat WHERE chat_id = $1",
		chatId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tmp int
		err := rows.Scan(&tmp)
		if err != nil {
			return nil, err
		}
		users_id = append(users_id, tmp)
	}

	return users_id, nil
}
