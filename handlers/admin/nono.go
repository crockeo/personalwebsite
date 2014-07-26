package admin

import (
	"github.com/crockeo/personalwebsite/helpers"
	"net/http"
)

// Serving the no-no page
func NonoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(503)
	helpers.SendPage(w, "nono", struct{}{})
}
