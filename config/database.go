package config

const (
	DbLoc             = DataDirectory + "personalwebsite.sqlite3.db" // The name of the database
	SchemaRoot string = "schema/"                                    // The location of the schema
	CreateName string = "create.sql"                                 // Create schema file name
	DropName   string = "drop.sql"                                   // Drop schema file name
)
