package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/eapastor/shop/testData"
)

func IndexHandler(c *gin.Context) {

	c.HTML(200, "pages/index", gin.H{
		"title":    "Index",
		"products": testData.Products,
	})

}
