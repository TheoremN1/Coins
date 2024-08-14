package envs

import "os"

var jwtSecret string

func init() {
	jwtSecret = os.Getenv("JWT_SECRET")
}

func GetJwtSecret() string {
	return jwtSecret
}
