package handlers

import (
	"github.com/crockeo/personalwebsite/helpers"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if !Check404(w, r, "/") {
		helpers.SendPage(w, "home", struct{}{})
	}
}
