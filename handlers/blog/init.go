package blog

import "github.com/go-martini/martini"

// Initializing the Blog handlers
func Init(m *martini.ClassicMartini) {
	m.Get("/blog/:id", PostHandler)
	m.Get("/blog", Handler)
}
