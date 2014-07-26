package project

import (
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/handlers/error"
	"github.com/crockeo/personalwebsite/helpers"
	"github.com/go-martini/martini"
	"net/http"
)

// Displaying an individual course
func ProjectCourseHandler(w http.ResponseWriter, r *http.Request, params martini.Params) {
	sertitle := params["name"]

	if sertitle == "" {
		error.ErrorHandler(w, r, 404)
	} else {
		db := database.QuickOpenDB()
		defer db.Close()

		course, err := database.GetCourseBySerTitle(db, sertitle)

		if err != nil {
			error.ErrorHandler(w, r, 404)
		} else {
			helpers.SendPage(w, "course", course)
		}
	}
}
