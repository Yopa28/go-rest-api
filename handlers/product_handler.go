package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"go-rest-api/models"
	"go-rest-api/repositories"
	"strconv"
)

func GetProducts(c *gin.Context) {
	products, err := repositories.GetProducts()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to retrieve products",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Products retrieved successfully",
		"data":    products,
	})
}
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
