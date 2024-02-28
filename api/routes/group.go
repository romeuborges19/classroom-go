package routes

import (
	"classroom/dto"
	"encoding/json"
	"log"
	"net/http"
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
