package api

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"shop/hash"
	jwtauth "shop/jwt"
	"shop/models"
	"shop/services"
)

type UserController struct {
	UserService   *services.UserService
	BasketService *services.BasketService
	JwtHandler    fiber.Handler
}

func NewUserController(service *services.UserService, basketService *services.BasketService, jwtHandler fiber.Handler) *UserController {
	return &UserController{UserService: service, BasketService: basketService, JwtHandler: jwtHandler}
}

func UserRegisterRoutes(r *fiber.App, userService *services.UserService, basketService *services.BasketService, jwtHandler fiber.Handler) {
	userController := NewUserController(userService, basketService, jwtHandler)

	routes := r.Group("/users")
	routes.Use(jwtHandler)
	routes.Post("/", userController.CreateUser)
	routes.Get("/", userController.GetUsers)
	routes.Get("/:id", userController.GetUser)
	routes.Post("/pattern-search", userController.PatternSearchUsers)
	routes.Put("/:id", userController.UpdateUser)
	routes.Delete("/:id", userController.DeleteUser)
	routes.Get("/basket/:id", userController.GetUserBasket)
	routes.Get("/login-search", userController.GetUserByLogin)
}

// @Summary Create an user
// @Tags User API
// @Description Create an user
// @ID create-user
// @Accept  json
// @Produce  json
// @Param input body models.SignUpUser true "user"
// @Success 201 {object} models.User "created"
// @Failure 400 {object} error "Bad request"
// @Router /users/ [post]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (controller UserController) CreateUser(c *fiber.Ctx) error {
	_, ok := jwtauth.JwtPayloadFromRequest(c)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	newUser := new(models.SignUpUser)

	if err := c.BodyParser(newUser); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

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

	ctx := context.Background()
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

// @Summary Get an user
// @Tags User API
// @Description Get an user
// @ID get-user
// @Accept  json
// @Produce  json
// @Param id path string true "user ID"
// @Success 200 {object} models.User "OK"
// @Failure 400 {object} error "Bad request"
// @Failure 404 {object} error "Not found"
// @Router /users/{id} [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (controller UserController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user *models.User
	var err error
	ctx := context.Background()
	if user, err = controller.UserService.GetUserById(id, ctx); err != nil {
		if errors.Is(err, services.NotFoundUserErr) {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		} else {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
	}
	return c.Status(http.StatusOK).JSON(user)
}

// @Summary Get Users
// @Tags User API
// @Description Get Users
// @ID get-users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User "OK"
// @Failure 400 {object} error "Bad request"
// @Router /users [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (controller UserController) GetUsers(c *fiber.Ctx) error {
	_, ok := jwtauth.JwtPayloadFromRequest(c)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	users := make([]models.User, 0)
	var err error

	limit := 1000
	ctx := context.Background()
	if users, err = controller.UserService.GetManyUsers(limit, ctx); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())

	}
	return c.Status(http.StatusOK).JSON(users)
}

// @Summary Update an user
// @Tags User API
// @Description Update an user
// @ID update-user
// @Accept  json
// @Produce  json
// @Param id path string true "user ID"
// @Param input body models.SignUpUser true "updated user"
// @Success 200 {object} models.User "OK"
// @Failure 400 {object} error "Bad request"
// @Failure 404 {object} error "Not found"
// @Router /users/{id} [put]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (controller UserController) UpdateUser(c *fiber.Ctx) error {
	updatedUser := new(models.SignUpUser)

	// get request body
	if err := c.BodyParser(updatedUser); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	id := c.Params("id")
	ctx := context.Background()

	previousUser, err := controller.UserService.GetUserById(id, ctx)
	if err != nil {
		if errors.Is(err, services.NotFoundUserErr) {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		} else {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
	}

	previousHashedPass := previousUser.Password
	IsPassSame := hash.CheckPasswordHash(updatedUser.Password, previousHashedPass)
	// password has changed
	// it has to be hashed
	if !IsPassSame {
		hashedPass, err := hash.HashPassword(updatedUser.Password)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
		updatedUser.Password = hashedPass
	}

	insertUser := &models.User{
		Id:           previousUser.Id,
		Name:         updatedUser.Name,
		Lastname:     updatedUser.Lastname,
		Password:     updatedUser.Password,
		Login:        updatedUser.Login,
		CreationDate: previousUser.CreationDate,
	}

	err = controller.UserService.UpdateUser(id, insertUser, ctx)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.SendStatus(http.StatusOK)
}

