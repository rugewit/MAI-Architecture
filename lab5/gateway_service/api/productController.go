package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

type ProductController struct {
	JwtHandler   fiber.Handler
	ProxyAddress string
}

func NewProductController(jwtHandler fiber.Handler, proxyAddress string) *ProductController {
	return &ProductController{JwtHandler: jwtHandler, ProxyAddress: proxyAddress}
}

func ProductRegisterRoutes(r *fiber.App, jwtHandler fiber.Handler, proxyAddress string) {
	productController := NewProductController(jwtHandler, proxyAddress)
	address := productController.ProxyAddress

	routes := r.Group("/products")
	routes.Use(jwtHandler)
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
}
