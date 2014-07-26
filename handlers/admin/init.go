package admin

import (
	"github.com/go-martini/martini"
)

// Initializing the Admin handlers
func Init(m *martini.ClassicMartini) {
	m.Post("/admin/login", LoginHandler)
	m.Post("/admin/new", NewBlogPostHandler)
	m.Post("/admin/update", UpdateHandler)
	m.Get("/admin/new", NewBlogPostHandler)
	m.Get("/admin/nono", NonoHandler)
	m.Get("/admin", Handler)
}
