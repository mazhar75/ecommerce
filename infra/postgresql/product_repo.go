package postgresql

import (
	"errors"
	"fmt"
	"github/ecommerce/config"
	"github/ecommerce/domain/product"
)

type ProductRepo struct {
	ProductList []product.Product
}

// compile-time check
var _ product.ProductRepository = &ProductRepo{}

func (r *ProductRepo) GetById(id int) (product.Product, error) {
	for _, p := range r.ProductList {
		if p.ProductId == id {
			return p, nil
		}
	}
	return product.Product{}, errors.New("no product found with given id")
}

func (r *ProductRepo) GetAll() ([]product.Product, error) {
	return r.ProductList, nil
}

func (r *ProductRepo) InsertProduct(p product.Product) error {

	dsn := config.DbString()
	db, err := DbConnection(dsn)
	if err != nil {
		fmt.Println("Error conecting db")
		return err
	}
	query := `
    INSERT INTO product (category_id, name, description, type, price, img_url)
    VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING product_id;
    `
	var id int
	err = db.QueryRow(query, p.CategoryId, p.Name, p.Description, p.Type, p.Price, p.ImgUrl).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Inserted Product ID:", id)
	return nil
}
