package handlers

import (
	"github.com/crockeo/personalwebsite/database"
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
		ErrorHandler(w, r, 404)
	} else {
		id := int(id64)

		db := database.QuickOpenDB()
		defer db.Close()

		post, err := database.GetPost(db, id)

		if err != nil {
			ErrorHandler(w, r, 404)
		} else {
			helpers.SendPage(w, "blog", struct{ Posts []template.HTML }{Posts: []template.HTML{post.Display()}})
		}
	}
}

// The blog display itself
func BlogHandler(w http.ResponseWriter, r *http.Request) {
	posts := database.QuickGetPosts()

	if len(posts) == 0 {
		helpers.SendPage(w, "noblog", struct{}{})
	} else {
		dposts := make([]template.HTML, len(posts))

		for i := 0; i < len(posts); i++ {
			dposts[i] = posts[i].Display()
		}

		helpers.SendPage(w, "blog", struct{ Posts []template.HTML }{Posts: dposts})
	}
}

// Initializing the Blog handlers
func InitBlogHandlers(m *martini.ClassicMartini) {
	m.Get("/blog/:id", PostHandler)
	m.Get("/blog", BlogHandler)
}
