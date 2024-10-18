package config

type AppConfig struct {
	ApiPort string
}

type MongoConfig struct {
	Port               string
	User               string
	Password           string
	Name               string
	Host               string
	Uri                string
	UsersCollection    string
	BasketsCollection  string
	ProductsCollection string
}

type RedisConfig struct {
	UseRedis          bool
	RedisAliveTimeSec string
	ConnectionUri     string
}
