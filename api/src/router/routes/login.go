package routes

import (
	"devBook/api/src/controllers"
	"net/http"
)

var rotaLogin = Route{
	URI: "/login",
	Method: http.MethodPost,
	Function: controllers.Login,
	RequiresAuthentication: false,
}
