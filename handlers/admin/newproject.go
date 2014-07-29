package admin

import (
	"github.com/crockeo/personalwebsite/config"
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/helpers"
	"net/http"
	"strings"
)

// GET request for a new project
func GetNewProjectHandler(w http.ResponseWriter) {
	helpers.SendPage(w, "newproject", struct{}{})
}

// POST request for a new project
func PostNewProjectHandler(w http.ResponseWriter, r *http.Request) {
	cauth, err := r.Cookie(config.AuthName)

	if err != nil {
		http.Redirect(w, r, "/admin", 301)
	} else {
		title := r.FormValue("title")
		sscreenshots := r.FormValue("screenshots")
		language := r.FormValue("language")
		shortdesc := r.FormValue("shortdesc")
		description := r.FormValue("description")

		if title != "" && language != "" && shortdesc != "" && description != "" {
			var screenshots []string
			if sscreenshots == "" {
				screenshots = []string{}
			} else {
				screenshots = strings.Split(sscreenshots, ",")
			}

			auth := database.QuickGetAuth()
			if auth.SecureString() == cauth.Value {
				database.QuickInsertProject(&database.Project{
					Title:       title,
					Screenshots: screenshots,
					Language:    language,
					ShortDesc:   shortdesc,
					Description: description,
				})

				http.Redirect(w, r, "/project/project/"+title, 301)
			}
		} else {
			http.Redirect(w, r, "/admin/newproject", 301)
		}
	}
}
