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

}
