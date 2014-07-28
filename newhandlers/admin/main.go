package admin

import (
	"github.com/crockeo/personalwebsite/config"
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/database/utils"
	"github.com/crockeo/personalwebsite/helpers"
	"github.com/crockeo/personalwebsite/newhandlers/errors"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request, db *database.DB) (int, string) {
	cauth, err := r.Cookie(config.AuthName)

	if err != nil {
		return 200, helpers.RenderPageUnsafe("login", struct{}{})
	}

	auth, err := utils.GetAuth(db)

	if err != nil {
		return errors.ErrorHandler(503, "Could not load saved auth.")
	}

	if utils.AuthString(auth) == cauth.Value {
		return 200, helpers.RenderPageUnsafe("apanel", struct{}{})
	}

	http.Redirect(w, r, "/admin/nono", 301)
	return 200, ""
}
