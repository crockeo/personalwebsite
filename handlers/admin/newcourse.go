package admin

import (
	"github.com/crockeo/personalwebsite/config"
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/helpers"
	"net/http"
)

// GET request for a new course
func GetNewCourseHandler(w http.ResponseWriter, r *http.Request) {
	helpers.SendPage(w, "newcourse", struct{}{})
}

// POST request for a new course
func PostNewCourseHandler(w http.ResponseWriter, r *http.Request) {
	cauth, err := r.Cookie(config.AuthName)

	if err != nil {
		http.Redirect(w, r, "/admin", 301)
	} else {
		title := r.FormValue("title")
		sertitle := r.FormValue("sertitle")
		inst := r.FormValue("inst")
		description := r.FormValue("description")
		comments := r.FormValue("comments")

		if title != "" && sertitle != "" && inst != "" && description != "" && comments != "" {
			auth := database.QuickGetAuth()

			if auth.SecureString() == cauth.Value {
				database.QuickInsertCourse(&database.Course{
					Title:       title,
					SerTitle:    sertitle,
					Inst:        inst,
					Description: description,
					Comments:    comments,
				})

				http.Redirect(w, r, "/project/course/"+sertitle, 301)
			} else {
				http.Redirect(w, r, "/admin/nono", 301)
			}
		} else {
			http.Redirect(w, r, "/admin/newcourse", 301)
		}
	}
}
