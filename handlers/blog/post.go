package blog

import (
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/handlers/error"
	"github.com/crockeo/personalwebsite/helpers"
	"github.com/go-martini/martini"
	"html/template"
	"net/http"
	"strconv"
)

// Displaying a single blogpost
func PostHandler(w http.ResponseWriter, r *http.Request, params martini.Params) {
	id64, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		error.ErrorHandler(w, r, 404)
	} else {
		id := int(id64)

		db := database.QuickOpenDB()
		defer db.Close()

		post, err := database.GetPost(db, id)

		if err != nil {
			error.ErrorHandler(w, r, 404)
		} else {
			helpers.SendPage(w, "blog", struct{ Posts []template.HTML }{Posts: []template.HTML{post.Display()}})
		}
	}
}
