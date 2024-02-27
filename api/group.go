package api

import (
	"classroom/dto"
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) handleCreateGroup(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return apiError{Err: "Invalid method.", Status: http.StatusMethodNotAllowed}
	}

	var group dto.CreateGroup

	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(group)

	return nil
}

type TestDTO struct {
	Name string
	Classes string
}
