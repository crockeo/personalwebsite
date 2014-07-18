package database

import (
	"database/sql"
	"github.com/crockeo/personalwebsite/config"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

// Opening the production database connection
func openPro(url string) (*sql.DB, error) {
	return sql.Open("postgres", url)
}

// Opening the dev database connection
func openDev(pat string) (*sql.DB, error) {
	return sql.Open("sqlite3", pat)
}

// Opening the default database
func OpenDefaultDatabase() (*sql.DB, error) {
	url := os.Getenv("DATABASE_URL")

	if url == "" {
		return openDev(config.DbLoc)
	} else {
		return openPro(url)
	}
}

// Creating the database schema
func CreateDatabaseSchema(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS auth (id SERIAL, username TEXT NOT NULL, password TEXT NOT NULL)")

	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT * FROM auth")

	if err != nil {
		rows.Close()
		return err
	} else if !rows.Next() {
		_, err = db.Exec("INSERT INTO auth(username, password) values('admin', 'password')")

		if err != nil {
			return err
		}
	} else {
		rows.Close()
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS posts (id SERIAL, title TEXT NOT NULL, author TEXT NOT NULL, body TEXT NOT NULL, written TEXT NOT NULL)")

	if err != nil {
		return err
	}

	return nil
}
