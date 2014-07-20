package database

import (
	"database/sql"
	"github.com/crockeo/personalwebsite/config"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

// Opening a database connection
func OpenDefaultDatabase() (*sql.DB, error) {
	url := os.Getenv("DATABASE_URL")

	if url == "" {
		return sql.Open("sqlite3", config.DbLoc)
	} else {
		return sql.Open("postgres", url)
	}
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
		_, err = db.Exec("INSERT INTO auth(id, username, password) values(1, 'admin', 'password')")

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
