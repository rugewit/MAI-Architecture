package api

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/models"
	"shop/services"
)

type BasketController struct {
	BasketService  *services.BasketService
	UserService    *services.UserService
	ProductService *services.ProductService
}

func NewBasketController(service *services.BasketService, userService *services.UserService, productService *services.ProductService) *BasketController {
	return &BasketController{BasketService: service, UserService: userService, ProductService: productService}
}

func BasketRegisterRoutes(r *gin.Engine, basketService *services.BasketService, userService *services.UserService, productService *services.ProductService) {
	basketController := NewBasketController(basketService, userService, productService)

	routes := r.Group("/baskets")
	//routes.POST("/", basketController.CreateBasket)
	routes.GET("/", basketController.GetBaskets)
	routes.GET("/:id", basketController.GetBasket)
	routes.PUT("/:id", basketController.UpdateBasket)
	//routes.DELETE("/:id", basketController.DeleteBasket)
	routes.POST("/add/:idbasket/:idproduct", basketController.AddProduct)
	routes.DELETE("/remove/:idbasket/:idproduct", basketController.RemoveProduct)
}

// @Summary Create a basket
// @Tags Basket API
// @Description Create a basket
// @ID create-basket
// @Accept  json
// @Produce  json
// @Param input body models.SignUpUser true "user"
// @Success 201 {object} models.User "created"
// @Failure 400 {object} error "Bad request"
// @Router /users/ [post]
/*
func (controller BasketController) CreateBasket(c *gin.Context) {
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
*/

// @Summary Get a basket
// @Tags Basket API
// @Description Get a basket
// @ID get-basket
// @Accept  json
// @Produce  json
// @Param id path string true "basket ID"
// @Success 200 {object} models.Basket "OK"
// @Failure 400 {object} error "Bad request"
// @Failure 404 {object} error "Not found"
// @Router /baskets/{id} [get]
func (controller BasketController) GetBasket(c *gin.Context) {
	id := c.Param("id")

	var basket *models.Basket
	var err error

	ctx := context.Background()

	if basket, err = controller.BasketService.GetBasketById(id, ctx); err != nil {
		if errors.Is(err, services.NotFoundBasketErr) {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
		return
	}
	c.JSON(http.StatusOK, basket)
}

// @Summary Get Baskets
// @Tags Basket API
// @Description Get Baskets
// @ID get-baskets
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Basket "OK"
// @Failure 400 {object} error "Bad request"
// @Router /baskets [get]
func (controller BasketController) GetBaskets(c *gin.Context) {
	baskets := make([]models.Basket, 0)
	var err error

	limit := 1000
	ctx := context.Background()
	if baskets, err = controller.BasketService.GetManyBaskets(limit, ctx); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, baskets)
}

// @Summary Update a basket
// @Tags Basket API
// @Description Update a basket
// @ID update-basket
// @Accept  json
// @Produce  json
// @Param id path string true "basket ID"
// @Param input body models.Basket true "updated basket"
// @Success 200 {object} models.Basket "OK"
// @Failure 400 {object} error "Bad request"
// @Failure 404 {object} error "Not found"
// @Router /basket/{id} [put]
func (controller BasketController) UpdateBasket(c *gin.Context) {
	updatedBasket := new(models.Basket)

	// get request body
	if err := c.BindJSON(updatedBasket); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id := c.Param("id")
	ctx := context.Background()

	err := controller.BasketService.UpdateBasket(id, updatedBasket, ctx)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

// @Summary Delete a basket
// @Tags Basket API
// @Description Delete a basket
// @ID delete-basket
// @Accept  json
// @Produce  json
// @Param id path string true "basket ID"
// @Success 200 {object} string "OK"
// @Failure 404 {object} error "Not found"
// @Router /baskets/{id} [delete]
/*
func (controller BasketController) DeleteBasket(c *gin.Context) {
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
*/

// @Summary Add a product
// @Tags Basket API
// @Description Add a product
// @ID add-product
// @Accept  json
// @Produce  json
// @Param idbasket path string true "Basket ID"
// @Param idproduct path string true "Product ID"
// @Success 201 {object} nil "created"
// @Failure 400 {object} error "Bad request"
// @Router /baskets/add/{idbasket}/{idproduct} [post]
func (controller BasketController) AddProduct(c *gin.Context) {
	idbasket := c.Param("idbasket")
	idproduct := c.Param("idproduct")

	var basket *models.Basket
	var err error

	ctx := context.Background()

	if basket, err = controller.BasketService.GetBasketById(idbasket, ctx); err != nil {
		if errors.Is(err, services.NotFoundBasketErr) {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
	}

	var product *models.Product
	//log.Println(idproduct)
	if product, err = controller.ProductService.GetProductById(idproduct, ctx); err != nil {
		if errors.Is(err, services.NotFoundProductErr) {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
	}

	basket.Products = append(basket.Products, *product)
	basket.TotalPrice += product.Price

	err = controller.BasketService.UpdateBasket(idbasket, basket, ctx)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, basket)
}

// @Summary Remove a product
// @Tags Basket API
// @Description Remove a product
// @ID remove-product
// @Accept  json
// @Produce  json
// @Param idbasket path string true "Basket ID"
// @Param idproduct path string true "Product ID"
// @Success 200 {object} models.Basket "OK"
// @Failure 400 {object} error "Bad request"
// @Failure 404 {object} error "Not found"
// @Router /baskets/remove/{idbasket}/{idproduct} [delete]
func (controller BasketController) RemoveProduct(c *gin.Context) {
	idbasket := c.Param("idbasket")
	idproduct := c.Param("idproduct")

	var basket *models.Basket
	var err error

	ctx := context.Background()

	if basket, err = controller.BasketService.GetBasketById(idbasket, ctx); err != nil {
		if errors.Is(err, services.NotFoundBasketErr) {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
		return
	}

	// Find and remove the product from the basket
	isRemoved := false
	for i, p := range basket.Products {
		if p.Id.Hex() == idproduct {
			basket.Products = append(basket.Products[:i], basket.Products[i+1:]...)
			isRemoved = true
			basket.TotalPrice -= p.Price
			break
		}
	}
	if !isRemoved {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	err = controller.BasketService.UpdateBasket(idbasket, basket, ctx)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, basket)
}
