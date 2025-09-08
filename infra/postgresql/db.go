package postgresql

import (
	"database/sql"
	"fmt"
	"github/ecommerce/config"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

// GetDB returns a singleton DB connection
func GetDB() *sql.DB {
	once.Do(func() {
		dsn := config.DbString()
		var err error
		db, err = sql.Open("postgres", *dsn)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		// Check the connection
		if err = db.Ping(); err != nil {
			log.Fatalf("Database unreachable: %v", err)
		}

		fmt.Println("Database connection established")
	})
	return db
}
