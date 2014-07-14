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

// Functions initialize the project structure
func PreRunInit() {
	os.Mkdir(config.DataDirectory, 775)
	os.Mkdir(config.PostsDir, 775)

	ioutil.WriteFile(config.AuthLoc, []byte(defaultAuth), 664)
	ioutil.WriteFile(config.PostIndexLoc, []byte("0"), 664)
}

// The main function
func main() {
	PreRunInit()
	handlers.InitHandlersAndStart()
}
