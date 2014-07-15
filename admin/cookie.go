package admin

import (
	"github.com/crockeo/personalwebsite/config"
	"net/http"
	"time"
)

// Making an authorizarion cookie
func MakeAuthCookie(auth Auth) *http.Cookie {
	cook := new(http.Cookie)

	exp := time.Now().Add(2 * 24 * time.Hour)

	cook.Name = config.AuthName
	cook.Value = auth.String()
	cook.Path = "/"
	cook.Domain = "127.0.0.1"
	cook.Expires = exp
	cook.RawExpires = exp.Format(time.UnixDate)

	cook.MaxAge = 2 * 24 * 60 * 60
	cook.Secure = false
	cook.HttpOnly = false
	cook.Raw = cook.Name + "=" + cook.Value
	cook.Unparsed = []string{cook.Raw}

	return cook
}
