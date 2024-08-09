package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"shop/hash"
	jwtauth "shop/jwt"
	"shop/models"
	"shop/services"
	"time"
)

type AuthController struct {
	UserService   *services.UserService
	BasketService *services.BasketService
}

func NewAuthController(userService *services.UserService, basketService *services.BasketService) *AuthController {
	return &AuthController{UserService: userService, BasketService: basketService}
}

func AuthRegisterRoutes(r *fiber.App, userService *services.UserService, basketService *services.BasketService) {
	basketController := NewAuthController(userService, basketService)

	routes := r.Group("/auth")
	routes.Post("/register", basketController.Register)
	routes.Post("/login", basketController.Login)
}

// @Summary Resiger an user
// @Tags Auth API
// @Description Resiger an user
// @ID register
// @Accept  json
// @Produce  json
// @Param input body models.SignUpUser true "user"
// @Success 200 {object} models.User "OK"
// @Failure 400 {object} error "Bad request"
// @Router /auth/register [post]
func (controller AuthController) Register(c *fiber.Ctx) error {
	regReq := models.SignUpUser{}
	if err := c.BodyParser(&regReq); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	ctx := context.Background()
	_, err := controller.UserService.GetUserByLogin(regReq.Login, ctx)
	if !(err != nil && errors.Is(err, services.NotFoundUserErr)) {
		return c.Status(http.StatusBadRequest).SendString("user already exists")
	}

	newUser := &regReq

	hashedPass, err := hash.HashPassword(newUser.Password)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	insertUser := &models.User{
		Name:     newUser.Name,
		Lastname: newUser.Lastname,
		Password: hashedPass,
		Login:    newUser.Login,
	}

	if err := controller.UserService.InsertUser(insertUser, ctx); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())

	}

	// create basket for user
	newBasket := new(models.Basket)
	newBasket.UserId = insertUser.Id
	newBasket.TotalPrice = 0

	if err := controller.BasketService.InsertBasket(newBasket, ctx); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	insertUser.BasketId = newBasket.Id

	return c.Status(http.StatusCreated).JSON(insertUser)
}

// @Summary Login an user
// @Tags Auth API
// @Description Login an user
// @ID login
// @Accept  json
// @Produce  json
// @Param input body models.LogInUser true "user"
// @Success 200 {object} string "OK"
// @Failure 400 {object} error "Bad request"
// @Failure 404 {object} error "Bad request"
// @Router /auth/login [post]
func (controller AuthController) Login(c *fiber.Ctx) error {
	regReq := models.LogInUser{}
	if err := c.BodyParser(&regReq); err != nil {
		return c.Status(http.StatusBadRequest).SendString("cannot parse request")
	}

	ctx := context.Background()
	user, err := controller.UserService.GetUserByLogin(regReq.Login, ctx)
	if err != nil && errors.Is(err, services.NotFoundUserErr) {
		return c.Status(http.StatusNotFound).SendString("user is not found")
	}

	reqPassword, _ := hash.HashPassword(regReq.Password)
	if hash.CheckPasswordHash(reqPassword, user.Password) {
		return c.Status(http.StatusBadRequest).SendString("wrong password")
	}

	payload := jwt.MapClaims{
		"sub": user.Login,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(jwtauth.JwtSecretKey)
	if err != nil {
		fmt.Errorf("%s JWT token signing\n", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(http.StatusOK).SendString(t)
}
