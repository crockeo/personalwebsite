package errors

import "github.com/crockeo/personalwebsite/helpers"

// A generic error handler
func ErrorHandler(status int, message string) (int, string) {
	return status, helpers.RenderPageUnsafe("error", struct {
		Status  int
		Message string
	}{
		Status:  status,
		Message: message,
	})
}
