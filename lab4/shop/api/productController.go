package api

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"shop/models"
	"shop/services"
)

type ProductController struct {
	ProductService *services.ProductService
	JwtHandler     fiber.Handler
}

func NewProductController(service *services.ProductService, jwtHandler fiber.Handler) *ProductController {
	return &ProductController{ProductService: service, JwtHandler: jwtHandler}
}

func ProductRegisterRoutes(r *fiber.App, productService *services.ProductService, jwtHandler fiber.Handler) {
	productController := NewProductController(productService, jwtHandler)

	routes := r.Group("/products")
	routes.Use(jwtHandler)
	routes.Post("/", productController.CreateProduct)
	routes.Get("/", productController.GetProducts)
	routes.Get("/:id", productController.GetProduct)
	routes.Put("/:id", productController.UpdateProduct)
	routes.Delete("/:id", productController.DeleteProduct)

}

// @Summary Create a product
// @Tags Product API
// @Description Create a product
// @ID create-product
// @Accept  json
// @Produce  json
// @Param input body models.Product true "product"
// @Success 201 {object} models.Product "created"
// @Failure 400 {object} error "Bad request"
// @Router /products/ [post]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (controller ProductController) CreateProduct(c *fiber.Ctx) error {
	newProduct := new(models.Product)

	if err := c.BodyParser(newProduct); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	ctx := context.Background()
	if err := controller.ProductService.InsertProduct(newProduct, ctx); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(http.StatusCreated).JSON(newProduct)
}

// @Summary Get a product
// @Tags Product API
// @Description Get a product
// @ID get-product
// @Accept  json
// @Produce  json
// @Param id path string true "product ID"
// @Success 200 {object} models.Product "OK"
// @Failure 400 {object} error "Bad request"
// @Failure 404 {object} error "Not found"
// @Router /products/{id} [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (controller ProductController) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product *models.Product
	var err error
	ctx := context.Background()
	if product, err = controller.ProductService.GetProductById(id, ctx); err != nil {
		if errors.Is(err, services.NotFoundProductErr) {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		} else {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
	}
	return c.Status(http.StatusOK).JSON(product)
}

// @Summary Get Products
// @Tags Product API
// @Description Get Products
// @ID get-products
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Product "OK"
// @Failure 400 {object} error "Bad request"
// @Router /products [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (controller ProductController) GetProducts(c *fiber.Ctx) error {
	products := make([]models.Product, 0)
	var err error

	limit := 1000
	ctx := context.Background()
	if products, err = controller.ProductService.GetManyProducts(limit, ctx); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(http.StatusOK).JSON(products)
}

// @Summary Update a product
// @Tags Product API
// @Description Update a product
// @ID update-product
// @Accept  json
// @Produce  json
// @Param id path string true "product ID"
// @Param input body models.Product true "updated product"
// @Success 200 {object} models.Product "OK"
// @Failure 400 {object} error "Bad request"
// @Failure 404 {object} error "Not found"
// @Router /products/{id} [put]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (controller ProductController) UpdateProduct(c *fiber.Ctx) error {
	updatedProduct := new(models.Product)

	// get request body
	if err := c.BodyParser(updatedProduct); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	id := c.Params("id")
	ctx := context.Background()

	err := controller.ProductService.UpdateProduct(id, updatedProduct, ctx)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	return c.SendStatus(http.StatusOK)
}

// @Summary Delete a product
// @Tags Product API
// @Description Delete a product
// @ID delete-product
// @Accept  json
// @Produce  json
// @Param id path string true "product ID"
// @Success 200 {object} string "OK"
// @Failure 404 {object} error "Not found"
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (controller ProductController) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := context.Background()
	if err := controller.ProductService.DeleteProduct(id, ctx); err != nil {
		if errors.Is(err, services.NotFoundProductErr) {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		} else {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
	}
	return c.SendStatus(http.StatusOK)
}
