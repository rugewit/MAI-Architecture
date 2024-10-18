package main

import (
	"account_service/additional"
	"account_service/api"
	"account_service/config"
	"account_service/db"
	"account_service/redis_setup"
	"account_service/services"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

func main() {
	r := fiber.New()

	err := additional.LoadViper("env/.env")
	if err != nil {
		log.Fatalln("cannot load viper")
		return
	}

	ctx := context.Background()
	mongoConfig := &config.MongoConfig{
		Port:               viper.GetString("DB_PORT"),
		User:               viper.GetString("DB_USER"),
		Password:           viper.GetString("DB_PASSWORD"),
		Name:               viper.GetString("DB_NAME"),
		Host:               viper.GetString("DB_HOST"),
		Uri:                viper.GetString("DB_URI"),
		UsersCollection:    viper.GetString("DB_USERS_COL"),
		BasketsCollection:  viper.GetString("DB_BASKETS_COL"),
		ProductsCollection: viper.GetString("DB_PRODUCTS_COL"),
	}
	mongoDb, err := db.NewMongoDb(mongoConfig)
	if err != nil {
		log.Fatalln(err)
	}
	defer mongoDb.Close(ctx)

	useRedis, err := strconv.ParseBool(viper.GetString("USE_REDIS"))
	if err != nil {
		log.Fatal(err)
	}
	_ = &config.RedisConfig{
		RedisAliveTimeSec: viper.GetString("REDIS_ALIVE_TIME_SEC"),
		ConnectionUri:     viper.GetString("CONNECTION_URI"),
		UseRedis:          useRedis,
	}

	redisAccount := redis_setup.NewRedisAccount()

	userService := services.NewUserService(mongoDb)
	basketService := services.NewBasketService(mongoDb)

	api.UserRegisterRoutes(r, userService, basketService, redisAccount)
	api.AuthRegisterRoutes(r, userService, basketService, redisAccount)

	appConfig := &config.AppConfig{ApiPort: viper.GetString("API_PORT")}

	apiPort := appConfig.ApiPort
	if err != nil {
		log.Fatal(err)
		return
	}
	r.Listen(apiPort)
}
