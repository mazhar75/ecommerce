package usecase

import (
	"errors"
	"github/ecommerce/domain/checkout"
	"github/ecommerce/domain/order"
)

type OrderService struct {
	OrderRepo     order.OrderRepository
	OrderItemRepo order.OrderItemsRepository
	RedisRepo     checkout.CheckoutRedisRepository
}

func NewOrderService(orderRepo order.OrderRepository, orderItemRepo order.OrderItemsRepository, redisRepo checkout.CheckoutRedisRepository) *OrderService {
	return &OrderService{OrderRepo: orderRepo, OrderItemRepo: orderItemRepo, RedisRepo: redisRepo}
}
func (s *OrderService) OrderPlacement(user_id int, cart_id int) (order.Orders, error) {
	isExist, p, err := s.RedisRepo.CheckRedisExistance(cart_id)
	if err != nil {
		return order.Orders{}, err
	}
	if !isExist {
		return order.Orders{}, errors.New("Please select items from cart and click checkout")
	}

	ord, err := s.OrderRepo.AddOrder(user_id, p.TotalPrice, "pending")
	if err != nil {
		return order.Orders{}, err
	}
	ids := make([]int, 0)
	q := make([]int, 0)
	prices := make([]float64, 0)
	for _, cf := range p.ProductList {
		ids = append(ids, cf.ProductId)
		q = append(q, cf.Quantity)
		prices = append(prices, float64(cf.Price))
	}
	err = s.OrderItemRepo.AddItems(ord.OrderId, ids, q, prices)
	if err != nil {
		return ord, err
	}
	return ord, nil

}
