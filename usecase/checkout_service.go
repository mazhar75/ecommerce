package usecase

import (
	"fmt"
	"github/ecommerce/domain/checkout"
)

type CheckoutService struct {
	PgRepo    checkout.CheckoutPgRepository
	RedisRepo checkout.CheckoutRedisRepository
}

func NewCheckoutService(pgrepo checkout.CheckoutPgRepository, redisrepo checkout.CheckoutRedisRepository) *CheckoutService {
	return &CheckoutService{PgRepo: pgrepo, RedisRepo: redisrepo}
}

func (s *CheckoutService) AddCheckoutfromPgtoRedis(cart_id int) (checkout.CheckoutFinal, []checkout.OutOfStock, error) {

	isExist, my_checkout, err := s.RedisRepo.CheckRedisExistance(cart_id)
	if err != nil {
		return checkout.CheckoutFinal{}, nil, err
	} else if isExist {
		return my_checkout, nil, err
	}

	p, out_stock, err := s.PgRepo.GetAllProductsFromCart(cart_id)
	if err != nil {
		return checkout.CheckoutFinal{}, out_stock, err
	}
	fmt.Println(p)
	fmt.Println(out_stock)
	finalcheckout, err := s.RedisRepo.InsertAllProductsIntoRedis(p, cart_id)
	if err != nil {
		return checkout.CheckoutFinal{}, out_stock, err
	}
	return finalcheckout, out_stock, nil
}
