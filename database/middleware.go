package database

import (
	"github.com/go-martini/martini"
	"log"
	"net/http"
)

// The Martini handler to use for injecting the database context
func Middleware() martini.Handler {
	db, err := OpenAndInit()

	if err != nil {
		log.Fatal(err)
	}

	return func(w http.ResponseWriter, r *http.Request, c martini.Context) {
		c.Map(&db)
	}
}
