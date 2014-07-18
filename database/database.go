package database

import (
	"database/sql"
	"github.com/crockeo/personalwebsite/config"
	_ "github.com/mattn/go-sqlite3"
)

// Opening a database connection
func OpenDatabase(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", name)

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// Opening the default database
func OpenDefaultDatabase() (*sql.DB, error) {
	return OpenDatabase(config.DbLoc)
}

// Creating the database schema
func CreateDatabaseSchema(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS auth (id INTEGER NOT NULL PRIMARY KEY, username TEXT NOT NULL, password TEXT NOT NULL)")

	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT * FROM auth")

	if err != nil {
		rows.Close()
		return err
	} else if !rows.Next() {
		_, err = db.Exec("INSERT INTO auth(username, password) values(\"admin\", \"password\")")

		if err != nil {
			return err
		}
	} else {
		rows.Close()
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS posts (id INTEGER NOT NULL PRIMARY KEY, title TEXT NOT NULL, author TEXT NOT NULL, body TEXT NOT NULL, written TEXT NOT NULL)")

	if err != nil {
		return err
	}

	return nil
}
