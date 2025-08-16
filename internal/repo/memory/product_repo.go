package memory

import (
	"errors"
	"github/ecommerce/internal/domain"
	"github/ecommerce/internal/port"
)

type ProductRepo struct {
	ProductList []domain.Product
}

// compile-time check
var _ port.ProductRepository = &ProductRepo{}

func (r *ProductRepo) GetById(id int) (domain.Product, error) {
	for _, p := range r.ProductList {
		if p.Id == id {
			return p, nil
		}
	}
	return domain.Product{}, errors.New("no product found with given id")
}

func (r *ProductRepo) GetAll() ([]domain.Product, error) {
	return r.ProductList, nil
}

func (r *ProductRepo) InsertProduct(p domain.Product) error {
	r.ProductList = append(r.ProductList, p)
	return nil
}
