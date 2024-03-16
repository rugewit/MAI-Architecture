package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"shop/additional"
	"shop/api"
	"shop/dataGenerator"
	"shop/db"
	"shop/services"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world1"})
	})

	r.GET("/gen", func(c *gin.Context) {
		data_generator.GenFakeData(10, 100)
		c.JSON(http.StatusOK, gin.H{"generated": "mb okay"})
	})

	err := additional.LoadViper("env/.env")
	if err != nil {
		log.Fatalln("cannot load viper")
		return
	}

	ctx := context.Background()
	postgres, err := db.NewPostgresDb(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer postgres.Close(ctx)

	userService := services.NewUserService(postgres)
	api.UserRegisterRoutes(r, userService)

	apiPort := viper.Get("API_PORT").(string)
	if err != nil {
		log.Fatal(err)
		return
	}
	r.Run(apiPort)
}
