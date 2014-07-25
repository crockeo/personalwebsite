package database

import (
	"database/sql"
	"errors"
	"github.com/crockeo/personalwebsite/config"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

type initFn func(db *sql.DB) error

var (
	// An error to desginate that a row doesn't exist
	RowDoesNotExistError error = errors.New("Row does not exist.")

	// The functions to initialize the database schema
	initFns []initFn = []initFn{
		makeAuthTable,
		makePostTable,
		makeCourseTable,
		makeProjectTable,
	}
)

// Opening a database connection
func OpenDB() (*sql.DB, error) {
	url := os.Getenv("DATABASE_URL")

	if url == "" {
		return sql.Open("sqlite3", config.DbLoc)
	} else {
		return sql.Open("postgres", url)
	}
}

// Quickly opening a database connection
func QuickOpenDB() *sql.DB {
	db, err := OpenDB()

	if err != nil {
		panic(err)
	}

	return db
}

// Creating the database schema
func CreateDatabaseSchema(db *sql.DB) error {
	var err error
	for i := 0; i < len(initFns); i++ {
		err = initFns[i](db)

		if err != nil {
			return err
		}
	}

	return nil
}
