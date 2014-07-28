package blog

import "github.com/go-martini/martini"

// Initializing all of the handlers for blog
func Init(m *martini.ClassicMartini) {
	m.Get("/blog", handler)
	m.Get("/blog/(?P<id>\\d+)", postHandler)
}
