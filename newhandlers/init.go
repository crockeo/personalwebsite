package handlers

import (
	"github.com/crockeo/personalwebsite/newhandlers/admin"
	"github.com/crockeo/personalwebsite/newhandlers/blog"
	"github.com/crockeo/personalwebsite/newhandlers/errors"
	"github.com/crockeo/personalwebsite/newhandlers/home"
	"github.com/go-martini/martini"
)

// Initializing all of the handlers
func InitHandlers(m *martini.ClassicMartini) {
	admin.Init(m)
	blog.Init(m)
	errors.Init(m)
	home.Init(m)
}
