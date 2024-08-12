package envs

import "os"

type ServerEnv struct {
	Host string
	Port string
}

var serverEnv ServerEnv

func init() {
	serverEnv.Host = os.Getenv("SERVER_HOST")
	serverEnv.Port = os.Getenv("SERVER_PORT")
}

func GetServerEnv() ServerEnv {
	return serverEnv
}
