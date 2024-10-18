package api

import (
	"gateway_service/additional"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

type UserController struct {
	JwtHandler   fiber.Handler
	ProxyAddress string
}

func NewUserController(jwtHandler fiber.Handler, proxyAddress string) *UserController {
	return &UserController{JwtHandler: jwtHandler, ProxyAddress: proxyAddress}
}

func UserRegisterRoutes(r *fiber.App, jwtHandler fiber.Handler, proxyAddress string) {
	userController := NewUserController(jwtHandler, proxyAddress)
	address := userController.ProxyAddress

	// use jwt or not
	err := additional.LoadViper("env/.env")
	if err != nil {
		log.Fatalln("cannot load viper")
		return
	}
	turnOffJwt, err := strconv.ParseBool(viper.GetString("TURN_OFF_ACCOUNT_JWT_FOR_TESTING"))
	if err != nil {
		log.Fatal(err)
		return
	}
	routes := r.Group("/users")
	if !turnOffJwt {
		routes.Use(jwtHandler)
	}
	routes.Post("/", proxy.Forward(address+"/"))
	routes.Get("/", proxy.Forward(address+"/"))
	routes.Get("/:id", func(c *fiber.Ctx) error {
		url := address + "/" + c.Params("id")
		if err := proxy.Do(c, url); err != nil {
			return err
		}
		// Remove Server header from response
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})
	routes.Post("/pattern-search", proxy.Forward(address+"/pattern-search"))
	routes.Put("/:id", func(c *fiber.Ctx) error {
		url := address + "/" + c.Params("id")
		if err := proxy.Do(c, url); err != nil {
			return err
		}
		// Remove Server header from response
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})
	routes.Delete("/:id", func(c *fiber.Ctx) error {
		url := address + "/" + c.Params("id")
		if err := proxy.Do(c, url); err != nil {
			return err
		}
		// Remove Server header from response
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})
	routes.Get("/basket/:id", func(c *fiber.Ctx) error {
		url := address + "/basket/" + c.Params("id")
		if err := proxy.Do(c, url); err != nil {
			return err
		}
		// Remove Server header from response
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})
	routes.Get("/login-search", proxy.Forward(address+"/login-search"))
}
