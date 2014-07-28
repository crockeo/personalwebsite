package main

import (
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/handlers"
	"github.com/go-martini/martini"
)

// The main function
func main() {
	m := martini.Classic()

	m.use(database.Injector())
	handlers.InitHandlers(m)

	m.Run()
}
