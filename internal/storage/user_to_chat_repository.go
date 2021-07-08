package storage

import "testSellerX/internal/model"

// Отношение многие ко многим между пользователями и чатами
// Работа с таблицей User_to_chat
type UserToChatRepository struct {
	storage *Storage
}

// Создание новой записи в User_to_chat с информацией о пользователях (user_id) и их чатами (chat_id)
func (ucr *UserToChatRepository) Create(c *model.Chat) error {
	for userId := range c.UsersID {
		err := ucr.storage.db.QueryRow(
			"INSERT INTO User_to_chat (chat_id, user_id) VALUES ($1, $2)",
			c.ID,
			userId,
		).Scan()
		if err != nil {
			return err
		}
	}
	return nil
}

// Вовзвращает список чатов (chat_id) в которых состоит пользователь u
func (ucr *UserToChatRepository) FindUserChats(u *model.User) ([]int, error) {
	var chats_id []int

	rows, err := ucr.storage.db.Query(
		"SELECT chat_id FROM User_to_chat WHERE user_id = $1",
		u.ID,
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
		chats_id = append(chats_id, tmp)
	}

	return chats_id, nil
}
