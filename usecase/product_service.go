package usecase

import (
	"github/ecommerce/domain/product"
)

type ProductService struct {
	Repo product.ProductRepository
}

// Constructor
func NewProductService(repo product.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

// GetAllProducts retrieves all products via the repository
func (s *ProductService) GetAllProducts() ([]product.Product, error) {
	return s.Repo.GetAll()
}

func (s *ProductService) GetByID(id int) (product.Product, error) {
	return s.Repo.GetById(id)
}
func (s *ProductService) InsertProduct(p product.Product) error {
	return s.Repo.InsertProduct(p)
}
func (s *ProductService) UpdateProduct(p product.Product) error {
	return s.Repo.UpdateProduct(p)
}
