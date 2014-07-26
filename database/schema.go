package database

import (
	"github.com/crockeo/personalwebsite/config"
	"io/ioutil"
)

// The Schema struct
type Schema struct {
	Create string // The code to be executed when you want to create the database
	Drop   string // The code to be executed when you want to drop the database
}

func (this *Schema) String() string { return "Create=" + this.Create + "|Drop=" + this.Drop }

// Loading a Schema from a file
func LoadSchema(name string) (*Schema, error) {
	genName := func(t string) string {
		return config.SchemaRoot + name + "/" + t
	}

	create, err := ioutil.ReadFile(genName(config.CreateName))

	if err != nil {
		return nil, err
	}

	drop, err := ioutil.ReadFile(genName(config.DropName))

	if err != nil {
		return nil, err
	}

	return &Schema{
		Create: string(create),
		Drop:   string(drop),
	}, nil
}

// Loading a Schema from a file and panicing on error
func MustLoadSchema(name string) *Schema {
	schema, err := LoadSchema(name)

	if err != nil {
		panic(err)
	}

	return schema
}
