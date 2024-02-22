package api

import (
	"classroom/service/google"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	mux *mux.Router
}

func NewServer() *Server{
	mux := mux.NewRouter()
	return &Server{
		mux: mux,
	}
}

func (s *Server) Start() error {
	c := service.NewClassroomService()
	g := service.NewGoogleService(c)

	// Testando serviço da API do Google
	g.GetCourses()

	s.mux.HandleFunc("/", s.handleHomePage)
	return http.ListenAndServe(":8000", s.mux)
}
