package redis

import (
	"encoding/json"
	"fmt"
	"github/ecommerce/domain/checkout"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type CheckoutRepo struct {
	Redis *redis.Client
}

func NewCheckoutRepo(redis *redis.Client) *CheckoutRepo {
	return &CheckoutRepo{Redis: redis}
}

var _ checkout.CheckoutRedisRepository = &CheckoutRepo{}

func (r *CheckoutRepo) InsertAllProductsIntoRedis(products []checkout.Checkout, cart_id int) error {
	key := "cart:" + strconv.Itoa(cart_id)
	exists, err := r.Redis.Exists(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	if exists == 1 {
		r.Redis.Del(ctx, key)
	}
	totalPrice := 0
	for _, p := range products {
		totalPrice += p.Price
	}
	checkoutFinal := checkout.CheckoutFinal{
		CartId:      cart_id,
		ProductList: products,
		TotalPrice:  float64(totalPrice),
		ExpireAt:    time.Now().Add(time.Hour),
	}
	data, err := json.Marshal(checkoutFinal)
	if err != nil {
		fmt.Println(err)
		return err
	}
	r.Redis.Set(ctx, key, data, time.Hour)
	fmt.Println("Inserted Successfull")
	return nil

}
