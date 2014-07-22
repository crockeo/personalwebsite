package main

import (
	"github.com/crockeo/personalwebsite/config"
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/handlers"
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
	PreRunInit()
	handlers.InitHandlersAndStart()
}
