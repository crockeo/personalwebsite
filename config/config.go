package config

const (
	DataDirectory string = "data/" // Directory that holds all of the other configurations
)

func InDir(path string) string { return DataDirectory + path }
