package blog

import (
	"io/ioutil"
	"strings"
)

type Auth struct {
	Username string
	Password string
}

// Loading an Auth from a file
func LoadAuth(path string) (*Auth, error) {
	val, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	vals := strings.Split(string(val), "\n")

	if len(vals) != 2 {
		return nil, nil
	} else {
		auth := new(Auth)

		auth.Username = vals[0]
		auth.Password = vals[1]

		return auth, nil
	}
}

// Loading the default Auth
func LoadDefaultAuth() (*Auth, error) {
	return LoadAuth(AuthLoc)
}

// Checking if two Auths are equal
func (auth *Auth) Equal(auth2 *Auth) bool {
	return auth.Username == auth2.Username &&
		auth.Password == auth2.Password
}

// Converting an Auth to a string
func (auth *Auth) String() string {
	return auth.Username + "\n" + auth.Password
}
