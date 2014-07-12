package handlers

import (
	"fmt"
	"net/http"
)

const (
	Ip   string = "localhost"
	Port string = "8000"
)

// Converting the Ip and Port variables into
// the actual string
func ConnectionString() string {
	return Ip + ":" + Port
}

// Initializing an starting the server
func InitHandlersAndStart() {
	// Alerting the user that the server is starting
	fmt.Println("Starting server on:", ConnectionString())

	// Adding all of the handlers
	http.HandleFunc("/static/", StaticHandler)
	http.HandleFunc("/blog/"  , BlogHandler)
	http.HandleFunc("/"       , HomeHandler)

	// Starting the server
	http.ListenAndServe(ConnectionString(), nil)
}
