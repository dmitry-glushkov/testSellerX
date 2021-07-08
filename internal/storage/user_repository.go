package storage

import "testSellerX/internal/model"

// Для работы с таблицей Users
type UserRepository struct {
	storage *Storage
}

func (ur *UserRepository) Create(u *model.User) error {
	return ur.storage.db.QueryRow(
		"INSERT INTO Users (username) VALUES $1 RETURNING id;",
		u.Username,
	).Scan(&u.ID)
}
