package api

import (
	"gateway_service/additional"
	"gateway_service/circuit_breaker"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strconv"
)

type UserController struct {
	JwtHandler     fiber.Handler
	ProxyAddress   string
	CircuitBreaker *circuit_breaker.CircuitBreaker
}

func NewUserController(jwtHandler fiber.Handler, proxyAddress string, circuitBreaker *circuit_breaker.CircuitBreaker) *UserController {
	return &UserController{JwtHandler: jwtHandler, ProxyAddress: proxyAddress, CircuitBreaker: circuitBreaker}
}

func UserRegisterRoutes(r *fiber.App, jwtHandler fiber.Handler, proxyAddress string, circuitBreaker *circuit_breaker.CircuitBreaker) {
	userController := NewUserController(jwtHandler, proxyAddress, circuitBreaker)
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

	//routes.Post("/", proxy.Forward(address+"/"))
	routes.Post("/", func(c *fiber.Ctx) error {
		url := address + "/"
		err := userController.CircuitBreaker.Call(func() error {
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
		err := userController.CircuitBreaker.Call(func() error {
			return proxy.Do(c, url)
		})
		if err != nil {
			log.Println("Service unavailable due to circuit breaker")
			return c.Status(http.StatusServiceUnavailable).SendString("Service unavailable due to circuit breaker")
		}
		// Remove Server header from response
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})

	routes.Get("/:id", func(c *fiber.Ctx) error {
		url := address + "/" + c.Params("id")
		err := userController.CircuitBreaker.Call(func() error {
			return proxy.Do(c, url)
		})
		if err != nil {
			log.Println("Service unavailable due to circuit breaker")
			return c.Status(http.StatusServiceUnavailable).SendString("Service unavailable due to circuit breaker")
		}
		// Remove Server header from response
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})

	//routes.Post("/pattern-search", proxy.Forward(address+"/pattern-search"))
	routes.Post("/pattern-search", func(c *fiber.Ctx) error {
		url := address + "/pattern-search"
		err := userController.CircuitBreaker.Call(func() error {
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
		err := userController.CircuitBreaker.Call(func() error {
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
		err := userController.CircuitBreaker.Call(func() error {
			return proxy.Do(c, url)
		})
		if err != nil {
			return c.Status(http.StatusServiceUnavailable).SendString("Service unavailable due to circuit breaker")
		}
		// Remove Server header from response
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})

	routes.Get("/basket/:id", func(c *fiber.Ctx) error {
		url := address + "/basket/" + c.Params("id")
		err := userController.CircuitBreaker.Call(func() error {
			return proxy.Do(c, url)
		})
		if err != nil {
			return c.Status(http.StatusServiceUnavailable).SendString("Service unavailable due to circuit breaker")
		}
		// Remove Server header from response
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})

	//routes.Get("/login-search", proxy.Forward(address+"/login-search"))
	routes.Get("/login-search", func(c *fiber.Ctx) error {
		url := address + "/login-search"
		err := userController.CircuitBreaker.Call(func() error {
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
