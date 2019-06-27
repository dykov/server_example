package config

type database struct {
	Host         string
	DatabaseName string
	Username     string
	Password     string
}

type Config struct {
	Port     string
	Database *database
}

func NewConfig() *Config {
	return &Config{
		Port: "8080",
		Database: &database{
			Host:         "127.0.0.1",
			DatabaseName: "DATABASE_NAME",
			Username:     "USERNAME",
			Password:     "PASSWORD",
		},
	}
}
