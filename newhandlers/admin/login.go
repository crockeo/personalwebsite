package admin

import (
	"fmt"
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/database/schema"
	"github.com/crockeo/personalwebsite/database/utils"
	"github.com/crockeo/personalwebsite/newhandlers/errors"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request, db *database.DB) (int, string) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username != "" && password != "" {
		auth, err := utils.GetAuth(db)

		if err != nil {
			return errors.ErrorHandler(503, err.Error())
		}

		nauth := schema.Auth{
			Username: username,
			Password: password,
		}

		if utils.AuthEquals(auth, nauth) {
			fmt.Println("SETTING COOKIE")
			http.SetCookie(w, utils.MakeCookie(nauth))
			return handler(w, r, db)
		}

		return nonoHandler()
	} else {
		return handler(w, r, db)
	}
}
