package handlers

import (
	"github.com/crockeo/personalwebsite/helpers"
	"github.com/go-martini/martini"
	"net/http"
)

type Error struct {
	ErrorCode int
	Path      string
}

func Check404(w http.ResponseWriter, r *http.Request, path string) bool {
	if r.URL.Path != path {
		ErrorHandler(w, r, 404)
		return true
	}

	return false
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	helpers.SendPage(w, "error", Error{ErrorCode: status, Path: r.URL.Path})
}

func Error404Handler(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, r, 404)
}

func InitErrorHandlers(m *martini.ClassicMartini) {
	m.NotFound(Error404Handler)
}
