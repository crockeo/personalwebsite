package handlers

import (
	"github.com/crockeo/personalwebsite/helpers"
	"github.com/go-martini/martini"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if !Check404(w, r, "/") {
		helpers.SendPage(w, "home", struct{}{})
	}
}

// Initializing the Home handlers
func InitHomeHandlers(m *martini.ClassicMartini) {
	m.Get("/", HomeHandler)
}
