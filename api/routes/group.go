package routes

import (
	"classroom/dto"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) handleCreateGroup(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return apiError{
			Err: "Invalid method.",
			Status: http.StatusMethodNotAllowed,
		}
	}

	var group dto.CreateGroup
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		return apiError{
			Err: "Error while decoding body",
			Status: http.StatusInternalServerError,
		}
	}

	resp, err := s.groupService.CreateGroup(group, s.db)
	if err != nil {
		log.Fatal(err)
	}

	writeJSON(w, http.StatusCreated, resp)

	return nil
}

func (s *Server) handleGetGroupById(w http.ResponseWriter,r *http.Request) error {
	if r.Method != http.MethodGet {
		return apiError{
			Err: "Invalid method.",
			Status: http.StatusMethodNotAllowed,
		}
	}

	vars := mux.Vars(r)
	id := vars["id"]

	resp, err := s.groupService.GetGroupById(id, s.db)
	if err != nil {	
		return APIError(err, http.StatusInternalServerError)
	}

	writeJSON(w, http.StatusOK, resp)
	return nil
}
