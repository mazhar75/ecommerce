package memory
import (
	"github/ecommerce/internal/port/productrepository"
	"github/ecommerce/internal/domain/product"
	"errors"
)


 type ProductRepo struct{
      ProductList []product.Product
 }

 // compile-time check
var _ port.ProductRepository = &ProductRepo{}

func (r *ProductRepo) GetById(id product.Product.Id)(product.Product,error){
     for i,p:=range r.ProductList{
		if p.Id==id {
			return  p,nil
		}
	 }
	 return nil,errors.New("No product found in given id")
}
func (r *ProductRepo)GetAll()([]product.Product,error){
	  return r.ProductList,nil
}

