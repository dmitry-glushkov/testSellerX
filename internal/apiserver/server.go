package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
}

// Создание и насройка сервера
func newServer() *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
	}

	s.configureRouter()

	return s
}

// Связывание хэндлеров и соответствующих им URL
func (s *server) configureRouter() {
	s.router.HandleFunc("/users/add", s.handleUserCreate()).Methods("POST")
	s.router.HandleFunc("/chats/add", s.handleChatCreate()).Methods("POST")
	s.router.HandleFunc("/messages/add", s.handleSendMessage()).Methods("POST")
	s.router.HandleFunc("/chats/get", s.handleGetUserChats()).Methods("POST")
	s.router.HandleFunc("/messages/get", s.handleGetChatMessages()).Methods("POST")
}
