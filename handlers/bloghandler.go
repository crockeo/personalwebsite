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
	title := r.FormValue("title")
	author := r.FormValue("author")
	body := r.FormValue("body")

	if title != "" && author != "" && body != "" {
		blog.SavePostNext(blog.NewPost(title, author, body))
		http.Redirect(w, r, "/blog/", 301)
	} else {
		helpers.SendPage(w, "newpost", struct{}{})
	}
}

// Authenticating to the blog
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, r, 404)
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
