package config

type AppConfig struct {
	ApiPort string
}

type PostgresConfig struct {
	Port     string
	User     string
	Password string
	Name     string
}
