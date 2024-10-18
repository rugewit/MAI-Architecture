package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

type AuthController struct {
	ProxyAddress string
}

func NewAuthController(proxyAddress string) *AuthController {
	return &AuthController{ProxyAddress: proxyAddress}
}

func AuthRegisterRoutes(r *fiber.App, proxyAddress string) {
	authController := NewAuthController(proxyAddress)
	address := authController.ProxyAddress

	routes := r.Group("/auth")
	routes.Post("/register", proxy.Forward(address+"/register"))
	routes.Post("/login", proxy.Forward(address+"/login"))
}
