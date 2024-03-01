package routes

import (
	"encoding/json"
	"net/http"
)

type apiError struct {
	Err string
	Status int
}

func (e apiError) Error() string {
	return e.Err
}

func APIError(err string, status int) apiError {
	return apiError{
		Err: err,
		Status: status,
	}
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

