package apiserver

import (
	"apiserver/internal/store"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	store  store.Store
}

func NewServer(st store.Store) *Server {
	s := Server{
		router: mux.NewRouter(),
		store:  st,
	}
	s.ConfigureRouter()
	return &s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) ConfigureRouter() {
	s.router.HandleFunc("/tasks", s.tasksHandler())
	s.router.HandleFunc("/users", s.usersHandler())
}

/*
CREATE - POST, валидация, проверка оригинальности, функция создания
FINDall - get

# Edit - PUT

DELETE - idk (ban user)
*/
func (s *Server) usersHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:

		case http.MethodPost:
			//создание юзера
		case http.MethodPut:
			//Редактирование данных юзера
		case http.MethodDelete:
			//Удаление аккаунта юзера
		default:
			// ошибка
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

/*
GET: GETalltasks
POST: создать задачу
PUT: -
DELETE:удалить задачу
*/
func (s *Server) tasksHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
		case http.MethodPost:
			w.WriteHeader(http.StatusContinue)
		case http.MethodPut:
			w.WriteHeader(http.StatusBadRequest)
		case http.MethodDelete:
			w.WriteHeader(http.StatusAccepted)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

/*
	хранить id решения у пользователя, при удалении задачи удалить решение
	создать решение

GET: -
POST:
PUT:
DELETE:
*/
func (s *Server) solutionsHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
		case http.MethodPost:
			w.WriteHeader(http.StatusContinue)
		case http.MethodPut:
			w.WriteHeader(http.StatusBadRequest)
		case http.MethodDelete:
			w.WriteHeader(http.StatusAccepted)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

/*
тесты к задаче, удаление его происходит с задачей, по id просмотр

GET:
POST:
PUT:
DELETE:
*/
func (s *Server) datatestsHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
		case http.MethodPost:
			w.WriteHeader(http.StatusContinue)
		case http.MethodPut:
			w.WriteHeader(http.StatusBadRequest)
		case http.MethodDelete:
			w.WriteHeader(http.StatusAccepted)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (s *Server) notificationsHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
		case http.MethodPost:
			w.WriteHeader(http.StatusContinue)
		case http.MethodPut:
			w.WriteHeader(http.StatusBadRequest)
		case http.MethodDelete:
			w.WriteHeader(http.StatusAccepted)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (s *Server) commentsHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
		case http.MethodPost:
			w.WriteHeader(http.StatusContinue)
		case http.MethodPut:
			w.WriteHeader(http.StatusBadRequest)
		case http.MethodDelete:
			w.WriteHeader(http.StatusAccepted)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (s *Server) tagsHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
		case http.MethodPost:
			w.WriteHeader(http.StatusContinue)
		case http.MethodPut:
			w.WriteHeader(http.StatusBadRequest)
		case http.MethodDelete:
			w.WriteHeader(http.StatusAccepted)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
