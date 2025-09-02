package postgresql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func DbConnection(dsn *string) (*sql.DB, error) {

	db, err := sql.Open("postgres", *dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL successfully 🚀")
	return db, nil
}
