package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Хранение информации о пользователях, сообщениях и чатах
type Storage struct {
	db                   *sql.DB
	userRepository       *UserRepository
	messageRepository    *MessageRepository
	chatRepository       *ChatRepository
	userToChatRepository *UserToChatRepository
}

// Инициализация
func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

// Для работы с таблицей пользователей
func (s *Storage) UserRepos() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		storage: s,
	}

	return s.userRepository
}

// Для работы с таблицей сообщений
func (s *Storage) MessageRepos() *MessageRepository {
	if s.userRepository != nil {
		return s.messageRepository
	}

	s.messageRepository = &MessageRepository{
		storage: s,
	}

	return s.messageRepository
}

// Для работы с таблицей чатов
func (s *Storage) ChatRepos() *ChatRepository {
	if s.userRepository != nil {
		return s.chatRepository
	}

	s.chatRepository = &ChatRepository{
		storage: s,
	}

	return s.chatRepository
}

// Для работы с таблицей User_to_chat
func (s *Storage) UserToChatRepos() *UserToChatRepository {
	if s.userRepository != nil {
		return s.userToChatRepository
	}

	s.userToChatRepository = &UserToChatRepository{
		storage: s,
	}

	return s.userToChatRepository
}
