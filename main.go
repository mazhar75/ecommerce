package main

import (
	"github/ecommerce/adapter/handlers/auth"
	"github/ecommerce/adapter/handlers/cart_handler"
	"github/ecommerce/adapter/handlers/cart_item_handler"
	"github/ecommerce/adapter/handlers/category_handlers"
	"github/ecommerce/adapter/handlers/health_handler"
	"github/ecommerce/adapter/handlers/product_handlers"
	"github/ecommerce/cmd"
	"github/ecommerce/domain/health"
	"github/ecommerce/infra/memory"
	"github/ecommerce/infra/postgresql"
	"github/ecommerce/usecase"
)

func main() {

	db := postgresql.GetDB()
	defer db.Close()

	//auth dependencies
	authRepo := postgresql.NewUserRepo(db)
	userService := usecase.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(userService)

	//health dependencies-in-memory
	_health := memory.HealthRepo{
		Health: health.Health{
			Status:  200,
			Message: "OK",
		},
	}
	healthservice := usecase.NewHealthService(&_health)
	healthHandler := health_handler.NewHealthHandler(healthservice)

	// category dependencies
	catRepo := postgresql.NewCategoryRepo(db)
	catService := usecase.NewCategoryService(catRepo)
	catHandler := category_handlers.NewProductHandler(catService)

	//product dependencies
	productRepo := postgresql.NewProductRepo(db)
	productservice := usecase.NewProductService(productRepo)
	productHandler := product_handlers.NewProductHandler(productservice)

	//cart dependencies
	cartRepo := postgresql.NewCartRepo(db)
	cartService := usecase.NewCartService(cartRepo)
	cartHandler := cart_handler.NewCartHandler(cartService)

	//cart items depedencies
	cartItemRepo := postgresql.NewCartItemRepo(db)
	cartItemService := usecase.NewCartItemService(cartItemRepo)
	cartItemHandler := cart_item_handler.NewCartItemHandler(cartItemService)

	cmd.CreateServer(catHandler, productHandler, authHandler, healthHandler, cartHandler, cartItemHandler)

}
