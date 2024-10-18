package api

import (
	"account_service/hash"
	jwtauth "account_service/jwt"
	"account_service/models"
	"account_service/redis_setup"
	"account_service/services"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"time"
)

type AuthController struct {
	UserService   *services.UserService
	BasketService *services.BasketService
	RedisAcc      *redis_setup.RedisAccount
}

func NewAuthController(userService *services.UserService, basketService *services.BasketService, redisAccount *redis_setup.RedisAccount) *AuthController {
	return &AuthController{UserService: userService, BasketService: basketService, RedisAcc: redisAccount}
}

func AuthRegisterRoutes(r *fiber.App, userService *services.UserService, basketService *services.BasketService, redisAccount *redis_setup.RedisAccount) {
	basketController := NewAuthController(userService, basketService, redisAccount)

	//routes := r.Group("/auth")
	r.Post("/register", basketController.Register)
	r.Post("/login", basketController.Login)
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

	var userId primitive.ObjectID
	if userId, err = controller.UserService.InsertUser(insertUser, ctx); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	insertUser.Id = userId

	// create basket for user
	newBasket := new(models.Basket)
	newBasket.UserId = insertUser.Id
	newBasket.TotalPrice = 0
	newBasket.Products = make([]models.Product, 0)

	var basketId primitive.ObjectID
	if basketId, err = controller.BasketService.InsertBasket(newBasket, ctx); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	newBasket.Id = basketId
	insertUser.BasketId = newBasket.Id

	// update basketID
	if err := controller.UserService.UpdateUser(insertUser.Id.Hex(), insertUser, ctx); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	// Write through (Сквозная запись)
	// push into redis
	if controller.RedisAcc.Config.UseRedis {
		jsonUser, err := json.Marshal(insertUser)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
		timeExpStr := controller.RedisAcc.Config.RedisAliveTimeSec
		timeExp, err := strconv.Atoi(timeExpStr)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
		controller.RedisAcc.Client.Set(ctx, insertUser.Id.Hex(), jsonUser, time.Duration(timeExp)*time.Second)
	}

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
