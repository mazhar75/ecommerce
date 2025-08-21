package product_handler

import (
	"github/ecommerce/internal/usecase"
)

type ProductHandler struct {
	Service *usecase.ProductService
}

func NewProductHandler(service *usecase.ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

// Handler method
