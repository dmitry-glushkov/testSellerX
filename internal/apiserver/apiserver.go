package apiserver

import (
	"database/sql"
	"net/http"
	"testSellerX/internal/storage"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	storage := storage.New(db)
	s := newServer(*storage)
	if err := http.ListenAndServe(config.BindAddr, s); err != nil {
		return err
	}
	return nil
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.Exec(
		`CREATE TABLE IF NOT EXISTS Users (
			id serial NOT NULL PRIMARY KEY,
			username varchar NOT NULL UNIQUE,
			created_at timestamp DEFAULT CURRENT_TIMESTAMP
		)`,
	)

	db.Exec(
		`CREATE TABLE IF NOT EXISTS Chat (
			id serial NOT NULL PRIMARY KEY,
			chat_name varchar NOT NULL,
			created_at timestamp DEFAULT CURRENT_TIMESTAMP
		)`,
	)

	db.Exec(
		`CREATE TABLE IF NOT EXISTS User_to_chat (
			user_id int REFERENCES Users (id),
			chat_id int REFERENCES Chat (id)
		)`,
	)

	db.Exec(
		`CREATE TABLE IF NOT EXISTS Messages (
			id serial NOT NULL PRIMARY KEY,
			chat_id int REFERENCES Chat (id),
			user_id int REFERENCES Users (id),
			message_text text NOT NULL,
			created_at timestamp DEFAULT CURRENT_TIMESTAMP
		)`,
	)

	return db, nil
}
