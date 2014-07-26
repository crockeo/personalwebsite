package project

import (
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/handlers/error"
	"github.com/crockeo/personalwebsite/helpers"
	"github.com/go-martini/martini"
	"net/http"
)

// Displaying an individual project
func ProjectProjectHandler(w http.ResponseWriter, r *http.Request, params martini.Params) {
	title := params["name"]

	if title == "" {
		error.ErrorHandler(w, r, 404)
	} else {
		db := database.QuickOpenDB()
		defer db.Close()

		project, err := database.GetProjectByTitle(db, title)

		if err != nil {
			error.ErrorHandler(w, r, 404)
		} else {
			helpers.SendPage(w, "project", toDisplayProject(project))
		}

	}
}
