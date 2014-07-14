package blog

import (
	"github.com/crockeo/personalwebsite/config"
	"io/ioutil"
	"strings"
)

type Auth struct {
	Username string
	Password string
}

// Creating a new Auth
func NewAuth(username string, password string) *Auth {
	auth := new(Auth)

	auth.Username = username
	auth.Password = password

	return auth
}

// Loading an Auth from a file
func LoadAuth(path string) (*Auth, error) {
	val, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	val = val[0 : len(val)-1]
	vals := strings.Split(string(val), "|")

	if len(vals) != 2 {
		return nil, nil
	} else {
		return NewAuth(vals[0], vals[1]), nil
	}
}

// Loading the default Auth
func LoadDefaultAuth() (*Auth, error) {
	return LoadAuth(config.AuthLoc)
}

// Checking if two Auths are equal
func (auth *Auth) Equal(auth2 *Auth) bool {
	return auth.Username == auth2.Username &&
		auth.Password == auth2.Password
}

// Converting an Auth to a string
func (auth *Auth) String() string {
	return auth.Username + "|" + auth.Password
}
