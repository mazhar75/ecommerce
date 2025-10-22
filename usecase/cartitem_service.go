package usecase

import "github/ecommerce/domain/cart"

type CartItemService struct {
	Repo cart.CartItemRepository
}

func NewCartItemService(repo cart.CartItemRepository) *CartItemService {
	return &CartItemService{Repo: repo}
}
func (s *CartItemService) AddProductToCart(user_id int, product_id int) error {
	return s.Repo.AddProductToCart(user_id, product_id)
}
func (s *CartItemService) ChangeQuantity(user_id int, product_id int, quantity int) error {
	return s.Repo.ChangeQuantity(user_id, product_id, quantity)
}
func (s *CartItemService) DeleteProductFromCart(user_id int, product_id int) error {
	return s.Repo.DeleteProductFromCart(user_id, product_id)
}
