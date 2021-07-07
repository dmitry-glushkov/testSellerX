package apiserver

import "net/http"

// Создает пользователя username. Возваращет model.User.ID
func (s *server) handleUserCreate() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// Создает чат name пользователей [user_id_1, user_id_2, ...]. Возвращает model.Chat.ID
func (s *server) handleChatCreate() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// Отправляет в чат chat_id сообщение text от лица пользователя user_id. Возвращает model.Message.ID
func (s *server) handleSendMessage() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// Возвращает список всех чатов со всеми полями (сортировка по времени последнего сообщения в чате) пользователя user_id
func (s *server) handleGetUserChats() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// Возвращает список всех сообщений со всеми полями (сортировка по времени) чата chat_id
func (s *server) handleGetChatMessages() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
