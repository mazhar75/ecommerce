package usecase

import (
    "github/ecommerce/internal/port"
    "github/ecommerce/internal/domain/product"
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
func (s *ProductService)GetByID(id int)(product.Product,error){
    return s.Repo.GetById()     
}