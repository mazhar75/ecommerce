package usecase

import (
	"github/ecommerce/internal/domain/product"
	"github/ecommerce/internal/port"
)

type ProductService struct {
	Repo port.ProductRepository
}

// Constructor
func NewProductService(repo port.ProductRepository) *ProductService {
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
