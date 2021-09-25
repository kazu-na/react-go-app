package products

import (
	"net/http"

	"api/domain/products"
	"api/utils/errors"

	"github.com/gin-gonic/gin"
)

// GetProduct - Get product by product id
func GetProduct(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not yet implement!")
}

// CreateProduct - Create product
func CreateProduct(c *gin.Context) {
	var product products.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusCreated, product)
}
