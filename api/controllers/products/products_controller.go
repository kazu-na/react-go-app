package products

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"api/domain/products"

	"github.com/gin-gonic/gin"
)

// GetProduct - Get product by product id
func GetProduct(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not yet implement!")
}

// CreateProduct - Create product
func CreateProduct(c *gin.Context) {
	var product products.Product
	// ginでは、ユーザーリクエストの情報がcontext(c)に入るため、c.Request.Bodyで取り出す
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if err := json.Unmarshal(bytes, &product); err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(string(bytes))
	fmt.Println(err)
	c.JSON(http.StatusOK, product)
}
