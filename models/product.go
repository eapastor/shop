package models

type Product struct {
	ID          string
	Image       string
	Title       string
	Description string
	Price       float32
}

type Products []Product

func (p *Products) Add(id, image, title, description string, price float32) {

	*p = append(*p, Product{id, image, title, description, price})

}
