package api

import (
	"classroom/dto"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type apiError struct {
	Err string
	Status int
}

func (e apiError) Error() string {
	return e.Err
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func MakeHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if e, ok := err.(apiError); ok {
				writeJSON(w, e.Status, e)
			}
			writeJSON(w, http.StatusMethodNotAllowed, apiError{Err: "Internal server error."})
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func (s *Server) handleGetCourses(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return apiError{Err: "Invalid method.", Status: http.StatusMethodNotAllowed}
	}

	ch := make(chan map[string]string)
	go s.googleService.GetCourses(ch)
	resp := <- ch 

	return writeJSON(w, http.StatusOK, resp)
}

func (s *Server) handleGetCourseStudentsData(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return apiError{Err: "Invalid method.", Status: http.StatusMethodNotAllowed}
	}
	vars := mux.Vars(r)
	id := vars["id"]

	ch := make(chan []dto.StudentInfo)
	go s.googleService.GetCourseData(id, ch)
	resp := <- ch

	return writeJSON(w, http.StatusOK, resp)
}

