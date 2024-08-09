package api

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"shop/models"
	"shop/services"
)

type BasketController struct {
	BasketService  *services.BasketService
	UserService    *services.UserService
	ProductService *services.ProductService
	JwtHandler     fiber.Handler
}

func NewBasketController(service *services.BasketService, userService *services.UserService, productService *services.ProductService,
	jwtHandler fiber.Handler) *BasketController {
	return &BasketController{BasketService: service, UserService: userService, ProductService: productService,
		JwtHandler: jwtHandler}
}

func BasketRegisterRoutes(r *fiber.App, basketService *services.BasketService, userService *services.UserService,
	productService *services.ProductService, jwtHandler fiber.Handler) {
	basketController := NewBasketController(basketService, userService, productService, jwtHandler)

	routes := r.Group("/baskets")
	routes.Use(jwtHandler)
	//routes.POST("/", basketController.CreateBasket)
	routes.Get("/", basketController.GetBaskets)
	routes.Get("/:id", basketController.GetBasket)
	routes.Put("/:id", basketController.UpdateBasket)
	//routes.DELETE("/:id", basketController.DeleteBasket)
	routes.Post("/add/:idbasket/:idproduct", basketController.AddProduct)
	routes.Delete("/remove/:idbasket/:idproduct", basketController.RemoveProduct)
}

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
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (controller BasketController) GetBasket(c *fiber.Ctx) error {
	id := c.Params("id")

	var basket *models.Basket
	var err error

	ctx := context.Background()

	if basket, err = controller.BasketService.GetBasketById(id, ctx); err != nil {
		if errors.Is(err, services.NotFoundBasketErr) {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		} else {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
	}
	return c.Status(http.StatusOK).JSON(basket)
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
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (controller BasketController) GetBaskets(c *fiber.Ctx) error {
	baskets := make([]models.Basket, 0)
	var err error

	limit := 1000
	ctx := context.Background()
	if baskets, err = controller.BasketService.GetManyBaskets(limit, ctx); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(http.StatusOK).JSON(baskets)
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
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (controller BasketController) UpdateBasket(c *fiber.Ctx) error {
	updatedBasket := new(models.Basket)

	// get request body
	if err := c.BodyParser(updatedBasket); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	id := c.Params("id")
	ctx := context.Background()

	err := controller.BasketService.UpdateBasket(id, updatedBasket, ctx)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.SendStatus(http.StatusOK)
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
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (controller BasketController) AddProduct(c *fiber.Ctx) error {
	idbasket := c.Params("idbasket")
	idproduct := c.Params("idproduct")

	var basket *models.Basket
	var err error

	ctx := context.Background()

	if basket, err = controller.BasketService.GetBasketById(idbasket, ctx); err != nil {
		if errors.Is(err, services.NotFoundBasketErr) {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		} else {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
	}

	var product *models.Product
	//log.Println(idproduct)
	if product, err = controller.ProductService.GetProductById(idproduct, ctx); err != nil {
		if errors.Is(err, services.NotFoundProductErr) {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		} else {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
	}

	basket.Products = append(basket.Products, *product)
	basket.TotalPrice += product.Price

	err = controller.BasketService.UpdateBasket(idbasket, basket, ctx)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(basket)
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
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (controller BasketController) RemoveProduct(c *fiber.Ctx) error {
	idbasket := c.Params("idbasket")
	idproduct := c.Params("idproduct")

	var basket *models.Basket
	var err error

	ctx := context.Background()

	if basket, err = controller.BasketService.GetBasketById(idbasket, ctx); err != nil {
		if errors.Is(err, services.NotFoundBasketErr) {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		} else {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
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
		return c.Status(http.StatusNotFound).SendString("product not found")
	}

	err = controller.BasketService.UpdateBasket(idbasket, basket, ctx)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(basket)
}
