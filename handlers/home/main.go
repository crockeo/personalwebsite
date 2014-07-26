package home

import (
	"github.com/crockeo/personalwebsite/helpers"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	helpers.SendPage(w, "home", struct{}{})
}
