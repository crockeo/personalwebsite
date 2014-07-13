package handlers

import (
	"github.com/crockeo/personalwebsite/blog"
	"github.com/crockeo/personalwebsite/helpers"
	"net/http"
	"net/http/cookiejar"
)

// Making a new post
func NewPostHandler(w http.ResponseWriter, r *http.Request) {
	jar, err := cookiejar.New(nil)

	if err != nil {
		ErrorHandler(w, r, 503)
	} else {
		cookies := jar.Cookies(r.URL)

		if len(cookies) == 0 || cookies[0] == nil {
			ErrorHandler(w, r, 0)
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
		posts, err := blog.LoadPosts("posts")

		if posts == nil || err != nil {
			helpers.SendPage(w, "noblog", struct{}{})
		} else {
			helpers.SendPage(w, "blog", struct{ Posts []*blog.Post }{Posts: posts})
		}
	}
}
