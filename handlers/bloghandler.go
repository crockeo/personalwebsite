package handlers

import (
	"net/http"
	"github.com/crockeo/personalwebsite/blog"
	"github.com/crockeo/personalwebsite/helpers"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	iserr := Check404(w, r, r.URL.Path[5:])

	if !iserr {
		posts, err := blog.LoadPosts("posts")

		if err != nil {
			helpers.SendPage(w, "noblog", struct{}{})
		} else {
			helpers.SendPage(w, "blog", struct{ Posts []blog.Post }{ Posts: posts })
		}
	}
}
