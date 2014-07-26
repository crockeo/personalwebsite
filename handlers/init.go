package handlers

import (
	"github.com/crockeo/personalwebsite/handlers/admin"
	"github.com/crockeo/personalwebsite/handlers/blog"
	"github.com/crockeo/personalwebsite/handlers/error"
	"github.com/crockeo/personalwebsite/handlers/home"
	"github.com/crockeo/personalwebsite/handlers/project"
	"github.com/go-martini/martini"
)

func InitHandlers(m *martini.ClassicMartini) {
	admin.Init(m)
	blog.Init(m)
	error.Init(m)
	home.Init(m)
	project.Init(m)
}
