package project

import (
	"github.com/crockeo/personalwebsite/handlers/error"
	"github.com/crockeo/personalwebsite/helpers"
	"net/http"
)

// The main project handler
func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	data, err := loadProjectRootData()

	if err != nil {
		error.ErrorHandler(w, r, 503)
	} else {
		helpers.SendPage(w, "projectroot", data)
	}
}
