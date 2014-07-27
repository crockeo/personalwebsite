package database

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	"github.com/crockeo/personalwebsite/newdatabase/schema"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

type DB struct {
	gorp.DbMap
}

// Opening the database
func open() (*DB, error) {
	url := os.Getenv("DATABASE_URL")

	if url == "" {
		db, err := sql.Open("sqlite3", "testdb.sqlite3.db")

		if err != nil {
			return nil, err
		}

		return &DB{gorp.DbMap{
			Db:      db,
			Dialect: gorp.SqliteDialect{},
		}}, nil
	} else {
		db, err := sql.Open("postgres", url)

		if err != nil {
			return nil, err
		}

		return &DB{gorp.DbMap{
			Db:      db,
			Dialect: gorp.PostgresDialect{},
		}}, nil
	}
}

// Initializing the database
func initialize(db *DB) error {
	db.AddTableWithName(schema.Auth{}, "auths").SetKeys(false, "Username", "Password")
	db.AddTableWithName(schema.Course{}, "courses").SetKeys(false, "SerTitle")
	db.AddTableWithName(schema.Post{}, "posts").SetKeys(true, "Id")
	db.AddTableWithName(schema.Screenshot{}, "screenshots").SetKeys(false, "Title")
	db.AddTableWithName(schema.Project{}, "projects").SetKeys(false, "Title")

	return db.CreateTablesIfNotExists()
}

// Opening an initialized database
func OpenAndInit() (*DB, error) {
	db, err := open()

	if err != nil {
		return nil, err
	}

	err = initialize(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}
