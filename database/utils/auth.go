package utils

import (
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/database/schema"
)

// Getting the current auth
func GetAuth(db *database.DB) (*schema.Auth, error) {
	auth := new(schema.Auth)

	err := db.SelectOne(auth, "SELECT * FROM auths;")

	if err != nil {
		return nil, err
	}

	return auth, err
}

// Changing the current auth
func ChangeAuth(db *database.DB, auth *schema.Auth) (*schema.Auth, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	tx.Exec("DELETE FROM auths;")
	tx.Insert(auth)

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return auth, nil
}
