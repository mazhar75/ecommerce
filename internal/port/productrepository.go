package port 
import (
	"github/ecommerce/internal/domain/product"
)

type ProductRepository interface{
	  GetAll()([]product.Product,error)
}