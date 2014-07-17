package database

import (
	"database/sql"
)

const (
	authTableName string = "auth"
	authTable     string = "CREATE TABLE " + authTableName + " (id INTEGER NOT NULL PRIMARY KEY, username TEXT, password TEXT)"
)

// The auth type
type Auth struct {
	Id       int    // The id for the auth
	Username string // The username for the auth
	Password string // The password for the auth
}

// Converting the auth to a string
func (auth *Auth) String() string {
	return auth.Username + "|" + auth.Password
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
func GetAuth(db *sql.DB) *Auth {
	row := db.QueryRow("SELECT * FROM auth")

	if row == nil {
		return nil
	}

	var id int
	var username string
	var password string

	row.Scan(&id, &username, &password)

	return &Auth{
		Id:       id,
		Username: username,
		Password: password,
	}
}

// Changing the current auth
func ChangeAuth(db *sql.DB, auth *Auth) error {
	_, err := db.Exec("DELETE FROM auth")

	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO auth(id, username, password) values(1, ?, ?)", auth.Username, auth.Password)

	return err
}
