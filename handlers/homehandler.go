package handlers

import (
	"net/http"
	"github.com/crockeo/personalwebsite/helpers"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	helpers.SendPage(w, "home", struct{}{})
}
