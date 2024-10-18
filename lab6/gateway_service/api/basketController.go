package api

import (
	"gateway_service/circuit_breaker"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"net/http"
)

type BasketController struct {
	JwtHandler     fiber.Handler
	ProxyAddress   string
	CircuitBreaker *circuit_breaker.CircuitBreaker
}

func NewBasketController(
	jwtHandler fiber.Handler, proxyAddress string, circuitBreaker *circuit_breaker.CircuitBreaker) *BasketController {
	return &BasketController{JwtHandler: jwtHandler, ProxyAddress: proxyAddress, CircuitBreaker: circuitBreaker}
}

func BasketRegisterRoutes(r *fiber.App, jwtHandler fiber.Handler, proxyAddress string, circuitBreaker *circuit_breaker.CircuitBreaker) {
	basketController := NewBasketController(jwtHandler, proxyAddress, circuitBreaker)
	address := basketController.ProxyAddress

	routes := r.Group("/baskets")
	routes.Use(jwtHandler)
	//routes.Get("/", proxy.Forward(address+"/"))
	routes.Get("/", func(c *fiber.Ctx) error {
		url := address + "/"
		err := basketController.CircuitBreaker.Call(func() error {
			return proxy.Do(c, url)
		})
		if err != nil {
			return c.Status(http.StatusServiceUnavailable).SendString("Service unavailable due to circuit breaker")
		}
		// Remove Server header from response
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})

	routes.Get("/:id", func(c *fiber.Ctx) error {
		url := address + "/" + c.Params("id")
		err := basketController.CircuitBreaker.Call(func() error {
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
		err := basketController.CircuitBreaker.Call(func() error {
			return proxy.Do(c, url)
		})
		if err != nil {
			return c.Status(http.StatusServiceUnavailable).SendString("Service unavailable due to circuit breaker")
		}
		// Remove Server header from response
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})

	routes.Post("/add/:idbasket/:idproduct", func(c *fiber.Ctx) error {
		url := address + "/add/" + c.Params("idbasket") + "/" + c.Params("idproduct")
		err := basketController.CircuitBreaker.Call(func() error {
			return proxy.Do(c, url)
		})
		if err != nil {
			return c.Status(http.StatusServiceUnavailable).SendString("Service unavailable due to circuit breaker")
		}
		// Remove Server header from response
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})

	routes.Delete("/remove/:idbasket/:idproduct", func(c *fiber.Ctx) error {
		url := address + "/remove/" + c.Params("idbasket") + "/" + c.Params("idproduct")
		err := basketController.CircuitBreaker.Call(func() error {
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
