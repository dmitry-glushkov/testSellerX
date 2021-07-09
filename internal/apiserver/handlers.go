package apiserver

import (
	"encoding/json"
	"net/http"
	"sort"
	"testSellerX/internal/model"
	"time"
)

// Создает пользователя username. Возваращет model.User.ID
func (s *server) handleUserCreate() http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Username: req.Username,
		}
		if err := s.storage.UserRepos().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, u.ID)
	})
}

// Создает чат с названием name между пользователей [user_id_1, user_id_2, ...]. Возвращает model.Chat.ID
func (s *server) handleChatCreate() http.HandlerFunc {
	type request struct {
		ChatName string `json:"name"`
		Users    []int  `json:"users"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		c := &model.Chat{
			Name:    req.ChatName,
			UsersID: req.Users,
		}
		if err := s.storage.ChatRepos().Create(c); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		if err := s.storage.UserToChatRepos().Create(c); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, c.ID)
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
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		m := &model.Message{
			ChatID:   req.ChatID,
			AuthorID: req.AuthorID,
			Text:     req.Text,
		}
		if err := s.storage.MessageRepos().Create(m); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, m.ID)
	})
}

// Возвращает список всех чатов со всеми полями (сортировка по времени последнего сообщения в чате) пользователя user_id
func (s *server) handleGetUserChats() http.HandlerFunc {
	type requset struct {
		UserID int `json:"user"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := &requset{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		var userChats []model.Chat
		userChats_id, err := s.storage.UserToChatRepos().FindUserChats(req.UserID)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		var latestMessagesInChats []time.Time
		for _, chatId := range userChats_id {
			chatMessages, err := s.storage.MessageRepos().GetChatMessages(chatId)
			if err != nil {
				s.error(w, r, http.StatusNotFound, err)
				return
			}
			latestMessagesInChats = append(latestMessagesInChats, chatMessages[len(chatMessages)-1].CreatedAt)
		}

		chatsMap := make(map[time.Time]int)
		for i := 0; i < len(latestMessagesInChats); i++ {
			chatsMap[latestMessagesInChats[i]] = userChats_id[i]
		}
		sort.Slice(latestMessagesInChats, func(i, j int) bool {
			return latestMessagesInChats[i].Before(latestMessagesInChats[j])
		})
		for i := 0; i < len(latestMessagesInChats); i++ {
			chat, err := s.storage.ChatRepos().GetChatById(chatsMap[latestMessagesInChats[i]])
			if err != nil {
				s.error(w, r, http.StatusNotFound, err)
				return
			}
			userChats = append(userChats, *chat)
		}

		s.respond(w, r, http.StatusOK, userChats)
	})
}

// Возвращает список всех сообщений со всеми полями (сортировка по времени) чата chat_id
func (s *server) handleGetChatMessages() http.HandlerFunc {
	type request struct {
		ChatID int `json:"chat"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		chatMessages, err := s.storage.MessageRepos().GetChatMessages(req.ChatID)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		sort.Slice(chatMessages, func(i, j int) bool {
			return chatMessages[i].CreatedAt.Before(chatMessages[j].CreatedAt)
		})

		s.respond(w, r, http.StatusOK, chatMessages)
	})
}
