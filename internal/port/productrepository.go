package port

import "github/ecommerce/internal/domain/product"

type ProductRepository interface {
	GetById(id int) (product.Product, error)
	GetAll() ([]product.Product, error)
	InsertProduct(p product.Product) error
}
