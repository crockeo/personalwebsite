package database

import (
	"errors"
	"github.com/crockeo/personalwebsite/config"
	"github.com/crockeo/personalwebsite/helpers"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func makeAuthTable(db *sqlx.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS auth (id INTEGER NOT NULL PRIMARY KEY, username TEXT NOT NULL, password TEXT NOT NULL)")

	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT * FROM auth")

	if err != nil {
		rows.Close()
		return err
	} else if !rows.Next() {
		_, err = db.Exec("INSERT INTO auth(id, username, password) values(1, 'admin', 'password')")

		if err != nil {
			return err
		}
	} else {
		rows.Close()
	}

	return nil
}

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
func GetAuth(db *sqlx.DB) (*Auth, error) {
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

// Quickly getting the current auth
func QuickGetAuth() *Auth {
	db := QuickOpenDB()
	defer db.Close()

	auth, err := GetAuth(db)

	if err != nil {
		panic(err)
	}

	return auth
}

// Changing the current auth
func ChangeAuth(db *sqlx.DB, auth *Auth) error {
	_, err := db.Exec("DELETE FROM auth")

	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO auth(id, username, password) values(1, $1, $2)", auth.Username, auth.Password)

	if err != nil {
		return err
	}

	return nil
}

// Quickly changing the current auth
func QuickChangeAuth(auth *Auth) {
	db := QuickOpenDB()
	defer db.Close()

	err := ChangeAuth(db, auth)

	if err != nil {
		panic(err)
	}
}
