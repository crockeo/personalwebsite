package database

import (
	"database/sql"
	"errors"
	"github.com/crockeo/personalwebsite/config"
	"github.com/crockeo/personalwebsite/helpers"
	"net/http"
)

// The auth type
type Auth struct {
	Id       int    // The id for the auth
	Username string // The username for the auth
	Password string // The password for the auth
}

// Checking if an auth equals another auth
func (auth1 *Auth) Equal(auth2 *Auth) bool {
	return auth1.Username == auth2.Username && auth1.Password == auth2.Password
}

// Converting the auth to a string
func (auth *Auth) String() string {
	return auth.Username + "|" + auth.Password
}

// Converting the auth to a secure string
func (auth *Auth) SecureString() string {
	return helpers.HashString(auth.String())
}

// Making a cookie from an auth
func (auth *Auth) MakeCookie() *http.Cookie {
	return &http.Cookie{
		Name:   config.AuthName,
		Value:  auth.SecureString(),
		Path:   "/",
		Domain: "",

		Secure:   false,
		HttpOnly: true,
		Raw:      config.AuthName + "=" + auth.String(),
		Unparsed: []string{config.AuthName + "=" + auth.String()},
	}
}

// Making a new auth
func MakeNewAuth(username string, password string) *Auth {
	return &Auth{
		Id:       1,
		Username: username,
		Password: password,
	}
}

// Getting the current auth
func GetAuth(db *sql.DB) (*Auth, error) {
	row := db.QueryRow("SELECT * FROM auth;")

	if row == nil {
		return nil, errors.New("Authorization does not exist.")
	}

	var id int
	var username string
	var password string

	row.Scan(&id, &username, &password)

	return &Auth{
		Id:       id,
		Username: username,
		Password: password,
	}, nil
}

// Changing the current auth
func ChangeAuth(db *sql.DB, auth *Auth) error {
	stmt, err := db.Prepare(`
	DELETE FROM auth;
	INSERT INTO auth(id, username, password) values(1, '?', '?');
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(auth.Username, auth.Password)

	return err
}
