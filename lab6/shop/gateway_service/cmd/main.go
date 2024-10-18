package main

import (
	"fmt"
	"gateway_service/additional"
	"gateway_service/api"
	"gateway_service/circuit_breaker"
	"gateway_service/config"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	_ "github.com/swaggo/files"
	"log"
	"net/http"
	"time"
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

	// URL: /swagger/index.html
	//r.Get("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.Get("/swagger/*", fiberSwagger.WrapHandler)

	//ctx := context.Background()

	//userService := services.NewUserService(mongoDb)

	JwtSecretKey := []byte("very-secret-key")
	ContextKeyUser := "user"

	jwtMiddleware := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: JwtSecretKey,
		},
		ContextKey: ContextKeyUser,
	})

	appConfig := &config.AppConfig{
		ApiPort:                 viper.GetString("API_PORT"),
		AccountServiceProxyPort: viper.GetString("ACCOUNT_SERVICE_PROXY_PORT"),
		BasketServiceProxyPort:  viper.GetString("BASKET_SERVICE_PROXY_PORT"),
		ProductServiceProxyPort: viper.GetString("PRODUCT_SERVICE_PROXY_PORT"),
	}

	userCircuitBreaker := circuit_breaker.NewCircuitBreaker(5, 3*time.Second, 3)
	basketCircuitBreaker := circuit_breaker.NewCircuitBreaker(5, 3*time.Second, 3)
	productCircuitBreaker := circuit_breaker.NewCircuitBreaker(5, 3*time.Second, 3)
	authCircuitBreaker := circuit_breaker.NewCircuitBreaker(5, 3*time.Second, 3)

	api.UserRegisterRoutes(r, jwtMiddleware, appConfig.AccountServiceProxyPort, userCircuitBreaker)
	api.BasketRegisterRoutes(r, jwtMiddleware, appConfig.BasketServiceProxyPort, basketCircuitBreaker)
	api.ProductRegisterRoutes(r, jwtMiddleware, appConfig.ProductServiceProxyPort, productCircuitBreaker)
	api.AuthRegisterRoutes(r, appConfig.AccountServiceProxyPort, authCircuitBreaker)

	apiPort := appConfig.ApiPort
	r.Listen(apiPort)
}
