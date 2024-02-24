package api

import (
	"classroom/service/google"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	mux *mux.Router
	googleService service.GoogleService
}

func NewServer(googleService service.GoogleService) *Server{
	mux := mux.NewRouter()
	return &Server{
		mux: mux,
		googleService: googleService,
	}
}

func (s *Server) Start() error {
	s.mux.HandleFunc("/courses", MakeHTTPHandler(s.handleGetCourses))
	s.mux.HandleFunc("/courses/{id}", MakeHTTPHandler(s.handleGetCourseStudentsData))
	s.mux.HandleFunc("/courses/get", MakeHTTPHandler(s.handleGetLisfOfCourseStudentsData))


	configuredRouter := LoggingMiddleware(s.mux)
	return http.ListenAndServe(":8080", configuredRouter)
}
