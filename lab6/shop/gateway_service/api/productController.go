package api

import (
	"gateway_service/circuit_breaker"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"net/http"
)

type ProductController struct {
	JwtHandler     fiber.Handler
	ProxyAddress   string
	CircuitBreaker *circuit_breaker.CircuitBreaker
}

func NewProductController(jwtHandler fiber.Handler, proxyAddress string, circuitBreaker *circuit_breaker.CircuitBreaker) *ProductController {
	return &ProductController{JwtHandler: jwtHandler, ProxyAddress: proxyAddress, CircuitBreaker: circuitBreaker}
}

func ProductRegisterRoutes(r *fiber.App, jwtHandler fiber.Handler, proxyAddress string, circuitBreaker *circuit_breaker.CircuitBreaker) {
	productController := NewProductController(jwtHandler, proxyAddress, circuitBreaker)
	address := productController.ProxyAddress

	routes := r.Group("/products")
	routes.Use(jwtHandler)

	//routes.Post("/", proxy.Forward(address+"/"))
	routes.Post("/", func(c *fiber.Ctx) error {
		url := address + "/"
		err := productController.CircuitBreaker.Call(func() error {
			return proxy.Do(c, url)
		})
		if err != nil {
			return c.Status(http.StatusServiceUnavailable).SendString("Service unavailable due to circuit breaker")
		}
		return nil
	})

	//routes.Get("/", proxy.Forward(address+"/"))
	routes.Get("/", func(c *fiber.Ctx) error {
		url := address + "/"
		err := productController.CircuitBreaker.Call(func() error {
			return proxy.Do(c, url)
		})
		if err != nil {
			return c.Status(http.StatusServiceUnavailable).SendString("Service unavailable due to circuit breaker")
		}
		return nil
	})

	routes.Get("/:id", func(c *fiber.Ctx) error {
		url := address + "/" + c.Params("id")
		err := productController.CircuitBreaker.Call(func() error {
			return proxy.Do(c, url)
		})
		if err != nil {
			return c.Status(http.StatusServiceUnavailable).SendString("Service unavailable due to circuit breaker")
		}
		// Remove Server header from response
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})

	routes.Put("/:id", func(c *fiber.Ctx) error {
		url := address + "/" + c.Params("id")
		err := productController.CircuitBreaker.Call(func() error {
			return proxy.Do(c, url)
		})
		if err != nil {
			return c.Status(http.StatusServiceUnavailable).SendString("Service unavailable due to circuit breaker")
		}
		// Remove Server header from response
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})

	routes.Delete("/:id", func(c *fiber.Ctx) error {
		url := address + "/" + c.Params("id")
		err := productController.CircuitBreaker.Call(func() error {
			return proxy.Do(c, url)
		})
		if err != nil {
			return c.Status(http.StatusServiceUnavailable).SendString("Service unavailable due to circuit breaker")
		}
		// Remove Server header from response
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})
}
