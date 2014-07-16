package admin

import (
	"github.com/crockeo/personalwebsite/config"
	"net/http"
)

// Making an authorizarion cookie
func MakeAuthCookie(auth Auth) *http.Cookie {
	return &http.Cookie{
		Name:   config.AuthName,
		Value:  auth.String(),
		Path:   "/",
		Domain: "",

		Secure:   false,
		HttpOnly: true,
		Raw:      config.AuthName + "=" + auth.String(),
		Unparsed: []string{config.AuthName + "=" + auth.String()}}
}

// Sending an auth cookie
func SendAuthCookie(w http.ResponseWriter, auth Auth) {
	http.SetCookie(w, MakeAuthCookie(auth))
}
