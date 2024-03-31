package api

import (
	"context"
	"net/http"
	"shop/hash"
	"shop/models"
	"shop/services"

	"github.com/gin-gonic/gin"
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
	routes.GET("/", userController.GetUsers)
	routes.GET("/:id", userController.GetUser)
	routes.PUT("/:id", userController.UpdateUser)
	routes.DELETE("/:id", userController.DeleteUser)
}

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
	insertUser := new(models.User)

	insertUser.Name = newUser.Name
	insertUser.Lastname = newUser.Lastname
	insertUser.Password = hashedPass

	ctx := context.Background()
	if err := controller.UserService.InsertUser(insertUser, ctx); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, insertUser)
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

	insertUser := new(models.User)
	insertUser.Name = updatedUser.Name
	insertUser.Lastname = updatedUser.Lastname
	insertUser.Password = updatedUser.Password
	insertUser.Id = previousUser.Id
	insertUser.BasketId = previousUser.BasketId
	insertUser.CreationDate = previousUser.CreationDate

	err = controller.UserService.UpdateUser(id, insertUser, ctx)
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
