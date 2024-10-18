package api

import (
	"gateway_service/circuit_breaker"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"net/http"
)

type AuthController struct {
	ProxyAddress   string
	CircuitBreaker *circuit_breaker.CircuitBreaker
}

func NewAuthController(proxyAddress string, circuitBreaker *circuit_breaker.CircuitBreaker) *AuthController {
	return &AuthController{ProxyAddress: proxyAddress, CircuitBreaker: circuitBreaker}
}

func AuthRegisterRoutes(r *fiber.App, proxyAddress string, circuitBreaker *circuit_breaker.CircuitBreaker) {
	authController := NewAuthController(proxyAddress, circuitBreaker)
	address := authController.ProxyAddress

	routes := r.Group("/auth")
	//routes.Post("/register", proxy.Forward(address+"/register"))
	routes.Post("/register", func(c *fiber.Ctx) error {
		url := address + "/register"
		err := authController.CircuitBreaker.Call(func() error {
			return proxy.Do(c, url)
		})
		if err != nil {
			return c.Status(http.StatusServiceUnavailable).SendString("Service unavailable due to circuit breaker")
		}
		return nil
	})

	//routes.Post("/login", proxy.Forward(address+"/login"))
	routes.Post("/login", func(c *fiber.Ctx) error {
		url := address + "/login"
		err := authController.CircuitBreaker.Call(func() error {
			return proxy.Do(c, url)
		})
		if err != nil {
			return c.Status(http.StatusServiceUnavailable).SendString("Service unavailable due to circuit breaker")
		}
		return nil
	})
}
