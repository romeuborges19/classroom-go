package routes

import (
	"classroom/api/middleware"
	"classroom/repository"
	"classroom/service"
	google "classroom/service/google"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	db *sql.DB
	cache *redis.Client
	mux *mux.Router
	googleService google.GoogleService
	groupService service.GroupService
}

func NewServer(googleService google.GoogleService, groupService service.GroupService, db *sql.DB) *Server{
	mux := mux.NewRouter()
	cache := repository.NewCache()

	return &Server{
		db: db,
		cache: cache,
		mux: mux,
		googleService: googleService,
		groupService: groupService,
	}
}

func (s *Server) Start() error {
	s.mux.HandleFunc("/courses", MakeHTTPHandler(s.handleGetCourses))
	s.mux.HandleFunc("/courses/{id}", MakeHTTPHandler(s.handleGetCourseStudentsData))
	s.mux.HandleFunc("/courses/", MakeHTTPHandler(s.handleGetLisfOfCourseStudentsData))
	s.mux.HandleFunc("/groups/create", MakeHTTPHandler(s.handleCreateGroup))

	configuredRouter := middleware.LoggingMiddleware(s.mux)
	return http.ListenAndServe(":8080", configuredRouter)
}
