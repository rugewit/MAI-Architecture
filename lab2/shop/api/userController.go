package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/hash"
	"shop/models"
	"shop/services"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{UserService: service}
}

func UserRegisterRoutes(r *gin.Engine, userService *services.UserService) {
	userController := NewUserController(userService)

	routes := r.Group("/users")
	routes.POST("/", userController.CreateUser)
}

func (controller UserController) CreateUser(c *gin.Context) {
	user := new(models.User)

	if err := c.BindJSON(user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	hashedPass, err := hash.HashPassword(user.Password)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user.Password = hashedPass

	ctx := context.Background()
	if err := controller.UserService.InsertUser(user, ctx); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (controller UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	var user *models.User
	var err error
	ctx := context.Background()
	if user, err = controller.UserService.GetUserById(id, ctx); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (controller UserController) GetUsers(c *gin.Context) {
	users := make([]models.User, 0)
	var err error

	limit := 1000
	ctx := context.Background()
	if users, err = controller.UserService.GetManyUsers(limit, ctx); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (controller UserController) UpdateUser(c *gin.Context) {
	updatedUser := new(models.User)

	// get request body
	if err := c.BindJSON(updatedUser); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id := c.Param("id")
	ctx := context.Background()

	previousUser, err := controller.UserService.GetUserById(id, ctx)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
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

	err = controller.UserService.UpdateUser(id, updatedUser, ctx)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (controller UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()
	if err := controller.UserService.DeleteUser(id, ctx); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}
