package handlers

import (
	"github.com/crockeo/personalwebsite/blog"
	"github.com/crockeo/personalwebsite/helpers"
	"html/template"
	"net/http"
	"time"
)

// A DisplayPost type
type DisplayPost struct {
	Id      int
	Title   string
	Author  string
	Body    template.HTML
	Written time.Time
}

func toDisplayPost(post *blog.Post) *DisplayPost {
	dp := new(DisplayPost)

	dp.Id = post.Id
	dp.Title = post.Title
	dp.Author = post.Author
	dp.Body = template.HTML(post.Body)
	dp.Written = post.Written

	return dp
}

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
		posts, err := blog.LoadPosts()

		if posts == nil || err != nil {
			helpers.SendPage(w, "noblog", struct{}{})
		} else {
			displayposts := make([]*DisplayPost, len(posts))
			for i := 0; i < len(posts); i++ {
				displayposts[i] = toDisplayPost(posts[i])
			}

			helpers.SendPage(w, "blog", struct{ Posts []*DisplayPost }{Posts: displayposts})
		}
	}
}
