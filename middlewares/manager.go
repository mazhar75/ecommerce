package middlewares

import "net/http"

type Middleware func(next http.Handler) http.Handler

type Manager struct {
	middlewaresList []Middleware
}

func NewManager() *Manager {
	return &Manager{
		middlewaresList: make([]Middleware, 0),
	}
}
func (mngr *Manager) Use(middleware ...Middleware) {
	mngr.middlewaresList = append(mngr.middlewaresList, middleware...)
}
func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next

	for _, middleware := range mngr.middlewaresList {
		n = middleware(n)
	}
	for _, middleware := range middlewares {
		n = middleware(n)
	}
	return n
}
