package routes

import (
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
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
