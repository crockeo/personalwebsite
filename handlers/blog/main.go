package blog

import (
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/helpers"
	"html/template"
	"net/http"
)

// The blog display itself
func Handler(w http.ResponseWriter, r *http.Request) {
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
