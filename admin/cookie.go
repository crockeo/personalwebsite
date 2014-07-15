package admin

import (
	"github.com/crockeo/personalwebsite/config"
	"net/http"
	"time"
)

// Making an authorizarion cookie
func MakeAuthCookie(auth Auth) *http.Cookie {
	exp := time.Now().Add(2 * 24 * time.Hour)

	return &http.Cookie{
		Name:       config.AuthName,
		Value:      auth.String(),
		Path:       "/",
		Domain:     "localhost",
		Expires:    exp,
		RawExpires: exp.Format(time.UnixDate),

		MaxAge:   2 * 24 * 60 * 60,
		Secure:   true,
		HttpOnly: true,
		Raw:      config.AuthName + "=" + auth.String(),
		Unparsed: []string{config.AuthName + "=" + auth.String()}}
}
