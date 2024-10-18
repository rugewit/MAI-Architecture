package main

import (
	"basket_service/additional"
	"basket_service/api"
	"basket_service/config"
	"basket_service/db"
	"basket_service/services"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, world")

	r := fiber.New()

	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	r.Get("/gen", func(c *fiber.Ctx) error {
		//data_generator.GenFakeData(10, 100)
		return c.SendStatus(http.StatusOK)
	})

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

	basketService := services.NewBasketService(mongoDb)
	userService := services.NewUserService(mongoDb)
	productService := services.NewProductService(mongoDb)

	api.BasketRegisterRoutes(r, basketService, userService, productService)

	appConfig := &config.AppConfig{ApiPort: viper.GetString("API_PORT")}

	apiPort := appConfig.ApiPort
	if err != nil {
		log.Fatal(err)
		return
	}
	r.Listen(apiPort)
}
