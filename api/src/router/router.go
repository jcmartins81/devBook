package router

import "github.com/gorilla/mux"

//Gerar vai retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigRouter(r)
}
