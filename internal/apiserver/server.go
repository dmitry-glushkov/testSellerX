package apiserver

import (
	"encoding/json"
	"net/http"
	"testSellerX/internal/storage"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router  *mux.Router
	logger  *logrus.Logger
	storage storage.Storage
}

// Создание и насройка сервера
func newServer(storage storage.Storage) *server {
	s := &server{
		router:  mux.NewRouter(),
		logger:  logrus.New(),
		storage: storage,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Связывание хэндлеров и соответствующих им URL
func (s *server) configureRouter() {
	s.router.HandleFunc("/users/add", s.handleUserCreate()).Methods("POST")
	s.router.HandleFunc("/chats/add", s.handleChatCreate()).Methods("POST")
	s.router.HandleFunc("/messages/add", s.handleSendMessage()).Methods("POST")
	s.router.HandleFunc("/chats/get", s.handleGetUserChats()).Methods("POST")
	s.router.HandleFunc("/messages/get", s.handleGetChatMessages()).Methods("POST")
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
