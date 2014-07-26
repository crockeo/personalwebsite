package admin

import (
	"github.com/crockeo/personalwebsite/config"
	"github.com/crockeo/personalwebsite/database"
	"net/http"
)

// Updating admin information
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	cauth, err := r.Cookie(config.AuthName)

	if err != nil {
		http.Redirect(w, r, "/admin", 301)
	} else {
		auth := database.QuickGetAuth()

		if auth.SecureString() == cauth.Value {
			newusername := r.FormValue("newusername")
			newpassword := r.FormValue("newpassword")
			cnewpassword := r.FormValue("cnewpassword")
			oldpassword := r.FormValue("oldpassword")

			if newusername != "" && newpassword != "" && cnewpassword != "" && oldpassword != "" {
				if oldpassword != auth.Password {
					http.Redirect(w, r, "/admin/nono", 301)
				} else {
					nauth := database.MakeNewAuth(newusername, newpassword)

					database.QuickChangeAuth(nauth)
					http.SetCookie(w, nauth.MakeCookie())
					http.Redirect(w, r, "/admin", 301)
				}
			} else {
				http.Redirect(w, r, "/admin", 301)
			}
		}
	}
}
