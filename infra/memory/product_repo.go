package memory

import (
	"errors"
	"github/ecommerce/domain/product"
)

type ProductRepo struct {
	ProductList []product.Product
}

// compile-time check
var _ product.ProductRepository = &ProductRepo{}

func (r *ProductRepo) GetById(id int) (product.Product, error) {
	for _, p := range r.ProductList {
		if p.ProductId == id {
			return p, nil
		}
	}
	return product.Product{}, errors.New("no product found with given id")
}

func (r *ProductRepo) GetAll() ([]product.Product, error) {
	return r.ProductList, nil
}

func (r *ProductRepo) InsertProduct(p product.Product) error {

	p.ProductId = len(r.ProductList) + 1
	r.ProductList = append(r.ProductList, p)
	return nil
}
