package api

import (
	"classroom/dto"
	"net/http"

	"github.com/gorilla/mux"
)


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

