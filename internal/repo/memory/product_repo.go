package memory

import (
	"errors"
	"github/ecommerce/internal/domain/product"
	"github/ecommerce/internal/port"
)

type ProductRepo struct {
	ProductList []product.Product
}

// compile-time check
var _ port.ProductRepository = &ProductRepo{}

func (r *ProductRepo) GetById(id int) (product.Product, error) {
	for _, p := range r.ProductList {
		if p.Id == id {
			return p, nil
		}
	}
	return product.Product{}, errors.New("no product found with given id")
}

func (r *ProductRepo) GetAll() ([]product.Product, error) {
	return r.ProductList, nil
}

func (r *ProductRepo) InsertProduct(p product.Product) error {
	r.ProductList = append(r.ProductList, p)
	return nil
}
