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

func (r *CheckoutRepo) CheckRedisExistance(cart_id int) (bool, checkout.CheckoutFinal, error) {
	key := "cart:" + strconv.Itoa(cart_id)
	exists, err := r.Redis.Exists(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return false, checkout.CheckoutFinal{}, err
	}
	if exists == 1 {
		val, err := rdb.Get(ctx, key).Result()
		if err == nil {
			var stored checkout.CheckoutFinal
			json.Unmarshal([]byte(val), &stored)
			fmt.Printf("Retrieved Cart: %+v\n", stored)
			return true, stored, nil
		}
		return true, checkout.CheckoutFinal{}, err
	}
	return false, checkout.CheckoutFinal{}, nil
}

func (r *CheckoutRepo) InsertAllProductsIntoRedis(products []checkout.Checkout, cart_id int) (checkout.CheckoutFinal, error) {
	key := "cart:" + strconv.Itoa(cart_id)
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
		return checkout.CheckoutFinal{}, err
	}
	r.Redis.Set(ctx, key, data, time.Hour)
	fmt.Println("Inserted Successfull")
	return checkoutFinal, nil

}
