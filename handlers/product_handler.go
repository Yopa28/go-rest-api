package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"go-rest-api/models"
	"go-rest-api/repositories"
	"strconv"
)

// GetProducts godoc
// @Summary Get all products
// @Description Get products with pagination, search and sorting
// @Tags Products
// @Produce json
// @Param page query int false "Page Number"
// @Param limit query int false "Items Per Page"
// @Param search query string false "Search Product Name"
// @Param sort query string false "Sort Field (id,name,price,stock)"
// @Param order query string false "Sort Order (asc,desc)"
// @Success 200 {object} map[string]interface{}
// @Security BearerAuth
// @Router /products [get]
func GetProducts(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	search := c.DefaultQuery("search", "")
	sort := c.DefaultQuery ("sort", "id")
	order := c.DefaultQuery ("order","asc")

	products, err := repositories.GetProducts(
	limit, offset, search, sort, order)
	
		allowedSorts := map[string]bool{
		"id" : true,
		"name" : true,
		"price" : true,
		"stock" : true,
	}
	if !allowedSorts [sort]{
		sort = "id"
	}

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to retrieve products",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Products retrieved successfully",
		"page":    page,
		"limit":   limit,
		"data":    products,
	})
}


// GetProductByID godoc
// @Summary Get product by ID
// @Description Retrieve a single product by its ID
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Security BearerAuth
// @Router /products/{id} [get]
func GetProductByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid product ID",
		})
		return
	}

	product, err := repositories.GetProductByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "Product not found",
			})
			return
		}
		c.JSON(500, gin.H{
			"message": "Failed to retrieve product",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Product retrieved successfully",
		"data":    product,
	})
}

// CreateProduct godoc
// @Summary Create Product
// @Description Create a new product (Admin Only)
// @Tags Products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product Data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Security BearerAuth
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	var newProduct models.Product

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	if newProduct.Name == "" {
		c.JSON(400, gin.H{
			"message": "Product name is required",
		})
		return
	}

	if newProduct.Price <= 0 {
		c.JSON(400, gin.H{
			"message": "Product price must be greater than 0",
		})
		return
	}

	if newProduct.Stock < 0 {
		c.JSON(400, gin.H{
			"message": "Product stock cannot be negative",
		})
		return
	}

	err := repositories.CreateProduct(&newProduct)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to create product",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Product created successfully",
		"data":    newProduct,
	})

}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Only admin can update product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Updated product data"
// @Success 200 {object} map[string]interface{}
// @Security BearerAuth
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid product ID",
		})
		return
	}

	var updatedProduct models.Product

	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	if updatedProduct.Name == "" {
		c.JSON(400, gin.H{
			"message": "Product name is required",
		})
		return
	}

	if updatedProduct.Price <= 0 {
		c.JSON(400, gin.H{
			"message": "Product price must be greater than 0",
		})
		return
	}

	if updatedProduct.Stock < 0 {
		c.JSON(400, gin.H{
			"message": "Product stock cannot be negative",
		})
		return
	}

	updatedProduct.ID = id

	err = repositories.UpdateProduct(&updatedProduct)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "Product not found",
			})
			return
		}
	}

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to update product",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Product updated successfully",
		"data":    updatedProduct,
	})
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Only admin can delete product
// @Tags Products
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Security BearerAuth
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid product ID",
		})
		return
	}

	err = repositories.DeleteProduct(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "Product not found",
			})
			return
		}
		c.JSON(500, gin.H{
			"message": "Failed to delete product",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Product deleted successfully",
	})
}
