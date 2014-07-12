package handlers

import (
	"os"
	"fmt"
	"net/http"
)

// Initializing an starting the server
func InitHandlersAndStart() {

	// Adding all of the handlers
	http.HandleFunc("/static/", StaticHandler)
	http.HandleFunc("/blog/"  , BlogHandler)
	http.HandleFunc("/"       , HomeHandler)

	// Starting the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Alerting the user that the server is starting
	fmt.Println("Starting server on port:", port)

	http.ListenAndServe(":" + port, nil)
}
