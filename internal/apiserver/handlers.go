package apiserver

import (
	"net/http"
)

// Создает пользователя username. Возваращет model.User.ID
func (s *server) handleUserCreate() http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// req := &request{}
		// if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		// 	s.error(w, r, http.StatusBadRequest, err)
		// 	return
		// }

		// u := &model.User{
		// 	Username: req.Username,
		// }
		// if err := s.storage.User().Create(u); err != nil {
		// 	s.error(w, r, http.StatusUnprocessableEntity, err)
		// 	return
		// }

		// s.respond(w, r, http.StatusCreated, u.ID)
	})
}

// Создает чат name пользователей [user_id_1, user_id_2, ...]. Возвращает model.Chat.ID
func (s *server) handleChatCreate() http.HandlerFunc {
	type request struct {
		ChatName string `json:"name"`
		Users    []int  `json:"users"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// Отправляет в чат chat_id сообщение text от лица пользователя user_id. Возвращает model.Message.ID
func (s *server) handleSendMessage() http.HandlerFunc {
	type request struct {
		ChatID   int    `json:"chat"`
		AuthorID int    `json:"author"`
		Text     string `json:"text"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// Возвращает список всех чатов со всеми полями (сортировка по времени последнего сообщения в чате) пользователя user_id
func (s *server) handleGetUserChats() http.HandlerFunc {
	type requset struct {
		UserID int `json:"user"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// Возвращает список всех сообщений со всеми полями (сортировка по времени) чата chat_id
func (s *server) handleGetChatMessages() http.HandlerFunc {
	type request struct {
		ChatID int `json:"chat"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
