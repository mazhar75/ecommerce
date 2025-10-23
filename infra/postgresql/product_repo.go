package postgresql

import (
	"database/sql"
	"errors"
	"fmt"

	"github/ecommerce/domain/product"
)

type ProductRepo struct {
	DB *sql.DB
}

// compile-time check
var _ product.ProductRepository = &ProductRepo{}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{DB: db}
}

func (r *ProductRepo) GetById(id int) (product.Product, error) {
	query := `
		SELECT product_id, category_id, name, description, type, price, img_url
		FROM product
		WHERE product_id = $1;
	`

	var p product.Product
	err := r.DB.QueryRow(query, id).Scan(
		&p.ProductId, &p.CategoryId, &p.Name, &p.Description, &p.Type, &p.Price, &p.ImgUrl,
	)
	if err == sql.ErrNoRows {
		return product.Product{}, errors.New("no product found with given id")
	}
	if err != nil {
		return product.Product{}, err
	}
	return p, nil
}

func (r *ProductRepo) GetAll() ([]product.Product, error) {
	query := `
		SELECT product_id, category_id, name, description, type, price, img_url
		FROM product;
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []product.Product
	for rows.Next() {
		var p product.Product
		err := rows.Scan(&p.ProductId, &p.CategoryId, &p.Name, &p.Description, &p.Type, &p.Price, &p.ImgUrl)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *ProductRepo) InsertProduct(p product.Product) error {

	query := `
		INSERT INTO product (category_id, name, description, type, price, img_url)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING product_id;
	`
	var id int
	err := r.DB.QueryRow(query, p.CategoryId, p.Name, p.Description, p.Type, p.Price, p.ImgUrl).Scan(&id)
	if err != nil {
		return err
	}
	query = `insert into inventory(product_id,category_id,quantity) values($1,$2,$3) returning inventory_id`
	var inv_id int
	err = r.DB.QueryRow(query, id, p.CategoryId, 10).Scan(&inv_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Inserted Product ID and inventory_id:", id, inv_id)
	return nil
}
func (r *ProductRepo) DeleteProduct(pId int) error {
	query := `DELETE FROM product WHERE product_id = $1;`
	result, err := r.DB.Exec(query, pId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("product not found")
	}
	return nil
}

func (r *ProductRepo) UpdateProduct(p product.Product) error {
	query := `
		UPDATE product
		SET category_id = $1, name = $2, description = $3, type = $4, price = $5, img_url = $6
		WHERE product_id = $7;
	`
	result, err := r.DB.Exec(query, p.CategoryId, p.Name, p.Description, p.Type, p.Price, p.ImgUrl, p.ProductId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("product not found")
	}
	return nil
}
