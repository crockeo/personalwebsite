package admin

import "github.com/go-martini/martini"

// Initializing all of the handlers for admin
func Init(m *martini.ClassicMartini) {
	m.Post("/admin/login", loginHandler)
	m.Post("/admin/new", postNewPostHandler)
	m.Post("/admin/update", updateHandler)
	m.Get("/admin/new", getNewPostHandler)
	m.Get("/admin/nono", nonoHandler)
	m.Get("/admin", handler)
}
