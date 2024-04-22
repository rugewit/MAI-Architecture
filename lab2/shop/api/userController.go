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
	routes.GET("/", userController.GetUsers)
	routes.GET("/:id", userController.GetUser)
	routes.PUT("/:id", userController.UpdateUser)
	routes.DELETE("/:id", userController.DeleteUser)
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
	}

	ctx := context.Background()
	if err := controller.UserService.InsertUser(insertUser, ctx); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
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
// @Failure 404 {object} error "Not found"
// @Router /users/{id} [get]
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

// @Summary Get Users
// @Tags User API
// @Description Get Users
// @ID get-users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User "OK"
// @Failure 404 {object} error "Not found"
// @Router /users [get]
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
		c.AbortWithError(http.StatusNotFound, err)
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
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}
