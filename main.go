package main

import (
	"github.com/crockeo/personalwebsite/config"
	"github.com/crockeo/personalwebsite/handlers"
	"io/ioutil"
	"os"
)

const (
	defaultAuth string = "admin|password"
)

// Checking if a file exists
func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// Functions initialize the project structure
func PreRunInit() {
	os.Mkdir(config.DataDirectory, 0775)
	os.Mkdir(config.PostsDir, 0775)

	if !exists(config.AuthLoc) {
		ioutil.WriteFile(config.AuthLoc, []byte(defaultAuth), 0664)
	}

	if !exists(config.PostIndexLoc) {
		ioutil.WriteFile(config.PostIndexLoc, []byte("0"), 0664)
	}
}

// The main function
func main() {
	PreRunInit()
	handlers.InitHandlersAndStart()
}
