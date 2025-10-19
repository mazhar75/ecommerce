package usecase

import "github/ecommerce/domain/cart"

type CartService struct {
	Repo cart.CartRepository
}

func NewCartService(repo cart.CartRepository) *CartService {
	return &CartService{Repo: repo}
}
func (s *CartService) GetCartByUserId(user_id int) ([]cart.CartItems, error) {
	return s.Repo.GetCartByUserId(user_id)
}
