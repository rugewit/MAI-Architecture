package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

type BasketController struct {
	JwtHandler   fiber.Handler
	ProxyAddress string
}

func NewBasketController(
	jwtHandler fiber.Handler, proxyAddress string) *BasketController {
	return &BasketController{JwtHandler: jwtHandler, ProxyAddress: proxyAddress}
}

func BasketRegisterRoutes(r *fiber.App, jwtHandler fiber.Handler, proxyAddress string) {
	basketController := NewBasketController(jwtHandler, proxyAddress)
	address := basketController.ProxyAddress

	routes := r.Group("/baskets")
	routes.Use(jwtHandler)
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
	routes.Put("/:id", func(c *fiber.Ctx) error {
		url := address + "/" + c.Params("id")
		if err := proxy.Do(c, url); err != nil {
			return err
		}
		// Remove Server header from response
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})
	routes.Post("/add/:idbasket/:idproduct", func(c *fiber.Ctx) error {
		url := address + "/add/" + c.Params("idbasket") + "/" + c.Params("idproduct")
		if err := proxy.Do(c, url); err != nil {
			return err
		}
		// Remove Server header from response
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})
	routes.Delete("/remove/:idbasket/:idproduct", func(c *fiber.Ctx) error {
		url := address + "/remove/" + c.Params("idbasket") + "/" + c.Params("idproduct")
		if err := proxy.Do(c, url); err != nil {
			return err
		}
		// Remove Server header from response
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})
}
