package storage

import (
	"fmt"
	"testSellerX/internal/model"
)

// Отношение многие ко многим между пользователями и чатами
// Работа с таблицей User_to_chat
type UserToChatRepository struct {
	storage *Storage
}

// Создание новой записи в User_to_chat с информацией о пользователях (user_id) и их чатами (chat_id)
func (ucr *UserToChatRepository) Create(c *model.Chat) error {
	for _, userId := range c.UsersID {
		fmt.Println()
		fmt.Print(userId)
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

// Вовзвращает список чатов (chat_id) в которых состоит пользователь u (user_id)
func (ucr *UserToChatRepository) FindUserChats(userId int) ([]int, error) {
	var chats_id []int

	rows, err := ucr.storage.db.Query(
		"SELECT chat_id FROM User_to_chat WHERE user_id = $1",
		userId,
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
