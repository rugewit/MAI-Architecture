package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
	"product_service/additional"
	"product_service/api"
	"product_service/config"
	"product_service/db"
	"product_service/services"
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

	productService := services.NewProductService(mongoDb)

	api.ProductRegisterRoutes(r, productService)

	appConfig := &config.AppConfig{ApiPort: viper.GetString("API_PORT")}

	apiPort := appConfig.ApiPort
	if err != nil {
		log.Fatal(err)
		return
	}
	r.Listen(apiPort)
}
