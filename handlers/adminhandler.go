package handlers

import (
	"fmt"
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
				http.Redirect(w, r, "/admin/nono/", 301)
			}
		}
	} else {
		http.Redirect(w, r, "/admin/", 301)
	}
}

// Updating admin information
func AdminUpdateHandler(w http.ResponseWriter, r *http.Request) {
	cauth, err := r.Cookie(config.AuthName)

	if err != nil {
		http.Redirect(w, r, "/admin/", 301)
	} else {
		auth, err := admin.LoadDefaultAuth()

		if err != nil {
			ErrorHandler(w, r, 503)
		} else if auth.String() == cauth.Value {
			newusername := r.FormValue("newusername")
			newpassword := r.FormValue("newpassword")
			cnewpassword := r.FormValue("cnewpassword")
			oldpassword := r.FormValue("oldpassword")

			if newusername != "" && newpassword != "" && cnewpassword != "" && oldpassword != "" {
				if oldpassword != auth.Password {
					http.Redirect(w, r, "/admin/nono/", 301)
				} else {
					fmt.Println("RRAWR")
					helpers.SendPage(w, r, "admin", struct{}{})
				}
			}
		} else {
			http.Redirect(w, r, "/admin/nono/", 301)
		}
	}
}

// Serving the no-no page
func AdminNonoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(503)
	helpers.SendPage(w, "nono", struct{}{})
}

// The base admin handler
func AdminHandler(w http.ResponseWriter, r *http.Request) {
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
