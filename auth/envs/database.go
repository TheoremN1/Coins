package envs

import "os"

type DatabaseEnv struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

var databaseEnv DatabaseEnv

func init() {
	databaseEnv.Host = os.Getenv("DATABASE_HOST")
	databaseEnv.Port = os.Getenv("DATABASE_PORT")
	databaseEnv.Name = os.Getenv("DATABASE_NAME")
	databaseEnv.User = os.Getenv("DATABASE_USER")
	databaseEnv.Password = os.Getenv("DATABASE_PASSWORD")
}

func GetDatabaseEnv() DatabaseEnv {
	return databaseEnv
}
