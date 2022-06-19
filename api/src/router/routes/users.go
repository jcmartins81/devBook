package rotas

import (
	"devBook/api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CriarUsuario,
		RequiresAuthntication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.BuscarUsuarios,
		RequiresAuthntication: false,
	},
	{
		uri:                   "/users/{userid}",
		method:                http.MethodGet,
		Function:              controllers.BuscarUsuario,
		requiresauthntication: false,
	},
	{
		uri:                   "/users/{userid}",
		method:                http.MethodPut,
		Function:              controllers.AtualizarUsuario,
		requiresauthntication: false,
	},
	{

		URI:                   "/users/{userid}",
		Method:                http.MethodDelete,
		Function:              controllers.DeletarUsuario,
		RequiresAuthntication: false,
	},
}
