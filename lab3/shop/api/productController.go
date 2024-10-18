package api

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/models"
	"shop/services"
)

type ProductController struct {
	ProductService *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{ProductService: service}
}

func ProductRegisterRoutes(r *gin.Engine, productService *services.ProductService) {
	productController := NewProductController(productService)

	routes := r.Group("/products")
	routes.POST("/", productController.CreateProduct)
	routes.GET("/", productController.GetProducts)
	routes.GET("/:id", productController.GetProduct)
	routes.PUT("/:id", productController.UpdateProduct)
	routes.DELETE("/:id", productController.DeleteProduct)
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
func (controller ProductController) CreateProduct(c *gin.Context) {
	newProduct := new(models.Product)

	if err := c.BindJSON(newProduct); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx := context.Background()
	if _, err := controller.ProductService.InsertProduct(newProduct, ctx); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, newProduct)
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
func (controller ProductController) GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product *models.Product
	var err error
	ctx := context.Background()
	if product, err = controller.ProductService.GetProductById(id, ctx); err != nil {
		if errors.Is(err, services.NotFoundProductErr) {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
		return
	}
	c.JSON(http.StatusOK, product)
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
func (controller ProductController) GetProducts(c *gin.Context) {
	products := make([]models.Product, 0)
	var err error

	limit := 1000
	ctx := context.Background()
	if products, err = controller.ProductService.GetManyProducts(limit, ctx); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, products)
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
func (controller ProductController) UpdateProduct(c *gin.Context) {
	updatedProduct := new(models.Product)

	// get request body
	if err := c.BindJSON(updatedProduct); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id := c.Param("id")
	ctx := context.Background()

	err := controller.ProductService.UpdateProduct(id, updatedProduct, ctx)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
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
func (controller ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()
	if err := controller.ProductService.DeleteProduct(id, ctx); err != nil {
		if errors.Is(err, services.NotFoundProductErr) {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
		return
	}
	c.JSON(http.StatusOK, nil)
}
