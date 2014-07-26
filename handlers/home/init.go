package home

import "github.com/go-martini/martini"

// Initializing the Home handlers
func Init(m *martini.ClassicMartini) {
	m.Get("/", HomeHandler)
}
