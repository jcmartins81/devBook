package routes

import (
	"devBook/api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CriarUsuario,
		RequiresAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.BuscarUsuarios,
		RequiresAuthentication: false,
	},
	{
		URI:                   "/users/{userid}",
		Method:                http.MethodGet,
		Function:              controllers.BuscarUsuario,
		RequiresAuthentication: false,
	},
	{
		URI:                   "/users/{userid}",
		Method:                http.MethodPut,
		Function:              controllers.AtualizarUsuario,
		RequiresAuthentication: false,
	},
	{

		URI:                   "/users/{userid}",
		Method:                http.MethodDelete,
		Function:              controllers.DeletarUsuario,
		RequiresAuthentication: false,
	},
}