// @Summary Delete a user
// @Tags User API
// @Description Delete a user
// @ID delete-user
// @Accept  json
// @Produce  json
// @Param id path string true "user ID"
// @Success 200 {object} string "OK"
// @Failure 404 {object} error "Not found"
// @Router /users/{id} [delete]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (controller UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := context.Background()
	if err := controller.UserService.DeleteUser(id, ctx); err != nil {
		if errors.Is(err, services.NotFoundUserErr) {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		} else {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
	}
	return c.SendStatus(http.StatusOK)
}

// @Summary Pattern Search
// @Tags User API
// @Description Pattern Search. % The percent sign represents zero, one, or multiple characters. _ The underscore sign represents one, single character
// @ID pattern-search-users
// @Accept  json
// @Produce  json
// @Param input body models.PatternSearchRequest true "pattern search request"
// @Success 200 {array} models.User "OK"
// @Failure 400 {object} error "Bad request"
// @Router /users/pattern-search [post]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (controller UserController) PatternSearchUsers(c *fiber.Ctx) error {
	users := make([]models.User, 0)
	var err error

	request := models.PatternSearchRequest{}

	// get request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	limit := 1000
	ctx := context.Background()
	if users, err = controller.UserService.PatternSearchUsers(request.NamePattern, request.LastNamePattern, limit, ctx); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(http.StatusOK).JSON(users)
}

// @Summary Get an user basket
// @Tags User API
// @Description Get an user basket
// @ID get-user-basket
// @Accept  json
// @Produce  json
// @Param id path string true "user ID"
// @Success 200 {object} models.Basket "OK"
// @Failure 400 {object} error "Bad request"
// @Failure 404 {object} error "Not found"
// @Router /users/basket/{id} [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (controller UserController) GetUserBasket(c *fiber.Ctx) error {
	id := c.Params("id")
	var user *models.User
	var err error
	ctx := context.Background()
	//log.Println("before getting user")
	if user, err = controller.UserService.GetUserById(id, ctx); err != nil {
		if errors.Is(err, services.NotFoundUserErr) {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		} else {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
	}

	var basket *models.Basket

	//log.Printf("id=%s\n", user.BasketId.Hex())
	//log.Println("before getting basket")
	if basket, err = controller.BasketService.GetBasketById(user.BasketId.Hex(), ctx); err != nil {
		if errors.Is(err, services.NotFoundBasketErr) {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		} else {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
	}

	return c.Status(http.StatusOK).JSON(basket)
}

// @Summary Get an user by login
// @Tags User API
// @Description Get an user by login
// @ID get-user-by-login
// @Accept  json
// @Produce  json
// @Param login query string true "user login"
// @Success 200 {object} models.User "OK"
// @Failure 400 {object} error "Bad request"
// @Failure 404 {object} error "Not found"
// @Router /users/login-search [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (controller UserController) GetUserByLogin(c *fiber.Ctx) error {
	login := c.Query("login")
	if login == "" {
		return c.Status(http.StatusBadRequest).SendString("login is required")
	}

	ctx := context.Background()
	user, err := controller.UserService.GetUserByLogin(login, ctx)
	if err != nil {
		if errors.Is(err, services.NotFoundUserErr) {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		} else {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
	}
	return c.Status(http.StatusOK).JSON(user)
}
