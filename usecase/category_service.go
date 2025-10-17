package usecase

import (
	"github/ecommerce/domain/product"
)

type CategoryService struct {
	Repo product.CategoryRepository
}

func NewCategoryService(repo product.CategoryRepository) *CategoryService {
	return &CategoryService{Repo: repo}
}

func (s *CategoryService) InsertCategory(c product.Category) (product.Category, error) {
	return s.Repo.InsertCategory(c)
}
