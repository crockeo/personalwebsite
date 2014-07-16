package handlers

import (
	"github.com/crockeo/personalwebsite/admin"
	"github.com/crockeo/personalwebsite/blog"
	"github.com/crockeo/personalwebsite/config"
	"github.com/crockeo/personalwebsite/helpers"
	"net/http"
	"strconv"
)

// The handler to deal with logging into the admin account
func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	if !Check404(w, r, r.URL.Path[12:]) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username != "" && password != "" {
			auth, err := admin.LoadDefaultAuth()

			if err != nil {
				ErrorHandler(w, r, 503)
			} else {
				nauth := admin.NewAuth(username, password)

				if auth.Equal(nauth) {
					http.SetCookie(w, admin.MakeAuthCookie(nauth))
					http.Redirect(w, r, "/admin/", 301)
				} else {
					http.Redirect(w, r, "/admin/nono/", 301)
				}
			}
		} else {
			http.Redirect(w, r, "/admin/", 301)
		}
	}
}

// Making a new blog post
func AdminNewBlogPostHandler(w http.ResponseWriter, r *http.Request) {
	if !Check404(w, r, r.URL.Path[10:]) {
		cauth, err := r.Cookie(config.AuthName)

		if err != nil {
			http.Redirect(w, r, "/admin/", 301)
		} else {
			auth, err := admin.LoadDefaultAuth()

			if err != nil {
				ErrorHandler(w, r, 503)
			} else if auth.String() == cauth.Value {
				title := r.FormValue("title")
				author := r.FormValue("author")
				body := r.FormValue("body")

				blog.SavePostNext(blog.MakePost(title, author, body))
				http.Redirect(w, r, "/blog/"+strconv.FormatInt(int64(blog.Posts()-1), 10), 301)
			} else {
				http.Redirect(w, r, "/admin/nono", 301)
			}
		}
	}
}

// Updating admin information
func AdminUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if !Check404(w, r, r.URL.Path[13:]) {
		cauth, err := r.Cookie(config.AuthName)

		if err != nil {
			http.Redirect(w, r, "/admin/", 301)
		} else {
			auth, err := admin.LoadDefaultAuth()

			if err != nil {
				ErrorHandler(w, r, 503)
			} else if auth.String() == cauth.Value {
				newusername := r.FormValue("newusername")
				newpassword := r.FormValue("newpassword")
				cnewpassword := r.FormValue("cnewpassword")
				oldpassword := r.FormValue("oldpassword")

				if newusername != "" && newpassword != "" && cnewpassword != "" && oldpassword != "" {
					if oldpassword != auth.Password {
						http.Redirect(w, r, "/admin/nono/", 301)
					} else {
						nauth := admin.NewAuth(newusername, newpassword)

						admin.SaveDefaultAuth(nauth)
						admin.SendAuthCookie(w, nauth)

						http.Redirect(w, r, "/admin/", 301)
					}
				} else {
					http.Redirect(w, r, "/admin/", 301)
				}
			} else {
				http.Redirect(w, r, "/admin/nono/", 301)
			}
		}
	}
}

// Serving the no-no page
func AdminNonoHandler(w http.ResponseWriter, r *http.Request) {
	if !Check404(w, r, r.URL.Path[11:]) {
		w.WriteHeader(503)
		helpers.SendPage(w, "nono", struct{}{})
	}
}

// The base admin handler
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	if !Check404(w, r, r.URL.Path[6:]) {
		cauth, err := r.Cookie(config.AuthName)

		if err != nil {
			helpers.SendPage(w, "login", struct{}{})
		} else {
			auth, err := admin.LoadDefaultAuth()

			if err != nil {
				ErrorHandler(w, r, 503)
			} else if auth.String() == cauth.Value {
				helpers.SendPage(w, "apanel", struct{}{})
			} else {
				http.Redirect(w, r, "/admin/nono/", 301)
			}
		}
	}
}
