package jwtauth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var ContextKeyUser = "user"
var JwtSecretKey = []byte("very-secret-key")

func JwtPayloadFromRequest(c *fiber.Ctx) (jwt.MapClaims, bool) {
	jwtToken, ok := c.Context().Value(ContextKeyUser).(*jwt.Token)
	if !ok {
		fmt.Println("wrong type of JWT token in context")
		return nil, false
	}

	payload, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("wrong type of JWT token claims")
		return nil, false
	}

	return payload, true
}
