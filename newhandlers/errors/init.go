package errors

import "github.com/go-martini/martini"

// Initializing all of the handlers for errors
func Init(m *martini.ClassicMartini) {
	m.NotFound(_404Handler)
}
