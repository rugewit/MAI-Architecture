package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "github.com/swaggo/files"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"shop/additional"
	"shop/api"
	"shop/config"
	"shop/dataGenerator"
	"shop/db"
	_ "shop/docs"
	"shop/services"
)

// @title Shop
// @version 1.0
// description shop

// @host localhost:8081
// @BasePath /
func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/gen", func(c *gin.Context) {
		data_generator.GenFakeData(10, 100)
		c.JSON(http.StatusOK, gin.H{"generated": "okay"})
	})

	// URL: /swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := additional.LoadViper("env/.env")
	if err != nil {
		log.Fatalln("cannot load viper")
		return
	}

	ctx := context.Background()
	pgConfig := &config.PostgresConfig{
		Port:     viper.GetString("DB_PORT"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		Name:     viper.GetString("DB_NAME"),
	}
	postgres, err := db.NewPostgresDb(ctx, pgConfig)
	if err != nil {
		log.Fatalln(err)
	}
	defer postgres.Close(ctx)

	userService := services.NewUserService(postgres)
	api.UserRegisterRoutes(r, userService)

	appConfig := &config.AppConfig{ApiPort: viper.GetString("API_PORT")}

	apiPort := appConfig.ApiPort
	if err != nil {
		log.Fatal(err)
		return
	}
	r.Run(apiPort)
}
