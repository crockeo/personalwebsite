package utils

import (
	"github.com/crockeo/personalwebsite/config"
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/database/schema"
	"net/http"
)

// Checking if two auths are equivalent
func AuthEquals(auth1 schema.Auth, auth2 schema.Auth) bool {
	return auth1.Username == auth2.Username &&
		auth1.Password == auth2.Password
}

// Converting an auth to a string
func AuthString(auth schema.Auth) string {
	return "username=" + auth.Username + "password=" + auth.Password
}

// Making an http.Cookie from an auth
func MakeCookie(auth schema.Auth) *http.Cookie {
	return &http.Cookie{
		Name:   config.AuthName,
		Value:  AuthString(auth),
		Path:   "/",
		Domain: "",
	}
}

// Getting the current auth
func GetAuth(db *database.DB) (schema.Auth, error) {
	var auth schema.Auth
	err := db.SelectOne(&auth, "SELECT * FROM auths;")

	if err != nil {
		return ChangeAuth(db, schema.Auth{
			Username: "admin",
			Password: "password",
		})
	}

	return auth, err
}

// Changing the current auth
func ChangeAuth(db *database.DB, auth schema.Auth) (schema.Auth, error) {
	tx, err := db.Begin()
	if err != nil {
		return schema.Auth{}, err
	}

	tx.Exec("DELETE FROM auths;")
	tx.Insert(&auth)

	err = tx.Commit()

	if err != nil {
		return schema.Auth{}, err
	}

	return auth, nil
}
