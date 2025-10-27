package usecase

import "github/ecommerce/domain/checkout"

type CheckoutService struct {
	PgRepo    checkout.CheckoutPgRepository
	RedisRepo checkout.CheckoutRedisRepository
}

func NewCheckoutService(pgrepo checkout.CheckoutPgRepository, redisrepo checkout.CheckoutRedisRepository) *CheckoutService {
	return &CheckoutService{PgRepo: pgrepo, RedisRepo: redisrepo}
}

func (s *CheckoutService) AddCheckoutfromPgtoRedis(cart_id int) ([]checkout.OutOfStock, error) {
	p, out_stock, err := s.PgRepo.GetAllProductsFromCart(cart_id)
	if err != nil {
		return out_stock, err
	}
	err = s.RedisRepo.InsertAllProductsIntoRedis(p)
	if err != nil {
		return out_stock, err
	}
	return out_stock, nil

}
