package config

type AppConfig struct {
	ApiPort                 string
	AccountServiceProxyPort string
	BasketServiceProxyPort  string
	ProductServiceProxyPort string
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
