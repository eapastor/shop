package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type cart struct{}

var Cart cart = cart{}

func (cart) Add(c *gin.Context) {

	id := c.PostForm("id")
	amp := c.PostForm("amp")

	fmt.Printf("\n ID is %s \n AMP is %s\n", id, amp)

	c.JSON(200, struct {
		Amp int `json:"amp"`
	}{22})
}

func (cart) Show(c *gin.Context) {

	c.HTML(200, "pages/cart", gin.H{"title": "Cart"})

}
