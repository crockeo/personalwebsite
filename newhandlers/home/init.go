package home

import "github.com/go-martini/martini"

// Initializing all of the handlers for home
func Init(m *martini.ClassicMartini) {
	m.Get("/", handler)
}
