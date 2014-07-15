package handlers

import (
	"github.com/crockeo/personalwebsite/blog"
	"github.com/crockeo/personalwebsite/helpers"
	"html/template"
	"net/http"
	"strconv"
)

// Making a new post
func NewPostHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	author := r.FormValue("author")
	body := r.FormValue("body")

	if title != "" && author != "" && body != "" {
		blog.SavePostNext(blog.MakePost(title, author, body))
		http.Redirect(w, r, "/blog/", 301)
	} else {
		helpers.SendPage(w, "newpost", struct{}{})
	}
}

// Authenticating to the blog
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, r, 404)
}

// Displaying a single blogpost
func postHandler(w http.ResponseWriter, r *http.Request, num int) {
	nposts := blog.Posts()

	if num < nposts {
		post, err := blog.LoadPost(num)

		if err != nil {
			ErrorHandler(w, r, 503)
		} else {
			helpers.SendPage(w, "post", struct{ Post template.HTML }{Post: post})
		}
	} else {
		ErrorHandler(w, r, 404)
	}
}

// The blog display itself
func BlogHandler(w http.ResponseWriter, r *http.Request) {
	num, err := strconv.ParseInt(r.URL.Path[6:], 10, 64)

	if err == nil {
		postHandler(w, r, int(num))
	} else {
		iserr := Check404(w, r, r.URL.Path[5:])

		if !iserr {
			posts, err := blog.LoadPosts()

			if posts == nil || err != nil {
				helpers.SendPage(w, "noblog", struct{}{})
			} else {
				helpers.SendPage(w, "blog", struct{ Posts []template.HTML }{Posts: posts})
			}
		}
	}
}
