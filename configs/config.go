package configs

type LaunchConf struct {
	Server   ServerConf
	Database DatabaseConf
}

type ServerConf struct {
	Host string
	Port int
}

type DatabaseConf struct {
	Host     string
	Name     string
	Username string
	Password string
}
