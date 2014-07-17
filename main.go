package main

import (
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/handlers"
)

// Functions initialize the project structure
func PreRunInit() {
	db, err := database.OpenDefaultDatabase()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = database.CreateDatabaseSchema(db)

	if err != nil {
		panic(err)
	}

	db.Close()
}

// The main function
func main() {
	PreRunInit()
	handlers.InitHandlersAndStart()
}
