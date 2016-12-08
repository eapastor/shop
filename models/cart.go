package models

import "github.com/eapastor/shop/utils"

type cart struct {
	ID       string
	Contents map[*Product]int
	Total    struct {
		Price float32
		Amp   int
	}
}

func NewCart() *cart {
	var c = new(cart)
	c.Contents = make(map[*Product]int)
	c.ID = utils.GenerateId()
	c.Total.Price = 0.00
	c.Total.Amp = 0
	return c
}

func (c *cart) AddProduct(product *Product, amp int) {
	c.Contents[product] += amp
	c.Total.Amp += amp
//	c.Total.Price += product.Price
}
