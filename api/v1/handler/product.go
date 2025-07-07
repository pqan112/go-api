package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"pqan.com/go-api/utils"
)

type ProductHandler struct{}

type GetProductParam struct {
	ID int `uri:"id" binding:"gt=0"`
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (u *ProductHandler) GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"messsage": "List all Products",
	})
}

func (u *ProductHandler) GetProduct(c *gin.Context) {
	var params GetProductParam
	if err := c.ShouldBindUri(&params); err != nil {
		message := "Error.Invalid_Param"
		c.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err, &message))
		return
	}

	limit := c.DefaultQuery("limit", "10")

	c.JSON(http.StatusOK, gin.H{
		"messsage": "Product detail",
		"limit":    limit,
	})
}

func (u *ProductHandler) CreateProduct(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"messsage": "Create Product",
	})
}

func (u *ProductHandler) UpdateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"messsage": "Update Product",
	})
}

func (u *ProductHandler) DeleteProduct(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"messsage": "Update Product",
	})
}
