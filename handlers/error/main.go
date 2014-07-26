package error

import (
	"github.com/crockeo/personalwebsite/helpers"
	"net/http"
)

type Error struct {
	ErrorCode int    // The HTTP error code
	Path      string // The path the error occurred upon
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	helpers.SendPage(w, "error", Error{ErrorCode: status, Path: r.URL.Path})
}
