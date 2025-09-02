// cmd/server.go
package cmd

import (
	"fmt"
	"github/ecommerce/adapter/routes"
	"github/ecommerce/config"
	"github/ecommerce/infra/postgresql"
	"github/ecommerce/middlewares"
	"log"
	"net/http"
	"strconv"
)

func CreateServer(handlers ...routes.RouteRegister) {

	cnf := config.NewConfig()

	// dsn := "postgres://postgres:1234@localhost:5432/ecommerce?sslmode=disable"
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cnf.DbUser, cnf.DbPass, cnf.DbHost, cnf.DbPort, cnf.DbName)
	conn, err := postgresql.DbConnection(&dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Server starting...at 9090")
	mux := http.NewServeMux()
	manager := middlewares.NewManager()
	manager.Use(middlewares.MiddlwareTest1, middlewares.MiddlewareTest2, middlewares.Logger, middlewares.CorsMiddleware)

	for _, h := range handlers {
		h.RegisterRoutes(mux, manager)
	}
	str := ":" + strconv.Itoa(cnf.HttpPORT)
	err = http.ListenAndServe(str, mux)
	if err != nil {
		log.Fatal(err)
	}
}
