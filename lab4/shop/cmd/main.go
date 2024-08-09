package main

import (
	"context"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	_ "github.com/swaggo/files"
	"log"
	"net/http"
	"shop/additional"
	"shop/api"
	"shop/config"
	"shop/dataGenerator"
	"shop/db"
	_ "shop/docs"
	jwtauth "shop/jwt"
	"shop/services"
)

// @title Shop
// @version 1.0
// description shop

// @host localhost:8081
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Token
func main() {
	r := fiber.New()

	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	r.Get("/gen", func(c *fiber.Ctx) error {
		data_generator.GenFakeData(10, 100)
		return c.SendStatus(http.StatusOK)
	})

	// URL: /swagger/index.html
	//r.Get("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Get("/swagger/*", fiberSwagger.WrapHandler)

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

	userService := services.NewUserService(mongoDb)
	basketService := services.NewBasketService(mongoDb)
	productService := services.NewProductService(mongoDb)

	jwtMiddleware := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: jwtauth.JwtSecretKey,
		},
		ContextKey: jwtauth.ContextKeyUser,
	})

	api.UserRegisterRoutes(r, userService, basketService, jwtMiddleware)
	api.BasketRegisterRoutes(r, basketService, userService, productService, jwtMiddleware)
	api.ProductRegisterRoutes(r, productService, jwtMiddleware)
	api.AuthRegisterRoutes(r, userService, basketService)

	appConfig := &config.AppConfig{ApiPort: viper.GetString("API_PORT")}

	apiPort := appConfig.ApiPort
	if err != nil {
		log.Fatal(err)
		return
	}
	r.Listen(apiPort)
}
