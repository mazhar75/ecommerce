// domain/service.go
package routes

import (
	"github/ecommerce/middlewares"
	"net/http"
)

type RouteRegister interface {
	RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager)
}
