package api

import (
	"classroom/repository"
	"classroom/service/google"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	db *sql.DB
	cache *redis.Client
	mux *mux.Router
	googleService service.GoogleService
}

func NewServer(googleService service.GoogleService) *Server{
	mux := mux.NewRouter()
	cache := repository.NewCache()

	db, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	

	return &Server{
		db: db,
		cache: cache,
		mux: mux,
		googleService: googleService,
	}
}

func (s *Server) Start() error {
	s.mux.HandleFunc("/courses", MakeHTTPHandler(s.handleGetCourses))
	s.mux.HandleFunc("/courses/{id}", MakeHTTPHandler(s.handleGetCourseStudentsData))
	s.mux.HandleFunc("/courses/", MakeHTTPHandler(s.handleGetLisfOfCourseStudentsData))

	configuredRouter := LoggingMiddleware(s.mux)
	return http.ListenAndServe(":8080", configuredRouter)
}
