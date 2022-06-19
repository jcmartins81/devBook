package rotas

import "net/http"

// Rota representa todas as rotas da API
type Route struct {
	Uri                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequiresAuthntication bool
}
