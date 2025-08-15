package port 
import (
	"github/ecommerce/internal/domain/product"
)

type ProductRepository interface{
	  GetById(id product.Product.Id)(product.Product,error)
	  GetAll()([]product.Product,error)
}