package error

import "net/http"

func _404Handler(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, r, 404)
}
