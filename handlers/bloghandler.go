package handlers

import (
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/helpers"
	"html/template"
	"net/http"
	"strconv"
)

// Displaying a single blogpost
func postHandler(w http.ResponseWriter, r *http.Request, num int) {
	db, err := database.OpenDefaultDatabase()

	if err != nil {
		ErrorHandler(w, r, 503)
	} else {
		post, err := database.GetPost(db, num)

		if err != nil {
			ErrorHandler(w, r, 404)
		} else {
			helpers.SendPage(w, "blog", struct{ Posts []template.HTML }{Posts: []template.HTML{post.Display()}})
		}
	}

	db.Close()
}

// The blog display itself
func BlogHandler(w http.ResponseWriter, r *http.Request) {
	num, err := strconv.ParseInt(r.URL.Path[6:], 10, 64)

	if err == nil {
		postHandler(w, r, int(num))
	} else {
		if !Check404(w, r, r.URL.Path[5:]) {
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
	}
}
