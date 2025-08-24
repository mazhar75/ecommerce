package product_handlers

import (
	"github/ecommerce/usecase"
)

type ProductHandler struct {
	Service *usecase.ProductService
}

// Constructor
func NewProductHandler(service *usecase.ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}
