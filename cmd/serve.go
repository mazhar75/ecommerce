// cmd/server.go
package cmd

import (
	"fmt"
	"github/ecommerce/adapter/routes"
	"github/ecommerce/middlewares"
	"log"
	"net/http"
	"strconv"
)

func CreateServer(handlers ...routes.RouteRegister) {

	fmt.Println("Server starting...at 9090")
	mux := http.NewServeMux()
	manager := middlewares.NewManager()
	manager.Use(middlewares.MiddlwareTest1, middlewares.MiddlewareTest2, middlewares.Logger, middlewares.CorsMiddleware)

	for _, h := range handlers {
		h.RegisterRoutes(mux, manager)
	}
	str := ":" + strconv.Itoa(9090)
	err := http.ListenAndServe(str, mux)
	if err != nil {
		log.Fatal(err)
	}
}
