package main

import (
	"github.com/crockeo/personalwebsite/config"
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/handlers"
	"github.com/go-martini/martini"
	"os"
)

// Functions initialize the project structure
func PreRunInit() {
	os.Mkdir(config.DataDirectory, 0775)
	db, err := database.OpenDefaultDatabase()

	if err != nil {
		panic(err)
	}

	err = database.CreateDatabaseSchema(db)

	if err != nil {
		panic(err)
	}

	db.Close()
}

func main() {
	go PreRunInit()

	m := martini.Classic()

	handlers.InitAdminHandlers(m)
	handlers.InitBlogHandlers(m)
	handlers.InitErrorHandlers(m)
	handlers.InitHomeHandlers(m)
	handlers.InitProjectHandlers(m)

	m.Run()
}
