package config

type Config struct {
	Db *DbConfig
}

type DbConfig struct {
	Dialect string
	Username string
	Password string
	Name string
	Charset string
}

func GetConfig() *Config {
	return &Config{
		Db: &DbConfig {
			Dialect: "mysql",
			Username: "root",
			Password: "password",
			Name: "richard_test",
			Charset: "utf8",
		},
	}
}