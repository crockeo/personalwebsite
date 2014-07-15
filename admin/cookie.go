package admin

import (
	"github.com/crockeo/personalwebsite/config"
	"net/http"
	"time"
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

// Making the delete version of a cookie
func ToDelete(cook *http.Cookie) *http.Cookie {
	exp := time.Now()

	cook.Expires = exp
	cook.RawExpires = exp.Format(time.UnixDate)
	cook.MaxAge = 0

	return cook
}
