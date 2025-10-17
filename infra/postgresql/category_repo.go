package postgresql

import (
	"database/sql"

	"github/ecommerce/domain/product"
)

type CategoryRepo struct {
	DB *sql.DB
}

var _ product.CategoryRepository = &CategoryRepo{}

func NewCategoryRepo(db *sql.DB) *CategoryRepo {
	return &CategoryRepo{DB: db}
}
func (r *CategoryRepo) InsertCategory(c product.Category) (product.Category, error) {
	query := `
		INSERT INTO category (name)
		VALUES ($1)
		RETURNING category_id;
	`
	var id int
	err := r.DB.QueryRow(query, c.Name).Scan(&id)
	if err != nil {
		return c, err
	}
	return c, nil
}
