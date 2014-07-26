package admin

import (
	"github.com/crockeo/personalwebsite/config"
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/helpers"
	"net/http"
)

// Making a new blog post
func NewBlogPostHandler(w http.ResponseWriter, r *http.Request) {
	cauth, err := r.Cookie(config.AuthName)

	if err != nil {
		http.Redirect(w, r, "/admin/", 301)
	} else {
		title := r.FormValue("title")
		author := r.FormValue("author")
		body := r.FormValue("body")

		if title != "" && author != "" && body != "" {
			auth := database.QuickGetAuth()

			if auth.SecureString() == cauth.Value {
				database.QuickInsertPost(database.MakeNewPost(title, author, body))
				http.Redirect(w, r, "/blog", 301)
			} else {
				http.Redirect(w, r, "/admin/nono", 301)
			}
		} else {
			helpers.SendPage(w, "newpost", struct{}{})
		}
	}
}
