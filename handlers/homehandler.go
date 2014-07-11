package handlers

import (
	"net/http"
	"github.com/crockeo/personalwebsite/helpers"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	iserr := Check404(w, r, r.URL.Path)

	if !iserr {
		helpers.SendPage(w, "home", struct{}{})
	}
}
