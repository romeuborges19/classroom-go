package api

import (
	"io"
	"net/http"
)

func (s *Server) handleHomePage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is my API!\n")
}
