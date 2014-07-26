package admin

import (
	"github.com/crockeo/personalwebsite/config"
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/helpers"
	"net/http"
)

// The base admin handler
func Handler(w http.ResponseWriter, r *http.Request) {
	cauth, err := r.Cookie(config.AuthName)

	if err != nil {
		helpers.SendPage(w, "login", struct{}{})
	} else {
		auth := database.QuickGetAuth()

		if auth.SecureString() == cauth.Value {
			helpers.SendPage(w, "apanel", struct{}{})
		} else {
			http.Redirect(w, r, "/admin/nono", 301)
		}
	}
}
