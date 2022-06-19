package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da API
type Route struct {
	Uri                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequiresAuthntication bool
}

func ConfigRouter(r *mux.Router) *mux.Router {
	routes := routesUsers

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
