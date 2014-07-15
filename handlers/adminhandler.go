package handlers

import (
	"github.com/crockeo/personalwebsite/admin"
	"github.com/crockeo/personalwebsite/config"
	"github.com/crockeo/personalwebsite/helpers"
	"net/http"
)

// The handler to deal with logging into the admin account
func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username != "" && password != "" {
		auth, err := admin.LoadDefaultAuth()

		if err != nil {
			ErrorHandler(w, r, 503)
		} else {
			nauth := admin.NewAuth(username, password)

			if auth.Equal(nauth) {
				http.SetCookie(w, admin.MakeAuthCookie(nauth))
				http.Redirect(w, r, "/admin/", 301)
			} else {
				helpers.SendPage(w, "nono", struct{}{})
			}
		}
	} else {
		http.Redirect(w, r, "/admin/", 301)
	}
}

// The handler to deal with logging out of the admin account
func AdminLogoutHandler(w http.ResponseWriter, r *http.Request) {
	cauth, err := r.Cookie(config.AuthName)

	if err != nil {
		http.Redirect(w, r, "/admin/", 301)
	} else {
		http.SetCookie(w, admin.ToDelete(cauth))
		http.Redirect(w, r, "/admin", 301)
	}
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	admin.NewAuth("admin", "password")
	cauth, err := r.Cookie(config.AuthName)

	if err != nil {
		helpers.SendPage(w, "login", struct{}{})
	} else {
		auth, err := admin.LoadDefaultAuth()

		if err != nil {
			ErrorHandler(w, r, 503)
		} else if auth.String() == cauth.Value {
			helpers.SendPage(w, "apanel", struct{}{})
		} else {
			helpers.SendPage(w, "home", struct{}{})
		}
	}
}
