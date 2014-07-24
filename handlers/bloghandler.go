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
		db, err := database.OpenDefaultDatabase()

		if err != nil {
			ErrorHandler(w, r, 503)
		} else {
			post, err := database.GetPost(db, id)

			if err != nil {
				ErrorHandler(w, r, 404)
			} else {
				helpers.SendPage(w, "blog", struct{ Posts []template.HTML }{Posts: []template.HTML{post.Display()}})
			}
		}

		db.Close()
	}
}

// The blog display itself
func BlogHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.OpenDefaultDatabase()

	if err != nil {
		ErrorHandler(w, r, 503)
	} else {
		posts, err := database.GetPosts(db)

		if len(posts) == 0 || err != nil {
			helpers.SendPage(w, "noblog", struct{}{})
		} else {
			dposts := make([]template.HTML, len(posts))

			for i := 0; i < len(posts); i++ {
				dposts[i] = posts[i].Display()
			}

			helpers.SendPage(w, "blog", struct{ Posts []template.HTML }{Posts: dposts})
		}

		db.Close()
	}
}

// Initializing the Blog handlers
func InitBlogHandlers(m *martini.ClassicMartini) {
	m.Get("/blog/:id", PostHandler)
	m.Get("/blog", BlogHandler)
}
