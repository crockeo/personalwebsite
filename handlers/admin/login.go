package admin

import (
	"github.com/crockeo/personalwebsite/database"
	"net/http"
)

// The handler to deal with logging into the admin account
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username != "" && password != "" {
		auth := database.QuickGetAuth()
		nauth := &database.Auth{
			Username: username,
			Password: password,
		}

		if auth.Equal(nauth) {
			http.SetCookie(w, nauth.MakeCookie())
			http.Redirect(w, r, "/admin", 301)
		} else {
			http.Redirect(w, r, "/admin/nono", 301)
		}
	} else {
		http.Redirect(w, r, "/admin", 301)
	}
}
