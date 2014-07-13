package handlers

import (
	"github.com/crockeo/personalwebsite/helpers"
	"net/http"
)

type Error struct {
	ErrorCode int
	Path      string
}

func Check404(w http.ResponseWriter, r *http.Request, path string) bool {
	if path != "/" {
		ErrorHandler(w, r, 404)
		return true
	}

	return false
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	path := r.URL.Path

	w.WriteHeader(status)
	helpers.SendPage(w, "error", Error{ErrorCode: status, Path: path})
}
