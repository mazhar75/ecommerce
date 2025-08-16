package port

import "github/ecommerce/internal/domain"

type ProductRepository interface {
	GetById(id int) (domain.Product, error)
	GetAll() ([]domain.Product, error)
}
