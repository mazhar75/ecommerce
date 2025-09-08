package product

type Category struct {
	CategoryId int    `json:"category_id"`
	Name       string `json:"name"`
}
type Product struct {
	ProductId   int     `json:"product_id"`
	CategoryId  int     `json:"category_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}

type ProductRepository interface {
	GetById(id int) (Product, error)
	GetAll() ([]Product, error)
	InsertProduct(p Product) error
}
