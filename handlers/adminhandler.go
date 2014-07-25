package handlers

import (
	"github.com/crockeo/personalwebsite/config"
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/helpers"
	"github.com/go-martini/martini"
	"net/http"
)

// The handler to deal with logging into the admin account
func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username != "" && password != "" {
		auth := database.QuickGetAuth()
		nauth := &database.Auth{
			Username: username,
			Password: password,
		}

		if auth.Equal(nauth) {
			http.SetCookie(w, nauth.MakeCookie())
			http.Redirect(w, r, "/admin", 301)
		} else {
			http.Redirect(w, r, "/admin/nono", 301)
		}
	} else {
		http.Redirect(w, r, "/admin/", 301)
	}
}

// Making a new blog post
func AdminNewBlogPostHandler(w http.ResponseWriter, r *http.Request) {
	cauth, err := r.Cookie(config.AuthName)

	if err != nil {
		http.Redirect(w, r, "/admin/", 301)
	} else {
		title := r.FormValue("title")
		author := r.FormValue("author")
		body := r.FormValue("body")

		if title != "" && author != "" && body != "" {
			auth := database.QuickGetAuth()

			if auth.SecureString() == cauth.Value {
				database.QuickInsertPost(database.MakeNewPost(title, author, body))
				http.Redirect(w, r, "/blog", 301)
			} else {
				http.Redirect(w, r, "/admin/nono", 301)
			}
		} else {
			helpers.SendPage(w, "newpost", struct{}{})
		}
	}
}

// Updating admin information
func AdminUpdateHandler(w http.ResponseWriter, r *http.Request) {
	cauth, err := r.Cookie(config.AuthName)

	if err != nil {
		http.Redirect(w, r, "/admin", 301)
	} else {
		auth := database.QuickGetAuth()

		if auth.SecureString() == cauth.Value {
			newusername := r.FormValue("newusername")
			newpassword := r.FormValue("newpassword")
			cnewpassword := r.FormValue("cnewpassword")
			oldpassword := r.FormValue("oldpassword")

			if newusername != "" && newpassword != "" && cnewpassword != "" && oldpassword != "" {
				if oldpassword != auth.Password {
					http.Redirect(w, r, "/admin/nono/", 301)
				} else {
					nauth := database.MakeNewAuth(newusername, newpassword)

					database.QuickChangeAuth(nauth)
					http.SetCookie(w, nauth.MakeCookie())
					http.Redirect(w, r, "/admin", 301)
				}
			} else {
				http.Redirect(w, r, "/admin", 301)
			}
		}
	}
}

// Serving the no-no page
func AdminNonoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(503)
	helpers.SendPage(w, "nono", struct{}{})
}

// The base admin handler
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	cauth, err := r.Cookie(config.AuthName)

	if err != nil {
		helpers.SendPage(w, "login", struct{}{})
	} else {
		auth := database.QuickGetAuth()

		if auth.SecureString() == cauth.Value {
			helpers.SendPage(w, "apanel", struct{}{})
		} else {
			http.Redirect(w, r, "/admin/nono/", 301)
		}
	}
}

// Initializing the Admin handlers
func InitAdminHandlers(m *martini.ClassicMartini) {
	m.Post("/admin/login", AdminLoginHandler)
	m.Post("/admin/new", AdminNewBlogPostHandler)
	m.Post("/admin/update", AdminUpdateHandler)
	m.Get("/admin/new", AdminNewBlogPostHandler)
	m.Get("/admin/nono", AdminNonoHandler)
	m.Get("/admin", AdminHandler)
}
