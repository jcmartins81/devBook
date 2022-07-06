package routes

import (
	"devBook/api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da API
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

func ConfigRouter(r *mux.Router) *mux.Router {
	routes := routesUsers

	routes = append(routes, rotaLogin)

	for _, route := range routes {
		if route.RequiresAuthentication {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Autenticar(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
