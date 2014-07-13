package handlers

import (
	"github.com/crockeo/personalwebsite/blog"
	"github.com/crockeo/personalwebsite/helpers"
	"net/http"
)

// Making a new post
func NewPostHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("login")

	if err != nil {
		helpers.SendPage(w, "notlogged", struct{}{})
	} else {
		auth, err := blog.LoadDefaultAuth()

		if err != nil {
			ErrorHandler(w, r, 503)
		} else if cookie.Value != auth.String() {
			helpers.SendPage(w, "invalidlogin", struct{}{})
		} else {
			helpers.SendPage(w, "newpost", struct{}{})
		}
	}
}

// Authenticating to the blog
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	helpers.SendPage(w, "login", struct{}{})
}

// The blog display itself
func BlogHandler(w http.ResponseWriter, r *http.Request) {
	iserr := Check404(w, r, r.URL.Path[5:])

	if !iserr {
		posts, err := blog.LoadDefaultPosts()

		if posts == nil || err != nil {
			helpers.SendPage(w, "noblog", struct{}{})
		} else {
			helpers.SendPage(w, "blog", struct{ Posts []*blog.Post }{Posts: posts})
		}
	}
}
