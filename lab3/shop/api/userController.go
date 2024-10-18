package api

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"shop/hash"
	"shop/models"
	"shop/services"
)

type UserController struct {
	UserService   *services.UserService
	BasketService *services.BasketService
}

func NewUserController(service *services.UserService, basketService *services.BasketService) *UserController {
	return &UserController{UserService: service, BasketService: basketService}
}

func UserRegisterRoutes(r *gin.Engine, userService *services.UserService, basketService *services.BasketService) {
	userController := NewUserController(userService, basketService)

	routes := r.Group("/users")
	routes.POST("/", userController.CreateUser)
	routes.GET("/", userController.GetUsers)
	routes.GET("/:id", userController.GetUser)
	routes.POST("/pattern-search", userController.PatternSearchUsers)
	routes.PUT("/:id", userController.UpdateUser)
	routes.DELETE("/:id", userController.DeleteUser)
	routes.GET("/basket/:id", userController.GetUserBasket)
	routes.GET("/login-search", userController.GetUserByLogin)
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
func (controller UserController) CreateUser(c *gin.Context) {
	newUser := new(models.SignUpUser)

	if err := c.BindJSON(newUser); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	hashedPass, err := hash.HashPassword(newUser.Password)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	insertUser := &models.User{
		Name:     newUser.Name,
		Lastname: newUser.Lastname,
		Password: hashedPass,
		Login:    newUser.Login,
	}

	ctx := context.Background()
	var userId primitive.ObjectID
	if userId, err = controller.UserService.InsertUser(insertUser, ctx); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	insertUser.Id = userId

	// create basket for user
	newBasket := new(models.Basket)
	newBasket.UserId = insertUser.Id
	newBasket.TotalPrice = 0
	newBasket.Products = make([]models.Product, 0)

	var basketId primitive.ObjectID
	if basketId, err = controller.BasketService.InsertBasket(newBasket, ctx); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	insertUser.BasketId = basketId
	newBasket.Id = basketId

	// update basketId
	if err := controller.UserService.UpdateUser(insertUser.Id.Hex(), insertUser, ctx); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusCreated, insertUser)
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
func (controller UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	var user *models.User
	var err error
	ctx := context.Background()
	if user, err = controller.UserService.GetUserById(id, ctx); err != nil {
		if errors.Is(err, services.NotFoundUserErr) {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
		return
	}
	c.JSON(http.StatusOK, user)
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
func (controller UserController) GetUsers(c *gin.Context) {
	users := make([]models.User, 0)
	var err error

	limit := 1000
	ctx := context.Background()
	if users, err = controller.UserService.GetManyUsers(limit, ctx); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, users)
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
func (controller UserController) UpdateUser(c *gin.Context) {
	updatedUser := new(models.SignUpUser)

	// get request body
	if err := c.BindJSON(updatedUser); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id := c.Param("id")
	ctx := context.Background()

	previousUser, err := controller.UserService.GetUserById(id, ctx)
	if err != nil {
		if errors.Is(err, services.NotFoundUserErr) {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
		return
	}

	previousHashedPass := previousUser.Password
	IsPassSame := hash.CheckPasswordHash(updatedUser.Password, previousHashedPass)
	// password has changed
	// it has to be hashed
	if !IsPassSame {
		hashedPass, err := hash.HashPassword(updatedUser.Password)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
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
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
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
func (controller UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()
	if err := controller.UserService.DeleteUser(id, ctx); err != nil {
		if errors.Is(err, services.NotFoundUserErr) {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
		return
	}
	c.JSON(http.StatusOK, nil)
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
func (controller UserController) PatternSearchUsers(c *gin.Context) {
	users := make([]models.User, 0)
	var err error

	request := models.PatternSearchRequest{}

	// get request body
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	limit := 1000
	ctx := context.Background()
	if users, err = controller.UserService.PatternSearchUsers(request.NamePattern, request.LastNamePattern, limit, ctx); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, users)
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
func (controller UserController) GetUserBasket(c *gin.Context) {
	id := c.Param("id")
	var user *models.User
	var err error
	ctx := context.Background()
	//log.Println("before getting user")
	if user, err = controller.UserService.GetUserById(id, ctx); err != nil {
		if errors.Is(err, services.NotFoundUserErr) {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
		return
	}

	var basket *models.Basket

	//log.Printf("id=%s\n", user.BasketId.Hex())
	//log.Println("before getting basket")
	if basket, err = controller.BasketService.GetBasketById(user.BasketId.Hex(), ctx); err != nil {
		if errors.Is(err, services.NotFoundBasketErr) {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
		return
	}

	c.JSON(http.StatusOK, basket)
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
func (controller UserController) GetUserByLogin(c *gin.Context) {
	login := c.Query("login")
	if login == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "login is required"})
		return
	}

	ctx := context.Background()
	user, err := controller.UserService.GetUserByLogin(login, ctx)
	if err != nil {
		if errors.Is(err, services.NotFoundUserErr) {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
		return
	}
	c.JSON(http.StatusOK, user)
}
