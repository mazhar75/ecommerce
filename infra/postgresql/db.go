package postgresql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func DbConnection(dsn *string) (*sql.DB, error) {

	if db != nil {
		return db, nil
	}
	db, err := sql.Open("postgres", *dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL successfully 🚀")
	return db, nil
}
