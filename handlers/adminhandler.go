package handlers

import (
	"github.com/crockeo/personalwebsite/config"
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/helpers"
	"net/http"
)

// The handler to deal with logging into the admin account
func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	if !Check404(w, r, r.URL.Path[12:]) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username != "" && password != "" {
			db, err := database.OpenDefaultDatabase()

			if err != nil {
				ErrorHandler(w, r, 503)
			} else {
				auth, err := database.GetAuth(db)

				if err != nil {
					ErrorHandler(w, r, 503)
				} else {
					nauth := database.MakeNewAuth(username, password)

					if auth.Equal(nauth) {
						http.SetCookie(w, nauth.MakeCookie())
						http.Redirect(w, r, "/admin/", 301)
					} else {
						http.Redirect(w, r, "/admin/nono/", 301)
					}
				}
			}

			db.Close()
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
			db, err := database.OpenDefaultDatabase()

			if err != nil {
				ErrorHandler(w, r, 503)
			} else {
				auth, err := database.GetAuth(db)

				if err != nil {
					ErrorHandler(w, r, 503)
				} else if auth.String() == cauth.Value {
					title := r.FormValue("title")
					author := r.FormValue("author")
					body := r.FormValue("body")

					database.InsertPost(db, database.MakeNewPost(title, author, body))
					http.Redirect(w, r, "/blog/", 301)
				} else {
					http.Redirect(w, r, "/admin/nono/", 301)
				}
			}

			db.Close()
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
			db, err := database.OpenDefaultDatabase()

			if err != nil {
				ErrorHandler(w, r, 503)
			} else {
				auth, err := database.GetAuth(db)

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
							nauth := database.MakeNewAuth(newusername, newpassword)

							database.ChangeAuth(db, nauth)
							http.SetCookie(w, nauth.MakeCookie())

							http.Redirect(w, r, "/admin/", 301)
						}
					} else {
						http.Redirect(w, r, "/admin/", 301)
					}
				}
			}

			db.Close()
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
			db, err := database.OpenDefaultDatabase()

			if err != nil {
				ErrorHandler(w, r, 503)
			} else {
				auth, err := database.GetAuth(db)

				if err != nil {
					ErrorHandler(w, r, 503)
				} else if auth.String() == cauth.Value {
					helpers.SendPage(w, "apanel", struct{}{})
				} else {
					http.Redirect(w, r, "/admin/nono/", 301)
				}
			}

			db.Close()
		}
	}
}
